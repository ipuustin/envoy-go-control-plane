// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3/address.proto

package envoy_config_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}

var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}

func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_73ff5d16d1e342ac, []int{1, 0}
}

type Pipe struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Mode                 uint32   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipe) Reset()         { *m = Pipe{} }
func (m *Pipe) String() string { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()    {}
func (*Pipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ff5d16d1e342ac, []int{0}
}

func (m *Pipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipe.Unmarshal(m, b)
}
func (m *Pipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipe.Marshal(b, m, deterministic)
}
func (m *Pipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipe.Merge(m, src)
}
func (m *Pipe) XXX_Size() int {
	return xxx_messageInfo_Pipe.Size(m)
}
func (m *Pipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipe.DiscardUnknown(m)
}

var xxx_messageInfo_Pipe proto.InternalMessageInfo

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Pipe) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.config.core.v3.SocketAddress_Protocol" json:"protocol,omitempty"`
	Address  string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier        isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	ResolverName         string                        `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	Ipv4Compat           bool                          `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SocketAddress) Reset()         { *m = SocketAddress{} }
func (m *SocketAddress) String() string { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()    {}
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ff5d16d1e342ac, []int{1}
}

func (m *SocketAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketAddress.Unmarshal(m, b)
}
func (m *SocketAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketAddress.Marshal(b, m, deterministic)
}
func (m *SocketAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketAddress.Merge(m, src)
}
func (m *SocketAddress) XXX_Size() int {
	return xxx_messageInfo_SocketAddress.Size(m)
}
func (m *SocketAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketAddress.DiscardUnknown(m)
}

var xxx_messageInfo_SocketAddress proto.InternalMessageInfo

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof"`
}

type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}

func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetIpv4Compat() bool {
	if m != nil {
		return m.Ipv4Compat
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

type TcpKeepalive struct {
	KeepaliveProbes      *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	KeepaliveTime        *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	KeepaliveInterval    *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TcpKeepalive) Reset()         { *m = TcpKeepalive{} }
func (m *TcpKeepalive) String() string { return proto.CompactTextString(m) }
func (*TcpKeepalive) ProtoMessage()    {}
func (*TcpKeepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ff5d16d1e342ac, []int{2}
}

func (m *TcpKeepalive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpKeepalive.Unmarshal(m, b)
}
func (m *TcpKeepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpKeepalive.Marshal(b, m, deterministic)
}
func (m *TcpKeepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpKeepalive.Merge(m, src)
}
func (m *TcpKeepalive) XXX_Size() int {
	return xxx_messageInfo_TcpKeepalive.Size(m)
}
func (m *TcpKeepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpKeepalive.DiscardUnknown(m)
}

var xxx_messageInfo_TcpKeepalive proto.InternalMessageInfo

func (m *TcpKeepalive) GetKeepaliveProbes() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveProbes
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveTime() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveTime
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveInterval() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveInterval
	}
	return nil
}

type BindConfig struct {
	SourceAddress        *SocketAddress      `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Freebind             *wrappers.BoolValue `protobuf:"bytes,2,opt,name=freebind,proto3" json:"freebind,omitempty"`
	SocketOptions        []*SocketOption     `protobuf:"bytes,3,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BindConfig) Reset()         { *m = BindConfig{} }
func (m *BindConfig) String() string { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()    {}
func (*BindConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ff5d16d1e342ac, []int{3}
}

func (m *BindConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindConfig.Unmarshal(m, b)
}
func (m *BindConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindConfig.Marshal(b, m, deterministic)
}
func (m *BindConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindConfig.Merge(m, src)
}
func (m *BindConfig) XXX_Size() int {
	return xxx_messageInfo_BindConfig.Size(m)
}
func (m *BindConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BindConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BindConfig proto.InternalMessageInfo

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *BindConfig) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *BindConfig) GetSocketOptions() []*SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address              isAddress_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ff5d16d1e342ac, []int{4}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof"`
}

type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}

func (*Address_Pipe) isAddress_Address() {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Address) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

