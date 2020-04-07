// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/dns/v2alpha/dns_table.proto

package envoy_data_dns_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DnsTable struct {
	ExternalRetryCount   uint32                       `protobuf:"varint,1,opt,name=external_retry_count,json=externalRetryCount,proto3" json:"external_retry_count,omitempty"`
	VirtualDomains       []*DnsTable_DnsVirtualDomain `protobuf:"bytes,2,rep,name=virtual_domains,json=virtualDomains,proto3" json:"virtual_domains,omitempty"`
	KnownSuffixes        []*matcher.StringMatcher     `protobuf:"bytes,3,rep,name=known_suffixes,json=knownSuffixes,proto3" json:"known_suffixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DnsTable) Reset()         { *m = DnsTable{} }
func (m *DnsTable) String() string { return proto.CompactTextString(m) }
func (*DnsTable) ProtoMessage()    {}
func (*DnsTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_74e6eda848825994, []int{0}
}

func (m *DnsTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable.Unmarshal(m, b)
}
func (m *DnsTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable.Marshal(b, m, deterministic)
}
func (m *DnsTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable.Merge(m, src)
}
func (m *DnsTable) XXX_Size() int {
	return xxx_messageInfo_DnsTable.Size(m)
}
func (m *DnsTable) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable proto.InternalMessageInfo

func (m *DnsTable) GetExternalRetryCount() uint32 {
	if m != nil {
		return m.ExternalRetryCount
	}
	return 0
}

func (m *DnsTable) GetVirtualDomains() []*DnsTable_DnsVirtualDomain {
	if m != nil {
		return m.VirtualDomains
	}
	return nil
}

func (m *DnsTable) GetKnownSuffixes() []*matcher.StringMatcher {
	if m != nil {
		return m.KnownSuffixes
	}
	return nil
}

type DnsTable_AddressList struct {
	Address              []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnsTable_AddressList) Reset()         { *m = DnsTable_AddressList{} }
func (m *DnsTable_AddressList) String() string { return proto.CompactTextString(m) }
func (*DnsTable_AddressList) ProtoMessage()    {}
func (*DnsTable_AddressList) Descriptor() ([]byte, []int) {
	return fileDescriptor_74e6eda848825994, []int{0, 0}
}

func (m *DnsTable_AddressList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_AddressList.Unmarshal(m, b)
}
func (m *DnsTable_AddressList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_AddressList.Marshal(b, m, deterministic)
}
func (m *DnsTable_AddressList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_AddressList.Merge(m, src)
}
func (m *DnsTable_AddressList) XXX_Size() int {
	return xxx_messageInfo_DnsTable_AddressList.Size(m)
}
func (m *DnsTable_AddressList) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_AddressList.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_AddressList proto.InternalMessageInfo

func (m *DnsTable_AddressList) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

type DnsTable_DnsEndpoint struct {
	// Types that are valid to be assigned to EndpointConfig:
	//	*DnsTable_DnsEndpoint_AddressList
	EndpointConfig       isDnsTable_DnsEndpoint_EndpointConfig `protobuf_oneof:"endpoint_config"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DnsTable_DnsEndpoint) Reset()         { *m = DnsTable_DnsEndpoint{} }
func (m *DnsTable_DnsEndpoint) String() string { return proto.CompactTextString(m) }
func (*DnsTable_DnsEndpoint) ProtoMessage()    {}
func (*DnsTable_DnsEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_74e6eda848825994, []int{0, 1}
}

func (m *DnsTable_DnsEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Unmarshal(m, b)
}
func (m *DnsTable_DnsEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Marshal(b, m, deterministic)
}
func (m *DnsTable_DnsEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_DnsEndpoint.Merge(m, src)
}
func (m *DnsTable_DnsEndpoint) XXX_Size() int {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Size(m)
}
func (m *DnsTable_DnsEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_DnsEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_DnsEndpoint proto.InternalMessageInfo

type isDnsTable_DnsEndpoint_EndpointConfig interface {
	isDnsTable_DnsEndpoint_EndpointConfig()
}

type DnsTable_DnsEndpoint_AddressList struct {
	AddressList *DnsTable_AddressList `protobuf:"bytes,1,opt,name=address_list,json=addressList,proto3,oneof"`
}

func (*DnsTable_DnsEndpoint_AddressList) isDnsTable_DnsEndpoint_EndpointConfig() {}

func (m *DnsTable_DnsEndpoint) GetEndpointConfig() isDnsTable_DnsEndpoint_EndpointConfig {
	if m != nil {
		return m.EndpointConfig
	}
	return nil
}

func (m *DnsTable_DnsEndpoint) GetAddressList() *DnsTable_AddressList {
	if x, ok := m.GetEndpointConfig().(*DnsTable_DnsEndpoint_AddressList); ok {
		return x.AddressList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DnsTable_DnsEndpoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DnsTable_DnsEndpoint_AddressList)(nil),
	}
}

