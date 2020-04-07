// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/metrics.proto

package envoy_admin_v2alpha

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

type SimpleMetric_Type int32

const (
	SimpleMetric_COUNTER SimpleMetric_Type = 0
	SimpleMetric_GAUGE   SimpleMetric_Type = 1
)

var SimpleMetric_Type_name = map[int32]string{
	0: "COUNTER",
	1: "GAUGE",
}

var SimpleMetric_Type_value = map[string]int32{
	"COUNTER": 0,
	"GAUGE":   1,
}

func (x SimpleMetric_Type) String() string {
	return proto.EnumName(SimpleMetric_Type_name, int32(x))
}

func (SimpleMetric_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_680a736ec6584458, []int{0, 0}
}

type SimpleMetric struct {
	Type                 SimpleMetric_Type `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.admin.v2alpha.SimpleMetric_Type" json:"type,omitempty"`
	Value                uint64            `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Name                 string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SimpleMetric) Reset()         { *m = SimpleMetric{} }
func (m *SimpleMetric) String() string { return proto.CompactTextString(m) }
func (*SimpleMetric) ProtoMessage()    {}
func (*SimpleMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_680a736ec6584458, []int{0}
}

func (m *SimpleMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMetric.Unmarshal(m, b)
}
func (m *SimpleMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMetric.Marshal(b, m, deterministic)
}
func (m *SimpleMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMetric.Merge(m, src)
}
func (m *SimpleMetric) XXX_Size() int {
	return xxx_messageInfo_SimpleMetric.Size(m)
}
func (m *SimpleMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMetric.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMetric proto.InternalMessageInfo

func (m *SimpleMetric) GetType() SimpleMetric_Type {
	if m != nil {
		return m.Type
	}
	return SimpleMetric_COUNTER
}

func (m *SimpleMetric) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *SimpleMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("envoy.admin.v2alpha.SimpleMetric_Type", SimpleMetric_Type_name, SimpleMetric_Type_value)
	proto.RegisterType((*SimpleMetric)(nil), "envoy.admin.v2alpha.SimpleMetric")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/metrics.proto", fileDescriptor_680a736ec6584458) }

var fileDescriptor_680a736ec6584458 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0xd4,
	0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06,
	0x2b, 0xd1, 0x03, 0x2b, 0xd1, 0x83, 0x2a, 0x91, 0x92, 0x2d, 0x4d, 0x29, 0x48, 0xd4, 0x4f, 0xcc,
	0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x85,
	0xea, 0x51, 0x9a, 0xc2, 0xc8, 0xc5, 0x13, 0x9c, 0x99, 0x5b, 0x90, 0x93, 0xea, 0x0b, 0x36, 0x4b,
	0xc8, 0x8a, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x4d,
	0x0f, 0x8b, 0x99, 0x7a, 0xc8, 0x1a, 0xf4, 0x42, 0x2a, 0x0b, 0x52, 0x83, 0xc0, 0x7a, 0x84, 0x44,
	0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x20, 0x1c,
	0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30,
	0x5b, 0x49, 0x8e, 0x8b, 0x05, 0xa4, 0x4f, 0x88, 0x9b, 0x8b, 0xdd, 0xd9, 0x3f, 0xd4, 0x2f, 0xc4,
	0x35, 0x48, 0x80, 0x41, 0x88, 0x93, 0x8b, 0xd5, 0xdd, 0x31, 0xd4, 0xdd, 0x55, 0x80, 0xd1, 0xc9,
	0x7a, 0x57, 0xc3, 0x89, 0x8b, 0x6c, 0x4c, 0x02, 0x4c, 0x5c, 0x8a, 0x99, 0xf9, 0x10, 0x37, 0x14,
	0x14, 0xe5, 0x57, 0x54, 0x62, 0x73, 0x8e, 0x13, 0x0f, 0xc4, 0x25, 0xc5, 0x01, 0x20, 0x1f, 0x05,
	0x30, 0x26, 0xb1, 0x81, 0xbd, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x18, 0xe3, 0xdd, 0x43,
	0x33, 0x01, 0x00, 0x00,
}
