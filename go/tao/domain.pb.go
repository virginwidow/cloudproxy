// Code generated by protoc-gen-go.
// source: domain.proto
// DO NOT EDIT!

/*
Package tao is a generated protocol buffer package.

It is generated from these files:
	domain.proto

It has these top-level messages:
	DomainDetails
	X509Details
	ACLGuardDetails
	DatalogGuardDetails
	TPMDetails
	TPM2Details
	DomainConfig
	DomainTemplate
*/
package tao

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

// TODO(jlm): Comments for arguments.
// Policy key should be compatible with library cipher suite.
type DomainDetails struct {
	// name of domain
	Name           *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PolicyKeysPath *string `protobuf:"bytes,2,opt,name=policy_keys_path,json=policyKeysPath" json:"policy_keys_path,omitempty"`
	GuardType      *string `protobuf:"bytes,3,opt,name=guard_type,json=guardType" json:"guard_type,omitempty"`
	// ??
	GuardNetwork     *string `protobuf:"bytes,4,opt,name=guard_network,json=guardNetwork" json:"guard_network,omitempty"`
	GuardAddress     *string `protobuf:"bytes,5,opt,name=guard_address,json=guardAddress" json:"guard_address,omitempty"`
	GuardTtl         *int64  `protobuf:"varint,6,opt,name=guard_ttl,json=guardTtl" json:"guard_ttl,omitempty"`
	CipherSuite      *string `protobuf:"bytes,7,opt,name=cipher_suite,json=cipherSuite" json:"cipher_suite,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DomainDetails) Reset()                    { *m = DomainDetails{} }
func (m *DomainDetails) String() string            { return proto.CompactTextString(m) }
func (*DomainDetails) ProtoMessage()               {}
func (*DomainDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DomainDetails) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DomainDetails) GetPolicyKeysPath() string {
	if m != nil && m.PolicyKeysPath != nil {
		return *m.PolicyKeysPath
	}
	return ""
}

func (m *DomainDetails) GetGuardType() string {
	if m != nil && m.GuardType != nil {
		return *m.GuardType
	}
	return ""
}

func (m *DomainDetails) GetGuardNetwork() string {
	if m != nil && m.GuardNetwork != nil {
		return *m.GuardNetwork
	}
	return ""
}

func (m *DomainDetails) GetGuardAddress() string {
	if m != nil && m.GuardAddress != nil {
		return *m.GuardAddress
	}
	return ""
}

func (m *DomainDetails) GetGuardTtl() int64 {
	if m != nil && m.GuardTtl != nil {
		return *m.GuardTtl
	}
	return 0
}

func (m *DomainDetails) GetCipherSuite() string {
	if m != nil && m.CipherSuite != nil {
		return *m.CipherSuite
	}
	return ""
}

type X509Details struct {
	CommonName         *string `protobuf:"bytes,1,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	Country            *string `protobuf:"bytes,2,opt,name=country" json:"country,omitempty"`
	State              *string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Organization       *string `protobuf:"bytes,4,opt,name=organization" json:"organization,omitempty"`
	OrganizationalUnit *string `protobuf:"bytes,5,opt,name=organizational_unit,json=organizationalUnit" json:"organizational_unit,omitempty"`
	SerialNumber       *int32  `protobuf:"varint,6,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *X509Details) Reset()                    { *m = X509Details{} }
func (m *X509Details) String() string            { return proto.CompactTextString(m) }
func (*X509Details) ProtoMessage()               {}
func (*X509Details) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *X509Details) GetCommonName() string {
	if m != nil && m.CommonName != nil {
		return *m.CommonName
	}
	return ""
}

func (m *X509Details) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *X509Details) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *X509Details) GetOrganization() string {
	if m != nil && m.Organization != nil {
		return *m.Organization
	}
	return ""
}

func (m *X509Details) GetOrganizationalUnit() string {
	if m != nil && m.OrganizationalUnit != nil {
		return *m.OrganizationalUnit
	}
	return ""
}

func (m *X509Details) GetSerialNumber() int32 {
	if m != nil && m.SerialNumber != nil {
		return *m.SerialNumber
	}
	return 0
}

type ACLGuardDetails struct {
	SignedAclsPath   *string `protobuf:"bytes,1,opt,name=signed_acls_path,json=signedAclsPath" json:"signed_acls_path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ACLGuardDetails) Reset()                    { *m = ACLGuardDetails{} }
func (m *ACLGuardDetails) String() string            { return proto.CompactTextString(m) }
func (*ACLGuardDetails) ProtoMessage()               {}
func (*ACLGuardDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ACLGuardDetails) GetSignedAclsPath() string {
	if m != nil && m.SignedAclsPath != nil {
		return *m.SignedAclsPath
	}
	return ""
}

type DatalogGuardDetails struct {
	SignedRulesPath  *string `protobuf:"bytes,2,opt,name=signed_rules_path,json=signedRulesPath" json:"signed_rules_path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DatalogGuardDetails) Reset()                    { *m = DatalogGuardDetails{} }
func (m *DatalogGuardDetails) String() string            { return proto.CompactTextString(m) }
func (*DatalogGuardDetails) ProtoMessage()               {}
func (*DatalogGuardDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DatalogGuardDetails) GetSignedRulesPath() string {
	if m != nil && m.SignedRulesPath != nil {
		return *m.SignedRulesPath
	}
	return ""
}

type TPMDetails struct {
	TpmPath *string `protobuf:"bytes,1,opt,name=tpm_path,json=tpmPath" json:"tpm_path,omitempty"`
	AikPath *string `protobuf:"bytes,2,opt,name=aik_path,json=aikPath" json:"aik_path,omitempty"`
	// A string representing the IDs of PCRs, like "17,18".
	Pcrs *string `protobuf:"bytes,3,opt,name=pcrs" json:"pcrs,omitempty"`
	// Path for AIK cert.
	AikCertPath      *string `protobuf:"bytes,4,opt,name=aik_cert_path,json=aikCertPath" json:"aik_cert_path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TPMDetails) Reset()                    { *m = TPMDetails{} }
func (m *TPMDetails) String() string            { return proto.CompactTextString(m) }
func (*TPMDetails) ProtoMessage()               {}
func (*TPMDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TPMDetails) GetTpmPath() string {
	if m != nil && m.TpmPath != nil {
		return *m.TpmPath
	}
	return ""
}

func (m *TPMDetails) GetAikPath() string {
	if m != nil && m.AikPath != nil {
		return *m.AikPath
	}
	return ""
}

func (m *TPMDetails) GetPcrs() string {
	if m != nil && m.Pcrs != nil {
		return *m.Pcrs
	}
	return ""
}

func (m *TPMDetails) GetAikCertPath() string {
	if m != nil && m.AikCertPath != nil {
		return *m.AikCertPath
	}
	return ""
}

type TPM2Details struct {
	Tpm2InfoDir      *string `protobuf:"bytes,1,opt,name=tpm2_info_dir,json=tpm2InfoDir" json:"tpm2_info_dir,omitempty"`
	Tpm2Device       *string `protobuf:"bytes,2,opt,name=tpm2_device,json=tpm2Device" json:"tpm2_device,omitempty"`
	Tpm2Pcrs         *string `protobuf:"bytes,3,opt,name=tpm2_pcrs,json=tpm2Pcrs" json:"tpm2_pcrs,omitempty"`
	Tpm2EkCert       *string `protobuf:"bytes,4,opt,name=tpm2_ek_cert,json=tpm2EkCert" json:"tpm2_ek_cert,omitempty"`
	Tpm2QuoteCert    *string `protobuf:"bytes,5,opt,name=tpm2_quote_cert,json=tpm2QuoteCert" json:"tpm2_quote_cert,omitempty"`
	Tpm2SealCert     *string `protobuf:"bytes,6,opt,name=tpm2_seal_cert,json=tpm2SealCert" json:"tpm2_seal_cert,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TPM2Details) Reset()                    { *m = TPM2Details{} }
func (m *TPM2Details) String() string            { return proto.CompactTextString(m) }
func (*TPM2Details) ProtoMessage()               {}
func (*TPM2Details) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TPM2Details) GetTpm2InfoDir() string {
	if m != nil && m.Tpm2InfoDir != nil {
		return *m.Tpm2InfoDir
	}
	return ""
}

func (m *TPM2Details) GetTpm2Device() string {
	if m != nil && m.Tpm2Device != nil {
		return *m.Tpm2Device
	}
	return ""
}

func (m *TPM2Details) GetTpm2Pcrs() string {
	if m != nil && m.Tpm2Pcrs != nil {
		return *m.Tpm2Pcrs
	}
	return ""
}

func (m *TPM2Details) GetTpm2EkCert() string {
	if m != nil && m.Tpm2EkCert != nil {
		return *m.Tpm2EkCert
	}
	return ""
}

func (m *TPM2Details) GetTpm2QuoteCert() string {
	if m != nil && m.Tpm2QuoteCert != nil {
		return *m.Tpm2QuoteCert
	}
	return ""
}

func (m *TPM2Details) GetTpm2SealCert() string {
	if m != nil && m.Tpm2SealCert != nil {
		return *m.Tpm2SealCert
	}
	return ""
}

type DomainConfig struct {
	DomainInfo       *DomainDetails       `protobuf:"bytes,1,opt,name=domain_info,json=domainInfo" json:"domain_info,omitempty"`
	X509Info         *X509Details         `protobuf:"bytes,2,opt,name=x509_info,json=x509Info" json:"x509_info,omitempty"`
	AclGuardInfo     *ACLGuardDetails     `protobuf:"bytes,3,opt,name=acl_guard_info,json=aclGuardInfo" json:"acl_guard_info,omitempty"`
	DatalogGuardInfo *DatalogGuardDetails `protobuf:"bytes,4,opt,name=datalog_guard_info,json=datalogGuardInfo" json:"datalog_guard_info,omitempty"`
	TpmInfo          *TPMDetails          `protobuf:"bytes,5,opt,name=tpm_info,json=tpmInfo" json:"tpm_info,omitempty"`
	Tpm2Info         *TPM2Details         `protobuf:"bytes,6,opt,name=tpm2_info,json=tpm2Info" json:"tpm2_info,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *DomainConfig) Reset()                    { *m = DomainConfig{} }
func (m *DomainConfig) String() string            { return proto.CompactTextString(m) }
func (*DomainConfig) ProtoMessage()               {}
func (*DomainConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DomainConfig) GetDomainInfo() *DomainDetails {
	if m != nil {
		return m.DomainInfo
	}
	return nil
}

func (m *DomainConfig) GetX509Info() *X509Details {
	if m != nil {
		return m.X509Info
	}
	return nil
}

func (m *DomainConfig) GetAclGuardInfo() *ACLGuardDetails {
	if m != nil {
		return m.AclGuardInfo
	}
	return nil
}

func (m *DomainConfig) GetDatalogGuardInfo() *DatalogGuardDetails {
	if m != nil {
		return m.DatalogGuardInfo
	}
	return nil
}

func (m *DomainConfig) GetTpmInfo() *TPMDetails {
	if m != nil {
		return m.TpmInfo
	}
	return nil
}

func (m *DomainConfig) GetTpm2Info() *TPM2Details {
	if m != nil {
		return m.Tpm2Info
	}
	return nil
}

type DomainTemplate struct {
	Config       *DomainConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	DatalogRules []string      `protobuf:"bytes,2,rep,name=datalog_rules,json=datalogRules" json:"datalog_rules,omitempty"`
	AclRules     []string      `protobuf:"bytes,3,rep,name=acl_rules,json=aclRules" json:"acl_rules,omitempty"`
	// The name of the host (used for policy statements)
	HostName          *string `protobuf:"bytes,4,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	HostPredicateName *string `protobuf:"bytes,5,opt,name=host_predicate_name,json=hostPredicateName" json:"host_predicate_name,omitempty"`
	// Program names (as paths to binaries)
	ProgramPaths         []string `protobuf:"bytes,6,rep,name=program_paths,json=programPaths" json:"program_paths,omitempty"`
	ProgramPredicateName *string  `protobuf:"bytes,7,opt,name=program_predicate_name,json=programPredicateName" json:"program_predicate_name,omitempty"`
	// Container names (as paths to images)
	ContainerPaths         []string `protobuf:"bytes,8,rep,name=container_paths,json=containerPaths" json:"container_paths,omitempty"`
	ContainerPredicateName *string  `protobuf:"bytes,9,opt,name=container_predicate_name,json=containerPredicateName" json:"container_predicate_name,omitempty"`
	// VM names (as paths to images)
	VmPaths         []string `protobuf:"bytes,10,rep,name=vm_paths,json=vmPaths" json:"vm_paths,omitempty"`
	VmPredicateName *string  `protobuf:"bytes,11,opt,name=vm_predicate_name,json=vmPredicateName" json:"vm_predicate_name,omitempty"`
	// LinuxHost names (as paths to images)
	LinuxHostPaths         []string `protobuf:"bytes,12,rep,name=linux_host_paths,json=linuxHostPaths" json:"linux_host_paths,omitempty"`
	LinuxHostPredicateName *string  `protobuf:"bytes,13,opt,name=linux_host_predicate_name,json=linuxHostPredicateName" json:"linux_host_predicate_name,omitempty"`
	// The name of the predicate to use for trusted guards.
	GuardPredicateName *string `protobuf:"bytes,14,opt,name=guard_predicate_name,json=guardPredicateName" json:"guard_predicate_name,omitempty"`
	// The name of the predicate to use for trusted TPMs.
	TpmPredicateName *string `protobuf:"bytes,15,opt,name=tpm_predicate_name,json=tpmPredicateName" json:"tpm_predicate_name,omitempty"`
	// The name of the predicate to use for trusted OSs.
	OsPredicateName *string `protobuf:"bytes,16,opt,name=os_predicate_name,json=osPredicateName" json:"os_predicate_name,omitempty"`
	// The name of the predicate to use for trusted TPM2s.
	Tpm2PredicateName *string `protobuf:"bytes,17,opt,name=tpm2_predicate_name,json=tpm2PredicateName" json:"tpm2_predicate_name,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *DomainTemplate) Reset()                    { *m = DomainTemplate{} }
func (m *DomainTemplate) String() string            { return proto.CompactTextString(m) }
func (*DomainTemplate) ProtoMessage()               {}
func (*DomainTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DomainTemplate) GetConfig() *DomainConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *DomainTemplate) GetDatalogRules() []string {
	if m != nil {
		return m.DatalogRules
	}
	return nil
}

func (m *DomainTemplate) GetAclRules() []string {
	if m != nil {
		return m.AclRules
	}
	return nil
}

func (m *DomainTemplate) GetHostName() string {
	if m != nil && m.HostName != nil {
		return *m.HostName
	}
	return ""
}

func (m *DomainTemplate) GetHostPredicateName() string {
	if m != nil && m.HostPredicateName != nil {
		return *m.HostPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetProgramPaths() []string {
	if m != nil {
		return m.ProgramPaths
	}
	return nil
}

func (m *DomainTemplate) GetProgramPredicateName() string {
	if m != nil && m.ProgramPredicateName != nil {
		return *m.ProgramPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetContainerPaths() []string {
	if m != nil {
		return m.ContainerPaths
	}
	return nil
}

func (m *DomainTemplate) GetContainerPredicateName() string {
	if m != nil && m.ContainerPredicateName != nil {
		return *m.ContainerPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetVmPaths() []string {
	if m != nil {
		return m.VmPaths
	}
	return nil
}

func (m *DomainTemplate) GetVmPredicateName() string {
	if m != nil && m.VmPredicateName != nil {
		return *m.VmPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetLinuxHostPaths() []string {
	if m != nil {
		return m.LinuxHostPaths
	}
	return nil
}

func (m *DomainTemplate) GetLinuxHostPredicateName() string {
	if m != nil && m.LinuxHostPredicateName != nil {
		return *m.LinuxHostPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetGuardPredicateName() string {
	if m != nil && m.GuardPredicateName != nil {
		return *m.GuardPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetTpmPredicateName() string {
	if m != nil && m.TpmPredicateName != nil {
		return *m.TpmPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetOsPredicateName() string {
	if m != nil && m.OsPredicateName != nil {
		return *m.OsPredicateName
	}
	return ""
}

func (m *DomainTemplate) GetTpm2PredicateName() string {
	if m != nil && m.Tpm2PredicateName != nil {
		return *m.Tpm2PredicateName
	}
	return ""
}

func init() {
	proto.RegisterType((*DomainDetails)(nil), "tao.DomainDetails")
	proto.RegisterType((*X509Details)(nil), "tao.X509Details")
	proto.RegisterType((*ACLGuardDetails)(nil), "tao.ACLGuardDetails")
	proto.RegisterType((*DatalogGuardDetails)(nil), "tao.DatalogGuardDetails")
	proto.RegisterType((*TPMDetails)(nil), "tao.TPMDetails")
	proto.RegisterType((*TPM2Details)(nil), "tao.TPM2Details")
	proto.RegisterType((*DomainConfig)(nil), "tao.DomainConfig")
	proto.RegisterType((*DomainTemplate)(nil), "tao.DomainTemplate")
}

/*
func init() { proto.RegisterFile("domain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 919 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x54, 0xdf, 0x6f, 0xe3, 0x44,
	0x10, 0x56, 0x9b, 0xa6, 0x49, 0xc6, 0x6e, 0x92, 0x6e, 0xab, 0x93, 0x4f, 0x08, 0x71, 0x04, 0x04,
	0xc7, 0x09, 0xca, 0xa9, 0x80, 0xc4, 0xc1, 0x53, 0xd5, 0xf0, 0x4b, 0x40, 0x15, 0x7c, 0x45, 0xe2,
	0xcd, 0x32, 0xf6, 0x5e, 0xba, 0xaa, 0xe3, 0x35, 0xeb, 0x75, 0xb8, 0xdc, 0x7f, 0xc9, 0x1b, 0x6f,
	0xfc, 0x0b, 0x3c, 0xf1, 0xcc, 0xec, 0x8c, 0x9d, 0xd8, 0xe9, 0x9b, 0xfd, 0xcd, 0x37, 0x3f, 0xf7,
	0x9b, 0x01, 0x3f, 0xd5, 0xab, 0x58, 0xe5, 0x17, 0x85, 0xd1, 0x56, 0x8b, 0x9e, 0x8d, 0xf5, 0xec,
	0xbf, 0x03, 0x38, 0x99, 0x13, 0x3a, 0x97, 0x36, 0x56, 0x59, 0x29, 0x04, 0x1c, 0xe5, 0xf1, 0x4a,
	0x06, 0x07, 0x4f, 0x0e, 0x9e, 0x8e, 0x42, 0xfa, 0x16, 0x4f, 0x61, 0x5a, 0xe8, 0x4c, 0x25, 0x9b,
	0xe8, 0x5e, 0x6e, 0xca, 0xa8, 0x88, 0xed, 0x5d, 0x70, 0x48, 0xf6, 0x31, 0xe3, 0x3f, 0x22, 0xbc,
	0x40, 0x54, 0xbc, 0x0d, 0xb0, 0xac, 0x62, 0x93, 0x46, 0x76, 0x53, 0xc8, 0xa0, 0x47, 0x9c, 0x11,
	0x21, 0xb7, 0x08, 0x88, 0xf7, 0xe0, 0x84, 0xcd, 0xb9, 0xb4, 0x7f, 0x6a, 0x73, 0x1f, 0x1c, 0x11,
	0xc3, 0x27, 0xf0, 0x86, 0xb1, 0x1d, 0x29, 0x4e, 0x53, 0x23, 0xcb, 0x32, 0xe8, 0xb7, 0x48, 0x57,
	0x8c, 0x89, 0xb7, 0x60, 0x54, 0x27, 0xb2, 0x59, 0x70, 0x8c, 0x84, 0x5e, 0x38, 0xe4, 0x3c, 0x36,
	0x13, 0xef, 0x82, 0x9f, 0xa8, 0xe2, 0x4e, 0x9a, 0xa8, 0xac, 0x94, 0x95, 0xc1, 0x80, 0x02, 0x78,
	0x8c, 0xbd, 0x74, 0xd0, 0xec, 0xef, 0x03, 0xf0, 0x7e, 0xfb, 0xe2, 0xf9, 0x8b, 0xa6, 0xed, 0x77,
	0xc0, 0x4b, 0xf4, 0x6a, 0xa5, 0xf3, 0xa8, 0xd5, 0x3d, 0x30, 0x74, 0xe3, 0x66, 0x10, 0xc0, 0x20,
	0xd1, 0x55, 0x6e, 0xcd, 0xa6, 0x6e, 0xbd, 0xf9, 0x15, 0xe7, 0xd0, 0x2f, 0x6d, 0x6c, 0x9b, 0x76,
	0xf9, 0x47, 0xcc, 0xc0, 0xd7, 0x66, 0x19, 0xe7, 0xea, 0x4d, 0x6c, 0x95, 0xce, 0x9b, 0x4e, 0xdb,
	0x98, 0xf8, 0x14, 0xce, 0xda, 0xff, 0x71, 0x16, 0x55, 0xb9, 0xb2, 0x75, 0xbf, 0xa2, 0x6b, 0xfa,
	0x15, 0x2d, 0x6e, 0x34, 0xa5, 0x34, 0x0a, 0x89, 0x79, 0xb5, 0xfa, 0x5d, 0x1a, 0xea, 0xbc, 0x1f,
	0xfa, 0x0c, 0xde, 0x10, 0x36, 0xfb, 0x1a, 0x26, 0x57, 0xd7, 0x3f, 0x7d, 0xe7, 0x86, 0xd1, 0x74,
	0x87, 0x0f, 0x58, 0xaa, 0x65, 0x2e, 0x71, 0xa6, 0x49, 0x56, 0x3f, 0x20, 0xb7, 0x38, 0x66, 0xfc,
	0x0a, 0x61, 0xf7, 0x80, 0xb3, 0x2b, 0x38, 0x9b, 0xc7, 0x36, 0xce, 0xf4, 0xb2, 0x13, 0xe0, 0x19,
	0x9c, 0xd6, 0x01, 0x4c, 0x95, 0xc9, 0x8e, 0x04, 0x26, 0x6c, 0x08, 0x1d, 0x4e, 0x21, 0xde, 0x00,
	0xdc, 0x2e, 0x7e, 0x6e, 0x3c, 0x1f, 0xc3, 0xd0, 0x16, 0xab, 0x76, 0xca, 0x01, 0xfe, 0x93, 0x58,
	0xd0, 0x14, 0xab, 0xfb, 0x76, 0xac, 0x01, 0xfe, 0x93, 0x09, 0x55, 0x58, 0x24, 0xa6, 0xac, 0x47,
	0x4a, 0xdf, 0x38, 0xd1, 0x13, 0x47, 0x4f, 0xa4, 0xb1, 0xec, 0xc3, 0x23, 0xf5, 0x10, 0xbc, 0x46,
	0x8c, 0x72, 0xff, 0x83, 0xcf, 0x8a, 0xc9, 0x2f, 0x9b, 0xec, 0xe8, 0x83, 0xd9, 0x2e, 0x23, 0x95,
	0xbf, 0xd2, 0x51, 0xaa, 0x4c, 0x5d, 0x82, 0xe7, 0xc0, 0x1f, 0x10, 0x9b, 0x2b, 0xe3, 0x9e, 0x9e,
	0x38, 0xa9, 0x5c, 0xab, 0x44, 0xd6, 0x95, 0x80, 0x83, 0xe6, 0x84, 0x38, 0xad, 0x11, 0xa1, 0x55,
	0x91, 0xeb, 0xe9, 0x72, 0xe1, 0xaa, 0x7a, 0x02, 0x3e, 0x19, 0x25, 0x57, 0x56, 0x17, 0x45, 0xee,
	0xdf, 0x50, 0x5d, 0xe2, 0x03, 0x98, 0x10, 0xe3, 0x8f, 0x4a, 0x5b, 0xc9, 0x24, 0x7e, 0x61, 0x2a,
	0xed, 0x17, 0x87, 0x12, 0xef, 0x7d, 0x18, 0x13, 0xaf, 0x94, 0xf8, 0xbe, 0x44, 0x3b, 0x66, 0xcd,
	0x38, 0xf4, 0x25, 0x82, 0x8e, 0x35, 0xfb, 0xeb, 0x10, 0x7c, 0xde, 0xd8, 0x6b, 0x9d, 0xbf, 0x52,
	0x4b, 0xf1, 0x19, 0x78, 0xbc, 0xd7, 0xd4, 0x24, 0x35, 0xe8, 0x5d, 0x8a, 0x0b, 0xdc, 0xee, 0x8b,
	0xce, 0x66, 0x87, 0xc0, 0x34, 0xd7, 0xb6, 0xf8, 0x04, 0x46, 0xaf, 0x51, 0xfd, 0xec, 0x72, 0x48,
	0x2e, 0x53, 0x72, 0x69, 0xed, 0x44, 0x38, 0x74, 0x14, 0xa2, 0x7f, 0x05, 0x63, 0x14, 0x4e, 0xc4,
	0x1b, 0x47, 0x3e, 0x3d, 0xf2, 0x39, 0x27, 0x9f, 0x3d, 0xb5, 0x85, 0x3e, 0x72, 0x09, 0x20, 0xdf,
	0x6f, 0x41, 0xa4, 0xac, 0xa8, 0xb6, 0xff, 0x11, 0xf9, 0x07, 0x5c, 0xe6, 0x43, 0xc1, 0x85, 0xd3,
	0xb4, 0x05, 0x52, 0x9c, 0x67, 0x2c, 0x24, 0xf2, 0xee, 0x93, 0xf7, 0x84, 0xbc, 0x77, 0x5a, 0x23,
	0x65, 0x35, 0xed, 0x6d, 0x9f, 0x9d, 0xa6, 0xd8, 0xb4, 0xd7, 0xd2, 0x06, 0xbf, 0xa1, 0xa3, 0xcf,
	0xfe, 0xed, 0xc3, 0x98, 0x67, 0x75, 0x2b, 0x57, 0x45, 0xe6, 0xd6, 0xf7, 0x23, 0x38, 0x4e, 0x68,
	0xbe, 0xf5, 0x40, 0x4f, 0x5b, 0x03, 0xe5, 0xc1, 0x87, 0x35, 0xc1, 0x2d, 0x65, 0xd3, 0x20, 0x2d,
	0x07, 0xce, 0xb3, 0xe7, 0x9e, 0xad, 0x06, 0x69, 0x31, 0x9c, 0x86, 0xdc, 0x04, 0x99, 0xd0, 0x23,
	0xc2, 0x10, 0x81, 0xad, 0xf1, 0x4e, 0x97, 0x96, 0x4f, 0x0f, 0x0b, 0x68, 0xe8, 0x00, 0x3a, 0x3c,
	0x17, 0x70, 0x46, 0xc6, 0xc2, 0xc8, 0x54, 0x25, 0x58, 0x1b, 0xd3, 0x58, 0x42, 0xa7, 0xce, 0xb4,
	0x68, 0x2c, 0xc4, 0xc7, 0x72, 0xf0, 0xc0, 0x2f, 0x4d, 0xcc, 0x4b, 0x57, 0x62, 0xff, 0x54, 0x4e,
	0x0d, 0xba, 0x35, 0x29, 0xc5, 0xe7, 0xf0, 0x68, 0x4b, 0xea, 0xc6, 0xe5, 0x5b, 0x79, 0xde, 0xb0,
	0x3b, 0xa1, 0x3f, 0x84, 0x09, 0xf6, 0x8c, 0xd3, 0xcb, 0xf1, 0xb4, 0x72, 0xf0, 0x21, 0x05, 0x1f,
	0x6f, 0x61, 0x0e, 0xff, 0x25, 0x04, 0x2d, 0x62, 0x37, 0xc1, 0x88, 0x12, 0x3c, 0xda, 0x79, 0x74,
	0x52, 0xe0, 0x4d, 0x58, 0x37, 0x85, 0x03, 0xc5, 0x1e, 0xac, 0xeb, 0x9a, 0xf1, 0x06, 0xad, 0x1f,
	0x94, 0xeb, 0xf1, 0x0d, 0x5a, 0xef, 0x55, 0x8a, 0x07, 0x2f, 0x53, 0x79, 0xf5, 0x3a, 0xe2, 0xd1,
	0x51, 0x38, 0x9f, 0x4b, 0x25, 0xfc, 0x7b, 0x37, 0x36, 0x8a, 0xfa, 0x02, 0x1e, 0xb7, 0x99, 0xdd,
	0xe8, 0x27, 0x5c, 0xeb, 0xce, 0xa5, 0x93, 0xe4, 0x39, 0x9c, 0xb3, 0xa2, 0xf7, 0xbc, 0xc6, 0x7c,
	0xbf, 0xc9, 0xd6, 0xf5, 0xf8, 0x18, 0x04, 0x1d, 0xc3, 0x2e, 0x7f, 0x42, 0xfc, 0xa9, 0x3b, 0x8b,
	0x1d, 0x36, 0x36, 0xac, 0xcb, 0x7d, 0xf2, 0x94, 0x1b, 0xd6, 0x65, 0x97, 0x8b, 0x2a, 0xe1, 0x1b,
	0xd5, 0x65, 0x9f, 0xb2, 0x4a, 0xe8, 0x5a, 0xb5, 0xf9, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf8,
	0x74, 0xe4, 0x10, 0x0c, 0x08, 0x00, 0x00,
}
*/
