// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/data_lifecycle/purge.proto

package data_lifecycle

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RunRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{0}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

type RunResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{1}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

type ConfigureRequest struct {
	Enabled              bool          `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty" toml:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	Recurrence           string        `protobuf:"bytes,2,opt,name=recurrence,proto3" json:"recurrence,omitempty" toml:"recurrence,omitempty" mapstructure:"recurrence,omitempty"`
	PolicyUpdate         *PolicyUpdate `protobuf:"bytes,3,opt,name=policy_update,json=policyUpdate,proto3" json:"policy_update,omitempty" toml:"policy_update,omitempty" mapstructure:"policy_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte        `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32         `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigureRequest) Reset()         { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()    {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{2}
}

func (m *ConfigureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureRequest.Unmarshal(m, b)
}
func (m *ConfigureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureRequest.Marshal(b, m, deterministic)
}
func (m *ConfigureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureRequest.Merge(m, src)
}
func (m *ConfigureRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigureRequest.Size(m)
}
func (m *ConfigureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureRequest proto.InternalMessageInfo

func (m *ConfigureRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ConfigureRequest) GetRecurrence() string {
	if m != nil {
		return m.Recurrence
	}
	return ""
}

func (m *ConfigureRequest) GetPolicyUpdate() *PolicyUpdate {
	if m != nil {
		return m.PolicyUpdate
	}
	return nil
}

type ConfigureResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigureResponse) Reset()         { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()    {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{3}
}

func (m *ConfigureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigureResponse.Unmarshal(m, b)
}
func (m *ConfigureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigureResponse.Marshal(b, m, deterministic)
}
func (m *ConfigureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigureResponse.Merge(m, src)
}
func (m *ConfigureResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigureResponse.Size(m)
}
func (m *ConfigureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigureResponse proto.InternalMessageInfo

type ShowRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ShowRequest) Reset()         { *m = ShowRequest{} }
func (m *ShowRequest) String() string { return proto.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()    {}
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{4}
}

func (m *ShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowRequest.Unmarshal(m, b)
}
func (m *ShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowRequest.Marshal(b, m, deterministic)
}
func (m *ShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowRequest.Merge(m, src)
}
func (m *ShowRequest) XXX_Size() int {
	return xxx_messageInfo_ShowRequest.Size(m)
}
func (m *ShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowRequest proto.InternalMessageInfo

type ShowResponse struct {
	InstanceName         string               `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty" toml:"instance_name,omitempty" mapstructure:"instance_name,omitempty"`
	WorkflowName         string               `protobuf:"bytes,2,opt,name=workflow_name,json=workflowName,proto3" json:"workflow_name,omitempty" toml:"workflow_name,omitempty" mapstructure:"workflow_name,omitempty"`
	EsPolicies           []*EsPolicy          `protobuf:"bytes,3,rep,name=es_policies,json=esPolicies,proto3" json:"es_policies,omitempty" toml:"es_policies,omitempty" mapstructure:"es_policies,omitempty"`
	PgPolicies           []*PgPolicy          `protobuf:"bytes,4,rep,name=pg_policies,json=pgPolicies,proto3" json:"pg_policies,omitempty" toml:"pg_policies,omitempty" mapstructure:"pg_policies,omitempty"`
	Recurrence           string               `protobuf:"bytes,5,opt,name=recurrence,proto3" json:"recurrence,omitempty" toml:"recurrence,omitempty" mapstructure:"recurrence,omitempty"`
	NextDueAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=next_due_at,json=nextDueAt,proto3" json:"next_due_at,omitempty" toml:"next_due_at,omitempty" mapstructure:"next_due_at,omitempty"`
	LastEnqueuedAt       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=last_enqueued_at,json=lastEnqueuedAt,proto3" json:"last_enqueued_at,omitempty" toml:"last_enqueued_at,omitempty" mapstructure:"last_enqueued_at,omitempty"`
	Enabled              bool                 `protobuf:"varint,8,opt,name=enabled,proto3" json:"enabled,omitempty" toml:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	LastStart            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=last_start,json=lastStart,proto3" json:"last_start,omitempty" toml:"last_start,omitempty" mapstructure:"last_start,omitempty"`
	LastEnd              *timestamp.Timestamp `protobuf:"bytes,10,opt,name=last_end,json=lastEnd,proto3" json:"last_end,omitempty" toml:"last_end,omitempty" mapstructure:"last_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ShowResponse) Reset()         { *m = ShowResponse{} }
func (m *ShowResponse) String() string { return proto.CompactTextString(m) }
func (*ShowResponse) ProtoMessage()    {}
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{5}
}

func (m *ShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowResponse.Unmarshal(m, b)
}
func (m *ShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowResponse.Marshal(b, m, deterministic)
}
func (m *ShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowResponse.Merge(m, src)
}
func (m *ShowResponse) XXX_Size() int {
	return xxx_messageInfo_ShowResponse.Size(m)
}
func (m *ShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowResponse proto.InternalMessageInfo

func (m *ShowResponse) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *ShowResponse) GetWorkflowName() string {
	if m != nil {
		return m.WorkflowName
	}
	return ""
}

func (m *ShowResponse) GetEsPolicies() []*EsPolicy {
	if m != nil {
		return m.EsPolicies
	}
	return nil
}

func (m *ShowResponse) GetPgPolicies() []*PgPolicy {
	if m != nil {
		return m.PgPolicies
	}
	return nil
}

func (m *ShowResponse) GetRecurrence() string {
	if m != nil {
		return m.Recurrence
	}
	return ""
}

func (m *ShowResponse) GetNextDueAt() *timestamp.Timestamp {
	if m != nil {
		return m.NextDueAt
	}
	return nil
}

func (m *ShowResponse) GetLastEnqueuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.LastEnqueuedAt
	}
	return nil
}

func (m *ShowResponse) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ShowResponse) GetLastStart() *timestamp.Timestamp {
	if m != nil {
		return m.LastStart
	}
	return nil
}

func (m *ShowResponse) GetLastEnd() *timestamp.Timestamp {
	if m != nil {
		return m.LastEnd
	}
	return nil
}

type PolicyUpdate struct {
	Es                   []*EsPolicyUpdate `protobuf:"bytes,1,rep,name=es,proto3" json:"es,omitempty" toml:"es,omitempty" mapstructure:"es,omitempty"`
	Pg                   []*PgPolicyUpdate `protobuf:"bytes,2,rep,name=pg,proto3" json:"pg,omitempty" toml:"pg,omitempty" mapstructure:"pg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PolicyUpdate) Reset()         { *m = PolicyUpdate{} }
func (m *PolicyUpdate) String() string { return proto.CompactTextString(m) }
func (*PolicyUpdate) ProtoMessage()    {}
func (*PolicyUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{6}
}

func (m *PolicyUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyUpdate.Unmarshal(m, b)
}
func (m *PolicyUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyUpdate.Marshal(b, m, deterministic)
}
func (m *PolicyUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyUpdate.Merge(m, src)
}
func (m *PolicyUpdate) XXX_Size() int {
	return xxx_messageInfo_PolicyUpdate.Size(m)
}
func (m *PolicyUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyUpdate proto.InternalMessageInfo

func (m *PolicyUpdate) GetEs() []*EsPolicyUpdate {
	if m != nil {
		return m.Es
	}
	return nil
}

func (m *PolicyUpdate) GetPg() []*PgPolicyUpdate {
	if m != nil {
		return m.Pg
	}
	return nil
}

type EsPolicy struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Index                string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty" toml:"index,omitempty" mapstructure:"index,omitempty"`
	OlderThanDays        int32    `protobuf:"varint,3,opt,name=older_than_days,json=olderThanDays,proto3" json:"older_than_days,omitempty" toml:"older_than_days,omitempty" mapstructure:"older_than_days,omitempty"`
	CustomPurgeField     string   `protobuf:"bytes,4,opt,name=custom_purge_field,json=customPurgeField,proto3" json:"custom_purge_field,omitempty" toml:"custom_purge_field,omitempty" mapstructure:"custom_purge_field,omitempty"`
	Disabled             bool     `protobuf:"varint,5,opt,name=disabled,proto3" json:"disabled,omitempty" toml:"disabled,omitempty" mapstructure:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EsPolicy) Reset()         { *m = EsPolicy{} }
func (m *EsPolicy) String() string { return proto.CompactTextString(m) }
func (*EsPolicy) ProtoMessage()    {}
func (*EsPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{7}
}

func (m *EsPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EsPolicy.Unmarshal(m, b)
}
func (m *EsPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EsPolicy.Marshal(b, m, deterministic)
}
func (m *EsPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EsPolicy.Merge(m, src)
}
func (m *EsPolicy) XXX_Size() int {
	return xxx_messageInfo_EsPolicy.Size(m)
}
func (m *EsPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_EsPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_EsPolicy proto.InternalMessageInfo

func (m *EsPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EsPolicy) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *EsPolicy) GetOlderThanDays() int32 {
	if m != nil {
		return m.OlderThanDays
	}
	return 0
}

func (m *EsPolicy) GetCustomPurgeField() string {
	if m != nil {
		return m.CustomPurgeField
	}
	return ""
}

func (m *EsPolicy) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type EsPolicyUpdate struct {
	PolicyName           string   `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	Disabled             bool     `protobuf:"varint,2,opt,name=disabled,proto3" json:"disabled,omitempty" toml:"disabled,omitempty" mapstructure:"disabled,omitempty"`
	OlderThanDays        int32    `protobuf:"varint,3,opt,name=older_than_days,json=olderThanDays,proto3" json:"older_than_days,omitempty" toml:"older_than_days,omitempty" mapstructure:"older_than_days,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *EsPolicyUpdate) Reset()         { *m = EsPolicyUpdate{} }
func (m *EsPolicyUpdate) String() string { return proto.CompactTextString(m) }
func (*EsPolicyUpdate) ProtoMessage()    {}
func (*EsPolicyUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{8}
}

func (m *EsPolicyUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EsPolicyUpdate.Unmarshal(m, b)
}
func (m *EsPolicyUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EsPolicyUpdate.Marshal(b, m, deterministic)
}
func (m *EsPolicyUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EsPolicyUpdate.Merge(m, src)
}
func (m *EsPolicyUpdate) XXX_Size() int {
	return xxx_messageInfo_EsPolicyUpdate.Size(m)
}
func (m *EsPolicyUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_EsPolicyUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_EsPolicyUpdate proto.InternalMessageInfo

func (m *EsPolicyUpdate) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *EsPolicyUpdate) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *EsPolicyUpdate) GetOlderThanDays() int32 {
	if m != nil {
		return m.OlderThanDays
	}
	return 0
}

type PgPolicy struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Disabled             bool     `protobuf:"varint,2,opt,name=disabled,proto3" json:"disabled,omitempty" toml:"disabled,omitempty" mapstructure:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PgPolicy) Reset()         { *m = PgPolicy{} }
func (m *PgPolicy) String() string { return proto.CompactTextString(m) }
func (*PgPolicy) ProtoMessage()    {}
func (*PgPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{9}
}

func (m *PgPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PgPolicy.Unmarshal(m, b)
}
func (m *PgPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PgPolicy.Marshal(b, m, deterministic)
}
func (m *PgPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PgPolicy.Merge(m, src)
}
func (m *PgPolicy) XXX_Size() int {
	return xxx_messageInfo_PgPolicy.Size(m)
}
func (m *PgPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_PgPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_PgPolicy proto.InternalMessageInfo

func (m *PgPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PgPolicy) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type PgPolicyUpdate struct {
	PolicyName           string   `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	Disabled             bool     `protobuf:"varint,2,opt,name=disabled,proto3" json:"disabled,omitempty" toml:"disabled,omitempty" mapstructure:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *PgPolicyUpdate) Reset()         { *m = PgPolicyUpdate{} }
func (m *PgPolicyUpdate) String() string { return proto.CompactTextString(m) }
func (*PgPolicyUpdate) ProtoMessage()    {}
func (*PgPolicyUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebdcc40a6dbce13, []int{10}
}

func (m *PgPolicyUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PgPolicyUpdate.Unmarshal(m, b)
}
func (m *PgPolicyUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PgPolicyUpdate.Marshal(b, m, deterministic)
}
func (m *PgPolicyUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PgPolicyUpdate.Merge(m, src)
}
func (m *PgPolicyUpdate) XXX_Size() int {
	return xxx_messageInfo_PgPolicyUpdate.Size(m)
}
func (m *PgPolicyUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_PgPolicyUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_PgPolicyUpdate proto.InternalMessageInfo

func (m *PgPolicyUpdate) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *PgPolicyUpdate) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func init() {
	proto.RegisterType((*RunRequest)(nil), "chef.automate.infra.data_lifecycle.api.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "chef.automate.infra.data_lifecycle.api.RunResponse")
	proto.RegisterType((*ConfigureRequest)(nil), "chef.automate.infra.data_lifecycle.api.ConfigureRequest")
	proto.RegisterType((*ConfigureResponse)(nil), "chef.automate.infra.data_lifecycle.api.ConfigureResponse")
	proto.RegisterType((*ShowRequest)(nil), "chef.automate.infra.data_lifecycle.api.ShowRequest")
	proto.RegisterType((*ShowResponse)(nil), "chef.automate.infra.data_lifecycle.api.ShowResponse")
	proto.RegisterType((*PolicyUpdate)(nil), "chef.automate.infra.data_lifecycle.api.PolicyUpdate")
	proto.RegisterType((*EsPolicy)(nil), "chef.automate.infra.data_lifecycle.api.EsPolicy")
	proto.RegisterType((*EsPolicyUpdate)(nil), "chef.automate.infra.data_lifecycle.api.EsPolicyUpdate")
	proto.RegisterType((*PgPolicy)(nil), "chef.automate.infra.data_lifecycle.api.PgPolicy")
	proto.RegisterType((*PgPolicyUpdate)(nil), "chef.automate.infra.data_lifecycle.api.PgPolicyUpdate")
}

func init() {
	proto.RegisterFile("api/interservice/data_lifecycle/purge.proto", fileDescriptor_eebdcc40a6dbce13)
}

var fileDescriptor_eebdcc40a6dbce13 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0x7c, 0xb4, 0xc9, 0x24, 0x29, 0x65, 0xe1, 0x60, 0xe5, 0x40, 0x23, 0x23, 0x55,
	0x91, 0x40, 0x0e, 0x4a, 0x0b, 0x6a, 0x7b, 0x2b, 0xb4, 0xbd, 0x81, 0x8a, 0x5b, 0x0e, 0x70, 0xb1,
	0x36, 0xf6, 0xc4, 0x59, 0xe1, 0xac, 0x5d, 0xef, 0x2e, 0x6d, 0x6e, 0x3c, 0x09, 0x47, 0xb8, 0xf0,
	0x6a, 0xbc, 0x03, 0xf2, 0xda, 0x6e, 0x6d, 0x40, 0x24, 0x81, 0x9b, 0x67, 0x76, 0xe6, 0xe7, 0xff,
	0xec, 0x8c, 0xc7, 0xf0, 0x84, 0xc6, 0x6c, 0xc4, 0xb8, 0xc4, 0x44, 0x60, 0xf2, 0x89, 0x79, 0x38,
	0xf2, 0xa9, 0xa4, 0x6e, 0xc8, 0xa6, 0xe8, 0x2d, 0xbc, 0x10, 0x47, 0xb1, 0x4a, 0x02, 0xb4, 0xe3,
	0x24, 0x92, 0x11, 0xd9, 0xf5, 0x66, 0x38, 0xb5, 0xa9, 0x92, 0xd1, 0x9c, 0x4a, 0xb4, 0x19, 0x9f,
	0x26, 0xd4, 0xae, 0xc6, 0xdb, 0x34, 0x66, 0xfd, 0x9d, 0x20, 0x8a, 0x82, 0x34, 0x37, 0xcd, 0x9a,
	0xa8, 0xe9, 0x48, 0xb2, 0x39, 0x0a, 0x49, 0xe7, 0x71, 0x06, 0xb2, 0xba, 0x00, 0x8e, 0xe2, 0x0e,
	0x5e, 0x29, 0x14, 0xd2, 0xea, 0x41, 0x47, 0x5b, 0x22, 0x8e, 0xb8, 0x40, 0xeb, 0x9b, 0x01, 0xdb,
	0xaf, 0x22, 0x3e, 0x65, 0x81, 0x4a, 0x30, 0x8f, 0x21, 0x26, 0x6c, 0x22, 0xa7, 0x93, 0x10, 0x7d,
	0xd3, 0x18, 0x18, 0xc3, 0x96, 0x53, 0x98, 0xe4, 0x11, 0x40, 0x82, 0x9e, 0x4a, 0x12, 0xe4, 0x1e,
	0x9a, 0xb5, 0x81, 0x31, 0x6c, 0x3b, 0x25, 0x0f, 0x79, 0x0f, 0xbd, 0x38, 0x0a, 0x99, 0xb7, 0x70,
	0x55, 0xec, 0x53, 0x89, 0x66, 0x7d, 0x60, 0x0c, 0x3b, 0xe3, 0x7d, 0x7b, 0xb5, 0x62, 0xec, 0x73,
	0x9d, 0xfc, 0x4e, 0xe7, 0x3a, 0xdd, 0xb8, 0x64, 0x59, 0x0f, 0xe0, 0x7e, 0x49, 0x68, 0x2e, 0xbf,
	0x07, 0x9d, 0x8b, 0x59, 0x74, 0x5d, 0x14, 0xf7, 0xbd, 0x01, 0xdd, 0xcc, 0xce, 0xce, 0xc9, 0x63,
	0xe8, 0x31, 0x2e, 0x24, 0xe5, 0x1e, 0xba, 0x9c, 0xce, 0x51, 0xd7, 0xd3, 0x76, 0xba, 0x85, 0xf3,
	0x0d, 0x9d, 0xeb, 0xa0, 0xeb, 0x28, 0xf9, 0x38, 0x0d, 0xa3, 0xeb, 0x2c, 0x28, 0xab, 0xab, 0x5b,
	0x38, 0x75, 0xd0, 0x5b, 0xe8, 0xa0, 0x70, 0xb5, 0x22, 0x86, 0xc2, 0xac, 0x0f, 0xea, 0xc3, 0xce,
	0xf8, 0xd9, 0xaa, 0x75, 0x9d, 0x8a, 0xac, 0x32, 0x07, 0x30, 0x7b, 0x62, 0x28, 0x52, 0x64, 0x1c,
	0xdc, 0x21, 0x1b, 0xeb, 0x21, 0xcf, 0x83, 0x02, 0x19, 0x07, 0xb7, 0xc8, 0x6a, 0x7f, 0x9a, 0xbf,
	0xf5, 0xe7, 0x08, 0x3a, 0x1c, 0x6f, 0xa4, 0xeb, 0x2b, 0x74, 0xa9, 0x34, 0x37, 0x74, 0x77, 0xfa,
	0x76, 0x36, 0x42, 0x76, 0x31, 0x42, 0xf6, 0x65, 0x31, 0x42, 0x4e, 0x3b, 0x0d, 0x3f, 0x51, 0x78,
	0x2c, 0xc9, 0x09, 0x6c, 0x87, 0x54, 0x48, 0x17, 0xf9, 0x95, 0x42, 0x85, 0x7e, 0x0a, 0xd8, 0x5c,
	0x0a, 0xd8, 0x4a, 0x73, 0x4e, 0xf3, 0x94, 0xe3, 0xca, 0x6c, 0xb5, 0xaa, 0xb3, 0x75, 0x08, 0xa0,
	0xf9, 0x42, 0xd2, 0x44, 0x9a, 0xed, 0xe5, 0xd2, 0xd2, 0xe8, 0x8b, 0x34, 0x98, 0x3c, 0x87, 0x56,
	0x2e, 0xcd, 0x37, 0x61, 0x69, 0xe2, 0x66, 0x26, 0xc9, 0xb7, 0xbe, 0x18, 0xd0, 0x2d, 0x4f, 0x1c,
	0x39, 0x83, 0x1a, 0x0a, 0xd3, 0xd0, 0x8d, 0x78, 0xb1, 0x6e, 0x6f, 0xf3, 0xa9, 0xad, 0xa1, 0x48,
	0x39, 0x71, 0x60, 0xd6, 0xd6, 0xe3, 0x14, 0x0d, 0x2d, 0x38, 0x71, 0x60, 0x7d, 0x35, 0xa0, 0x55,
	0xe0, 0x09, 0x81, 0x46, 0x69, 0x84, 0xf5, 0x33, 0x79, 0x08, 0x4d, 0xc6, 0x7d, 0xbc, 0xc9, 0x47,
	0x36, 0x33, 0xc8, 0x2e, 0xdc, 0x8b, 0x42, 0x1f, 0x13, 0x57, 0xce, 0x28, 0x77, 0x7d, 0xba, 0x10,
	0xfa, 0x3b, 0x6c, 0x3a, 0x3d, 0xed, 0xbe, 0x9c, 0x51, 0x7e, 0x42, 0x17, 0x82, 0x3c, 0x05, 0xe2,
	0x29, 0x21, 0xa3, 0xb9, 0xab, 0x17, 0x8f, 0x3b, 0x65, 0x18, 0xfa, 0x66, 0x43, 0xa3, 0xb6, 0xb3,
	0x93, 0xf3, 0xf4, 0xe0, 0x2c, 0xf5, 0x93, 0x3e, 0xb4, 0x7c, 0x26, 0xb2, 0xd6, 0x35, 0x75, 0xeb,
	0x6e, 0x6d, 0x4b, 0xc1, 0x56, 0xf5, 0x1a, 0xc8, 0x0e, 0x74, 0xf2, 0x4d, 0x50, 0x12, 0x0d, 0x99,
	0x4b, 0x7f, 0x50, 0x65, 0x5c, 0xad, 0x8a, 0x5b, 0xb5, 0x00, 0xeb, 0x08, 0x5a, 0xc5, 0xad, 0xfd,
	0xf1, 0x7a, 0xfe, 0xf2, 0x0e, 0xeb, 0x35, 0x6c, 0x55, 0x6f, 0xfc, 0xbf, 0x24, 0x8f, 0x7f, 0xd4,
	0xa0, 0xa9, 0x2f, 0x8b, 0x70, 0xa8, 0x3b, 0x8a, 0x93, 0xf1, 0xaa, 0x7d, 0xbf, 0x5b, 0xce, 0xfd,
	0xbd, 0xb5, 0x72, 0xf2, 0x1d, 0xf7, 0xd9, 0x80, 0xf6, 0xed, 0x66, 0x24, 0x07, 0xab, 0x22, 0x7e,
	0xdd, 0xfa, 0xfd, 0xc3, 0x7f, 0xc8, 0xcc, 0x25, 0x5c, 0x41, 0x23, 0x5d, 0xbb, 0x64, 0x65, 0xfd,
	0xa5, 0xa5, 0xdd, 0xdf, 0x5f, 0x2f, 0x29, 0x7b, 0xe5, 0xcb, 0xa3, 0x0f, 0x07, 0x01, 0x93, 0x33,
	0x35, 0xb1, 0xbd, 0x68, 0x3e, 0x4a, 0x09, 0xa3, 0x82, 0x30, 0x5a, 0xf2, 0x9b, 0x9d, 0x6c, 0xe8,
	0xa5, 0xb0, 0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xbf, 0xf3, 0x86, 0x90, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PurgeClient is the client API for Purge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PurgeClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
}

type purgeClient struct {
	cc grpc.ClientConnInterface
}

func NewPurgeClient(cc grpc.ClientConnInterface) PurgeClient {
	return &purgeClient{cc}
}

func (c *purgeClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.infra.data_lifecycle.api.Purge/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purgeClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.infra.data_lifecycle.api.Purge/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purgeClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.infra.data_lifecycle.api.Purge/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PurgeServer is the server API for Purge service.
type PurgeServer interface {
	Run(context.Context, *RunRequest) (*RunResponse, error)
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
}

// UnimplementedPurgeServer can be embedded to have forward compatible implementations.
type UnimplementedPurgeServer struct {
}

func (*UnimplementedPurgeServer) Run(ctx context.Context, req *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedPurgeServer) Configure(ctx context.Context, req *ConfigureRequest) (*ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedPurgeServer) Show(ctx context.Context, req *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}

func RegisterPurgeServer(s *grpc.Server, srv PurgeServer) {
	s.RegisterService(&_Purge_serviceDesc, srv)
}

func _Purge_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurgeServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.infra.data_lifecycle.api.Purge/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurgeServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purge_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurgeServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.infra.data_lifecycle.api.Purge/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurgeServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purge_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurgeServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.infra.data_lifecycle.api.Purge/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurgeServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Purge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.infra.data_lifecycle.api.Purge",
	HandlerType: (*PurgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Purge_Run_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Purge_Configure_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Purge_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/data_lifecycle/purge.proto",
}
