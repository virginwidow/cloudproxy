// Code generated by protoc-gen-go.
// source: hosted_programs.proto
// DO NOT EDIT!

package tao

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HostedProgram struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	HashAlg          *string `protobuf:"bytes,2,req,name=hash_alg" json:"hash_alg,omitempty"`
	Hash             []byte  `protobuf:"bytes,3,req,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HostedProgram) Reset()         { *m = HostedProgram{} }
func (m *HostedProgram) String() string { return proto.CompactTextString(m) }
func (*HostedProgram) ProtoMessage()    {}

func (m *HostedProgram) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *HostedProgram) GetHashAlg() string {
	if m != nil && m.HashAlg != nil {
		return *m.HashAlg
	}
	return ""
}

func (m *HostedProgram) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Whitelist struct {
	Programs         []*HostedProgram `protobuf:"bytes,1,rep,name=programs" json:"programs,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Whitelist) Reset()         { *m = Whitelist{} }
func (m *Whitelist) String() string { return proto.CompactTextString(m) }
func (*Whitelist) ProtoMessage()    {}

func (m *Whitelist) GetPrograms() []*HostedProgram {
	if m != nil {
		return m.Programs
	}
	return nil
}

type SignedWhitelist struct {
	SerializedWhitelist *string `protobuf:"bytes,1,req,name=serialized_whitelist" json:"serialized_whitelist,omitempty"`
	Signature           []byte  `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *SignedWhitelist) Reset()         { *m = SignedWhitelist{} }
func (m *SignedWhitelist) String() string { return proto.CompactTextString(m) }
func (*SignedWhitelist) ProtoMessage()    {}

func (m *SignedWhitelist) GetSerializedWhitelist() string {
	if m != nil && m.SerializedWhitelist != nil {
		return *m.SerializedWhitelist
	}
	return ""
}

func (m *SignedWhitelist) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type HostedProgramState struct {
	Hp               *HostedProgram `protobuf:"bytes,1,req,name=hp" json:"hp,omitempty"`
	Secret           []byte         `protobuf:"bytes,2,req,name=secret" json:"secret,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *HostedProgramState) Reset()         { *m = HostedProgramState{} }
func (m *HostedProgramState) String() string { return proto.CompactTextString(m) }
func (*HostedProgramState) ProtoMessage()    {}

func (m *HostedProgramState) GetHp() *HostedProgram {
	if m != nil {
		return m.Hp
	}
	return nil
}

func (m *HostedProgramState) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

type Secrets struct {
	ProgramSecrets   []*HostedProgramState `protobuf:"bytes,1,rep,name=program_secrets" json:"program_secrets,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Secrets) Reset()         { *m = Secrets{} }
func (m *Secrets) String() string { return proto.CompactTextString(m) }
func (*Secrets) ProtoMessage()    {}

func (m *Secrets) GetProgramSecrets() []*HostedProgramState {
	if m != nil {
		return m.ProgramSecrets
	}
	return nil
}

func init() {
}