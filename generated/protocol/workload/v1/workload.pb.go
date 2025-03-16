// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: protocol/workload/v1/workload.proto

package workloadv1

import (
	_ "github.com/gandelm/gandelm/generated/protocol/catalog/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CatalogId     string                 `protobuf:"bytes,1,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Workloads     []*Workload            `protobuf:"bytes,1,rep,name=workloads,proto3" json:"workloads,omitempty"`
	Deployments   []*Deployment          `protobuf:"bytes,2,rep,name=deployments,proto3" json:"deployments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetWorkloads() []*Workload {
	if x != nil {
		return x.Workloads
	}
	return nil
}

func (x *ListResponse) GetDeployments() []*Deployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CatalogId     string                 `protobuf:"bytes,1,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	WorkloadId    string                 `protobuf:"bytes,2,opt,name=workload_id,json=workloadId,proto3" json:"workload_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

func (x *GetRequest) GetWorkloadId() string {
	if x != nil {
		return x.WorkloadId
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Workload      *Workload              `protobuf:"bytes,1,opt,name=workload,proto3" json:"workload,omitempty"`
	Deployments   []*Deployment          `protobuf:"bytes,2,rep,name=deployments,proto3" json:"deployments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetWorkload() *Workload {
	if x != nil {
		return x.Workload
	}
	return nil
}

func (x *GetResponse) GetDeployments() []*Deployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type Workload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkloadId    string                 `protobuf:"bytes,1,opt,name=workload_id,json=workloadId,proto3" json:"workload_id,omitempty"`
	Endpoint      string                 `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Entrypoint    string                 `protobuf:"bytes,3,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	ExternalLinks []*ExternalLink        `protobuf:"bytes,4,rep,name=external_links,json=externalLinks,proto3" json:"external_links,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Workload) Reset() {
	*x = Workload{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Workload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workload) ProtoMessage() {}

func (x *Workload) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workload.ProtoReflect.Descriptor instead.
func (*Workload) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{4}
}

func (x *Workload) GetWorkloadId() string {
	if x != nil {
		return x.WorkloadId
	}
	return ""
}

func (x *Workload) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Workload) GetEntrypoint() string {
	if x != nil {
		return x.Entrypoint
	}
	return ""
}

func (x *Workload) GetExternalLinks() []*ExternalLink {
	if x != nil {
		return x.ExternalLinks
	}
	return nil
}

type ExternalLink struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExternalLink) Reset() {
	*x = ExternalLink{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExternalLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalLink) ProtoMessage() {}

func (x *ExternalLink) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalLink.ProtoReflect.Descriptor instead.
func (*ExternalLink) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{5}
}

func (x *ExternalLink) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ExternalLink) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Deployment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Containers    []*Container           `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	ReplicaStatus *ReplicaStatus         `protobuf:"bytes,2,opt,name=replica_status,json=replicaStatus,proto3" json:"replica_status,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{6}
}

func (x *Deployment) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *Deployment) GetReplicaStatus() *ReplicaStatus {
	if x != nil {
		return x.ReplicaStatus
	}
	return nil
}