type DnsTable_DnsVirtualDomain struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Endpoint             *DnsTable_DnsEndpoint `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AnswerTtl            *duration.Duration    `protobuf:"bytes,3,opt,name=answer_ttl,json=answerTtl,proto3" json:"answer_ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DnsTable_DnsVirtualDomain) Reset()         { *m = DnsTable_DnsVirtualDomain{} }
func (m *DnsTable_DnsVirtualDomain) String() string { return proto.CompactTextString(m) }
func (*DnsTable_DnsVirtualDomain) ProtoMessage()    {}
func (*DnsTable_DnsVirtualDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_74e6eda848825994, []int{0, 2}
}

func (m *DnsTable_DnsVirtualDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Unmarshal(m, b)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Marshal(b, m, deterministic)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_DnsVirtualDomain.Merge(m, src)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Size() int {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Size(m)
}
func (m *DnsTable_DnsVirtualDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_DnsVirtualDomain.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_DnsVirtualDomain proto.InternalMessageInfo

func (m *DnsTable_DnsVirtualDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsTable_DnsVirtualDomain) GetEndpoint() *DnsTable_DnsEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *DnsTable_DnsVirtualDomain) GetAnswerTtl() *duration.Duration {
	if m != nil {
		return m.AnswerTtl
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsTable)(nil), "envoy.data.dns.v2alpha.DnsTable")
	proto.RegisterType((*DnsTable_AddressList)(nil), "envoy.data.dns.v2alpha.DnsTable.AddressList")
	proto.RegisterType((*DnsTable_DnsEndpoint)(nil), "envoy.data.dns.v2alpha.DnsTable.DnsEndpoint")
	proto.RegisterType((*DnsTable_DnsVirtualDomain)(nil), "envoy.data.dns.v2alpha.DnsTable.DnsVirtualDomain")
}

func init() {
	proto.RegisterFile("envoy/data/dns/v2alpha/dns_table.proto", fileDescriptor_74e6eda848825994)
}

