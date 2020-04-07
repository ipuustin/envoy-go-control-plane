// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/lua/v2/lua.proto

package envoy_config_filter_http_lua_v2

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

type Lua struct {
	InlineCode           string   `protobuf:"bytes,1,opt,name=inline_code,json=inlineCode,proto3" json:"inline_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lua) Reset()         { *m = Lua{} }
func (m *Lua) String() string { return proto.CompactTextString(m) }
func (*Lua) ProtoMessage()    {}
func (*Lua) Descriptor() ([]byte, []int) {
	return fileDescriptor_f59dca3e63e33613, []int{0}
}

func (m *Lua) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lua.Unmarshal(m, b)
}
func (m *Lua) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lua.Marshal(b, m, deterministic)
}
func (m *Lua) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lua.Merge(m, src)
}
func (m *Lua) XXX_Size() int {
	return xxx_messageInfo_Lua.Size(m)
}
func (m *Lua) XXX_DiscardUnknown() {
	xxx_messageInfo_Lua.DiscardUnknown(m)
}

var xxx_messageInfo_Lua proto.InternalMessageInfo

func (m *Lua) GetInlineCode() string {
	if m != nil {
		return m.InlineCode
	}
	return ""
}

func init() {
	proto.RegisterType((*Lua)(nil), "envoy.config.filter.http.lua.v2.Lua")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/lua/v2/lua.proto", fileDescriptor_f59dca3e63e33613)
}

var fileDescriptor_f59dca3e63e33613 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x29, 0x4d, 0xd4, 0x2f, 0x33, 0x02, 0x51, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xf2, 0x60, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x10, 0xa5, 0x7a, 0x20, 0xa5,
	0x7a, 0x20, 0x35, 0x65, 0x46, 0x52, 0x72, 0xa5, 0x29, 0x05, 0x89, 0xfa, 0x89, 0x79, 0x79, 0xf9,
	0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xfa, 0xb9, 0x99, 0xe9, 0x45, 0x89, 0x25, 0xa9, 0x10,
	0x03, 0xa4, 0x64, 0x31, 0xe4, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0xa1, 0xd2, 0xe2, 0x65, 0x89,
	0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x42, 0x49, 0x9f, 0x8b, 0xd9, 0xa7,
	0x34, 0x51, 0x48, 0x83, 0x8b, 0x3b, 0x33, 0x2f, 0x27, 0x33, 0x2f, 0x35, 0x3e, 0x39, 0x3f, 0x25,
	0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x89, 0xfd, 0x97, 0x13, 0x4b, 0x11, 0x93, 0x02, 0x63,
	0x10, 0x17, 0x44, 0xce, 0x39, 0x3f, 0x25, 0xd5, 0x29, 0xff, 0xd3, 0x8c, 0x7f, 0xfd, 0xac, 0x6a,
	0x42, 0x2a, 0x10, 0x17, 0xa7, 0x56, 0x94, 0xa4, 0xe6, 0x15, 0x83, 0x2c, 0x84, 0xba, 0xba, 0x18,
	0xc9, 0xd9, 0xc6, 0xbb, 0x1a, 0x4e, 0x5c, 0x64, 0x63, 0x12, 0x60, 0xe2, 0xd2, 0xcd, 0xcc, 0xd7,
	0x03, 0x6b, 0x28, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0xe0, 0x5b, 0x27, 0x0e, 0x9f, 0xd2, 0xc4,
	0x00, 0x90, 0xfb, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x0e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0x10, 0x8c, 0xe1, 0x4e, 0x01, 0x00, 0x00,
}
