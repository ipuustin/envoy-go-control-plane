// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v3/rbac.proto

package envoy_config_rbac_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v32 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
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

type RBAC_Action int32

const (
	RBAC_ALLOW RBAC_Action = 0
	RBAC_DENY  RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}

func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{0, 0}
}

type RBAC struct {
	Action               RBAC_Action        `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.rbac.v3.RBAC_Action" json:"action,omitempty"`
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type Policy struct {
	Permissions          []*Permission  `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Principals           []*Principal   `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	Condition            *v1alpha1.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetCondition() *v1alpha1.Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

type Permission struct {
	// Types that are valid to be assigned to Rule:
	//	*Permission_AndRules
	//	*Permission_OrRules
	//	*Permission_Any
	//	*Permission_Header
	//	*Permission_UrlPath
	//	*Permission_DestinationIp
	//	*Permission_DestinationPort
	//	*Permission_Metadata
	//	*Permission_NotRule
	//	*Permission_RequestedServerName
	Rule                 isPermission_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{2}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

type isPermission_Rule interface {
	isPermission_Rule()
}

type Permission_AndRules struct {
	AndRules *Permission_Set `protobuf:"bytes,1,opt,name=and_rules,json=andRules,proto3,oneof"`
}

type Permission_OrRules struct {
	OrRules *Permission_Set `protobuf:"bytes,2,opt,name=or_rules,json=orRules,proto3,oneof"`
}

type Permission_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Permission_Header struct {
	Header *v3.HeaderMatcher `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type Permission_UrlPath struct {
	UrlPath *v31.PathMatcher `protobuf:"bytes,10,opt,name=url_path,json=urlPath,proto3,oneof"`
}

type Permission_DestinationIp struct {
	DestinationIp *v32.CidrRange `protobuf:"bytes,5,opt,name=destination_ip,json=destinationIp,proto3,oneof"`
}

type Permission_DestinationPort struct {
	DestinationPort uint32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,proto3,oneof"`
}

type Permission_Metadata struct {
	Metadata *v31.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Permission_NotRule struct {
	NotRule *Permission `protobuf:"bytes,8,opt,name=not_rule,json=notRule,proto3,oneof"`
}

type Permission_RequestedServerName struct {
	RequestedServerName *v31.StringMatcher `protobuf:"bytes,9,opt,name=requested_server_name,json=requestedServerName,proto3,oneof"`
}

func (*Permission_AndRules) isPermission_Rule() {}

func (*Permission_OrRules) isPermission_Rule() {}

func (*Permission_Any) isPermission_Rule() {}

func (*Permission_Header) isPermission_Rule() {}

func (*Permission_UrlPath) isPermission_Rule() {}

func (*Permission_DestinationIp) isPermission_Rule() {}

func (*Permission_DestinationPort) isPermission_Rule() {}

func (*Permission_Metadata) isPermission_Rule() {}

func (*Permission_NotRule) isPermission_Rule() {}

func (*Permission_RequestedServerName) isPermission_Rule() {}

func (m *Permission) GetRule() isPermission_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Permission) GetAndRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_AndRules); ok {
		return x.AndRules
	}
	return nil
}

func (m *Permission) GetOrRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_OrRules); ok {
		return x.OrRules
	}
	return nil
}

func (m *Permission) GetAny() bool {
	if x, ok := m.GetRule().(*Permission_Any); ok {
		return x.Any
	}
	return false
}

