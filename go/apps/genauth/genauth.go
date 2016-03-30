// Binary genauth generates a C++ implementation and serialization from the Go
// version.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

var primitives = map[string]bool{
	"string": true,
	"bytes": true,
	"int": true,
	"int64": true,
	"bool": true,
}

type FieldType int

const (
	IdentType FieldType = iota
	StarType
	ArrayType
)

type Field struct {
	Name string
	Type FieldType
	TypeName string
}

type TypeVisitor struct {
	ConcreteTypes map[string][]Field
	InterfaceTypes map[string]bool
}

type FieldVisitor struct {
	ConcreteTypes map[string][]Field
	InterfaceTypes map[string]bool
	Name string
}

func (tv *TypeVisitor) Visit(n ast.Node) ast.Visitor {
	ts, ok := n.(*ast.TypeSpec)
	if !ok {
		return tv
	}

	name := ts.Name.Name
	switch ts.Type.(type) {
	case *ast.StructType:
		if _, ok := tv.ConcreteTypes[name]; !ok {
			tv.ConcreteTypes[name] = make([]Field, 0)
		}

		return &FieldVisitor{
			ConcreteTypes: tv.ConcreteTypes,
			InterfaceTypes: tv.InterfaceTypes,
			Name: name,
		}
	case *ast.ArrayType:
		at := ts.Type.(*ast.ArrayType)
		elt, ok := at.Elt.(*ast.Ident)
		if ok {
			tv.ConcreteTypes[name] = []Field{Field{"elt", ArrayType, elt.Name}}
		}
	case *ast.Ident:
		id := ts.Type.(*ast.Ident)
		tv.ConcreteTypes[name] = []Field{Field{"value", IdentType, id.Name}}
	case *ast.InterfaceType:
		tv.InterfaceTypes[name] = true
	}

	return tv
}

func (fv *FieldVisitor) Visit(n ast.Node) ast.Visitor {
	f, ok := n.(*ast.Field)
	if !ok {
		return fv
	}

	if len(f.Names) == 0 {
		return nil
	}

	name := strings.ToLower(f.Names[0].Name)
	st := fv.ConcreteTypes
	switch f.Type.(type) {
	case *ast.Ident:
		ident := f.Type.(*ast.Ident)
		st[fv.Name] = append(st[fv.Name], Field{name, IdentType, ident.Name})
	case *ast.StarExpr:
		star := f.Type.(*ast.StarExpr)
		ident, ok := star.X.(*ast.Ident)
		if ok {
			st[fv.Name] = append(st[fv.Name], Field{name, StarType, ident.Name})
		}
	case *ast.ArrayType:
		atype := f.Type.(*ast.ArrayType)
		elt, ok := atype.Elt.(*ast.Ident)
		if ok {
			st[fv.Name] = append(st[fv.Name], Field{name, ArrayType, elt.Name})
		}
	default:
		return fv
	}

	return &TypeVisitor{
		ConcreteTypes: fv.ConcreteTypes,
		InterfaceTypes: fv.InterfaceTypes,
	}
}

type FuncReceiverWalker struct {
	types map[string]bool
	name string
}

func (fw *FuncReceiverWalker) Visit(n ast.Node) ast.Visitor {
	fd, ok := n.(*ast.FuncDecl)
	if !ok {
		return fw
	}

	if fd.Name.Name != fw.name {
		return fw
	}

	// Record the name of the type of the first receiver.
	if len(fd.Recv.List) == 0 {
		return fw
	}

	field := fd.Recv.List[0]
	ident, ok := field.Type.(*ast.Ident)
	if !ok {
		return fw
	}

	fw.types[ident.Name] = true
	return fw
}

type Constant struct {
	Name string
	Value int
}

type ConstantVisitor struct {
	Constants []Constant
}

func (tv *ConstantVisitor) Visit(n ast.Node) ast.Visitor {
	vs, ok := n.(*ast.ValueSpec)
	if !ok {
		return tv
	}

	if len(vs.Names) == 0 {
		return tv
	}

	ident := vs.Names[0]
	if ident.Name == "_" {
		return tv
	}

	if ident.Obj == nil || ident.Obj.Data == nil || ident.Obj.Kind != ast.Con {
		return tv
	}

	value, ok := ident.Obj.Data.(int)
	if !ok {
		return tv
	}

	// Turn the constant name into a C++ constant name.
	name := "k" + strings.Title(ident.Name)

	tv.Constants = append(tv.Constants, Constant{name, value})
	return tv
}

