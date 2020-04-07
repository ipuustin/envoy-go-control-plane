// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/number.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
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

type DoubleMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*DoubleMatcher_Range
	//	*DoubleMatcher_Exact
	MatchPattern         isDoubleMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DoubleMatcher) Reset()         { *m = DoubleMatcher{} }
func (m *DoubleMatcher) String() string { return proto.CompactTextString(m) }
func (*DoubleMatcher) ProtoMessage()    {}
func (*DoubleMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de077d68a31b59d, []int{0}
}

func (m *DoubleMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleMatcher.Unmarshal(m, b)
}
func (m *DoubleMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleMatcher.Marshal(b, m, deterministic)
}
func (m *DoubleMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleMatcher.Merge(m, src)
}
func (m *DoubleMatcher) XXX_Size() int {
	return xxx_messageInfo_DoubleMatcher.Size(m)
}
func (m *DoubleMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleMatcher proto.InternalMessageInfo

type isDoubleMatcher_MatchPattern interface {
	isDoubleMatcher_MatchPattern()
}

type DoubleMatcher_Range struct {
	Range *v3.DoubleRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type DoubleMatcher_Exact struct {
	Exact float64 `protobuf:"fixed64,2,opt,name=exact,proto3,oneof"`
}

func (*DoubleMatcher_Range) isDoubleMatcher_MatchPattern() {}

func (*DoubleMatcher_Exact) isDoubleMatcher_MatchPattern() {}

func (m *DoubleMatcher) GetMatchPattern() isDoubleMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *DoubleMatcher) GetRange() *v3.DoubleRange {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Range); ok {
		return x.Range
	}
	return nil
}

func (m *DoubleMatcher) GetExact() float64 {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Exact); ok {
		return x.Exact
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DoubleMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DoubleMatcher_Range)(nil),
		(*DoubleMatcher_Exact)(nil),
	}
}

func init() {
	proto.RegisterType((*DoubleMatcher)(nil), "envoy.type.matcher.v3.DoubleMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/number.proto", fileDescriptor_9de077d68a31b59d) }

var fileDescriptor_9de077d68a31b59d = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f,
	0x33, 0xd6, 0xcf, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x05, 0xab, 0xd1, 0x03, 0xa9, 0xd1, 0x83, 0xaa, 0xd1, 0x2b, 0x33, 0x96, 0x92, 0x44, 0xd2, 0x5a,
	0x66, 0xac, 0x5f, 0x94, 0x98, 0x97, 0x9e, 0x0a, 0xd1, 0x21, 0x25, 0x5b, 0x9a, 0x52, 0x90, 0xa8,
	0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0xac, 0x5f, 0x5c, 0x92, 0x58,
	0x52, 0x5a, 0x0c, 0x95, 0x56, 0xc4, 0x90, 0x2e, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xcb, 0xcc,
	0x4b, 0x87, 0x2a, 0x11, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20,
	0x12, 0x4a, 0xb3, 0x18, 0xb9, 0x78, 0x5d, 0xf2, 0x4b, 0x93, 0x72, 0x52, 0x7d, 0x21, 0x4e, 0x11,
	0x32, 0xe2, 0x62, 0x05, 0xdb, 0x2d, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x87, 0xe4,
	0xdc, 0x32, 0x63, 0x3d, 0x88, 0xe2, 0x20, 0x90, 0x0a, 0x0f, 0x86, 0x20, 0x88, 0x52, 0x21, 0x31,
	0x2e, 0xd6, 0xd4, 0x8a, 0xc4, 0xe4, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x46, 0x90, 0x38, 0x98,
	0x6b, 0xa5, 0x3e, 0xeb, 0x68, 0x87, 0x9c, 0x12, 0x97, 0x02, 0x16, 0x1f, 0xa3, 0x58, 0xea, 0x24,
	0xc2, 0xc5, 0x0b, 0x96, 0x88, 0x2f, 0x48, 0x2c, 0x29, 0x49, 0x2d, 0xca, 0x13, 0x62, 0xfe, 0xe1,
	0xc4, 0xe8, 0x64, 0xb3, 0xab, 0xe1, 0xc4, 0x45, 0x36, 0x26, 0x01, 0x66, 0x2e, 0xe5, 0xcc, 0x7c,
	0x88, 0x3b, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0xb0, 0x86, 0xa0, 0x13, 0xb7, 0x1f, 0x38, 0x98,
	0x03, 0x40, 0x1e, 0x0b, 0x60, 0x4c, 0x62, 0x03, 0xfb, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0x87, 0x18, 0x4f, 0x94, 0x01, 0x00, 0x00,
}