func (m *Permission) GetHeader() *v3.HeaderMatcher {
	if x, ok := m.GetRule().(*Permission_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission) GetUrlPath() *v31.PathMatcher {
	if x, ok := m.GetRule().(*Permission_UrlPath); ok {
		return x.UrlPath
	}
	return nil
}

func (m *Permission) GetDestinationIp() *v32.CidrRange {
	if x, ok := m.GetRule().(*Permission_DestinationIp); ok {
		return x.DestinationIp
	}
	return nil
}

func (m *Permission) GetDestinationPort() uint32 {
	if x, ok := m.GetRule().(*Permission_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *Permission) GetMetadata() *v31.MetadataMatcher {
	if x, ok := m.GetRule().(*Permission_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Permission) GetNotRule() *Permission {
	if x, ok := m.GetRule().(*Permission_NotRule); ok {
		return x.NotRule
	}
	return nil
}

func (m *Permission) GetRequestedServerName() *v31.StringMatcher {
	if x, ok := m.GetRule().(*Permission_RequestedServerName); ok {
		return x.RequestedServerName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Permission) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Permission_AndRules)(nil),
		(*Permission_OrRules)(nil),
		(*Permission_Any)(nil),
		(*Permission_Header)(nil),
		(*Permission_UrlPath)(nil),
		(*Permission_DestinationIp)(nil),
		(*Permission_DestinationPort)(nil),
		(*Permission_Metadata)(nil),
		(*Permission_NotRule)(nil),
		(*Permission_RequestedServerName)(nil),
	}
}

type Permission_Set struct {
	Rules                []*Permission `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Permission_Set) Reset()         { *m = Permission_Set{} }
func (m *Permission_Set) String() string { return proto.CompactTextString(m) }
func (*Permission_Set) ProtoMessage()    {}
func (*Permission_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{2, 0}
}

func (m *Permission_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Set.Unmarshal(m, b)
}
func (m *Permission_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Set.Marshal(b, m, deterministic)
}
func (m *Permission_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Set.Merge(m, src)
}
func (m *Permission_Set) XXX_Size() int {
	return xxx_messageInfo_Permission_Set.Size(m)
}
func (m *Permission_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Set proto.InternalMessageInfo

func (m *Permission_Set) GetRules() []*Permission {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Principal struct {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds
	//	*Principal_OrIds
	//	*Principal_Any
	//	*Principal_Authenticated_
	//	*Principal_SourceIp
	//	*Principal_Header
	//	*Principal_UrlPath
	//	*Principal_Metadata
	//	*Principal_NotId
	Identifier           isPrincipal_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{3}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

type isPrincipal_Identifier interface {
	isPrincipal_Identifier()
}

type Principal_AndIds struct {
	AndIds *Principal_Set `protobuf:"bytes,1,opt,name=and_ids,json=andIds,proto3,oneof"`
}

type Principal_OrIds struct {
	OrIds *Principal_Set `protobuf:"bytes,2,opt,name=or_ids,json=orIds,proto3,oneof"`
}

type Principal_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Principal_Authenticated_ struct {
	Authenticated *Principal_Authenticated `protobuf:"bytes,4,opt,name=authenticated,proto3,oneof"`
}

type Principal_SourceIp struct {
	SourceIp *v32.CidrRange `protobuf:"bytes,5,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

type Principal_Header struct {
	Header *v3.HeaderMatcher `protobuf:"bytes,6,opt,name=header,proto3,oneof"`
}

type Principal_UrlPath struct {
	UrlPath *v31.PathMatcher `protobuf:"bytes,9,opt,name=url_path,json=urlPath,proto3,oneof"`
}

type Principal_Metadata struct {
	Metadata *v31.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Principal_NotId struct {
	NotId *Principal `protobuf:"bytes,8,opt,name=not_id,json=notId,proto3,oneof"`
}

func (*Principal_AndIds) isPrincipal_Identifier() {}

func (*Principal_OrIds) isPrincipal_Identifier() {}

func (*Principal_Any) isPrincipal_Identifier() {}

func (*Principal_Authenticated_) isPrincipal_Identifier() {}

func (*Principal_SourceIp) isPrincipal_Identifier() {}

func (*Principal_Header) isPrincipal_Identifier() {}

func (*Principal_UrlPath) isPrincipal_Identifier() {}

func (*Principal_Metadata) isPrincipal_Identifier() {}

func (*Principal_NotId) isPrincipal_Identifier() {}

func (m *Principal) GetIdentifier() isPrincipal_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Principal) GetAndIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_AndIds); ok {
		return x.AndIds
	}
	return nil
}

func (m *Principal) GetOrIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_OrIds); ok {
		return x.OrIds
	}
	return nil
}

func (m *Principal) GetAny() bool {
	if x, ok := m.GetIdentifier().(*Principal_Any); ok {
		return x.Any
	}
	return false
}

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if x, ok := m.GetIdentifier().(*Principal_Authenticated_); ok {
		return x.Authenticated
	}
	return nil
}

func (m *Principal) GetSourceIp() *v32.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_SourceIp); ok {
		return x.SourceIp
	}
	return nil
}

