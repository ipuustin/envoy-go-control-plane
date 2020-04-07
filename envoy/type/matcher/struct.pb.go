// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/struct.proto

package envoy_type_matcher

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

type StructMatcher struct {
	Path                 []*StructMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Value                *ValueMatcher                `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StructMatcher) Reset()         { *m = StructMatcher{} }
func (m *StructMatcher) String() string { return proto.CompactTextString(m) }
func (*StructMatcher) ProtoMessage()    {}
func (*StructMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_80addaed29127401, []int{0}
}

func (m *StructMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructMatcher.Unmarshal(m, b)
}
func (m *StructMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructMatcher.Marshal(b, m, deterministic)
}
func (m *StructMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructMatcher.Merge(m, src)
}
func (m *StructMatcher) XXX_Size() int {
	return xxx_messageInfo_StructMatcher.Size(m)
}
func (m *StructMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StructMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StructMatcher proto.InternalMessageInfo

func (m *StructMatcher) GetPath() []*StructMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *StructMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

type StructMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*StructMatcher_PathSegment_Key
	Segment              isStructMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *StructMatcher_PathSegment) Reset()         { *m = StructMatcher_PathSegment{} }
func (m *StructMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*StructMatcher_PathSegment) ProtoMessage()    {}
func (*StructMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_80addaed29127401, []int{0, 0}
}

func (m *StructMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructMatcher_PathSegment.Unmarshal(m, b)
}
func (m *StructMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructMatcher_PathSegment.Marshal(b, m, deterministic)
}
func (m *StructMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructMatcher_PathSegment.Merge(m, src)
}
func (m *StructMatcher_PathSegment) XXX_Size() int {
	return xxx_messageInfo_StructMatcher_PathSegment.Size(m)
}
func (m *StructMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_StructMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_StructMatcher_PathSegment proto.InternalMessageInfo

type isStructMatcher_PathSegment_Segment interface {
	isStructMatcher_PathSegment_Segment()
}

type StructMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*StructMatcher_PathSegment_Key) isStructMatcher_PathSegment_Segment() {}

func (m *StructMatcher_PathSegment) GetSegment() isStructMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *StructMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*StructMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StructMatcher_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StructMatcher_PathSegment_Key)(nil),
	}
}

func init() {
	proto.RegisterType((*StructMatcher)(nil), "envoy.type.matcher.StructMatcher")
	proto.RegisterType((*StructMatcher_PathSegment)(nil), "envoy.type.matcher.StructMatcher.PathSegment")
}

func init() { proto.RegisterFile("envoy/type/matcher/struct.proto", fileDescriptor_80addaed29127401) }

var fileDescriptor_80addaed29127401 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0x9d, 0xcd, 0xfd, 0x73, 0x83, 0x72, 0xa4, 0xf1, 0x38, 0x51, 0x17, 0xab, 0x6b, 0xdc,
	0x85, 0xb3, 0xd2, 0x4a, 0xb6, 0x12, 0x44, 0x08, 0x39, 0xb0, 0x5f, 0xef, 0x16, 0x13, 0xbc, 0xcb,
	0x86, 0x64, 0x12, 0x4c, 0x67, 0x6d, 0xe9, 0xc7, 0xf1, 0x13, 0xd8, 0xda, 0xfa, 0x49, 0xe4, 0x2a,
	0xc9, 0x6e, 0x04, 0xc5, 0x74, 0xcb, 0xce, 0x7b, 0xbf, 0xf7, 0x66, 0xe8, 0x89, 0x4e, 0x2b, 0x53,
	0x0b, 0xac, 0x33, 0x2d, 0x36, 0x0a, 0x97, 0xb1, 0xce, 0x45, 0x81, 0x79, 0xb9, 0x44, 0x9e, 0xe5,
	0x06, 0x4d, 0x10, 0x58, 0x01, 0x6f, 0x04, 0xbc, 0x15, 0x4c, 0x8f, 0x3b, 0x4c, 0x95, 0x5a, 0x97,
	0xda, 0x79, 0xa6, 0x47, 0xe5, 0x2a, 0x53, 0x42, 0xa5, 0xa9, 0x41, 0x85, 0x89, 0x49, 0x0b, 0x51,
	0xa0, 0xc2, 0xb2, 0x68, 0xc7, 0x07, 0x95, 0x5a, 0x27, 0x2b, 0x85, 0x5a, 0xfc, 0x3c, 0xdc, 0xe0,
	0xf4, 0x13, 0xe8, 0xde, 0xc2, 0x86, 0xdf, 0x3a, 0x6a, 0x70, 0x43, 0x7b, 0x99, 0xc2, 0x78, 0x42,
	0x98, 0x37, 0xf3, 0xe7, 0x67, 0xfc, 0x7f, 0x19, 0xfe, 0xc7, 0xc0, 0x43, 0x85, 0xf1, 0x42, 0x3f,
	0x6c, 0x74, 0x8a, 0x72, 0xb4, 0x95, 0xfd, 0x57, 0x20, 0x23, 0x88, 0x2c, 0x24, 0xb8, 0xa2, 0x7d,
	0xdb, 0x72, 0xe2, 0x31, 0x98, 0xf9, 0x73, 0xd6, 0x45, 0xbb, 0x6b, 0x04, 0x2d, 0xcc, 0x02, 0x5e,
	0x80, 0x8c, 0x21, 0x72, 0xc6, 0xe9, 0x25, 0xf5, 0x7f, 0x05, 0x04, 0x87, 0xd4, 0x7b, 0xd4, 0xf5,
	0x04, 0x18, 0xcc, 0x76, 0xe5, 0x70, 0x2b, 0x7b, 0x39, 0x61, 0x70, 0xbd, 0x13, 0x35, 0xbf, 0x72,
	0x9f, 0x0e, 0x8b, 0x56, 0xe7, 0x7d, 0x49, 0x90, 0x17, 0x6f, 0xcf, 0xef, 0x1f, 0x03, 0x32, 0x26,
	0x94, 0x25, 0xc6, 0x45, 0x67, 0xb9, 0x79, 0xaa, 0x3b, 0x5a, 0x48, 0xdf, 0x2d, 0x15, 0x36, 0x57,
	0x09, 0xe1, 0x7e, 0x60, 0xcf, 0x73, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x63, 0xcc, 0xe7,
	0xad, 0x01, 0x00, 0x00,
}
