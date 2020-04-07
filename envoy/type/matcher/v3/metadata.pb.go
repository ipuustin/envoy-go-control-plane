// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/metadata.proto

package envoy_type_matcher_v3

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
	return fileDescriptor_ea39646b3de596af, []int{0}
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
	return fileDescriptor_ea39646b3de596af, []int{0, 0}
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
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.v3.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.v3.MetadataMatcher.PathSegment")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3/metadata.proto", fileDescriptor_ea39646b3de596af)
}

var fileDescriptor_ea39646b3de596af = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0xbd, 0xa4, 0xff, 0xbc, 0xa2, 0x96, 0x80, 0x58, 0x2a, 0x6a, 0xfa, 0x67, 0xa8, 0xcb,
	0x9d, 0xb4, 0xb8, 0x74, 0x70, 0x38, 0x17, 0x97, 0x42, 0x88, 0xe0, 0x7e, 0xda, 0xb3, 0x0d, 0xb6,
	0x77, 0x31, 0x79, 0x1b, 0xcc, 0xe6, 0x28, 0x8e, 0x8e, 0x7e, 0x14, 0x77, 0xc1, 0xd5, 0x4f, 0xe1,
	0x57, 0x90, 0x4e, 0x92, 0xe4, 0x0a, 0x52, 0xa3, 0x6e, 0xc7, 0xbd, 0xcf, 0xf3, 0x7b, 0xde, 0x87,
	0x17, 0x77, 0x84, 0x8c, 0x54, 0x4c, 0x21, 0xf6, 0x05, 0x9d, 0x71, 0xb8, 0x9a, 0x88, 0x80, 0x46,
	0x7d, 0x3a, 0x13, 0xc0, 0x47, 0x1c, 0x38, 0xf1, 0x03, 0x05, 0xca, 0xda, 0x4e, 0x55, 0x24, 0x51,
	0x11, 0xad, 0x22, 0x51, 0xbf, 0xd1, 0xcc, 0x37, 0x47, 0x7c, 0x3a, 0x17, 0x99, 0xb3, 0xb1, 0x37,
	0x1f, 0xf9, 0x9c, 0x72, 0x29, 0x15, 0x70, 0xf0, 0x94, 0x0c, 0x69, 0x08, 0x1c, 0xe6, 0xa1, 0x1e,
	0x37, 0x7f, 0x8c, 0x23, 0x11, 0x84, 0x9e, 0x92, 0x9e, 0x1c, 0x6b, 0xc9, 0x4e, 0xc4, 0xa7, 0xde,
	0x88, 0x83, 0xa0, 0xcb, 0x47, 0x36, 0x68, 0x7d, 0x18, 0x78, 0x6b, 0xa8, 0xf7, 0x1c, 0x66, 0xe9,
	0xd6, 0x01, 0x2e, 0x5d, 0x7b, 0x53, 0x10, 0x41, 0x1d, 0xd9, 0xa8, 0xbb, 0xce, 0xca, 0x0b, 0x56,
	0x08, 0x0c, 0x1b, 0xb9, 0xfa, 0xdb, 0x72, 0x70, 0xc1, 0xe7, 0x30, 0xa9, 0x1b, 0xb6, 0xd9, 0xad,
	0xf6, 0x7a, 0x24, 0xb7, 0x18, 0x59, 0xc1, 0x12, 0x87, 0xc3, 0xe4, 0x5c, 0x8c, 0x67, 0x42, 0x02,
	0xab, 0x2c, 0x58, 0xf1, 0x09, 0x19, 0x15, 0xe4, 0xa6, 0x24, 0xeb, 0x14, 0x17, 0xd3, 0xc2, 0x75,
	0xd3, 0x46, 0xdd, 0x6a, 0xaf, 0xfd, 0x0b, 0xf2, 0x22, 0xd1, 0x68, 0x5e, 0xca, 0x78, 0x44, 0x46,
	0x0d, 0xb9, 0x99, 0xb7, 0x71, 0x8b, 0xab, 0xdf, 0x32, 0xac, 0x5d, 0x6c, 0xde, 0x88, 0x78, 0xa5,
	0xc3, 0xd9, 0x9a, 0x9b, 0xfc, 0x0e, 0x8e, 0x9f, 0x5f, 0x1f, 0xf6, 0x8f, 0x70, 0x5e, 0xce, 0x5f,
	0x7b, 0x6f, 0xe2, 0x72, 0xa8, 0xf1, 0xe6, 0x27, 0x43, 0x83, 0xc3, 0x04, 0xd3, 0xc1, 0xad, 0xff,
	0x31, 0xec, 0xe4, 0xe5, 0xfe, 0xed, 0xbd, 0x64, 0xd4, 0x4c, 0xdc, 0xf6, 0x54, 0x96, 0xeb, 0x07,
	0xea, 0x2e, 0xce, 0xaf, 0xca, 0x36, 0x96, 0x7e, 0x27, 0xb9, 0x93, 0x83, 0x2e, 0x4b, 0xe9, 0xc1,
	0xfa, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x74, 0x69, 0x96, 0x6d, 0x02, 0x00, 0x00,
}