func (m *Principal) GetHeader() *v3.HeaderMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Principal) GetUrlPath() *v31.PathMatcher {
	if x, ok := m.GetIdentifier().(*Principal_UrlPath); ok {
		return x.UrlPath
	}
	return nil
}

func (m *Principal) GetMetadata() *v31.MetadataMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Principal) GetNotId() *Principal {
	if x, ok := m.GetIdentifier().(*Principal_NotId); ok {
		return x.NotId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Principal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Principal_AndIds)(nil),
		(*Principal_OrIds)(nil),
		(*Principal_Any)(nil),
		(*Principal_Authenticated_)(nil),
		(*Principal_SourceIp)(nil),
		(*Principal_Header)(nil),
		(*Principal_UrlPath)(nil),
		(*Principal_Metadata)(nil),
		(*Principal_NotId)(nil),
	}
}

type Principal_Set struct {
	Ids                  []*Principal `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Principal_Set) Reset()         { *m = Principal_Set{} }
func (m *Principal_Set) String() string { return proto.CompactTextString(m) }
func (*Principal_Set) ProtoMessage()    {}
func (*Principal_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{3, 0}
}

func (m *Principal_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Set.Unmarshal(m, b)
}
func (m *Principal_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Set.Marshal(b, m, deterministic)
}
func (m *Principal_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Set.Merge(m, src)
}
func (m *Principal_Set) XXX_Size() int {
	return xxx_messageInfo_Principal_Set.Size(m)
}
func (m *Principal_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Set proto.InternalMessageInfo

func (m *Principal_Set) GetIds() []*Principal {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Principal_Authenticated struct {
	PrincipalName        *v31.StringMatcher `protobuf:"bytes,2,opt,name=principal_name,json=principalName,proto3" json:"principal_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_17db9cb79f9c2859, []int{3, 1}
}

