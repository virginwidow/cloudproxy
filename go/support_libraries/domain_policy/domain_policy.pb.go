// Code generated by protoc-gen-go.
// source: domain_policy.proto
// DO NOT EDIT!

/*
Package domain_policy is a generated protocol buffer package.

It is generated from these files:
	domain_policy.proto

It has these top-level messages:
	DomainCertRequest
	DomainCertResponse
*/
package domain_policy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This is used to request a signed cert from the domain service.
// attestation is the marshaled attestation.
// program_key is the der encoded program public key.
// cert chain is any supporting certificates.
type DomainCertRequest struct {
	Attestation      []byte   `protobuf:"bytes,1,opt,name=attestation" json:"attestation,omitempty"`
	KeyType          *string  `protobuf:"bytes,2,opt,name=key_type,json=keyType" json:"key_type,omitempty"`
	SubjectPublicKey []byte   `protobuf:"bytes,3,opt,name=subject_public_key,json=subjectPublicKey" json:"subject_public_key,omitempty"`
	CertChain        [][]byte `protobuf:"bytes,4,rep,name=cert_chain,json=certChain" json:"cert_chain,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DomainCertRequest) Reset()                    { *m = DomainCertRequest{} }
func (m *DomainCertRequest) String() string            { return proto.CompactTextString(m) }
func (*DomainCertRequest) ProtoMessage()               {}
func (*DomainCertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DomainCertRequest) GetAttestation() []byte {
	if m != nil {
		return m.Attestation
	}
	return nil
}

func (m *DomainCertRequest) GetKeyType() string {
	if m != nil && m.KeyType != nil {
		return *m.KeyType
	}
	return ""
}

func (m *DomainCertRequest) GetSubjectPublicKey() []byte {
	if m != nil {
		return m.SubjectPublicKey
	}
	return nil
}

func (m *DomainCertRequest) GetCertChain() [][]byte {
	if m != nil {
		return m.CertChain
	}
	return nil
}

// This is the response from the domain service.
// signed_cert is the signed program cert.
// cert_chain is the supporting certificate chain.
type DomainCertResponse struct {
	Error            *int32   `protobuf:"varint,1,req,name=error" json:"error,omitempty"`
	SignedCert       []byte   `protobuf:"bytes,2,opt,name=signed_cert,json=signedCert" json:"signed_cert,omitempty"`
	CertChain        [][]byte `protobuf:"bytes,3,rep,name=cert_chain,json=certChain" json:"cert_chain,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DomainCertResponse) Reset()                    { *m = DomainCertResponse{} }
func (m *DomainCertResponse) String() string            { return proto.CompactTextString(m) }
func (*DomainCertResponse) ProtoMessage()               {}
func (*DomainCertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DomainCertResponse) GetError() int32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return 0
}

func (m *DomainCertResponse) GetSignedCert() []byte {
	if m != nil {
		return m.SignedCert
	}
	return nil
}

func (m *DomainCertResponse) GetCertChain() [][]byte {
	if m != nil {
		return m.CertChain
	}
	return nil
}

func init() {
	proto.RegisterType((*DomainCertRequest)(nil), "domain_policy.DomainCertRequest")
	proto.RegisterType((*DomainCertResponse)(nil), "domain_policy.DomainCertResponse")
}

func init() { proto.RegisterFile("domain_policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8e, 0xb1, 0x4e, 0x03, 0x31,
	0x10, 0x44, 0x75, 0x39, 0x22, 0xc8, 0x26, 0x48, 0xb0, 0x50, 0x1c, 0x05, 0xe2, 0x94, 0x2a, 0x05,
	0xe2, 0x27, 0x42, 0x47, 0x83, 0x2c, 0x7a, 0xeb, 0xe2, 0xac, 0xc0, 0x01, 0x6c, 0x63, 0xef, 0x15,
	0xfe, 0x18, 0xfe, 0x95, 0xb5, 0x43, 0x91, 0xa4, 0xf3, 0xbc, 0xf1, 0xec, 0x0c, 0xdc, 0x6c, 0xfd,
	0xf7, 0x60, 0x9d, 0x0e, 0xfe, 0xcb, 0x9a, 0xfc, 0x14, 0xa2, 0x67, 0x8f, 0x97, 0x47, 0x70, 0xf9,
	0xdb, 0xc0, 0xf5, 0x73, 0x25, 0x6b, 0x8a, 0xac, 0xe8, 0x67, 0xa4, 0xc4, 0xd8, 0xc3, 0x7c, 0x60,
	0x96, 0xd7, 0xc0, 0xd6, 0xbb, 0xae, 0xe9, 0x9b, 0xd5, 0x42, 0x1d, 0x22, 0xbc, 0x83, 0x8b, 0x4f,
	0xca, 0x9a, 0x73, 0xa0, 0x6e, 0x22, 0xf6, 0x4c, 0x9d, 0x8b, 0x7e, 0x13, 0x89, 0x8f, 0x80, 0x69,
	0xdc, 0xec, 0xc8, 0xb0, 0x0e, 0xe3, 0x46, 0x5a, 0xb4, 0x38, 0x5d, 0x5b, 0x6f, 0x5c, 0xfd, 0x3b,
	0xaf, 0xd5, 0x78, 0xa1, 0x8c, 0xf7, 0x00, 0x46, 0x9a, 0xb5, 0xf9, 0x90, 0x0d, 0xdd, 0x59, 0xdf,
	0xca, 0xaf, 0x59, 0x21, 0xeb, 0x02, 0x96, 0x3b, 0xc0, 0xc3, 0x79, 0x29, 0x78, 0x97, 0x08, 0x6f,
	0x61, 0x4a, 0x31, 0xfa, 0x28, 0xcb, 0x26, 0xab, 0xa9, 0xda, 0x0b, 0x7c, 0x80, 0x79, 0xb2, 0xef,
	0x8e, 0xb6, 0xba, 0xe4, 0xeb, 0xac, 0x85, 0x82, 0x3d, 0x2a, 0xf1, 0x93, 0xae, 0xf6, 0xa4, 0xeb,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x94, 0xac, 0x12, 0xca, 0x30, 0x01, 0x00, 0x00,
}