func writeHeader(constants []Constant, types map[string][]Field, interfaces map[string]bool, formTypes map[string]bool, termTypes map[string]bool) []string {
	header := []string{
		"#include <memory>",
		"#include <vector>",
		"",
		"#include <google/protobuf/io/coded_stream.h>",
		"",
		"class LogicElement {",
		" public:",
		"  virtual bool Marshal(google::protobuf::io::CodedOutputStream* output) = 0;",
		"};",
		"",
		"class Form : public LogicElement {};",
		"class Term : public LogicElement {};",
		"",
	}

	constructor := "  %s(google::protobuf::io::CodedInputStream* input);"
	marshal := "  bool Marshal(google::protobuf::io::CodedOutputStream* output)"

	header = append(header, "enum class BinaryTags {")
	for i, constant := range constants {
		value := fmt.Sprintf("  %s = %d", constant.Name, constant.Value)
		if i < len(constants) - 1 {
			value += ","
		}
		header = append(header, value)
	}
	header = append(header, "};", "")

	for name, fields := range types {
		class := fmt.Sprintf("class %s", name)
		if _, isForm := formTypes[name]; isForm {
			class += ": public Form"
		} else if _, isTerm := termTypes[name]; isTerm {
			class += ": public Term"
		}

		class += " {"
		header = append(header, class, " public:")

		header = append(header, fmt.Sprintf(constructor, name))
		header = append(header, fmt.Sprintf("%s override;", marshal))

		if len(fields) > 0 {
			header = append(header, " private:")
		}

		for _, info := range fields {
			typeName := info.TypeName

			if primitives[typeName] {
				if info.Type == StarType {
					typeName = typeName + "*"
				}
				header = append(header, fmt.Sprintf("  %s %s_;", typeName, info.Name))
				continue
			}

			switch info.Type {
			case IdentType, StarType:
				header = append(header, fmt.Sprintf("  std::unique_ptr<%s> %s_;", info.TypeName, info.Name))
			case ArrayType:
				if info.TypeName == "byte" {
					header = append(header, fmt.Sprintf("  string %s_;", info.Name))
					continue
				}

				typeName := info.TypeName
				if interfaces[typeName] {
					typeName = "std::unique_ptr<" + typeName + ">"
				}
				header = append(header, fmt.Sprintf("  std::vector<%s> %ss_;", typeName, info.Name))
			}
		}

		header = append(header, "};", "")
	}

	return header
}

func main() {
	output := os.Stdout

	tset := token.NewFileSet()
	tf, err := parser.ParseFile(tset, "../../tao/auth/binary.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ast.Print(tset, tf)

	constantVisitor := &ConstantVisitor{
		Constants: make([]Constant, 0),
	}

	ast.Walk(constantVisitor, tf)

	//fmt.Fprintf(output, "The constants are as follows: %v\n", constantVisitor.Constants)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../tao/auth/ast.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	tv := &TypeVisitor{
		ConcreteTypes: make(map[string][]Field),
		InterfaceTypes: make(map[string]bool),
	}
	ast.Walk(tv, f)

	//fmt.Printf("The full set of types is %+v\n", tv)

	formWalker := &FuncReceiverWalker{
		types: make(map[string]bool),
		name: "isForm",
	}
	ast.Walk(formWalker, f)

	//fmt.Printf("The following types are Forms: %v\n", formWalker.types)

	termWalker := &FuncReceiverWalker{
		types: make(map[string]bool),
		name: "isTerm",
	}
	ast.Walk(termWalker, f)

	//fmt.Printf("The following types are Terms: %v\n", termWalker.types)


	header := writeHeader(constantVisitor.Constants, tv.ConcreteTypes, tv.InterfaceTypes, formWalker.types, termWalker.types)
	for _, line := range header {
		fmt.Fprintln(output, line)
	}

	// Generate the class declarations with inheritance that puts the Marshal method on everything. Everything should have a constructor based on the protobuf streams.
	//
	// Then generate the marshalling/demarshalling code.

	// The following line prints the full (immense) AST.
	//ast.Print(fset, f)
}