func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (m *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(m, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetPrincipalName() *v31.StringMatcher {
	if m != nil {
		return m.PrincipalName
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.rbac.v3.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v3.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v3.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v3.Policy")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v3.Permission")
	proto.RegisterType((*Permission_Set)(nil), "envoy.config.rbac.v3.Permission.Set")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v3.Principal")
	proto.RegisterType((*Principal_Set)(nil), "envoy.config.rbac.v3.Principal.Set")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v3.Principal.Authenticated")
}

func init() { proto.RegisterFile("envoy/config/rbac/v3/rbac.proto", fileDescriptor_17db9cb79f9c2859) }

var fileDescriptor_17db9cb79f9c2859 = []byte{
	// 1050 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0xe3, 0x44,
	0x18, 0x8d, 0x9d, 0xc4, 0x71, 0xbe, 0x2a, 0x25, 0x0c, 0x20, 0x4c, 0x4a, 0xbb, 0xa9, 0xb7, 0x5d,
	0x95, 0xd5, 0xe2, 0x68, 0x53, 0x09, 0xed, 0x06, 0x28, 0xd4, 0xdd, 0x4a, 0x29, 0x74, 0x97, 0xca,
	0x15, 0x42, 0xcb, 0x25, 0x9a, 0xda, 0xb3, 0xcd, 0x40, 0x32, 0x63, 0xc6, 0x93, 0xa8, 0xb9, 0x81,
	0xb8, 0x70, 0xe6, 0x82, 0xc4, 0x95, 0xff, 0x82, 0x0b, 0x27, 0x24, 0xae, 0xfc, 0x37, 0xa8, 0x97,
	0x45, 0x33, 0x76, 0x7e, 0x2d, 0x69, 0x43, 0x61, 0x4f, 0xb1, 0x32, 0xef, 0xbd, 0xf9, 0xe6, 0xfb,
	0xde, 0x1b, 0x1b, 0x6e, 0x11, 0x36, 0xe4, 0xa3, 0x46, 0xc8, 0xd9, 0x33, 0x7a, 0xde, 0x10, 0x67,
	0x38, 0x6c, 0x0c, 0x77, 0xf5, 0xaf, 0x17, 0x0b, 0x2e, 0x39, 0x7a, 0x5d, 0x03, 0xbc, 0x14, 0xe0,
	0xe9, 0x85, 0xe1, 0x6e, 0xcd, 0x9d, 0xa3, 0x85, 0x5c, 0x10, 0x45, 0xc3, 0x51, 0x24, 0x48, 0x92,
	0xa4, 0xcc, 0xda, 0xbd, 0x79, 0x69, 0x3e, 0x90, 0x1a, 0xa4, 0x1f, 0x3a, 0x21, 0xef, 0xc7, 0x9c,
	0x11, 0x26, 0xc7, 0xe8, 0xad, 0x14, 0x2d, 0x47, 0x31, 0x69, 0xf4, 0xb1, 0x0c, 0xbb, 0x44, 0x28,
	0x74, 0x9f, 0x48, 0x1c, 0x61, 0x89, 0x33, 0x54, 0x7d, 0x31, 0x2a, 0xc6, 0xb2, 0x9b, 0x21, 0xdc,
	0xc5, 0x88, 0x44, 0x0a, 0xca, 0xce, 0x33, 0xcc, 0xf6, 0x39, 0xe7, 0xe7, 0x3d, 0xd2, 0xc0, 0x31,
	0x6d, 0x90, 0x8b, 0x58, 0x34, 0x86, 0xf7, 0x71, 0x2f, 0xee, 0xe2, 0xfb, 0x8d, 0x64, 0xc4, 0x24,
	0xbe, 0xc8, 0x60, 0xeb, 0x83, 0x28, 0xc6, 0x0d, 0xcc, 0x18, 0x97, 0x58, 0x52, 0xce, 0x92, 0x46,
	0x22, 0xb1, 0x1c, 0x8c, 0x2b, 0xde, 0xfc, 0xc7, 0xf2, 0x90, 0x88, 0x84, 0x72, 0x36, 0xdd, 0xe8,
	0xcd, 0x21, 0xee, 0xd1, 0x08, 0xab, 0x83, 0x67, 0x0f, 0xe9, 0x82, 0xfb, 0x8b, 0x09, 0x85, 0xc0,
	0xdf, 0x3f, 0x40, 0x0f, 0xc1, 0xc2, 0xa1, 0x62, 0x3b, 0x46, 0xdd, 0xd8, 0x59, 0x6d, 0x6e, 0x7a,
	0x8b, 0xfa, 0xed, 0x29, 0xac, 0xb7, 0xaf, 0x81, 0x41, 0x46, 0x40, 0x8f, 0xc0, 0x8e, 0x79, 0x8f,
	0x86, 0x94, 0x24, 0x8e, 0x59, 0xcf, 0xef, 0xac, 0x34, 0x77, 0xae, 0x21, 0x9f, 0x64, 0xd0, 0x43,
	0x26, 0xc5, 0x28, 0x98, 0x30, 0x6b, 0x4f, 0xa1, 0x32, 0xb7, 0x84, 0xaa, 0x90, 0xff, 0x9a, 0x8c,
	0x74, 0x39, 0xe5, 0x40, 0x3d, 0xa2, 0x26, 0x14, 0x87, 0xb8, 0x37, 0x20, 0x8e, 0x59, 0x37, 0x76,
	0x56, 0x9a, 0x6f, 0x2f, 0xde, 0x45, 0xab, 0x8c, 0x82, 0x14, 0xda, 0x32, 0x1f, 0x18, 0xee, 0x3a,
	0x58, 0x69, 0xc9, 0xa8, 0x0c, 0xc5, 0xfd, 0xe3, 0xe3, 0xcf, 0xbe, 0xa8, 0xe6, 0x90, 0x0d, 0x85,
	0x47, 0x87, 0x4f, 0x9e, 0x56, 0x8d, 0x56, 0xfd, 0xe7, 0xdf, 0x7f, 0xd8, 0x58, 0x83, 0xb7, 0x16,
	0xa8, 0x35, 0x75, 0xcd, 0xee, 0xf7, 0x26, 0x58, 0xa9, 0x2c, 0x3a, 0x86, 0x95, 0x98, 0x88, 0x3e,
	0x4d, 0x54, 0x83, 0x13, 0xc7, 0xd0, 0xe7, 0xad, 0x5f, 0x51, 0xc9, 0x04, 0xe8, 0xdb, 0x97, 0x7e,
	0xf1, 0x47, 0xc3, 0xb4, 0x8d, 0x60, 0x96, 0x8e, 0x8e, 0x00, 0x62, 0x41, 0x59, 0x48, 0x63, 0xdc,
	0x1b, 0x37, 0xef, 0xd6, 0x15, 0x62, 0x63, 0xdc, 0x8c, 0xd6, 0x0c, 0x19, 0x7d, 0x00, 0xe5, 0x90,
	0xb3, 0x88, 0xea, 0x19, 0xe6, 0x75, 0x83, 0x36, 0xbc, 0xd4, 0x5f, 0x1e, 0x8e, 0xa9, 0xa7, 0xfc,
	0xe5, 0x8d, 0xfd, 0xe5, 0x1d, 0x5e, 0xc4, 0x22, 0x98, 0x12, 0x5a, 0xae, 0xea, 0xc1, 0x3a, 0xac,
	0x2d, 0xec, 0x41, 0x7a, 0x74, 0xf7, 0x37, 0x0b, 0x60, 0x7a, 0x24, 0x74, 0x00, 0x65, 0xcc, 0xa2,
	0x8e, 0x18, 0xf4, 0x48, 0xa2, 0xa7, 0xb4, 0xd2, 0xdc, 0x5a, 0xd6, 0x07, 0xef, 0x94, 0xc8, 0x76,
	0x2e, 0xb0, 0x31, 0x8b, 0x02, 0xc5, 0x43, 0xfb, 0x60, 0x73, 0x91, 0x69, 0x98, 0x37, 0xd2, 0x28,
	0x71, 0x91, 0x4a, 0xac, 0x41, 0x1e, 0xb3, 0x91, 0x3e, 0xb2, 0xed, 0x97, 0x2e, 0xfd, 0xc2, 0x57,
	0xa6, 0x6d, 0xb4, 0x73, 0x81, 0xfa, 0x17, 0xed, 0x81, 0xd5, 0x25, 0x38, 0x22, 0xc2, 0x29, 0x2c,
	0x54, 0x57, 0x77, 0x80, 0x92, 0x6f, 0x6b, 0xd0, 0xe3, 0x34, 0xa7, 0xed, 0x5c, 0x90, 0xb1, 0xd0,
	0x47, 0x60, 0x0f, 0x44, 0xaf, 0xa3, 0x72, 0xed, 0x80, 0x56, 0x70, 0x33, 0x05, 0x15, 0x6c, 0x2f,
	0x0b, 0xb6, 0x2e, 0x10, 0xcb, 0xee, 0x94, 0x5f, 0x1a, 0x88, 0x9e, 0xfa, 0x07, 0xb5, 0x61, 0x35,
	0x22, 0x89, 0xa4, 0x4c, 0x47, 0xb3, 0x43, 0x63, 0xa7, 0xa8, 0x65, 0x5e, 0x98, 0xb2, 0xba, 0xb9,
	0x94, 0xca, 0x01, 0x8d, 0x44, 0x80, 0xd9, 0x39, 0x69, 0xe7, 0x82, 0xca, 0x0c, 0xf1, 0x28, 0x46,
	0xef, 0x41, 0x75, 0x56, 0x29, 0xe6, 0x42, 0x3a, 0x56, 0xdd, 0xd8, 0xa9, 0xf8, 0xe5, 0x4b, 0xdf,
	0xba, 0x5b, 0x70, 0x9e, 0x3f, 0xcf, 0xb7, 0x73, 0xc1, 0x2b, 0x33, 0xa0, 0x13, 0x2e, 0xa4, 0x8a,
	0xe7, 0xf8, 0xf2, 0x72, 0x4a, 0x7a, 0xef, 0x3b, 0x57, 0x1c, 0xe1, 0x71, 0x06, 0x9b, 0x1e, 0x63,
	0xc2, 0x44, 0x1f, 0x82, 0xcd, 0xb8, 0xd4, 0x93, 0x72, 0x6c, 0xad, 0xb2, 0xd4, 0xf4, 0xaa, 0x0d,
	0x8c, 0x4b, 0x35, 0x25, 0xf4, 0x25, 0xbc, 0x21, 0xc8, 0x37, 0x03, 0x92, 0x48, 0x12, 0x75, 0x12,
	0x22, 0x86, 0x44, 0x74, 0x18, 0xee, 0x13, 0xa7, 0x3c, 0x37, 0x96, 0x17, 0x2b, 0x3a, 0xd5, 0xb7,
	0xe5, 0xb4, 0x9e, 0xd7, 0x26, 0x22, 0xa7, 0x5a, 0xe3, 0x09, 0xee, 0x93, 0x5a, 0x02, 0xf9, 0x53,
	0x22, 0xd1, 0xc7, 0x50, 0x1c, 0x7b, 0xf1, 0xa6, 0x99, 0x4c, 0x89, 0xad, 0xbb, 0x2a, 0x04, 0xdb,
	0x70, 0x7b, 0x71, 0x08, 0xe6, 0x0c, 0xd8, 0xba, 0xa3, 0xb0, 0x9b, 0xd9, 0x6b, 0xeb, 0x6a, 0xac,
	0xbf, 0x02, 0x05, 0x25, 0x8e, 0xf2, 0x7f, 0xf9, 0x86, 0xfb, 0x5d, 0x09, 0xca, 0x93, 0x1c, 0xa3,
	0x3d, 0x28, 0xa9, 0x00, 0xd1, 0x68, 0x1c, 0x9f, 0xdb, 0x4b, 0x92, 0x9f, 0x39, 0xdf, 0xc2, 0x2c,
	0x3a, 0x8a, 0x54, 0xe2, 0x2d, 0x2e, 0x34, 0xdd, 0xbc, 0x09, 0xbd, 0xc8, 0x85, 0x62, 0x5f, 0x1b,
	0x9b, 0xcf, 0xa1, 0x82, 0x07, 0xb2, 0x4b, 0x98, 0xa4, 0x21, 0x96, 0x24, 0xca, 0xd2, 0xf3, 0xee,
	0xb2, 0x1d, 0xf6, 0x67, 0x49, 0xca, 0xc2, 0x73, 0x2a, 0x68, 0x0f, 0xca, 0x09, 0x1f, 0x88, 0x90,
	0xdc, 0x28, 0x07, 0x76, 0xca, 0x39, 0x8a, 0x67, 0xd2, 0x6c, 0xfd, 0xef, 0x34, 0x97, 0xff, 0x4b,
	0x9a, 0x5f, 0x4e, 0x96, 0x1e, 0x80, 0xa5, 0xb2, 0x44, 0xa3, 0x2c, 0x49, 0xcb, 0x6e, 0x7c, 0x35,
	0x34, 0xc6, 0xe5, 0x51, 0x54, 0xeb, 0xa7, 0x56, 0x7f, 0x1f, 0xf2, 0xa9, 0x6b, 0x6e, 0xf8, 0xbe,
	0x50, 0xac, 0xd6, 0x3b, 0xca, 0xb9, 0x5b, 0xe0, 0x2e, 0x76, 0xee, 0xac, 0x59, 0x6a, 0x3f, 0x19,
	0x50, 0x99, 0x1b, 0x29, 0xfa, 0x14, 0x56, 0x27, 0xef, 0x9c, 0x34, 0xc0, 0xe6, 0xbf, 0x0f, 0x70,
	0x50, 0x99, 0x70, 0x55, 0x70, 0x5b, 0xbb, 0xaa, 0x12, 0x0f, 0xee, 0x2d, 0xa9, 0x64, 0xae, 0x82,
	0x4f, 0x0a, 0xb6, 0x51, 0x35, 0x5b, 0xdb, 0x8a, 0x5a, 0x87, 0x8d, 0xeb, 0xa9, 0xfe, 0xab, 0x00,
	0x34, 0x52, 0xdc, 0x67, 0x94, 0x08, 0x9d, 0x41, 0xff, 0xe1, 0xaf, 0xdf, 0xfe, 0xf1, 0xa7, 0x65,
	0x56, 0xf3, 0xe0, 0x52, 0x9e, 0x56, 0x1d, 0x0b, 0x7e, 0x31, 0x5a, 0xd8, 0x45, 0xbf, 0x1c, 0x9c,
	0xe1, 0xf0, 0x44, 0x7d, 0x2a, 0x9d, 0x18, 0x67, 0x96, 0xfe, 0x66, 0xda, 0xfd, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0xc9, 0x3b, 0xd1, 0xac, 0x0a, 0x00, 0x00,
}
