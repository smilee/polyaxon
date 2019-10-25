// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/run.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Run specification
type Run struct {
	// UUID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Optional name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional Tags of this entity
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional if the entity has been deleted
	Deleted bool `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Required name of user started this entity
	User string `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	// Required name of owner of this entity
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	// Required project name
	Project string `protobuf:"bytes,8,opt,name=project,proto3" json:"project,omitempty"`
	// Optional time when the entityt was created
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional last time the entity was updated
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Optional last time the entity was started
	StartedAt *timestamp.Timestamp `protobuf:"bytes,11,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Optional last time the entity was started
	FinishedAt *timestamp.Timestamp `protobuf:"bytes,12,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	// Optional flag to tell if this entity is managed by the platform
	IsManaged string `protobuf:"bytes,13,opt,name=is_managed,json=isManaged,proto3" json:"is_managed,omitempty"`
	// Optional content of the entity's spec
	Content string `protobuf:"bytes,14,opt,name=content,proto3" json:"content,omitempty"`
	// Optional latest status of this entity
	Status string `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	// Optional run kind
	Kind string `protobuf:"bytes,16,opt,name=kind,proto3" json:"kind,omitempty"`
	// Optional a readme text describing this entity
	Readme string `protobuf:"bytes,17,opt,name=readme,proto3" json:"readme,omitempty"`
	// Optional if this entity was bookmarked
	Bookmarked bool `protobuf:"varint,18,opt,name=bookmarked,proto3" json:"bookmarked,omitempty"`
	// Optional inputs of this entity
	Inputs map[string]string `protobuf:"bytes,19,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional outputs of this entity
	Outputs map[string]string `protobuf:"bytes,20,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional run environment tracked
	RunEnv map[string]string `protobuf:"bytes,21,rep,name=run_env,json=runEnv,proto3" json:"run_env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Is resume
	IsResume bool `protobuf:"varint,22,opt,name=is_resume,json=isResume,proto3" json:"is_resume,omitempty"`
	// Is clone
	IsClone bool `protobuf:"varint,23,opt,name=is_clone,json=isClone,proto3" json:"is_clone,omitempty"`
	// Optional if this run was restarted/copied/resumed
	CloningStrategy string `protobuf:"bytes,24,opt,name=cloning_strategy,json=cloningStrategy,proto3" json:"cloning_strategy,omitempty"`
	// Optional uuid of the original run and pipeline
	Pipeline             string   `protobuf:"bytes,25,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Original             string   `protobuf:"bytes,26,opt,name=original,proto3" json:"original,omitempty"`
	PipelineName         string   `protobuf:"bytes,27,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	OriginalName         string   `protobuf:"bytes,28,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{0}
}

func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Run) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Run) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Run) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Run) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Run) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Run) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Run) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Run) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Run) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Run) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *Run) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

func (m *Run) GetIsManaged() string {
	if m != nil {
		return m.IsManaged
	}
	return ""
}

func (m *Run) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Run) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Run) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Run) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *Run) GetBookmarked() bool {
	if m != nil {
		return m.Bookmarked
	}
	return false
}

func (m *Run) GetInputs() map[string]string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Run) GetOutputs() map[string]string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Run) GetRunEnv() map[string]string {
	if m != nil {
		return m.RunEnv
	}
	return nil
}

func (m *Run) GetIsResume() bool {
	if m != nil {
		return m.IsResume
	}
	return false
}

func (m *Run) GetIsClone() bool {
	if m != nil {
		return m.IsClone
	}
	return false
}

func (m *Run) GetCloningStrategy() string {
	if m != nil {
		return m.CloningStrategy
	}
	return ""
}

func (m *Run) GetPipeline() string {
	if m != nil {
		return m.Pipeline
	}
	return ""
}

func (m *Run) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *Run) GetPipelineName() string {
	if m != nil {
		return m.PipelineName
	}
	return ""
}

func (m *Run) GetOriginalName() string {
	if m != nil {
		return m.OriginalName
	}
	return ""
}

