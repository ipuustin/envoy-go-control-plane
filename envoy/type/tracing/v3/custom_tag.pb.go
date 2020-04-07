// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/tracing/v3/custom_tag.proto

package envoy_type_tracing_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/metadata/v3"
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

type CustomTag struct {
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*CustomTag_Literal_
	//	*CustomTag_Environment_
	//	*CustomTag_RequestHeader
	//	*CustomTag_Metadata_
	Type                 isCustomTag_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomTag) Reset()         { *m = CustomTag{} }
func (m *CustomTag) String() string { return proto.CompactTextString(m) }
func (*CustomTag) ProtoMessage()    {}
func (*CustomTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb41267f4395b6, []int{0}
}

func (m *CustomTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag.Unmarshal(m, b)
}
func (m *CustomTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag.Marshal(b, m, deterministic)
}
func (m *CustomTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag.Merge(m, src)
}
func (m *CustomTag) XXX_Size() int {
	return xxx_messageInfo_CustomTag.Size(m)
}
func (m *CustomTag) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag proto.InternalMessageInfo

func (m *CustomTag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type isCustomTag_Type interface {
	isCustomTag_Type()
}

type CustomTag_Literal_ struct {
	Literal *CustomTag_Literal `protobuf:"bytes,2,opt,name=literal,proto3,oneof"`
}

type CustomTag_Environment_ struct {
	Environment *CustomTag_Environment `protobuf:"bytes,3,opt,name=environment,proto3,oneof"`
}

type CustomTag_RequestHeader struct {
	RequestHeader *CustomTag_Header `protobuf:"bytes,4,opt,name=request_header,json=requestHeader,proto3,oneof"`
}

type CustomTag_Metadata_ struct {
	Metadata *CustomTag_Metadata `protobuf:"bytes,5,opt,name=metadata,proto3,oneof"`
}

func (*CustomTag_Literal_) isCustomTag_Type() {}

func (*CustomTag_Environment_) isCustomTag_Type() {}

func (*CustomTag_RequestHeader) isCustomTag_Type() {}

func (*CustomTag_Metadata_) isCustomTag_Type() {}

func (m *CustomTag) GetType() isCustomTag_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CustomTag) GetLiteral() *CustomTag_Literal {
	if x, ok := m.GetType().(*CustomTag_Literal_); ok {
		return x.Literal
	}
	return nil
}

func (m *CustomTag) GetEnvironment() *CustomTag_Environment {
	if x, ok := m.GetType().(*CustomTag_Environment_); ok {
		return x.Environment
	}
	return nil
}

func (m *CustomTag) GetRequestHeader() *CustomTag_Header {
	if x, ok := m.GetType().(*CustomTag_RequestHeader); ok {
		return x.RequestHeader
	}
	return nil
}

func (m *CustomTag) GetMetadata() *CustomTag_Metadata {
	if x, ok := m.GetType().(*CustomTag_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomTag) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomTag_Literal_)(nil),
		(*CustomTag_Environment_)(nil),
		(*CustomTag_RequestHeader)(nil),
		(*CustomTag_Metadata_)(nil),
	}
}

type CustomTag_Literal struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Literal) Reset()         { *m = CustomTag_Literal{} }
func (m *CustomTag_Literal) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Literal) ProtoMessage()    {}
func (*CustomTag_Literal) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb41267f4395b6, []int{0, 0}
}

func (m *CustomTag_Literal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Literal.Unmarshal(m, b)
}
func (m *CustomTag_Literal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Literal.Marshal(b, m, deterministic)
}
func (m *CustomTag_Literal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Literal.Merge(m, src)
}
func (m *CustomTag_Literal) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Literal.Size(m)
}
func (m *CustomTag_Literal) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Literal.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Literal proto.InternalMessageInfo

func (m *CustomTag_Literal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CustomTag_Environment struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Environment) Reset()         { *m = CustomTag_Environment{} }
func (m *CustomTag_Environment) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Environment) ProtoMessage()    {}
func (*CustomTag_Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb41267f4395b6, []int{0, 1}
}

func (m *CustomTag_Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Environment.Unmarshal(m, b)
}
func (m *CustomTag_Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Environment.Marshal(b, m, deterministic)
}
func (m *CustomTag_Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Environment.Merge(m, src)
}
func (m *CustomTag_Environment) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Environment.Size(m)
}
func (m *CustomTag_Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Environment proto.InternalMessageInfo

func (m *CustomTag_Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomTag_Environment) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type CustomTag_Header struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Header) Reset()         { *m = CustomTag_Header{} }
func (m *CustomTag_Header) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Header) ProtoMessage()    {}
func (*CustomTag_Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb41267f4395b6, []int{0, 2}
}

