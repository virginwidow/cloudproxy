// Code generated by protoc-gen-go.
// source: ca.proto
// DO NOT EDIT!

package tao

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CAType int32

const (
	CAType_ERROR          CAType = 0
	CAType_ATTESTATION    CAType = 1
	CAType_DATALOG_POLICY CAType = 2
	CAType_ACL_POLICY     CAType = 3
	CAType_UNDEFINED      CAType = 4
)

var CAType_name = map[int32]string{
	0: "ERROR",
	1: "ATTESTATION",
	2: "DATALOG_POLICY",
	3: "ACL_POLICY",
	4: "UNDEFINED",
}
var CAType_value = map[string]int32{
	"ERROR":          0,
	"ATTESTATION":    1,
	"DATALOG_POLICY": 2,
	"ACL_POLICY":     3,
	"UNDEFINED":      4,
}

func (x CAType) Enum() *CAType {
	p := new(CAType)
	*p = x
	return p
}
func (x CAType) String() string {
	return proto.EnumName(CAType_name, int32(x))
}
func (x *CAType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CAType_value, data, "CAType")
	if err != nil {
		return err
	}
	*x = CAType(value)
	return nil
}

type CARequest struct {
	Type             *CAType      `protobuf:"varint,1,req,name=type,enum=tao.CAType" json:"type,omitempty"`
	Attestation      *Attestation `protobuf:"bytes,2,opt,name=attestation" json:"attestation,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CARequest) Reset()         { *m = CARequest{} }
func (m *CARequest) String() string { return proto.CompactTextString(m) }
func (*CARequest) ProtoMessage()    {}

func (m *CARequest) GetType() CAType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CAType_ERROR
}

func (m *CARequest) GetAttestation() *Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

type CAResponse struct {
	Type               *CAType             `protobuf:"varint,1,req,name=type,enum=tao.CAType" json:"type,omitempty"`
	Attestation        *Attestation        `protobuf:"bytes,2,opt,name=attestation" json:"attestation,omitempty"`
	SignedDatalogRules *SignedDatalogRules `protobuf:"bytes,3,opt,name=signed_datalog_rules" json:"signed_datalog_rules,omitempty"`
	SignedAclSet       *SignedACLSet       `protobuf:"bytes,4,opt,name=signed_acl_set" json:"signed_acl_set,omitempty"`
	XXX_unrecognized   []byte              `json:"-"`
}

func (m *CAResponse) Reset()         { *m = CAResponse{} }
func (m *CAResponse) String() string { return proto.CompactTextString(m) }
func (*CAResponse) ProtoMessage()    {}

func (m *CAResponse) GetType() CAType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CAType_ERROR
}

func (m *CAResponse) GetAttestation() *Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

func (m *CAResponse) GetSignedDatalogRules() *SignedDatalogRules {
	if m != nil {
		return m.SignedDatalogRules
	}
	return nil
}

func (m *CAResponse) GetSignedAclSet() *SignedACLSet {
	if m != nil {
		return m.SignedAclSet
	}
	return nil
}

func init() {
	proto.RegisterEnum("tao.CAType", CAType_name, CAType_value)
}