var fileDescriptor_74e6eda848825994 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8e, 0xd3, 0x3c,
	0x10, 0x5e, 0x27, 0xdd, 0xdd, 0xd4, 0xf9, 0xdb, 0xad, 0xa2, 0x5f, 0x4b, 0xa8, 0xc4, 0x52, 0x56,
	0x08, 0x55, 0x08, 0x39, 0x50, 0x0e, 0x9c, 0x31, 0x45, 0xf4, 0x00, 0xd2, 0x92, 0x5d, 0x71, 0x8d,
	0xdc, 0xda, 0xed, 0x5a, 0xa4, 0x76, 0x64, 0x3b, 0xdd, 0xf6, 0xc6, 0x33, 0xec, 0x63, 0xf0, 0x08,
	0x9c, 0xf6, 0xc8, 0x95, 0x57, 0xe0, 0x29, 0x50, 0x4f, 0x28, 0x76, 0x02, 0x15, 0x02, 0xc1, 0x29,
	0x33, 0xfe, 0xe6, 0x9b, 0x6f, 0xe6, 0xcb, 0xc0, 0x07, 0x4c, 0xac, 0xe4, 0x26, 0xa1, 0xc4, 0x90,
	0x84, 0x0a, 0x9d, 0xac, 0x46, 0x24, 0x2f, 0x2e, 0x6d, 0x9c, 0x19, 0x32, 0xcd, 0x19, 0x2a, 0x94,
	0x34, 0x32, 0x3a, 0xb6, 0x75, 0xa8, 0xaa, 0x43, 0x54, 0x68, 0x54, 0xd7, 0xf5, 0xef, 0x3a, 0xbe,
	0xd9, 0x14, 0x2c, 0x59, 0x12, 0x33, 0xbb, 0x64, 0x2a, 0xd1, 0x46, 0x71, 0xb1, 0x70, 0xc4, 0xfe,
	0xc9, 0x42, 0xca, 0x45, 0xce, 0x12, 0x9b, 0x4d, 0xcb, 0x79, 0x42, 0x4b, 0x45, 0x0c, 0x97, 0xa2,
	0xc6, 0xef, 0x94, 0xb4, 0x20, 0x09, 0x11, 0x42, 0x1a, 0xfb, 0xac, 0x13, 0x6d, 0x88, 0x29, 0x75,
	0x0d, 0xdf, 0x5a, 0x91, 0x9c, 0x53, 0x62, 0x58, 0xd2, 0x04, 0x0e, 0x38, 0xfd, 0xda, 0x82, 0xc1,
	0x58, 0xe8, 0x8b, 0x6a, 0xc6, 0xe8, 0x31, 0xfc, 0x9f, 0xad, 0x0d, 0x53, 0x82, 0xe4, 0x99, 0x62,
	0x46, 0x6d, 0xb2, 0x99, 0x2c, 0x85, 0x89, 0xc1, 0x00, 0x0c, 0x3b, 0x69, 0xd4, 0x60, 0x69, 0x05,
	0xbd, 0xa8, 0x90, 0x88, 0xc2, 0xa3, 0x15, 0x57, 0xa6, 0x24, 0x79, 0x46, 0xe5, 0x92, 0x70, 0xa1,
	0x63, 0x6f, 0xe0, 0x0f, 0xc3, 0xd1, 0x13, 0xf4, 0xfb, 0x4d, 0x51, 0x23, 0x56, 0x05, 0xef, 0x1c,
	0x75, 0x6c, 0x99, 0x38, 0xd8, 0xe2, 0xfd, 0x6b, 0xe0, 0x05, 0x20, 0xed, 0xae, 0x76, 0x01, 0x1d,
	0x4d, 0x60, 0xf7, 0xbd, 0x90, 0x57, 0x22, 0xd3, 0xe5, 0x7c, 0xce, 0xd7, 0x4c, 0xc7, 0xbe, 0x15,
	0xb9, 0x57, 0x8b, 0x54, 0xb6, 0xa1, 0xda, 0x36, 0x74, 0x6e, 0x6d, 0x7b, 0xe3, 0xb2, 0xb4, 0x63,
	0x89, 0xe7, 0x35, 0xaf, 0xff, 0x0c, 0x86, 0xcf, 0x29, 0x55, 0x4c, 0xeb, 0xd7, 0x5c, 0x9b, 0x68,
	0x08, 0x0f, 0x89, 0x4b, 0x63, 0x30, 0xf0, 0x87, 0x6d, 0xdc, 0xdd, 0xe2, 0xf0, 0x1a, 0x04, 0x01,
	0x38, 0x6d, 0x29, 0xaf, 0xe7, 0xa7, 0x0d, 0xdc, 0x5f, 0xc3, 0x70, 0x2c, 0xf4, 0x4b, 0x41, 0x0b,
	0xc9, 0x85, 0x89, 0xde, 0xc2, 0xff, 0x6a, 0x24, 0xcb, 0xb9, 0x76, 0x0e, 0x85, 0xa3, 0x47, 0x7f,
	0x5d, 0x7a, 0x47, 0x7c, 0xb2, 0x97, 0x86, 0xe4, 0x67, 0x8a, 0x8f, 0xe1, 0x11, 0xab, 0xdb, 0x67,
	0x33, 0x29, 0xe6, 0x7c, 0x11, 0xf9, 0xdf, 0x30, 0xe8, 0xdf, 0x00, 0xd8, 0xfb, 0xd5, 0xab, 0xe8,
	0x04, 0xb6, 0x04, 0x59, 0x32, 0xab, 0xdb, 0xc6, 0x70, 0x8b, 0x0f, 0xd5, 0x7e, 0xcf, 0xbb, 0x01,
	0x20, 0xb5, 0xef, 0xd1, 0x04, 0x06, 0x4d, 0xb3, 0xd8, 0xfb, 0xc7, 0xd9, 0x76, 0xf6, 0x4b, 0x7f,
	0xb0, 0x23, 0x0c, 0x21, 0x11, 0xfa, 0x8a, 0xa9, 0xcc, 0x98, 0x3c, 0xf6, 0x6d, 0xaf, 0xdb, 0xc8,
	0x5d, 0x23, 0x6a, 0xae, 0x11, 0x8d, 0xeb, 0x6b, 0xb4, 0x3f, 0xf1, 0x23, 0xf0, 0x1e, 0xee, 0xa5,
	0x6d, 0x47, 0xbb, 0x30, 0x39, 0x7e, 0xf5, 0xe9, 0xc3, 0xe7, 0x2f, 0x07, 0x5e, 0x00, 0xdc, 0xb7,
	0xe7, 0xc1, 0xfb, 0x5c, 0xba, 0x79, 0x0a, 0x25, 0xd7, 0x9b, 0x3f, 0x8c, 0x86, 0x3b, 0xcd, 0x6c,
	0x67, 0x95, 0xce, 0x19, 0x98, 0x1e, 0x58, 0xc1, 0xa7, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x96,
	0xc1, 0x64, 0x68, 0x6f, 0x03, 0x00, 0x00,
}
