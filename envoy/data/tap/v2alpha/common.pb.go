// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/common.proto

package envoy_data_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Body struct {
	// Types that are valid to be assigned to BodyType:
	//	*Body_AsBytes
	//	*Body_AsString
	BodyType             isBody_BodyType `protobuf_oneof:"body_type"`
	Truncated            bool            `protobuf:"varint,3,opt,name=truncated,proto3" json:"truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Body) Reset()         { *m = Body{} }
func (m *Body) String() string { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()    {}
func (*Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_a560f1e2899ebe7a, []int{0}
}

func (m *Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Body.Unmarshal(m, b)
}
func (m *Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Body.Marshal(b, m, deterministic)
}
func (m *Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Body.Merge(m, src)
}
func (m *Body) XXX_Size() int {
	return xxx_messageInfo_Body.Size(m)
}
func (m *Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Body.DiscardUnknown(m)
}

var xxx_messageInfo_Body proto.InternalMessageInfo

type isBody_BodyType interface {
	isBody_BodyType()
}

type Body_AsBytes struct {
	AsBytes []byte `protobuf:"bytes,1,opt,name=as_bytes,json=asBytes,proto3,oneof"`
}

type Body_AsString struct {
	AsString string `protobuf:"bytes,2,opt,name=as_string,json=asString,proto3,oneof"`
}

func (*Body_AsBytes) isBody_BodyType() {}

func (*Body_AsString) isBody_BodyType() {}

func (m *Body) GetBodyType() isBody_BodyType {
	if m != nil {
		return m.BodyType
	}
	return nil
}

func (m *Body) GetAsBytes() []byte {
	if x, ok := m.GetBodyType().(*Body_AsBytes); ok {
		return x.AsBytes
	}
	return nil
}

func (m *Body) GetAsString() string {
	if x, ok := m.GetBodyType().(*Body_AsString); ok {
		return x.AsString
	}
	return ""
}

func (m *Body) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Body) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Body_AsBytes)(nil),
		(*Body_AsString)(nil),
	}
}

func init() {
	proto.RegisterType((*Body)(nil), "envoy.data.tap.v2alpha.Body")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/common.proto", fileDescriptor_a560f1e2899ebe7a)
}

var fileDescriptor_a560f1e2899ebe7a = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x14, 0x45, 0x77, 0xa2, 0xac, 0xc9, 0xac, 0x85, 0xa4, 0x90, 0xa0, 0x2e, 0x04, 0xb5, 0x48, 0x35,
	0x03, 0x5a, 0xdb, 0x8c, 0xcd, 0x96, 0x4b, 0xfc, 0x80, 0xf0, 0xb2, 0x13, 0x34, 0x60, 0xe6, 0x0d,
	0x79, 0x2f, 0xc1, 0xe9, 0xfc, 0x2e, 0xbf, 0xc0, 0xd6, 0x3f, 0x92, 0x8c, 0x82, 0x8d, 0xed, 0xe1,
	0x5c, 0x38, 0x57, 0xde, 0x74, 0x6e, 0xc6, 0xa0, 0x2d, 0x30, 0x68, 0x06, 0xaf, 0xe7, 0x3b, 0x78,
	0xf5, 0x2f, 0xa0, 0x0f, 0x38, 0x0c, 0xe8, 0x94, 0x1f, 0x91, 0x31, 0x3f, 0x8f, 0x92, 0x5a, 0x24,
	0xc5, 0xe0, 0xd5, 0xaf, 0x74, 0xb1, 0x9d, 0xac, 0x07, 0x0d, 0xce, 0x21, 0x03, 0xf7, 0xe8, 0x48,
	0x13, 0x03, 0x4f, 0xf4, 0x33, 0xbb, 0x1e, 0xe4, 0xb1, 0x41, 0x1b, 0xf2, 0x4b, 0x99, 0x02, 0x35,
	0x6d, 0xe0, 0x8e, 0x0a, 0x51, 0x8a, 0xea, 0x74, 0xb7, 0xaa, 0x4f, 0x80, 0xcc, 0x02, 0xf2, 0xad,
	0xcc, 0x80, 0x1a, 0xe2, 0xb1, 0x77, 0xcf, 0x45, 0x52, 0x8a, 0x2a, 0xdb, 0xad, 0xea, 0x14, 0xe8,
	0x29, 0x92, 0xfc, 0x4a, 0x66, 0x3c, 0x4e, 0xee, 0x00, 0xdc, 0xd9, 0xe2, 0xa8, 0x14, 0x55, 0x5a,
	0xff, 0x01, 0xb3, 0x91, 0x59, 0x8b, 0x36, 0x34, 0x1c, 0x7c, 0x67, 0x1e, 0x3e, 0xde, 0x3f, 0xbf,
	0xd6, 0xc9, 0x59, 0x22, 0x6f, 0x7b, 0x54, 0x31, 0xd9, 0x8f, 0xf8, 0x16, 0xd4, 0xff, 0xf5, 0x66,
	0xf3, 0x18, 0x3f, 0xee, 0x97, 0xd6, 0xbd, 0x68, 0xd7, 0x31, 0xfa, 0xfe, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x34, 0x90, 0x36, 0xbb, 0x12, 0x01, 0x00, 0x00,
}
