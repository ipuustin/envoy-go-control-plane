// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: envoy/extensions/certificate_providers/local_certificate/v3/local_certificate.proto

package local_certificatev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LocalCertificate_Pkey int32

const (
	LocalCertificate_UNSPECIFIED LocalCertificate_Pkey = 0
	LocalCertificate_RSA_2048    LocalCertificate_Pkey = 1
	LocalCertificate_RSA_3072    LocalCertificate_Pkey = 2
	LocalCertificate_RSA_4096    LocalCertificate_Pkey = 3
	LocalCertificate_ECDSA_P256  LocalCertificate_Pkey = 4
)

// Enum value maps for LocalCertificate_Pkey.
var (
	LocalCertificate_Pkey_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "RSA_2048",
		2: "RSA_3072",
		3: "RSA_4096",
		4: "ECDSA_P256",
	}
	LocalCertificate_Pkey_value = map[string]int32{
		"UNSPECIFIED": 0,
		"RSA_2048":    1,
		"RSA_3072":    2,
		"RSA_4096":    3,
		"ECDSA_P256":  4,
	}
)

func (x LocalCertificate_Pkey) Enum() *LocalCertificate_Pkey {
	p := new(LocalCertificate_Pkey)
	*p = x
	return p
}

func (x LocalCertificate_Pkey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalCertificate_Pkey) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_enumTypes[0].Descriptor()
}

func (LocalCertificate_Pkey) Type() protoreflect.EnumType {
	return &file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_enumTypes[0]
}

func (x LocalCertificate_Pkey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalCertificate_Pkey.Descriptor instead.
func (LocalCertificate_Pkey) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescGZIP(), []int{0, 0}
}

// [#extension: envoy.certificate_providers.local_certificate]
// [#next-free-field: 8]
type LocalCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key and cert of root ca used to sign certificates.
	RootcaCert *v3.DataSource `protobuf:"bytes,1,opt,name=rootca_cert,json=rootcaCert,proto3" json:"rootca_cert,omitempty"`
	RootcaKey  *v3.DataSource `protobuf:"bytes,2,opt,name=rootca_key,json=rootcaKey,proto3" json:"rootca_key,omitempty"`
	// Default Identity key and cert used for TLS handshake
	DefaultIdentityCert *v3.DataSource `protobuf:"bytes,3,opt,name=default_identity_cert,json=defaultIdentityCert,proto3" json:"default_identity_cert,omitempty"`
	DefaultIdentityKey  *v3.DataSource `protobuf:"bytes,4,opt,name=default_identity_key,json=defaultIdentityKey,proto3" json:"default_identity_key,omitempty"`
	// Indicates the time at which the certificate expires.
	ExpirationTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	// The pkey type and size. If not specified, the type/size of original server cert will be copied.
	// It supports RSA_2048, RSA_3072, RSA_4096, ECDSA_P256. Otherwise, defaults to RSA_2048.
	Pkey LocalCertificate_Pkey `protobuf:"varint,6,opt,name=pkey,proto3,enum=envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate_Pkey" json:"pkey,omitempty"`
	// The maximum number of hosts that the cache will hold. If not specified defaults to 1024.
	CacheSize *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=cache_size,json=cacheSize,proto3" json:"cache_size,omitempty"`
}

func (x *LocalCertificate) Reset() {
	*x = LocalCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalCertificate) ProtoMessage() {}

func (x *LocalCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalCertificate.ProtoReflect.Descriptor instead.
func (*LocalCertificate) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *LocalCertificate) GetRootcaCert() *v3.DataSource {
	if x != nil {
		return x.RootcaCert
	}
	return nil
}

func (x *LocalCertificate) GetRootcaKey() *v3.DataSource {
	if x != nil {
		return x.RootcaKey
	}
	return nil
}

func (x *LocalCertificate) GetDefaultIdentityCert() *v3.DataSource {
	if x != nil {
		return x.DefaultIdentityCert
	}
	return nil
}

func (x *LocalCertificate) GetDefaultIdentityKey() *v3.DataSource {
	if x != nil {
		return x.DefaultIdentityKey
	}
	return nil
}

func (x *LocalCertificate) GetExpirationTime() *timestamp.Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *LocalCertificate) GetPkey() LocalCertificate_Pkey {
	if x != nil {
		return x.Pkey
	}
	return LocalCertificate_UNSPECIFIED
}

func (x *LocalCertificate) GetCacheSize() *wrappers.UInt32Value {
	if x != nil {
		return x.CacheSize
	}
	return nil
}

var File_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto protoreflect.FileDescriptor

var file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDesc = []byte{
	0x0a, 0x53, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x05, 0x0a,
	0x10, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x41, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x63, 0x61,
	0x43, 0x65, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x54, 0x0a, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x13, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x65, 0x72, 0x74, 0x12, 0x52, 0x0a, 0x14, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x12, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12,
	0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x66, 0x0a, 0x04, 0x70, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x52, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x2e, 0x50, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x70, 0x6b, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x0a,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x51, 0x0a, 0x04, 0x50, 0x6b, 0x65, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x53, 0x41, 0x5f, 0x32, 0x30, 0x34, 0x38, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x53, 0x41,
	0x5f, 0x33, 0x30, 0x37, 0x32, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x53, 0x41, 0x5f, 0x34,
	0x30, 0x39, 0x36, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x43, 0x44, 0x53, 0x41, 0x5f, 0x50,
	0x32, 0x35, 0x36, 0x10, 0x04, 0x42, 0xe4, 0x01, 0x0a, 0x49, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x33, 0x42, 0x15, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x76, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x33,
	0x3b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescOnce sync.Once
	file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescData = file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDesc
)

func file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescGZIP() []byte {
	file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescData)
	})
	return file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDescData
}

var file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_goTypes = []interface{}{
	(LocalCertificate_Pkey)(0),   // 0: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.Pkey
	(*LocalCertificate)(nil),     // 1: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate
	(*v3.DataSource)(nil),        // 2: envoy.config.core.v3.DataSource
	(*timestamp.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*wrappers.UInt32Value)(nil), // 4: google.protobuf.UInt32Value
}
var file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.rootca_cert:type_name -> envoy.config.core.v3.DataSource
	2, // 1: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.rootca_key:type_name -> envoy.config.core.v3.DataSource
	2, // 2: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.default_identity_cert:type_name -> envoy.config.core.v3.DataSource
	2, // 3: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.default_identity_key:type_name -> envoy.config.core.v3.DataSource
	3, // 4: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.expiration_time:type_name -> google.protobuf.Timestamp
	0, // 5: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.pkey:type_name -> envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.Pkey
	4, // 6: envoy.extensions.certificate_providers.local_certificate.v3.LocalCertificate.cache_size:type_name -> google.protobuf.UInt32Value
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_init()
}
func file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_init() {
	if File_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalCertificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_msgTypes,
	}.Build()
	File_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto = out.File
	file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_rawDesc = nil
	file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_goTypes = nil
	file_envoy_extensions_certificate_providers_local_certificate_v3_local_certificate_proto_depIdxs = nil
}