// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v2/csrf.proto

package envoy_config_filter_http_csrf_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
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

type CsrfPolicy struct {
	FilterEnabled        *core.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	ShadowEnabled        *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	AdditionalOrigins    []*matcher.StringMatcher       `protobuf:"bytes,3,rep,name=additional_origins,json=additionalOrigins,proto3" json:"additional_origins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CsrfPolicy) Reset()         { *m = CsrfPolicy{} }
func (m *CsrfPolicy) String() string { return proto.CompactTextString(m) }
func (*CsrfPolicy) ProtoMessage()    {}
func (*CsrfPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9146cdf92353980, []int{0}
}

func (m *CsrfPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsrfPolicy.Unmarshal(m, b)
}
func (m *CsrfPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsrfPolicy.Marshal(b, m, deterministic)
}
func (m *CsrfPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfPolicy.Merge(m, src)
}
func (m *CsrfPolicy) XXX_Size() int {
	return xxx_messageInfo_CsrfPolicy.Size(m)
}
func (m *CsrfPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfPolicy proto.InternalMessageInfo

func (m *CsrfPolicy) GetFilterEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetShadowEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.ShadowEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetAdditionalOrigins() []*matcher.StringMatcher {
	if m != nil {
		return m.AdditionalOrigins
	}
	return nil
}

func init() {
	proto.RegisterType((*CsrfPolicy)(nil), "envoy.config.filter.http.csrf.v2.CsrfPolicy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/csrf/v2/csrf.proto", fileDescriptor_a9146cdf92353980)
}

var fileDescriptor_a9146cdf92353980 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x91, 0x8c, 0x4d, 0x2b, 0x63, 0xe3, 0x6a, 0x53, 0x63, 0xfa, 0xe3, 0x16, 0x4a, 0x0d,
	0x86, 0x19, 0x90, 0xdf, 0x40, 0xa5, 0xdd, 0x95, 0x0a, 0x75, 0x5b, 0x30, 0xd7, 0xd2, 0x48, 0x1e,
	0x90, 0x67, 0xc4, 0xcc, 0xb5, 0x6a, 0xed, 0xb2, 0xce, 0x26, 0x59, 0xe6, 0x59, 0xf2, 0x04, 0xd9,
	0xe6, 0x55, 0xb2, 0xcc, 0x22, 0x04, 0xe9, 0xca, 0x49, 0x20, 0x8b, 0x90, 0x95, 0x24, 0xee, 0x39,
	0xdf, 0x5c, 0x9d, 0x33, 0xde, 0x52, 0xa8, 0x4a, 0xd7, 0x3c, 0xd1, 0x2a, 0x93, 0x39, 0xcf, 0x64,
	0x81, 0xc2, 0xf0, 0x2d, 0x62, 0xc9, 0x13, 0x6b, 0x32, 0x5e, 0x05, 0xed, 0x93, 0x95, 0x46, 0xa3,
	0xf6, 0xe7, 0xad, 0x98, 0x91, 0x98, 0x91, 0x98, 0x35, 0x62, 0xd6, 0x8a, 0xaa, 0x60, 0xf6, 0x81,
	0x70, 0x50, 0xca, 0xd6, 0xaa, 0x8d, 0xe0, 0x1b, 0xb0, 0x82, 0xfc, 0xb3, 0xcf, 0x34, 0xc5, 0xba,
	0x14, 0x7c, 0x07, 0x98, 0x6c, 0x85, 0xe1, 0x16, 0x8d, 0x54, 0x79, 0x27, 0xf8, 0xb4, 0x4f, 0x4b,
	0xe0, 0xa0, 0x94, 0x46, 0x40, 0xa9, 0x95, 0xe5, 0x3b, 0x99, 0x1b, 0xc0, 0x23, 0xe0, 0xe3, 0xb3,
	0xb9, 0x45, 0xc0, 0xbd, 0xed, 0xc6, 0xef, 0x2b, 0x28, 0x64, 0x0a, 0x28, 0xf8, 0xf1, 0x85, 0x06,
	0x5f, 0xcf, 0x5d, 0xcf, 0xfb, 0x61, 0x4d, 0x16, 0xe9, 0x42, 0x26, 0xb5, 0xff, 0xcf, 0x1b, 0xd3,
	0xf2, 0x6b, 0xa1, 0x60, 0x53, 0x88, 0x74, 0xea, 0xcc, 0x9d, 0xc5, 0x30, 0x58, 0x32, 0xfa, 0x41,
	0x28, 0x25, 0xab, 0x02, 0xd6, 0xac, 0xcf, 0xe2, 0xbd, 0x42, 0xb9, 0x13, 0xbf, 0x0c, 0x24, 0xcd,
	0x89, 0x50, 0x44, 0xc2, 0x24, 0x42, 0x61, 0xf8, 0xe6, 0x36, 0xec, 0x9f, 0x3a, 0xee, 0xc4, 0x89,
	0x47, 0x04, 0xfb, 0x49, 0x2c, 0x3f, 0xf6, 0xc6, 0x76, 0x0b, 0xa9, 0xfe, 0xff, 0x40, 0x77, 0x5f,
	0x4d, 0x8f, 0x47, 0x84, 0x38, 0x32, 0x23, 0xcf, 0x87, 0x34, 0x95, 0xa4, 0x59, 0x6b, 0x23, 0x73,
	0xa9, 0xec, 0xb4, 0x37, 0xef, 0x2d, 0x86, 0xc1, 0x97, 0x8e, 0xdb, 0xc4, 0xca, 0xba, 0x58, 0xd9,
	0xdf, 0x36, 0xd6, 0xdf, 0xf4, 0x15, 0xbf, 0x7b, 0x34, 0xff, 0x21, 0x6f, 0x68, 0x6e, 0x2e, 0xee,
	0xce, 0xfa, 0xdf, 0xfd, 0x6f, 0x64, 0x16, 0x07, 0x14, 0xca, 0x36, 0x91, 0x76, 0xbd, 0xda, 0xa7,
	0xc5, 0xae, 0x2e, 0x4f, 0xae, 0xae, 0x07, 0xee, 0xc4, 0xf5, 0x98, 0xd4, 0x74, 0x5c, 0x69, 0xf4,
	0xa1, 0x66, 0x2f, 0x5d, 0x88, 0xf0, 0x6d, 0x1b, 0x7b, 0x53, 0x42, 0xe4, 0x6c, 0x06, 0x6d, 0x1b,
	0xab, 0xfb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x08, 0x28, 0xa3, 0x75, 0x02, 0x00, 0x00,
}
