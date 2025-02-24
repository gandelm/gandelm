// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: protocol/github_webhook/v1/github_webhook.proto

package github_webhookv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GitHub Webhook のリクエスト
type GitHubWebhookRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	EventType  string                 `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`    // "workflow_run", "push" などのイベント名
	DeliveryId string                 `protobuf:"bytes,2,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"` // GitHub が発行するリクエストの一意 ID
	// Types that are valid to be assigned to Payload:
	//
	//	*GitHubWebhookRequest_WorkflowRun
	//	*GitHubWebhookRequest_Push
	Payload       isGitHubWebhookRequest_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GitHubWebhookRequest) Reset() {
	*x = GitHubWebhookRequest{}
	mi := &file_protocol_github_webhook_v1_github_webhook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitHubWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitHubWebhookRequest) ProtoMessage() {}

func (x *GitHubWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_github_webhook_v1_github_webhook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitHubWebhookRequest.ProtoReflect.Descriptor instead.
func (*GitHubWebhookRequest) Descriptor() ([]byte, []int) {
	return file_protocol_github_webhook_v1_github_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *GitHubWebhookRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *GitHubWebhookRequest) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

func (x *GitHubWebhookRequest) GetPayload() isGitHubWebhookRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *GitHubWebhookRequest) GetWorkflowRun() *WorkflowRunEvent {
	if x != nil {
		if x, ok := x.Payload.(*GitHubWebhookRequest_WorkflowRun); ok {
			return x.WorkflowRun
		}
	}
	return nil
}

func (x *GitHubWebhookRequest) GetPush() *PushEvent {
	if x != nil {
		if x, ok := x.Payload.(*GitHubWebhookRequest_Push); ok {
			return x.Push
		}
	}
	return nil
}

type isGitHubWebhookRequest_Payload interface {
	isGitHubWebhookRequest_Payload()
}

type GitHubWebhookRequest_WorkflowRun struct {
	WorkflowRun *WorkflowRunEvent `protobuf:"bytes,3,opt,name=workflow_run,json=workflowRun,proto3,oneof"`
}

type GitHubWebhookRequest_Push struct {
	Push *PushEvent `protobuf:"bytes,4,opt,name=push,proto3,oneof"`
}

func (*GitHubWebhookRequest_WorkflowRun) isGitHubWebhookRequest_Payload() {}

func (*GitHubWebhookRequest_Push) isGitHubWebhookRequest_Payload() {}

// workflow_run のペイロード
type WorkflowRunEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        string                 `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`         // "completed", "in_progress" など
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`         // "completed", "in_progress"
	Conclusion    string                 `protobuf:"bytes,3,opt,name=conclusion,proto3" json:"conclusion,omitempty"` // "success", "failure"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowRunEvent) Reset() {
	*x = WorkflowRunEvent{}
	mi := &file_protocol_github_webhook_v1_github_webhook_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowRunEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRunEvent) ProtoMessage() {}

func (x *WorkflowRunEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_github_webhook_v1_github_webhook_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRunEvent.ProtoReflect.Descriptor instead.
func (*WorkflowRunEvent) Descriptor() ([]byte, []int) {
	return file_protocol_github_webhook_v1_github_webhook_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowRunEvent) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *WorkflowRunEvent) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WorkflowRunEvent) GetConclusion() string {
	if x != nil {
		return x.Conclusion
	}
	return ""
}

// push のペイロード
type PushEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ref           string                 `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"` // "refs/heads/main"
	Before        string                 `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	After         string                 `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushEvent) Reset() {
	*x = PushEvent{}
	mi := &file_protocol_github_webhook_v1_github_webhook_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushEvent) ProtoMessage() {}

func (x *PushEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_github_webhook_v1_github_webhook_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushEvent.ProtoReflect.Descriptor instead.
func (*PushEvent) Descriptor() ([]byte, []int) {
	return file_protocol_github_webhook_v1_github_webhook_proto_rawDescGZIP(), []int{2}
}

func (x *PushEvent) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *PushEvent) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

func (x *PushEvent) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