type CidrRange struct {
	AddressPrefix        string                `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"`
	PrefixLen            *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CidrRange) Reset()         { *m = CidrRange{} }
func (m *CidrRange) String() string { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()    {}
func (*CidrRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ff5d16d1e342ac, []int{5}
}

func (m *CidrRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrRange.Unmarshal(m, b)
}
func (m *CidrRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrRange.Marshal(b, m, deterministic)
}
func (m *CidrRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrRange.Merge(m, src)
}
func (m *CidrRange) XXX_Size() int {
	return xxx_messageInfo_CidrRange.Size(m)
}
func (m *CidrRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrRange.DiscardUnknown(m)
}

var xxx_messageInfo_CidrRange proto.InternalMessageInfo

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
	proto.RegisterType((*Pipe)(nil), "envoy.config.core.v3.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.config.core.v3.SocketAddress")
	proto.RegisterType((*TcpKeepalive)(nil), "envoy.config.core.v3.TcpKeepalive")
	proto.RegisterType((*BindConfig)(nil), "envoy.config.core.v3.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.config.core.v3.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.config.core.v3.CidrRange")
}

func init() { proto.RegisterFile("envoy/config/core/v3/address.proto", fileDescriptor_73ff5d16d1e342ac) }

var fileDescriptor_73ff5d16d1e342ac = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcb, 0x8e, 0xe3, 0x44,
	0x14, 0x6d, 0xc7, 0x99, 0x3c, 0x6e, 0x77, 0x42, 0x28, 0xf1, 0x30, 0xe9, 0x74, 0x26, 0xed, 0x16,
	0x28, 0x6a, 0x21, 0x1b, 0x25, 0x88, 0x45, 0xc4, 0x66, 0x1c, 0x10, 0xdd, 0x9a, 0x11, 0x58, 0x26,
	0xc3, 0xd6, 0x54, 0xec, 0x4a, 0x28, 0x8d, 0xe3, 0x2a, 0x95, 0x1d, 0x33, 0xb3, 0x1b, 0xb1, 0x42,
	0xb3, 0x62, 0xcd, 0x07, 0xf0, 0x11, 0xb0, 0x46, 0x62, 0xcb, 0xcf, 0x20, 0xd4, 0x9b, 0x46, 0x55,
	0x7e, 0xa4, 0x87, 0x44, 0xcc, 0xb0, 0x4b, 0xdd, 0x7b, 0xcf, 0xa9, 0x73, 0x8f, 0x4f, 0x0a, 0x4c,
	0x12, 0x67, 0xec, 0x99, 0x1d, 0xb0, 0x78, 0x45, 0xd7, 0x76, 0xc0, 0x04, 0xb1, 0xb3, 0xa9, 0x8d,
	0xc3, 0x50, 0x90, 0x24, 0xb1, 0xb8, 0x60, 0x29, 0x43, 0x6f, 0xa9, 0x19, 0x2b, 0x9f, 0xb1, 0xe4,
	0x8c, 0x95, 0x4d, 0xfb, 0xe3, 0x83, 0xc8, 0x84, 0x05, 0x4f, 0x48, 0xea, 0x33, 0x9e, 0x52, 0x16,
	0xe7, 0xf8, 0xfe, 0x70, 0xcd, 0xd8, 0x3a, 0x22, 0xb6, 0x3a, 0x2d, 0xb7, 0x2b, 0xfb, 0x7b, 0x81,
	0x39, 0x27, 0xa2, 0xe0, 0xef, 0x9f, 0x6d, 0x43, 0x8e, 0x6d, 0x1c, 0xc7, 0x2c, 0xc5, 0x12, 0x96,
	0xd8, 0x49, 0x8a, 0xd3, 0x6d, 0xd9, 0x3e, 0xdf, 0x6b, 0x67, 0x44, 0x24, 0x94, 0xc5, 0x34, 0x5e,
	0x17, 0x23, 0xef, 0x66, 0x38, 0xa2, 0x21, 0x4e, 0x89, 0x5d, 0xfe, 0xc8, 0x1b, 0xe6, 0xb7, 0x50,
	0x77, 0x29, 0x27, 0xe8, 0x14, 0xea, 0x1c, 0xa7, 0xdf, 0x19, 0xda, 0x48, 0x1b, 0xb7, 0x9d, 0xe6,
	0x8d, 0x53, 0x17, 0xb5, 0x91, 0xe6, 0xa9, 0x22, 0x1a, 0x40, 0x7d, 0xc3, 0x42, 0x62, 0xd4, 0x46,
	0xda, 0xb8, 0xe3, 0xb4, 0x6e, 0x9c, 0x7b, 0x97, 0xba, 0x71, 0xab, 0x7b, 0xaa, 0x3a, 0x3b, 0xfb,
	0xf9, 0xf7, 0x1f, 0x87, 0x06, 0xbc, 0x93, 0x9b, 0x80, 0x39, 0xb5, 0xb2, 0x49, 0x6e, 0x82, 0x64,
	0x36, 0xff, 0xaa, 0x41, 0xe7, 0x6b, 0xb5, 0xf4, 0x83, 0xdc, 0x34, 0xe4, 0x41, 0x4b, 0x5d, 0x1e,
	0xb0, 0x48, 0xdd, 0xd7, 0x9d, 0x7c, 0x68, 0x1d, 0x72, 0xd0, 0x7a, 0x09, 0x66, 0xb9, 0x05, 0x46,
	0x09, 0xf8, 0x41, 0xab, 0xf5, 0x34, 0xaf, 0xe2, 0x41, 0xe7, 0xd0, 0x2c, 0xbe, 0x89, 0x52, 0x79,
	0x67, 0x85, 0xb2, 0x8e, 0x2e, 0x01, 0x38, 0x13, 0xa9, 0x9f, 0xe1, 0x68, 0x4b, 0x0c, 0x5d, 0xed,
	0xd2, 0xbe, 0x71, 0x1a, 0x97, 0x75, 0xe3, 0xf6, 0x56, 0xbf, 0x3a, 0xf2, 0xda, 0xb2, 0xfd, 0x8d,
	0xec, 0xa2, 0xfb, 0x00, 0x31, 0xde, 0x90, 0xd0, 0x97, 0x25, 0xa3, 0x2e, 0x19, 0xe5, 0x80, 0xaa,
	0xb9, 0x4c, 0xa4, 0xe8, 0x02, 0x3a, 0x82, 0x24, 0x2c, 0xca, 0x88, 0xf0, 0x65, 0xd5, 0xb8, 0x27,
	0x67, 0xbc, 0x93, 0xb2, 0xf8, 0x25, 0xde, 0x48, 0x96, 0x63, 0xca, 0xb3, 0x8f, 0xfd, 0x80, 0x6d,
	0x38, 0x4e, 0x8d, 0xc6, 0x48, 0x1b, 0xb7, 0x3c, 0x90, 0xa5, 0xb9, 0xaa, 0x98, 0x03, 0x68, 0x95,
	0x5b, 0xa1, 0x26, 0xe8, 0x8b, 0xb9, 0xdb, 0x3b, 0x92, 0x3f, 0x1e, 0x7f, 0xe6, 0xf6, 0xb4, 0xd9,
	0x07, 0xd2, 0xd8, 0x73, 0xb8, 0xbf, 0x6f, 0xec, 0x4b, 0xc6, 0x38, 0x6f, 0x43, 0x57, 0x2d, 0x96,
	0x70, 0x12, 0xd0, 0x15, 0x25, 0x02, 0xe9, 0x7f, 0x3b, 0x9a, 0xf9, 0x53, 0x0d, 0x4e, 0x16, 0x01,
	0x7f, 0x48, 0x08, 0xc7, 0x11, 0xcd, 0x08, 0xfa, 0x02, 0x7a, 0x4f, 0xca, 0x83, 0xcf, 0x05, 0x5b,
	0x92, 0x44, 0xf9, 0x7f, 0x3c, 0x19, 0x58, 0x79, 0x02, 0xad, 0x32, 0x81, 0xd6, 0xe3, 0xeb, 0x38,
	0x9d, 0x4e, 0x94, 0x19, 0xde, 0x1b, 0x15, 0xca, 0x55, 0x20, 0x34, 0x87, 0xee, 0x8e, 0x28, 0xa5,
	0x9b, 0x3c, 0x19, 0xaf, 0xa2, 0xe9, 0x54, 0x98, 0x05, 0xdd, 0x10, 0xf4, 0x10, 0xd0, 0x8e, 0x84,
	0xc6, 0x29, 0x11, 0x19, 0x8e, 0xd4, 0x67, 0x79, 0x15, 0xd1, 0x9b, 0x15, 0xee, 0xba, 0x80, 0xcd,
	0xde, 0x97, 0x56, 0x8d, 0x60, 0xb8, 0x6f, 0xd5, 0x5d, 0x07, 0xcc, 0x17, 0x35, 0x00, 0x87, 0xc6,
	0xe1, 0x5c, 0xe5, 0x0c, 0x2d, 0xa0, 0x9b, 0xb0, 0xad, 0x08, 0x88, 0x5f, 0x66, 0x27, 0xb7, 0xe3,
	0xe2, 0x35, 0xe2, 0xa8, 0x52, 0xf8, 0x42, 0xa5, 0xb0, 0x93, 0x93, 0x94, 0xf1, 0xfe, 0x04, 0x5a,
	0x2b, 0x41, 0xc8, 0x92, 0xc6, 0x61, 0xe1, 0x4b, 0x7f, 0x6f, 0x1d, 0x87, 0xb1, 0x28, 0x5f, 0xa6,
	0x9a, 0x45, 0xd7, 0x52, 0xcd, 0x9d, 0xc7, 0x21, 0x31, 0xf4, 0x91, 0x3e, 0x3e, 0x9e, 0x98, 0xff,
	0xa5, 0xe6, 0x2b, 0x35, 0x2a, 0x25, 0xec, 0x4e, 0xc9, 0xec, 0x42, 0xda, 0x31, 0x84, 0xc1, 0xbe,
	0x1d, 0xbb, 0xed, 0xcd, 0xdf, 0x34, 0x68, 0x96, 0x9a, 0x1f, 0x55, 0x77, 0xff, 0x7f, 0x27, 0xae,
	0x8e, 0xca, 0xeb, 0x4b, 0xb6, 0x8f, 0xa0, 0xce, 0x29, 0x27, 0xd5, 0xf6, 0x07, 0x39, 0xe4, 0xe3,
	0x70, 0x75, 0xe4, 0xa9, 0xc9, 0xd9, 0x48, 0x0a, 0x3e, 0x85, 0xf7, 0xf6, 0x05, 0x97, 0x76, 0x77,
	0xab, 0x3f, 0x78, 0x9e, 0xee, 0x5f, 0x34, 0x68, 0xcf, 0x69, 0x28, 0x3c, 0x1c, 0xaf, 0x09, 0xb2,
	0xa0, 0x5b, 0x74, 0x7d, 0x2e, 0xc8, 0x8a, 0x3e, 0xfd, 0xf7, 0x43, 0xd6, 0x29, 0xda, 0xae, 0xea,
	0xa2, 0xcf, 0x01, 0xf2, 0x39, 0x3f, 0x22, 0xf1, 0xeb, 0xa4, 0xb7, 0x7c, 0xf5, 0x9e, 0x6b, 0x5e,
	0x3b, 0x47, 0x3e, 0x22, 0xf1, 0xcc, 0x94, 0xb2, 0xcf, 0xe0, 0x74, 0x5f, 0x76, 0x25, 0xcd, 0xf9,
	0xf4, 0xd7, 0xe7, 0x7f, 0xfc, 0xd9, 0xa8, 0xf5, 0x74, 0x30, 0x29, 0xcb, 0xad, 0xe0, 0x82, 0x3d,
	0x7d, 0x76, 0xd0, 0x15, 0xe7, 0xe4, 0x41, 0xa9, 0x93, 0xa5, 0xcc, 0xd5, 0x96, 0x0d, 0x25, 0x66,
	0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xc3, 0x29, 0x36, 0x87, 0x06, 0x00, 0x00,
}
