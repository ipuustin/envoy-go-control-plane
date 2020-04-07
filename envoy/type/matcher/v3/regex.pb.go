// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/regex.proto

package envoy_type_matcher_v3

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

type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType           isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	Regex                string                    `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{0}
}

func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher.Unmarshal(m, b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher.Size(m)
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

type RegexMatcher_GoogleRE2 struct {
	MaxProgramSize       *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_program_size,json=maxProgramSize,proto3" json:"max_program_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{0, 0}
}

func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Unmarshal(m, b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Size(m)
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

func (m *RegexMatcher_GoogleRE2) GetMaxProgramSize() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxProgramSize
	}
	return nil
}

type RegexMatchAndSubstitute struct {
	Pattern              *RegexMatcher `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Substitution         string        `protobuf:"bytes,2,opt,name=substitution,proto3" json:"substitution,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RegexMatchAndSubstitute) Reset()         { *m = RegexMatchAndSubstitute{} }
func (m *RegexMatchAndSubstitute) String() string { return proto.CompactTextString(m) }
func (*RegexMatchAndSubstitute) ProtoMessage()    {}
func (*RegexMatchAndSubstitute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7de4f215ebc85482, []int{1}
}

func (m *RegexMatchAndSubstitute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatchAndSubstitute.Unmarshal(m, b)
}
func (m *RegexMatchAndSubstitute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatchAndSubstitute.Marshal(b, m, deterministic)
}
func (m *RegexMatchAndSubstitute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatchAndSubstitute.Merge(m, src)
}
func (m *RegexMatchAndSubstitute) XXX_Size() int {
	return xxx_messageInfo_RegexMatchAndSubstitute.Size(m)
}
func (m *RegexMatchAndSubstitute) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatchAndSubstitute.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatchAndSubstitute proto.InternalMessageInfo

func (m *RegexMatchAndSubstitute) GetPattern() *RegexMatcher {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *RegexMatchAndSubstitute) GetSubstitution() string {
	if m != nil {
		return m.Substitution
	}
	return ""
}

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "envoy.type.matcher.v3.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "envoy.type.matcher.v3.RegexMatcher.GoogleRE2")
	proto.RegisterType((*RegexMatchAndSubstitute)(nil), "envoy.type.matcher.v3.RegexMatchAndSubstitute")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/regex.proto", fileDescriptor_7de4f215ebc85482) }

var fileDescriptor_7de4f215ebc85482 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0x9d, 0xad, 0x6d, 0xcd, 0xdb, 0x22, 0x61, 0x41, 0x5a, 0x82, 0xad, 0x49, 0x0a, 0x52,
	0xff, 0xcd, 0xe8, 0xe6, 0x16, 0xf1, 0xe0, 0x80, 0xff, 0x0e, 0x42, 0x98, 0xa2, 0x78, 0x0b, 0x13,
	0xf3, 0xba, 0x0e, 0x24, 0x33, 0xc3, 0xec, 0xec, 0x9a, 0xf4, 0xe4, 0x45, 0x10, 0x3f, 0x82, 0x5f,
	0xc2, 0xbb, 0x77, 0x41, 0xf0, 0xe4, 0xb7, 0x91, 0x9e, 0x64, 0x77, 0xb2, 0x6a, 0x71, 0xd1, 0xde,
	0x86, 0xf7, 0x7d, 0x9e, 0x87, 0xdf, 0xfb, 0x0c, 0xf4, 0x50, 0x17, 0x66, 0xc9, 0xfc, 0xd2, 0x22,
	0x9b, 0x4b, 0xff, 0xf2, 0x35, 0x3a, 0x56, 0x0c, 0x98, 0xc3, 0x14, 0x17, 0xd4, 0x3a, 0xe3, 0x4d,
	0x7c, 0xa9, 0x92, 0xd0, 0x52, 0x42, 0x57, 0x12, 0x5a, 0x0c, 0x3a, 0xfb, 0xa9, 0x31, 0xe9, 0x0c,
	0x59, 0x25, 0x9a, 0xe4, 0xaf, 0xd8, 0x1b, 0x27, 0xad, 0x45, 0x97, 0x05, 0x5b, 0x67, 0x2f, 0x9f,
	0x5a, 0xc9, 0xa4, 0xd6, 0xc6, 0x4b, 0xaf, 0x8c, 0xce, 0x58, 0xe6, 0xa5, 0xcf, 0xeb, 0x75, 0xef,
	0xaf, 0x75, 0x81, 0x2e, 0x53, 0x46, 0x2b, 0x9d, 0xae, 0x24, 0x3b, 0x85, 0x9c, 0xa9, 0xa9, 0xf4,
	0xc8, 0xea, 0x47, 0x58, 0xf4, 0xbf, 0x45, 0xb0, 0x2d, 0x4a, 0xc2, 0xa7, 0x01, 0x27, 0x7e, 0x01,
	0x10, 0x68, 0xc6, 0x0e, 0x93, 0x5d, 0xd2, 0x25, 0x87, 0x5b, 0xc9, 0x2d, 0xda, 0xc8, 0x4d, 0xff,
	0x34, 0xd2, 0x47, 0x95, 0x4b, 0x3c, 0x48, 0xf8, 0x85, 0x13, 0xbe, 0xfe, 0x81, 0x44, 0x6d, 0xf2,
	0xf8, 0x9c, 0x68, 0x85, 0x30, 0x81, 0x49, 0xbc, 0x07, 0xeb, 0x55, 0x17, 0xbb, 0x51, 0x97, 0x1c,
	0xb6, 0xf8, 0xe6, 0x09, 0x3f, 0xef, 0xa2, 0x2e, 0x11, 0x61, 0xda, 0x79, 0x47, 0xa0, 0xf5, 0x2b,
	0x23, 0x7e, 0x08, 0xed, 0xb9, 0x5c, 0x8c, 0xad, 0x33, 0xa9, 0x93, 0xf3, 0x71, 0xa6, 0x8e, 0x71,
	0x05, 0x73, 0x99, 0x86, 0x48, 0x5a, 0xb7, 0x45, 0x9f, 0x3d, 0xd1, 0x7e, 0x90, 0x3c, 0x97, 0xb3,
	0x1c, 0xc5, 0xc5, 0xb9, 0x5c, 0x8c, 0x82, 0xe9, 0x48, 0x1d, 0xe3, 0xf0, 0xf6, 0xc7, 0x2f, 0xef,
	0xf7, 0x6f, 0xc0, 0xb5, 0x86, 0x03, 0x9a, 0xe9, 0x87, 0x57, 0x4b, 0x47, 0x0f, 0xae, 0xfc, 0xc7,
	0xc1, 0x63, 0xd8, 0x42, 0x9d, 0x2a, 0x8d, 0xe3, 0x52, 0x13, 0xaf, 0xfd, 0xe0, 0xa4, 0xff, 0x89,
	0xc0, 0xce, 0x6f, 0xd1, 0x7d, 0x3d, 0x3d, 0xca, 0x27, 0x99, 0x57, 0x3e, 0xf7, 0x18, 0xdf, 0x83,
	0x4d, 0x2b, 0xbd, 0x47, 0xa7, 0x57, 0x87, 0x1c, 0x9c, 0xa1, 0x55, 0x51, 0x7b, 0xe2, 0x3e, 0x6c,
	0x67, 0x75, 0x98, 0x32, 0x3a, 0x94, 0x28, 0x4e, 0xcd, 0x86, 0x77, 0x4a, 0xf4, 0x9b, 0x70, 0xfd,
	0x9f, 0xe8, 0xa7, 0xa8, 0xf8, 0xdd, 0xcf, 0x6f, 0xbf, 0x7e, 0xdf, 0x88, 0xda, 0x6b, 0x70, 0xa0,
	0x4c, 0x00, 0xb2, 0xce, 0x2c, 0x96, 0xcd, 0x6c, 0x1c, 0xaa, 0x9c, 0x51, 0xd9, 0xfc, 0x88, 0x4c,
	0x36, 0xaa, 0x2f, 0x18, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xd2, 0xec, 0xda, 0xfa, 0x02,
	0x00, 0x00,
}