var File_protocol_github_webhook_v1_github_webhook_proto protoreflect.FileDescriptor

var file_protocol_github_webhook_v1_github_webhook_proto_rawDesc = string([]byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x14, 0x47,
	0x69, 0x74, 0x48, 0x75, 0x62, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x72, 0x75, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x61, 0x6e, 0x64,
	0x65, 0x6c, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x75, 0x6e, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x75, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x70, 0x75, 0x73,
	0x68, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x62, 0x0a, 0x10,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x4b, 0x0a, 0x09, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x32, 0x70, 0x0a,
	0x14, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x2f, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x87, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x42, 0x12, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x64,
	0x65, 0x6c, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x47, 0x58, 0xaa,
	0x02, 0x18, 0x47, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x47, 0x61, 0x6e,
	0x64, 0x65, 0x6c, 0x6d, 0x5c, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x47, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x5c,
	0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x47,
	0x61, 0x6e, 0x64, 0x65, 0x6c, 0x6d, 0x3a, 0x3a, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_protocol_github_webhook_v1_github_webhook_proto_rawDescOnce sync.Once
	file_protocol_github_webhook_v1_github_webhook_proto_rawDescData []byte
)

func file_protocol_github_webhook_v1_github_webhook_proto_rawDescGZIP() []byte {
	file_protocol_github_webhook_v1_github_webhook_proto_rawDescOnce.Do(func() {
		file_protocol_github_webhook_v1_github_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protocol_github_webhook_v1_github_webhook_proto_rawDesc), len(file_protocol_github_webhook_v1_github_webhook_proto_rawDesc)))
	})
	return file_protocol_github_webhook_v1_github_webhook_proto_rawDescData
}

var file_protocol_github_webhook_v1_github_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocol_github_webhook_v1_github_webhook_proto_goTypes = []any{
	(*GitHubWebhookRequest)(nil), // 0: gandelm.github_webhook.v1.GitHubWebhookRequest
	(*WorkflowRunEvent)(nil),     // 1: gandelm.github_webhook.v1.WorkflowRunEvent
	(*PushEvent)(nil),            // 2: gandelm.github_webhook.v1.PushEvent
	(*emptypb.Empty)(nil),        // 3: google.protobuf.Empty
}
var file_protocol_github_webhook_v1_github_webhook_proto_depIdxs = []int32{
	1, // 0: gandelm.github_webhook.v1.GitHubWebhookRequest.workflow_run:type_name -> gandelm.github_webhook.v1.WorkflowRunEvent
	2, // 1: gandelm.github_webhook.v1.GitHubWebhookRequest.push:type_name -> gandelm.github_webhook.v1.PushEvent
	0, // 2: gandelm.github_webhook.v1.GitHubWebhookService.HandleWebhook:input_type -> gandelm.github_webhook.v1.GitHubWebhookRequest
	3, // 3: gandelm.github_webhook.v1.GitHubWebhookService.HandleWebhook:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_github_webhook_v1_github_webhook_proto_init() }
func file_protocol_github_webhook_v1_github_webhook_proto_init() {
	if File_protocol_github_webhook_v1_github_webhook_proto != nil {
		return
	}
	file_protocol_github_webhook_v1_github_webhook_proto_msgTypes[0].OneofWrappers = []any{
		(*GitHubWebhookRequest_WorkflowRun)(nil),
		(*GitHubWebhookRequest_Push)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protocol_github_webhook_v1_github_webhook_proto_rawDesc), len(file_protocol_github_webhook_v1_github_webhook_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_github_webhook_v1_github_webhook_proto_goTypes,
		DependencyIndexes: file_protocol_github_webhook_v1_github_webhook_proto_depIdxs,
		MessageInfos:      file_protocol_github_webhook_v1_github_webhook_proto_msgTypes,
	}.Build()
	File_protocol_github_webhook_v1_github_webhook_proto = out.File
	file_protocol_github_webhook_v1_github_webhook_proto_goTypes = nil
	file_protocol_github_webhook_v1_github_webhook_proto_depIdxs = nil
}