// Request data to create run
type RunBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project where the experiement will be assigned
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Run object
	Run                  *Run     `protobuf:"bytes,3,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunBodyRequest) Reset()         { *m = RunBodyRequest{} }
func (m *RunBodyRequest) String() string { return proto.CompactTextString(m) }
func (*RunBodyRequest) ProtoMessage()    {}
func (*RunBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{1}
}

func (m *RunBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunBodyRequest.Unmarshal(m, b)
}
func (m *RunBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunBodyRequest.Marshal(b, m, deterministic)
}
func (m *RunBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunBodyRequest.Merge(m, src)
}
func (m *RunBodyRequest) XXX_Size() int {
	return xxx_messageInfo_RunBodyRequest.Size(m)
}
func (m *RunBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunBodyRequest proto.InternalMessageInfo

func (m *RunBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RunBodyRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RunBodyRequest) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// Request data to update run
type EntityRunBodyRequest struct {
	// Enitity Run
	Entity *EntityResourceRequest `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// Run object
	Run                  *Run     `protobuf:"bytes,2,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityRunBodyRequest) Reset()         { *m = EntityRunBodyRequest{} }
func (m *EntityRunBodyRequest) String() string { return proto.CompactTextString(m) }
func (*EntityRunBodyRequest) ProtoMessage()    {}
func (*EntityRunBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{2}
}

func (m *EntityRunBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityRunBodyRequest.Unmarshal(m, b)
}
func (m *EntityRunBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityRunBodyRequest.Marshal(b, m, deterministic)
}
func (m *EntityRunBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityRunBodyRequest.Merge(m, src)
}
func (m *EntityRunBodyRequest) XXX_Size() int {
	return xxx_messageInfo_EntityRunBodyRequest.Size(m)
}
func (m *EntityRunBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityRunBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntityRunBodyRequest proto.InternalMessageInfo

func (m *EntityRunBodyRequest) GetEntity() *EntityResourceRequest {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityRunBodyRequest) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

// Contains list runs
type ListRunsResponse struct {
	// Count of the entities
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// List of all entities
	Results []*Run `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Previous page
	Previous string `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
	// Next page
	Next                 string   `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRunsResponse) Reset()         { *m = ListRunsResponse{} }
func (m *ListRunsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRunsResponse) ProtoMessage()    {}
func (*ListRunsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ec9b83e76db539, []int{3}
}

