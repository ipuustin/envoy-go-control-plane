// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/alts/v3/alts.proto

package envoy_extensions_transport_sockets_alts_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type Alts struct {
	HandshakerService    string   `protobuf:"bytes,1,opt,name=handshaker_service,json=handshakerService,proto3" json:"handshaker_service,omitempty"`
	PeerServiceAccounts  []string `protobuf:"bytes,2,rep,name=peer_service_accounts,json=peerServiceAccounts,proto3" json:"peer_service_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alts) Reset()         { *m = Alts{} }
func (m *Alts) String() string { return proto.CompactTextString(m) }
func (*Alts) ProtoMessage()    {}
func (*Alts) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b780b78c099503e, []int{0}
}

func (m *Alts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alts.Unmarshal(m, b)
}
func (m *Alts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alts.Marshal(b, m, deterministic)
}
func (m *Alts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alts.Merge(m, src)
}
func (m *Alts) XXX_Size() int {
	return xxx_messageInfo_Alts.Size(m)
}
func (m *Alts) XXX_DiscardUnknown() {
	xxx_messageInfo_Alts.DiscardUnknown(m)
}

var xxx_messageInfo_Alts proto.InternalMessageInfo

func (m *Alts) GetHandshakerService() string {
	if m != nil {
		return m.HandshakerService
	}
	return ""
}

func (m *Alts) GetPeerServiceAccounts() []string {
	if m != nil {
		return m.PeerServiceAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Alts)(nil), "envoy.extensions.transport_sockets.alts.v3.Alts")
}

func init() {
	proto.RegisterFile("envoy/extensions/transport_sockets/alts/v3/alts.proto", fileDescriptor_3b780b78c099503e)
}

var fileDescriptor_3b780b78c099503e = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x1c, 0xc5, 0x95, 0xb6, 0xea, 0xa7, 0x7a, 0xfa, 0x08, 0x42, 0x54, 0x95, 0x40, 0x81, 0xa9, 0x62,
	0xb0, 0x45, 0x23, 0x2a, 0xc4, 0xd6, 0x5c, 0x80, 0xaa, 0x1c, 0x20, 0xfa, 0x93, 0x98, 0xc6, 0x6a,
	0x64, 0x5b, 0xf6, 0x3f, 0x56, 0xbb, 0x31, 0x72, 0x06, 0x8e, 0xc0, 0x11, 0xd8, 0x91, 0x58, 0xb9,
	0x0e, 0x13, 0x8a, 0x13, 0x94, 0x21, 0x0b, 0x93, 0x2d, 0xbd, 0xf7, 0x7b, 0xf6, 0x7b, 0xe4, 0x86,
	0x4b, 0xa7, 0x0e, 0x8c, 0xef, 0x91, 0x4b, 0x2b, 0x94, 0xb4, 0x0c, 0x0d, 0x48, 0xab, 0x95, 0xc1,
	0xd4, 0xaa, 0x6c, 0xc7, 0xd1, 0x32, 0x28, 0xd1, 0x32, 0x17, 0xfb, 0x93, 0x6a, 0xa3, 0x50, 0x85,
	0x57, 0x1e, 0xa3, 0x1d, 0x46, 0x7b, 0x18, 0xf5, 0x76, 0x17, 0xcf, 0xce, 0xaa, 0x5c, 0x03, 0x03,
	0x29, 0x15, 0x02, 0xfa, 0x27, 0x2c, 0x02, 0x56, 0x6d, 0xd4, 0xec, 0xa2, 0x27, 0x3b, 0x6e, 0xea,
	0x4c, 0x21, 0xb7, 0xad, 0xe5, 0xd4, 0x41, 0x29, 0x72, 0x40, 0xce, 0x7e, 0x2f, 0x8d, 0x70, 0xf9,
	0x16, 0x90, 0xd1, 0xaa, 0x44, 0x1b, 0x2e, 0x49, 0x58, 0x80, 0xcc, 0x6d, 0x01, 0x3b, 0x6e, 0x52,
	0xcb, 0x8d, 0x13, 0x19, 0x9f, 0x06, 0x51, 0x30, 0x9f, 0x24, 0xff, 0xbe, 0x93, 0x91, 0x19, 0x44,
	0xc1, 0xe6, 0xa8, 0xb3, 0x3c, 0x34, 0x8e, 0x70, 0x41, 0x4e, 0x34, 0xef, 0x88, 0x14, 0xb2, 0x4c,
	0x55, 0x12, 0xed, 0x74, 0x10, 0x0d, 0xe7, 0x93, 0xcd, 0x71, 0x2d, 0xb6, 0xde, 0x55, 0x2b, 0xdd,
	0x2d, 0x5f, 0x3f, 0x5e, 0xce, 0xaf, 0x09, 0x6b, 0x26, 0xc8, 0x94, 0x7c, 0x12, 0xdb, 0x5e, 0xfd,
	0xb6, 0xfd, 0x02, 0x4a, 0x5d, 0x00, 0xad, 0xff, 0x98, 0xdc, 0xbf, 0x3f, 0x7f, 0x7e, 0x8d, 0x07,
	0xff, 0x87, 0xe4, 0x56, 0x28, 0xea, 0x69, 0x6d, 0xd4, 0xfe, 0x40, 0xff, 0xbe, 0x65, 0x32, 0xa9,
	0x93, 0xd6, 0x75, 0xf7, 0x75, 0xf0, 0x38, 0xf6, 0x23, 0xc4, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xe9, 0x2b, 0x12, 0xe1, 0xc4, 0x01, 0x00, 0x00,
}
