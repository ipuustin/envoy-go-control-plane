// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/v3/token_bucket.proto

package envoy_type_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type TokenBucket struct {
	MaxTokens            uint32                `protobuf:"varint,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	TokensPerFill        *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=tokens_per_fill,json=tokensPerFill,proto3" json:"tokens_per_fill,omitempty"`
	FillInterval         *duration.Duration    `protobuf:"bytes,3,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TokenBucket) Reset()         { *m = TokenBucket{} }
func (m *TokenBucket) String() string { return proto.CompactTextString(m) }
func (*TokenBucket) ProtoMessage()    {}
func (*TokenBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b88f84e3f7d4627, []int{0}
}

func (m *TokenBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenBucket.Unmarshal(m, b)
}
func (m *TokenBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenBucket.Marshal(b, m, deterministic)
}
func (m *TokenBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenBucket.Merge(m, src)
}
func (m *TokenBucket) XXX_Size() int {
	return xxx_messageInfo_TokenBucket.Size(m)
}
func (m *TokenBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenBucket.DiscardUnknown(m)
}

var xxx_messageInfo_TokenBucket proto.InternalMessageInfo

func (m *TokenBucket) GetMaxTokens() uint32 {
	if m != nil {
		return m.MaxTokens
	}
	return 0
}

func (m *TokenBucket) GetTokensPerFill() *wrappers.UInt32Value {
	if m != nil {
		return m.TokensPerFill
	}
	return nil
}

func (m *TokenBucket) GetFillInterval() *duration.Duration {
	if m != nil {
		return m.FillInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenBucket)(nil), "envoy.type.v3.TokenBucket")
}

func init() { proto.RegisterFile("envoy/type/v3/token_bucket.proto", fileDescriptor_5b88f84e3f7d4627) }

var fileDescriptor_5b88f84e3f7d4627 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x3f, 0x4b, 0x03, 0x31,
	0x18, 0xc6, 0x7b, 0x6d, 0x69, 0x35, 0xf5, 0xb0, 0xdc, 0xa0, 0xe7, 0x9f, 0x96, 0xd3, 0x41, 0x4a,
	0x87, 0x04, 0x7a, 0x93, 0x8e, 0x41, 0x84, 0x0a, 0x42, 0x29, 0xea, 0x7a, 0xa4, 0x36, 0x2d, 0xa1,
	0x69, 0x12, 0x72, 0xb9, 0xb3, 0xdd, 0x1c, 0xfd, 0x0c, 0x7e, 0x04, 0x3f, 0x82, 0xbb, 0xe0, 0xea,
	0x27, 0x71, 0xef, 0x24, 0x97, 0xbb, 0x62, 0xa5, 0x5b, 0xe0, 0xf9, 0x3d, 0x4f, 0xde, 0xe7, 0x7d,
	0x41, 0x40, 0x45, 0x2a, 0x97, 0xc8, 0x2c, 0x15, 0x45, 0x69, 0x88, 0x8c, 0x9c, 0x51, 0x11, 0x8d,
	0x92, 0xa7, 0x19, 0x35, 0x50, 0x69, 0x69, 0xa4, 0xe7, 0x5a, 0x02, 0x66, 0x04, 0x4c, 0xc3, 0xe3,
	0xf6, 0x54, 0xca, 0x29, 0xa7, 0xc8, 0x8a, 0xa3, 0x64, 0x82, 0xc6, 0x89, 0x26, 0x86, 0x49, 0x91,
	0xe3, 0xdb, 0xfa, 0xb3, 0x26, 0x4a, 0x51, 0x1d, 0x17, 0x7a, 0x2b, 0x19, 0x2b, 0x82, 0x88, 0x10,
	0xd2, 0x58, 0x5b, 0x8c, 0x62, 0x43, 0x4c, 0xb2, 0x96, 0xcf, 0xb6, 0xe4, 0x94, 0xea, 0x98, 0x49,
	0xc1, 0xc4, 0xb4, 0x40, 0x0e, 0x53, 0xc2, 0xd9, 0x98, 0x18, 0x8a, 0xd6, 0x8f, 0x5c, 0x38, 0xff,
	0x71, 0x40, 0xe3, 0x3e, 0x2b, 0x80, 0xed, 0xfc, 0xde, 0x05, 0x00, 0x73, 0xb2, 0x88, 0x6c, 0xa7,
	0xd8, 0x77, 0x02, 0xa7, 0xe3, 0xe2, 0xfa, 0x0a, 0x57, 0xbb, 0xe5, 0xa0, 0x34, 0xdc, 0x9d, 0x93,
	0x85, 0x85, 0x63, 0xef, 0x0e, 0xec, 0xe7, 0x4c, 0xa4, 0xa8, 0x8e, 0x26, 0x8c, 0x73, 0xbf, 0x1c,
	0x38, 0x9d, 0x46, 0xef, 0x14, 0xe6, 0x65, 0xe0, 0xba, 0x0c, 0x7c, 0xe8, 0x0b, 0x13, 0xf6, 0x1e,
	0x09, 0x4f, 0xe8, 0x5f, 0x94, 0x9b, 0xbb, 0x07, 0x54, 0xdf, 0x30, 0xce, 0xbd, 0x5b, 0xe0, 0x66,
	0x19, 0x11, 0x13, 0x86, 0xea, 0x94, 0x70, 0xbf, 0x62, 0xc3, 0x8e, 0xb6, 0xc2, 0xae, 0x8b, 0xcd,
	0x61, 0xb0, 0xc2, 0xf5, 0x77, 0xa7, 0xba, 0xe3, 0x74, 0x4b, 0xc3, 0xbd, 0xcc, 0xdb, 0x2f, 0xac,
	0x57, 0xad, 0xb7, 0xcf, 0xd7, 0xb6, 0x0f, 0x0e, 0x36, 0x6e, 0xb0, 0xd1, 0x10, 0x5f, 0x7e, 0xbc,
	0x7c, 0x7d, 0xd7, 0xca, 0xcd, 0x0a, 0x38, 0x61, 0x12, 0x5a, 0x48, 0x69, 0xb9, 0x58, 0xc2, 0x7f,
	0x37, 0xc3, 0xcd, 0x0d, 0xcf, 0x20, 0xfb, 0x7d, 0xe0, 0x8c, 0x6a, 0x76, 0x8c, 0xf0, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x72, 0x23, 0x6f, 0x82, 0x01, 0x02, 0x00, 0x00,
}