func (m *ListRunsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunsResponse.Unmarshal(m, b)
}
func (m *ListRunsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunsResponse.Marshal(b, m, deterministic)
}
func (m *ListRunsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunsResponse.Merge(m, src)
}
func (m *ListRunsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRunsResponse.Size(m)
}
func (m *ListRunsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunsResponse proto.InternalMessageInfo

func (m *ListRunsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListRunsResponse) GetResults() []*Run {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListRunsResponse) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *ListRunsResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func init() {
	proto.RegisterType((*Run)(nil), "v1.Run")
	proto.RegisterMapType((map[string]string)(nil), "v1.Run.InputsEntry")
	proto.RegisterMapType((map[string]string)(nil), "v1.Run.OutputsEntry")
	proto.RegisterMapType((map[string]string)(nil), "v1.Run.RunEnvEntry")
	proto.RegisterType((*RunBodyRequest)(nil), "v1.RunBodyRequest")
	proto.RegisterType((*EntityRunBodyRequest)(nil), "v1.EntityRunBodyRequest")
	proto.RegisterType((*ListRunsResponse)(nil), "v1.ListRunsResponse")
}

func init() { proto.RegisterFile("v1/run.proto", fileDescriptor_46ec9b83e76db539) }

var fileDescriptor_46ec9b83e76db539 = []byte{
	// 761 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x73, 0x1b, 0x45,
	0x10, 0x2d, 0x49, 0xb6, 0x3e, 0x7a, 0xed, 0xc4, 0x4c, 0x4c, 0x18, 0x2b, 0x01, 0x84, 0x73, 0x31,
	0x45, 0x22, 0x95, 0xcc, 0x05, 0x87, 0x93, 0xa1, 0x7c, 0xa0, 0x8a, 0x8f, 0xaa, 0x85, 0x1b, 0x07,
	0xd5, 0x78, 0xb7, 0xb3, 0x0c, 0x96, 0x66, 0x96, 0x99, 0x9e, 0x4d, 0x54, 0xfc, 0x67, 0x7e, 0x03,
	0x35, 0x5f, 0x42, 0x49, 0x85, 0x72, 0xe5, 0xa4, 0x79, 0xaf, 0xdf, 0x9b, 0xd6, 0xf6, 0x74, 0x37,
	0x1c, 0x75, 0xcb, 0x85, 0x71, 0x6a, 0xde, 0x1a, 0x4d, 0x9a, 0xf5, 0xbb, 0xe5, 0xf4, 0x69, 0xa3,
	0x75, 0xb3, 0xc6, 0x85, 0x68, 0xe5, 0x42, 0x28, 0xa5, 0x49, 0x90, 0xd4, 0xca, 0x46, 0xc5, 0x2e,
	0x1a, 0xd0, 0xad, 0x7b, 0xb5, 0xb0, 0x64, 0x5c, 0x45, 0x29, 0xfa, 0xf9, 0xbb, 0x51, 0x92, 0x1b,
	0xb4, 0x24, 0x36, 0x6d, 0x12, 0x3c, 0x0f, 0x3f, 0xd5, 0x8b, 0x06, 0xd5, 0x0b, 0xfb, 0x5a, 0x34,
	0x0d, 0x9a, 0x85, 0x6e, 0x43, 0x82, 0xf7, 0x24, 0x3b, 0xee, 0x96, 0x8b, 0x5b, 0x61, 0x31, 0xc2,
	0xf3, 0x7f, 0xc6, 0x30, 0x28, 0x9d, 0x62, 0x0c, 0x0e, 0x9c, 0x93, 0x35, 0xef, 0xcd, 0x7a, 0x17,
	0x93, 0x32, 0x9c, 0x3d, 0xa7, 0xc4, 0x06, 0x79, 0x3f, 0x72, 0xfe, 0xcc, 0x66, 0x50, 0xd4, 0x68,
	0x2b, 0x23, 0x43, 0x02, 0x3e, 0x08, 0xa1, 0x7d, 0xca, 0xbb, 0x48, 0x34, 0x96, 0x1f, 0xcc, 0x06,
	0xde, 0xe5, 0xcf, 0x8c, 0xc3, 0xa8, 0xc6, 0x35, 0x12, 0xd6, 0xfc, 0x70, 0xd6, 0xbb, 0x18, 0x97,
	0x19, 0x86, 0xbc, 0x16, 0x0d, 0x1f, 0xa6, 0xbc, 0x16, 0x0d, 0x3b, 0x85, 0x43, 0xfd, 0x5a, 0xa1,
	0xe1, 0xa3, 0x40, 0x46, 0xe0, 0xef, 0x68, 0x8d, 0xfe, 0x13, 0x2b, 0xe2, 0xe3, 0xc0, 0x67, 0xc8,
	0xae, 0x00, 0x2a, 0x83, 0x82, 0xb0, 0x5e, 0x09, 0xe2, 0x93, 0x59, 0xef, 0xa2, 0xb8, 0x9c, 0xce,
	0x63, 0xd9, 0xe6, 0xb9, 0x6c, 0xf3, 0xdf, 0x72, 0xd9, 0xca, 0x49, 0x52, 0x5f, 0x07, 0xab, 0x6b,
	0xeb, 0x6c, 0x85, 0xfb, 0xad, 0x49, 0x1d, 0xad, 0x96, 0x84, 0x49, 0xd6, 0xe2, 0x7e, 0x6b, 0x52,
	0x5f, 0x13, 0xfb, 0x16, 0x8a, 0x57, 0x52, 0x49, 0xfb, 0x47, 0xf4, 0x1e, 0xdd, 0xeb, 0x85, 0x2c,
	0xbf, 0x26, 0xf6, 0x29, 0x80, 0xb4, 0xab, 0x8d, 0x50, 0xa2, 0xc1, 0x9a, 0x1f, 0x87, 0x52, 0x4c,
	0xa4, 0xfd, 0x29, 0x12, 0xbe, 0x4c, 0x95, 0x56, 0x84, 0x8a, 0xf8, 0x83, 0x58, 0xa6, 0x04, 0xd9,
	0x63, 0x18, 0x5a, 0x12, 0xe4, 0x2c, 0x7f, 0x18, 0x02, 0x09, 0xf9, 0x27, 0xb8, 0x93, 0xaa, 0xe6,
	0x27, 0xf1, 0x09, 0xfc, 0xd9, 0x6b, 0x0d, 0x8a, 0x7a, 0x83, 0xfc, 0xa3, 0xa8, 0x8d, 0x88, 0x7d,
	0x06, 0x70, 0xab, 0xf5, 0xdd, 0x46, 0x98, 0x3b, 0xac, 0x39, 0x0b, 0x6f, 0xb9, 0xc7, 0xb0, 0xaf,
	0x60, 0x28, 0x55, 0xeb, 0xc8, 0xf2, 0x47, 0xb3, 0xc1, 0x45, 0x71, 0xf9, 0x68, 0xde, 0x2d, 0xe7,
	0xa5, 0x53, 0xf3, 0x1f, 0x02, 0x7b, 0xa3, 0xc8, 0x6c, 0xcb, 0x24, 0x61, 0x73, 0x18, 0x69, 0x47,
	0x41, 0x7d, 0x1a, 0xd4, 0xa7, 0x59, 0xfd, 0x4b, 0xa4, 0xa3, 0x3c, 0x8b, 0xd8, 0x73, 0x18, 0x19,
	0xa7, 0x56, 0xa8, 0x3a, 0xfe, 0xf1, 0xdb, 0xb7, 0x97, 0x4e, 0xdd, 0xa8, 0x2e, 0xdd, 0x6e, 0x02,
	0x60, 0x4f, 0x60, 0x22, 0xed, 0xca, 0xa0, 0x75, 0x1b, 0xe4, 0x8f, 0xc3, 0x3f, 0x1d, 0x4b, 0x5b,
	0x06, 0xcc, 0xce, 0x60, 0x2c, 0xed, 0xaa, 0x5a, 0x6b, 0x85, 0xfc, 0x93, 0xd8, 0x91, 0xd2, 0x7e,
	0xef, 0x21, 0xfb, 0x12, 0x4e, 0x3c, 0x2f, 0x55, 0xb3, 0xb2, 0x64, 0x04, 0x61, 0xb3, 0xe5, 0x3c,
	0x14, 0xe1, 0x61, 0xe2, 0x7f, 0x4d, 0x34, 0x9b, 0xc2, 0xb8, 0x95, 0x2d, 0xae, 0xa5, 0x42, 0x7e,
	0x16, 0x24, 0x3b, 0xec, 0x63, 0xda, 0xc8, 0x46, 0x2a, 0xb1, 0xe6, 0xd3, 0x18, 0xcb, 0x98, 0x3d,
	0x83, 0xe3, 0xac, 0x5b, 0x85, 0x09, 0x7b, 0x12, 0x04, 0x47, 0x99, 0xfc, 0xd9, 0x4f, 0xda, 0x33,
	0x38, 0xce, 0x86, 0x28, 0x7a, 0x1a, 0x45, 0x99, 0xf4, 0xa2, 0xe9, 0x15, 0x14, 0x7b, 0x95, 0x65,
	0x27, 0x30, 0xb8, 0xc3, 0x6d, 0x1a, 0x62, 0x7f, 0xf4, 0xb3, 0xd4, 0x89, 0xb5, 0xcb, 0x43, 0x1c,
	0xc1, 0xcb, 0xfe, 0x37, 0xbd, 0xe9, 0x4b, 0x38, 0xda, 0x2f, 0xf3, 0x07, 0x79, 0xaf, 0xa0, 0xd8,
	0x2b, 0xf9, 0x87, 0x58, 0xcf, 0x7f, 0x87, 0x07, 0xa5, 0x53, 0xdf, 0xe9, 0x7a, 0x5b, 0xe2, 0x5f,
	0x0e, 0x2d, 0xfd, 0x37, 0xee, 0xbd, 0xff, 0x19, 0xf7, 0xfe, 0xdb, 0xe3, 0x7e, 0x06, 0x03, 0xe3,
	0xe2, 0xea, 0x29, 0x2e, 0x47, 0xa9, 0x05, 0x4a, 0xcf, 0x9d, 0xd7, 0x70, 0x7a, 0xa3, 0x48, 0xd2,
	0xf6, 0x9d, 0x14, 0x4b, 0x18, 0x62, 0xe0, 0x43, 0x8e, 0xe2, 0xf2, 0xcc, 0xbb, 0x92, 0x12, 0xad,
	0x76, 0xa6, 0xc2, 0x24, 0x2d, 0x93, 0x30, 0x67, 0xe9, 0xbf, 0x27, 0xcb, 0xdf, 0x70, 0xf2, 0xa3,
	0xb4, 0x54, 0x3a, 0xe5, 0xdb, 0xa9, 0xd5, 0xca, 0xa2, 0xff, 0x88, 0x4a, 0x3b, 0x45, 0x21, 0xc1,
	0x61, 0x19, 0x01, 0xfb, 0x02, 0x46, 0xbe, 0x01, 0xd7, 0x64, 0x79, 0x3f, 0x74, 0xec, 0xee, 0xa2,
	0xcc, 0x87, 0x1e, 0x32, 0xd8, 0x49, 0xed, 0x6c, 0xda, 0xa6, 0x3b, 0x1c, 0x16, 0x30, 0xbe, 0x21,
	0x7e, 0x90, 0x16, 0x30, 0xbe, 0xa1, 0xdb, 0x61, 0x58, 0x0f, 0x5f, 0xff, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xb9, 0x82, 0x21, 0xa5, 0x65, 0x06, 0x00, 0x00,
}