func (m *CustomTag_Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Header.Unmarshal(m, b)
}
func (m *CustomTag_Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Header.Marshal(b, m, deterministic)
}
func (m *CustomTag_Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Header.Merge(m, src)
}
func (m *CustomTag_Header) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Header.Size(m)
}
func (m *CustomTag_Header) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Header.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Header proto.InternalMessageInfo

func (m *CustomTag_Header) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomTag_Header) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type CustomTag_Metadata struct {
	Kind                 *v3.MetadataKind `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	MetadataKey          *v3.MetadataKey  `protobuf:"bytes,2,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
	DefaultValue         string           `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomTag_Metadata) Reset()         { *m = CustomTag_Metadata{} }
func (m *CustomTag_Metadata) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Metadata) ProtoMessage()    {}
func (*CustomTag_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb41267f4395b6, []int{0, 3}
}

func (m *CustomTag_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Metadata.Unmarshal(m, b)
}
func (m *CustomTag_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Metadata.Marshal(b, m, deterministic)
}
func (m *CustomTag_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Metadata.Merge(m, src)
}
func (m *CustomTag_Metadata) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Metadata.Size(m)
}
func (m *CustomTag_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Metadata proto.InternalMessageInfo

func (m *CustomTag_Metadata) GetKind() *v3.MetadataKind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *CustomTag_Metadata) GetMetadataKey() *v3.MetadataKey {
	if m != nil {
		return m.MetadataKey
	}
	return nil
}

func (m *CustomTag_Metadata) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomTag)(nil), "envoy.type.tracing.v3.CustomTag")
	proto.RegisterType((*CustomTag_Literal)(nil), "envoy.type.tracing.v3.CustomTag.Literal")
	proto.RegisterType((*CustomTag_Environment)(nil), "envoy.type.tracing.v3.CustomTag.Environment")
	proto.RegisterType((*CustomTag_Header)(nil), "envoy.type.tracing.v3.CustomTag.Header")
	proto.RegisterType((*CustomTag_Metadata)(nil), "envoy.type.tracing.v3.CustomTag.Metadata")
}

func init() {
	proto.RegisterFile("envoy/type/tracing/v3/custom_tag.proto", fileDescriptor_75fb41267f4395b6)
}

var fileDescriptor_75fb41267f4395b6 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xee, 0xb4, 0xdd, 0x7e, 0xbc, 0xd9, 0x2e, 0x32, 0x20, 0xd6, 0xc8, 0x6a, 0x6b, 0x75, 0xb7,
	0xab, 0x6e, 0x02, 0xed, 0x45, 0xf6, 0x22, 0xc4, 0xaf, 0x82, 0x0a, 0x25, 0x88, 0x78, 0x2b, 0xe3,
	0x66, 0xac, 0x61, 0xdb, 0x99, 0x3a, 0x99, 0x04, 0x73, 0x52, 0xf0, 0xe2, 0x6f, 0xf0, 0xa7, 0x78,
	0xf2, 0x22, 0xec, 0x55, 0xf0, 0x17, 0xf8, 0x2b, 0x64, 0x4f, 0x32, 0x93, 0x8f, 0x8d, 0xb6, 0x18,
	0xbc, 0x25, 0xef, 0xfb, 0x3c, 0xcf, 0xfb, 0xbc, 0xcf, 0x24, 0x03, 0x7b, 0x94, 0x45, 0x3c, 0xb6,
	0x65, 0xbc, 0xa2, 0xb6, 0x14, 0xe4, 0xd8, 0x67, 0x73, 0x3b, 0x1a, 0xdb, 0xc7, 0x61, 0x20, 0xf9,
	0x72, 0x26, 0xc9, 0xdc, 0x5a, 0x09, 0x2e, 0x39, 0xbe, 0xa8, 0x71, 0x96, 0xc2, 0x59, 0x29, 0xce,
	0x8a, 0xc6, 0xe6, 0xcd, 0x02, 0x7d, 0x49, 0x25, 0xf1, 0x88, 0x24, 0x8a, 0x9f, 0x3d, 0x27, 0x6c,
	0x73, 0x37, 0xf4, 0x56, 0xc4, 0x26, 0x8c, 0x71, 0x49, 0xa4, 0xcf, 0x59, 0x60, 0x07, 0x92, 0xc8,
	0x30, 0x48, 0xdb, 0xfd, 0xb5, 0x76, 0x44, 0x45, 0xe0, 0x73, 0xa6, 0xa6, 0x24, 0x90, 0x4b, 0x11,
	0x59, 0xf8, 0x1e, 0x91, 0xd4, 0xce, 0x1e, 0x92, 0xc6, 0xf5, 0x1f, 0x4d, 0x68, 0xdf, 0xd7, 0x6e,
	0x9f, 0x93, 0x39, 0xbe, 0x0c, 0x35, 0x49, 0xe6, 0x5d, 0xd4, 0x43, 0xc3, 0xb6, 0xd3, 0x3c, 0x73,
	0xea, 0xa2, 0xda, 0x43, 0xae, 0xaa, 0xe1, 0x07, 0xd0, 0x5c, 0xf8, 0x92, 0x0a, 0xb2, 0xe8, 0x56,
	0x7b, 0x68, 0x68, 0x8c, 0x86, 0xd6, 0xc6, 0x9d, 0xac, 0x5c, 0xcd, 0x7a, 0x9a, 0xe0, 0x27, 0x15,
	0x37, 0xa3, 0xe2, 0x29, 0x18, 0x94, 0x45, 0xbe, 0xe0, 0x6c, 0x49, 0x99, 0xec, 0xd6, 0xb4, 0xd2,
	0x9d, 0x52, 0xa5, 0x87, 0xe7, 0x9c, 0x49, 0xc5, 0x2d, 0x4a, 0xe0, 0x29, 0xec, 0x08, 0xfa, 0x36,
	0xa4, 0x81, 0x9c, 0xbd, 0xa1, 0xc4, 0xa3, 0xa2, 0x5b, 0xd7, 0xa2, 0xfb, 0xa5, 0xa2, 0x13, 0x0d,
	0x9f, 0x54, 0xdc, 0x4e, 0x2a, 0x90, 0x14, 0xf0, 0x63, 0x68, 0x65, 0xf9, 0x77, 0xb7, 0xb4, 0xd6,
	0x41, 0xa9, 0xd6, 0xb3, 0x94, 0x30, 0xa9, 0xb8, 0x39, 0xd9, 0x7c, 0x09, 0xcd, 0x34, 0x02, 0xbc,
	0x0b, 0x5b, 0x11, 0x59, 0x84, 0xf4, 0xef, 0x68, 0x93, 0xea, 0x91, 0xf5, 0xf9, 0xdb, 0xa7, 0xab,
	0x07, 0xb0, 0xbf, 0x69, 0xcc, 0x68, 0x3d, 0x51, 0xf3, 0x23, 0x02, 0xa3, 0x90, 0x09, 0xbe, 0x02,
	0x75, 0x46, 0x96, 0x6b, 0xea, 0xba, 0x88, 0x07, 0xd0, 0xf1, 0xe8, 0x6b, 0x12, 0x2e, 0xe4, 0x2c,
	0xf1, 0xa0, 0xce, 0xaf, 0xed, 0x6e, 0xa7, 0xc5, 0x17, 0xda, 0xc1, 0x48, 0x39, 0x38, 0x84, 0xdb,
	0x65, 0x0e, 0x0a, 0x53, 0xcd, 0xf7, 0xd0, 0x48, 0x23, 0xeb, 0xff, 0x31, 0xbf, 0x73, 0xe6, 0x80,
	0x68, 0xf5, 0xd0, 0x57, 0x84, 0x4e, 0x51, 0xe5, 0x7f, 0x5c, 0x1c, 0x2a, 0x17, 0xc3, 0xf4, 0xaf,
	0xfa, 0x87, 0x8b, 0x64, 0xac, 0xf9, 0x13, 0x41, 0x2b, 0x4b, 0x1e, 0xdf, 0x85, 0xfa, 0x89, 0xcf,
	0x3c, 0xed, 0xc1, 0x18, 0xdd, 0x28, 0x1e, 0x59, 0xfe, 0x3b, 0x45, 0xe3, 0xfc, 0xa4, 0x9e, 0xf8,
	0xcc, 0x73, 0x35, 0x03, 0x3f, 0x82, 0xed, 0x0c, 0x31, 0x3b, 0xa1, 0x71, 0xfa, 0x7d, 0x0f, 0x4a,
	0x15, 0x68, 0xec, 0x1a, 0xcb, 0xf3, 0x97, 0xf5, 0x15, 0x6b, 0x1b, 0x56, 0xb4, 0xd5, 0x8a, 0xb7,
	0x60, 0x58, 0xb6, 0x62, 0x36, 0xe5, 0x68, 0x4f, 0x11, 0xfa, 0x70, 0xad, 0x84, 0xe0, 0x18, 0x50,
	0x57, 0x4d, 0x5c, 0xfb, 0xe5, 0x20, 0xe7, 0xde, 0x97, 0x0f, 0xa7, 0xdf, 0x1b, 0xd5, 0x0b, 0x35,
	0x18, 0xf8, 0x3c, 0x59, 0x64, 0x25, 0xf8, 0xbb, 0x78, 0xf3, 0x87, 0xec, 0xec, 0xe4, 0x32, 0x53,
	0x75, 0x2b, 0x4c, 0xd1, 0xab, 0x86, 0xbe, 0x1e, 0xc6, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x5f, 0x7a, 0x23, 0xe1, 0x04, 0x00, 0x00,
}
