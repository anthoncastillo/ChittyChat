// This proto file represents the API for 'ChittyChat';
// a distributed service that enables its clients to chat.
// The service will be using gRPC for communication.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: Chitty-Chat/chittychat/chittychat.proto

package chittychat

import (
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

// Message definition for Chatting
type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientInfo *ClientInfo `protobuf:"bytes,1,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	Content    string      `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_Chitty_Chat_chittychat_chittychat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Response for publishing a message
type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	LamportTime int64 `protobuf:"varint,2,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_Chitty_Chat_chittychat_chittychat_proto_rawDescGZIP(), []int{1}
}

func (x *PublishResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PublishResponse) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

// Client information (used for Join and Leave RPCs)
type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    int64  `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientName  string `protobuf:"bytes,2,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	LamportTime int64  `protobuf:"varint,3,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_Chitty_Chat_chittychat_chittychat_proto_rawDescGZIP(), []int{2}
}

func (x *ClientInfo) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *ClientInfo) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *ClientInfo) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

// Response when a client joins the chat
type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success        bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	LamportTime    int64  `protobuf:"varint,2,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
	WelcomeMessage string `protobuf:"bytes,3,opt,name=welcome_message,json=welcomeMessage,proto3" json:"welcome_message,omitempty"`
	ClientId       int64  `protobuf:"varint,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_Chitty_Chat_chittychat_chittychat_proto_rawDescGZIP(), []int{3}
}

func (x *JoinResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *JoinResponse) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *JoinResponse) GetWelcomeMessage() string {
	if x != nil {
		return x.WelcomeMessage
	}
	return ""
}

func (x *JoinResponse) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

// Response when a client leaves the chat
type LeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	LamportTime int64 `protobuf:"varint,2,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
}

func (x *LeaveResponse) Reset() {
	*x = LeaveResponse{}
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveResponse) ProtoMessage() {}

func (x *LeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Chitty_Chat_chittychat_chittychat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveResponse.ProtoReflect.Descriptor instead.
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return file_Chitty_Chat_chittychat_chittychat_proto_rawDescGZIP(), []int{4}
}

func (x *LeaveResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LeaveResponse) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

var File_Chitty_Chat_chittychat_chittychat_proto protoreflect.FileDescriptor

var file_Chitty_Chat_chittychat_chittychat_proto_rawDesc = []byte{
	0x0a, 0x27, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x2d, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68,
	0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x69, 0x74, 0x74,
	0x79, 0x63, 0x68, 0x61, 0x74, 0x22, 0x60, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x69, 0x74,
	0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77,
	0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x0d, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x8a, 0x02, 0x0a, 0x0a, 0x43, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x16, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x17, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74,
	0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x30, 0x01, 0x42, 0x18, 0x5a, 0x16, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x2d,
	0x43, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Chitty_Chat_chittychat_chittychat_proto_rawDescOnce sync.Once
	file_Chitty_Chat_chittychat_chittychat_proto_rawDescData = file_Chitty_Chat_chittychat_chittychat_proto_rawDesc
)

func file_Chitty_Chat_chittychat_chittychat_proto_rawDescGZIP() []byte {
	file_Chitty_Chat_chittychat_chittychat_proto_rawDescOnce.Do(func() {
		file_Chitty_Chat_chittychat_chittychat_proto_rawDescData = protoimpl.X.CompressGZIP(file_Chitty_Chat_chittychat_chittychat_proto_rawDescData)
	})
	return file_Chitty_Chat_chittychat_chittychat_proto_rawDescData
}

var file_Chitty_Chat_chittychat_chittychat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Chitty_Chat_chittychat_chittychat_proto_goTypes = []any{
	(*ChatMessage)(nil),     // 0: chittychat.ChatMessage
	(*PublishResponse)(nil), // 1: chittychat.PublishResponse
	(*ClientInfo)(nil),      // 2: chittychat.ClientInfo
	(*JoinResponse)(nil),    // 3: chittychat.JoinResponse
	(*LeaveResponse)(nil),   // 4: chittychat.LeaveResponse
}
var file_Chitty_Chat_chittychat_chittychat_proto_depIdxs = []int32{
	2, // 0: chittychat.ChatMessage.client_info:type_name -> chittychat.ClientInfo
	2, // 1: chittychat.ChittyChat.Join:input_type -> chittychat.ClientInfo
	2, // 2: chittychat.ChittyChat.Leave:input_type -> chittychat.ClientInfo
	0, // 3: chittychat.ChittyChat.PublishMessage:input_type -> chittychat.ChatMessage
	2, // 4: chittychat.ChittyChat.Subscribe:input_type -> chittychat.ClientInfo
	3, // 5: chittychat.ChittyChat.Join:output_type -> chittychat.JoinResponse
	4, // 6: chittychat.ChittyChat.Leave:output_type -> chittychat.LeaveResponse
	1, // 7: chittychat.ChittyChat.PublishMessage:output_type -> chittychat.PublishResponse
	0, // 8: chittychat.ChittyChat.Subscribe:output_type -> chittychat.ChatMessage
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Chitty_Chat_chittychat_chittychat_proto_init() }
func file_Chitty_Chat_chittychat_chittychat_proto_init() {
	if File_Chitty_Chat_chittychat_chittychat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Chitty_Chat_chittychat_chittychat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Chitty_Chat_chittychat_chittychat_proto_goTypes,
		DependencyIndexes: file_Chitty_Chat_chittychat_chittychat_proto_depIdxs,
		MessageInfos:      file_Chitty_Chat_chittychat_chittychat_proto_msgTypes,
	}.Build()
	File_Chitty_Chat_chittychat_chittychat_proto = out.File
	file_Chitty_Chat_chittychat_chittychat_proto_rawDesc = nil
	file_Chitty_Chat_chittychat_chittychat_proto_goTypes = nil
	file_Chitty_Chat_chittychat_chittychat_proto_depIdxs = nil
}
