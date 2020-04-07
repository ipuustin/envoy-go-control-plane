// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/metadata.proto

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

type MetadataMatcher struct {
	Filter               string                         `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Path                 []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Value                *ValueMatcher                  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MetadataMatcher) Reset()         { *m = MetadataMatcher{} }
func (m *MetadataMatcher) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher) ProtoMessage()    {}
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_865eaf6a1e9e266d, []int{0}
}

func (m *MetadataMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataMatcher.Unmarshal(m, b)
}
func (m *MetadataMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher.Merge(m, src)
}
func (m *MetadataMatcher) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher.Size(m)
}
func (m *MetadataMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher proto.InternalMessageInfo

func (m *MetadataMatcher) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *MetadataMatcher) GetPath() []*MetadataMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *MetadataMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

type MetadataMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*MetadataMatcher_PathSegment_Key
	Segment              isMetadataMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MetadataMatcher_PathSegment) Reset()         { *m = MetadataMatcher_PathSegment{} }
func (m *MetadataMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher_PathSegment) ProtoMessage()    {}
func (*MetadataMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_865eaf6a1e9e266d, []int{0, 0}
}

func (m *MetadataMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Unmarshal(m, b)
}
func (m *MetadataMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher_PathSegment.Merge(m, src)
}
func (m *MetadataMatcher_PathSegment) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Size(m)
}
func (m *MetadataMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher_PathSegment proto.InternalMessageInfo

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*MetadataMatcher_PathSegment_Key) isMetadataMatcher_PathSegment_Segment() {}

func (m *MetadataMatcher_PathSegment) GetSegment() isMetadataMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MetadataMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*MetadataMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetadataMatcher_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
}

func init() {
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.MetadataMatcher.PathSegment")
}

func init() { proto.RegisterFile("envoy/type/matcher/metadata.proto", fileDescriptor_865eaf6a1e9e266d) }

var fileDescriptor_865eaf6a1e9e266d = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xb1, 0xd3, 0x7f, 0xb8, 0x02, 0x2a, 0x2f, 0x54, 0x45, 0x80, 0x61, 0xea, 0x64, 0x4b,
	0x65, 0x83, 0x05, 0x79, 0x62, 0xa9, 0x54, 0x05, 0x89, 0xfd, 0xa0, 0x86, 0x46, 0x24, 0x71, 0x94,
	0x5c, 0x22, 0xb2, 0x31, 0x33, 0xf2, 0x71, 0xf8, 0x04, 0xac, 0x7c, 0x1b, 0xd4, 0x05, 0x94, 0xc4,
	0x95, 0x50, 0xc9, 0x66, 0xf9, 0xde, 0xfb, 0xbd, 0x77, 0xc7, 0xce, 0x4c, 0x5c, 0xd8, 0x52, 0x61,
	0x99, 0x18, 0x15, 0x01, 0x3e, 0xac, 0x4c, 0xaa, 0x22, 0x83, 0xb0, 0x04, 0x04, 0x99, 0xa4, 0x16,
	0x2d, 0xe7, 0xb5, 0x44, 0x56, 0x12, 0xe9, 0x24, 0x93, 0x93, 0x16, 0x5b, 0x01, 0x61, 0x6e, 0x1a,
	0xcf, 0xe4, 0x38, 0x5f, 0x26, 0xa0, 0x20, 0x8e, 0x2d, 0x02, 0x06, 0x36, 0xce, 0x54, 0x86, 0x80,
	0x79, 0xe6, 0xc6, 0x87, 0x05, 0x84, 0xc1, 0x12, 0xd0, 0xa8, 0xcd, 0xa3, 0x19, 0x9c, 0xff, 0x10,
	0x76, 0x30, 0x77, 0xf1, 0xf3, 0x86, 0xcb, 0x4f, 0x59, 0xef, 0x31, 0x08, 0xd1, 0xa4, 0x63, 0x22,
	0xc8, 0x74, 0x57, 0xf7, 0xd7, 0xba, 0x93, 0x52, 0x41, 0x7c, 0xf7, 0xcd, 0xe7, 0xac, 0x93, 0x00,
	0xae, 0xc6, 0x54, 0x78, 0xd3, 0xe1, 0x4c, 0xc9, 0xff, 0x7d, 0xe5, 0x16, 0x53, 0x2e, 0x00, 0x57,
	0xb7, 0xe6, 0x29, 0x32, 0x31, 0xea, 0xc1, 0x5a, 0x77, 0xdf, 0x09, 0x1d, 0x10, 0xbf, 0xc6, 0xf0,
	0x6b, 0xd6, 0xad, 0x57, 0x19, 0x7b, 0x82, 0x4c, 0x87, 0x33, 0xd1, 0xc6, 0xbb, 0xab, 0x04, 0x0e,
	0x56, 0x03, 0xde, 0x08, 0x1d, 0x11, 0xbf, 0x31, 0x4e, 0x2e, 0xd9, 0xf0, 0x4f, 0x00, 0x3f, 0x62,
	0xde, 0xb3, 0x29, 0xb7, 0xda, 0xdf, 0xec, 0xf8, 0xd5, 0xaf, 0xde, 0x67, 0xfd, 0xcc, 0xe9, 0xbc,
	0x6f, 0x4d, 0xf4, 0xd5, 0xc7, 0xeb, 0xe7, 0x57, 0x8f, 0x8e, 0x28, 0x13, 0x81, 0x6d, 0xa2, 0x93,
	0xd4, 0xbe, 0x94, 0x2d, 0x2d, 0xf4, 0xde, 0x66, 0xad, 0x45, 0x75, 0xbc, 0x05, 0xb9, 0xef, 0xd5,
	0x57, 0xbc, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xe3, 0x5f, 0x0c, 0xd6, 0x01, 0x00, 0x00,
}
