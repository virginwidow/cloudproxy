// Copyright (c) 2016, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protected_objects

import (
	"fmt"
	"container/list"
	"testing"
	"time"

	"github.com/jlmucb/cloudproxy/go/protected_objects"
	// "github.com/jlmucb/cloudproxy/go/tpm2"
	// "github.com/golang/protobuf/proto"
)

func TestBasicObject(t *testing.T) {
	// make it
	obj_type := "file"
	status := "active"
	notBefore := time.Now()
	validFor := 365*24*time.Hour
	notAfter := notBefore.Add(validFor)

	obj, err := protected_objects.CreateObject("/jlm/file/file1", 1,
		&obj_type, &status, &notBefore, &notAfter, nil)
	if err != nil {
		t.Fatal("Can't create object")
	}
	fmt.Printf("Obj: %s\n", *obj.NotBefore)
	obj_type = "key"
	obj_2, _:= protected_objects.CreateObject("/jlm/key/key1", 1,
		&obj_type, &status, &notBefore, &notAfter, nil)
	obj_3, _:= protected_objects.CreateObject("/jlm/key/key2", 1,
		&obj_type, &status, &notBefore, &notAfter, nil)

	// add it to object list
	obj_list := list.New()
	err = protected_objects.AddObject(obj_list, *obj)
	if err != nil {
		t.Fatal("Can't add object")
	}
	_ = protected_objects.AddObject(obj_list, *obj_2)
	_ = protected_objects.AddObject(obj_list, *obj_3)

	o3 := protected_objects.FindObject(obj_list, *obj.ObjId.ObjName, *obj.ObjId.ObjEpoch)
	fmt.Printf("Found object\n")
	protected_objects.PrintObject(o3)

	protectorKeys := []byte{
		0,1,2,3,4,5,6,7,8,9,0xa,0xb,0xc,0xd,0xe,0xf,
		0,1,2,3,4,5,6,7,8,9,0xa,0xb,0xc,0xd,0xe,0xf,
	}
	p_obj, err := protected_objects.MakeProtectedObject(*obj, "/jlm/key/key1", 1, protectorKeys)
	if err != nil {
		t.Fatal("Can't make protected object")
	}
	if p_obj == nil {
		t.Fatal("Bad protected object")
	}
	protected_objects.PrintProtectedObject(p_obj)

	obj2, err := protected_objects.RecoverProtectedObject(p_obj, protectorKeys)
	if err != nil {
		t.Fatal("Can't recover protected object")
	}
	if *obj.ObjId.ObjName != *obj2.ObjId.ObjName {
		t.Fatal("objects don't match")
	}
	//protected_obj_list := list.New()

	err = protected_objects.SaveObjects(obj_list, "tmptest/s1")
	if err != nil {
		t.Fatal("Can't save objects")
	}
	r := protected_objects.LoadObjects("tmptest/s1")
	if r == nil {
		t.Fatal("Can't Load objects")
	}
	e := r.Front()
	o2 := e.Value.(protected_objects.ObjectMessage)
	fmt.Printf("Recovered object\n")
	protected_objects.PrintObject(&o2)

	n1 := protected_objects.MakeNode("/jlm/key/key1", 1, "/jlm/file/file1", 1);
	if n1 == nil {
		t.Fatal("Can't make node")
	}

	n2 := protected_objects.MakeNode("/jlm/key/key2", 1, "/jlm/key/key1", 1);
	if n2 == nil {
		t.Fatal("Can't make node")
	}

	n_list := list.New()
	err = protected_objects.AddNode(n_list, *n1)
	if n1 == nil {
		t.Fatal("Can't add node")
	}
	err = protected_objects.AddNode(n_list, *n2)
	if n2 == nil {
		t.Fatal("Can't add node")
	}

	pr_list1 := protected_objects.FindProtectorNodes(n_list, "/jlm/key/key1", 1)
	if pr_list1 == nil {
		t.Fatal("FindProtectorNodes fails")
	}
	fmt.Printf("Protecting:\n")
	for e := pr_list1.Front(); e != nil; e = e.Next() {
		o := e.Value.(protected_objects.NodeMessage)
		protected_objects.PrintNode(&o)
	}
	fmt.Printf("\n")
	fmt.Printf("Protected:\n")
	pr_list2 := protected_objects.FindProtectedNodes(n_list, "/jlm/key/key1", 1)
	if pr_list2 == nil {
		t.Fatal("FindProtectedNodes fails")
	}
	for e := pr_list2.Front(); e != nil; e = e.Next() {
		o := e.Value.(protected_objects.NodeMessage)
		protected_objects.PrintNode(&o)
	}
	fmt.Printf("\n")

	seen_list := list.New()
	chain, err := protected_objects.ConstructProtectorChain(obj_list,
		"/jlm/file/file1", 1, nil, nil, seen_list, n_list)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		t.Fatal("Can't ConstructProtectorChain ")
	}
	fmt.Printf("Chain:\n")
	for e := chain.Front(); e != nil; e = e.Next() {
		o := e.Value.(protected_objects.ObjectMessage)
		protected_objects.PrintObject(&o)
	}
}

func TestConstructChain(t *testing.T) {
	// base object
	// ancestors
	// construct chain with latest epoch
	// validate chain
}