func (x *Deployment) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Container struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Image         string                 `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsInit        bool                   `protobuf:"varint,3,opt,name=is_init,json=isInit,proto3" json:"is_init,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Container) Reset() {
	*x = Container{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{7}
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetIsInit() bool {
	if x != nil {
		return x.IsInit
	}
	return false
}

type ReplicaStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Desired       uint32                 `protobuf:"varint,1,opt,name=desired,proto3" json:"desired,omitempty"`
	Available     uint32                 `protobuf:"varint,2,opt,name=available,proto3" json:"available,omitempty"`
	Ready         uint32                 `protobuf:"varint,3,opt,name=ready,proto3" json:"ready,omitempty"`
	Updated       uint32                 `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReplicaStatus) Reset() {
	*x = ReplicaStatus{}
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplicaStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicaStatus) ProtoMessage() {}

func (x *ReplicaStatus) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_workload_v1_workload_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicaStatus.ProtoReflect.Descriptor instead.
func (*ReplicaStatus) Descriptor() ([]byte, []int) {
	return file_protocol_workload_v1_workload_proto_rawDescGZIP(), []int{8}
}

func (x *ReplicaStatus) GetDesired() uint32 {
	if x != nil {
		return x.Desired
	}
	return 0
}

func (x *ReplicaStatus) GetAvailable() uint32 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *ReplicaStatus) GetReady() uint32 {
	if x != nil {
		return x.Ready
	}
	return 0
}

func (x *ReplicaStatus) GetUpdated() uint32 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_protocol_workload_v1_workload_proto protoreflect.FileDescriptor

var file_protocol_workload_v1_workload_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x4c, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67,
	0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x6e,
	0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x08, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x48, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x6e,
	0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0d, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x36, 0x0a, 0x0c,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0xb1, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c,
	0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x12, 0x49, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x61,
	0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x22, 0x77, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x73, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x32, 0xa8, 0x01, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e,
	0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x48, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x61, 0x6e, 0x64,
	0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x6e,
	0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xdc, 0x01, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2f, 0x67, 0x61,
	0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x47, 0x57, 0x58, 0xaa, 0x02, 0x13, 0x47, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47, 0x61,
	0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x5c, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1f, 0x47, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x5c, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x47, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x3a, 0x3a, 0x57,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_protocol_workload_v1_workload_proto_rawDescOnce sync.Once
	file_protocol_workload_v1_workload_proto_rawDescData []byte
)

func file_protocol_workload_v1_workload_proto_rawDescGZIP() []byte {
	file_protocol_workload_v1_workload_proto_rawDescOnce.Do(func() {
		file_protocol_workload_v1_workload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protocol_workload_v1_workload_proto_rawDesc), len(file_protocol_workload_v1_workload_proto_rawDesc)))
	})
	return file_protocol_workload_v1_workload_proto_rawDescData
}

var file_protocol_workload_v1_workload_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protocol_workload_v1_workload_proto_goTypes = []any{
	(*ListRequest)(nil),   // 0: gandelm.workload.v1.ListRequest
	(*ListResponse)(nil),  // 1: gandelm.workload.v1.ListResponse
	(*GetRequest)(nil),    // 2: gandelm.workload.v1.GetRequest
	(*GetResponse)(nil),   // 3: gandelm.workload.v1.GetResponse
	(*Workload)(nil),      // 4: gandelm.workload.v1.Workload
	(*ExternalLink)(nil),  // 5: gandelm.workload.v1.ExternalLink
	(*Deployment)(nil),    // 6: gandelm.workload.v1.Deployment
	(*Container)(nil),     // 7: gandelm.workload.v1.Container
	(*ReplicaStatus)(nil), // 8: gandelm.workload.v1.ReplicaStatus
}
var file_protocol_workload_v1_workload_proto_depIdxs = []int32{
	4, // 0: gandelm.workload.v1.ListResponse.workloads:type_name -> gandelm.workload.v1.Workload
	6, // 1: gandelm.workload.v1.ListResponse.deployments:type_name -> gandelm.workload.v1.Deployment
	4, // 2: gandelm.workload.v1.GetResponse.workload:type_name -> gandelm.workload.v1.Workload
	6, // 3: gandelm.workload.v1.GetResponse.deployments:type_name -> gandelm.workload.v1.Deployment
	5, // 4: gandelm.workload.v1.Workload.external_links:type_name -> gandelm.workload.v1.ExternalLink
	7, // 5: gandelm.workload.v1.Deployment.containers:type_name -> gandelm.workload.v1.Container
	8, // 6: gandelm.workload.v1.Deployment.replica_status:type_name -> gandelm.workload.v1.ReplicaStatus
	0, // 7: gandelm.workload.v1.WorkloadService.List:input_type -> gandelm.workload.v1.ListRequest
	2, // 8: gandelm.workload.v1.WorkloadService.Get:input_type -> gandelm.workload.v1.GetRequest
	1, // 9: gandelm.workload.v1.WorkloadService.List:output_type -> gandelm.workload.v1.ListResponse
	3, // 10: gandelm.workload.v1.WorkloadService.Get:output_type -> gandelm.workload.v1.GetResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protocol_workload_v1_workload_proto_init() }
func file_protocol_workload_v1_workload_proto_init() {
	if File_protocol_workload_v1_workload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protocol_workload_v1_workload_proto_rawDesc), len(file_protocol_workload_v1_workload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_workload_v1_workload_proto_goTypes,
		DependencyIndexes: file_protocol_workload_v1_workload_proto_depIdxs,
		MessageInfos:      file_protocol_workload_v1_workload_proto_msgTypes,
	}.Build()
	File_protocol_workload_v1_workload_proto = out.File
	file_protocol_workload_v1_workload_proto_goTypes = nil
	file_protocol_workload_v1_workload_proto_depIdxs = nil
}
