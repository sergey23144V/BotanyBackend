// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: ecomorphs.proto

package ecomorph

import (
	resource "github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ecomorph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *resource.Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Title       string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	UserId      *resource.Identifier   `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Ecomorph) Reset() {
	*x = Ecomorph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecomorphs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ecomorph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ecomorph) ProtoMessage() {}

func (x *Ecomorph) ProtoReflect() protoreflect.Message {
	mi := &file_ecomorphs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ecomorph.ProtoReflect.Descriptor instead.
func (*Ecomorph) Descriptor() ([]byte, []int) {
	return file_ecomorphs_proto_rawDescGZIP(), []int{0}
}

func (x *Ecomorph) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Ecomorph) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ecomorph) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Ecomorph) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Ecomorph) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Ecomorph) GetUserId() *resource.Identifier {
	if x != nil {
		return x.UserId
	}
	return nil
}

type EcomorphsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ecomorph []*Ecomorph `protobuf:"bytes,1,rep,name=ecomorph,proto3" json:"ecomorph,omitempty"`
}

func (x *EcomorphsList) Reset() {
	*x = EcomorphsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecomorphs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcomorphsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcomorphsList) ProtoMessage() {}

func (x *EcomorphsList) ProtoReflect() protoreflect.Message {
	mi := &file_ecomorphs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcomorphsList.ProtoReflect.Descriptor instead.
func (*EcomorphsList) Descriptor() ([]byte, []int) {
	return file_ecomorphs_proto_rawDescGZIP(), []int{1}
}

func (x *EcomorphsList) GetEcomorph() []*Ecomorph {
	if x != nil {
		return x.Ecomorph
	}
	return nil
}

type InputFormEcomorph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InputFormEcomorph) Reset() {
	*x = InputFormEcomorph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecomorphs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputFormEcomorph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputFormEcomorph) ProtoMessage() {}

func (x *InputFormEcomorph) ProtoReflect() protoreflect.Message {
	mi := &file_ecomorphs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputFormEcomorph.ProtoReflect.Descriptor instead.
func (*InputFormEcomorph) Descriptor() ([]byte, []int) {
	return file_ecomorphs_proto_rawDescGZIP(), []int{2}
}

func (x *InputFormEcomorph) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InputFormEcomorph) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type InputEcomorph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload *InputFormEcomorph   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InputEcomorph) Reset() {
	*x = InputEcomorph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecomorphs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputEcomorph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputEcomorph) ProtoMessage() {}

func (x *InputEcomorph) ProtoReflect() protoreflect.Message {
	mi := &file_ecomorphs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputEcomorph.ProtoReflect.Descriptor instead.
func (*InputEcomorph) Descriptor() ([]byte, []int) {
	return file_ecomorphs_proto_rawDescGZIP(), []int{3}
}

func (x *InputEcomorph) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *InputEcomorph) GetPayload() *InputFormEcomorph {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_ecomorphs_proto protoreflect.FileDescriptor

var file_ecomorphs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x1a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x72, 0x6d, 0x40, 0x76, 0x31, 0x2e, 0x31, 0x2e, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x40, 0x76, 0x31, 0x2e, 0x31,
	0x2e, 0x32, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x72, 0x79, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02,
	0x0a, 0x08, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x0e, 0xba,
	0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x47, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x17, 0xba, 0xb9, 0x19, 0x13, 0x0a, 0x11, 0x12,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x7a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01,
	0x22, 0x3d, 0x0a, 0x0d, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x63, 0x6f,
	0x6d, 0x6f, 0x72, 0x70, 0x68, 0x52, 0x08, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x22,
	0x4b, 0x0a, 0x11, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x45, 0x63, 0x6f, 0x6d,
	0x6f, 0x72, 0x70, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x0d,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x12, 0x25, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08,
	0x01, 0x32, 0xce, 0x02, 0x0a, 0x0f, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45,
	0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x1a, 0x10,
	0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68,
	0x12, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e,
	0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x12, 0x39, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x12, 0x15, 0x2e, 0x62, 0x6f, 0x74,
	0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70,
	0x68, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6d, 0x6f,
	0x72, 0x70, 0x68, 0x12, 0x4d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x63, 0x6f,
	0x6d, 0x6f, 0x72, 0x70, 0x68, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61,
	0x6e, 0x79, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62,
	0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72,
	0x70, 0x68, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x63, 0x6f,
	0x6d, 0x6f, 0x72, 0x70, 0x68, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6f,
	0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72,
	0x70, 0x68, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ecomorphs_proto_rawDescOnce sync.Once
	file_ecomorphs_proto_rawDescData = file_ecomorphs_proto_rawDesc
)

func file_ecomorphs_proto_rawDescGZIP() []byte {
	file_ecomorphs_proto_rawDescOnce.Do(func() {
		file_ecomorphs_proto_rawDescData = protoimpl.X.CompressGZIP(file_ecomorphs_proto_rawDescData)
	})
	return file_ecomorphs_proto_rawDescData
}

var file_ecomorphs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ecomorphs_proto_goTypes = []interface{}{
	(*Ecomorph)(nil),              // 0: botany.Ecomorph
	(*EcomorphsList)(nil),         // 1: botany.EcomorphsList
	(*InputFormEcomorph)(nil),     // 2: botany.InputFormEcomorph
	(*InputEcomorph)(nil),         // 3: botany.InputEcomorph
	(*resource.Identifier)(nil),   // 4: atlas.rpc.Identifier
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*api.IdRequest)(nil),         // 6: botany.IdRequest
	(*api.EmptyRequest)(nil),      // 7: botany.EmptyRequest
	(*api.BoolResponse)(nil),      // 8: botany.BoolResponse
}
var file_ecomorphs_proto_depIdxs = []int32{
	4,  // 0: botany.Ecomorph.id:type_name -> atlas.rpc.Identifier
	5,  // 1: botany.Ecomorph.created_at:type_name -> google.protobuf.Timestamp
	5,  // 2: botany.Ecomorph.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 3: botany.Ecomorph.user_id:type_name -> atlas.rpc.Identifier
	0,  // 4: botany.EcomorphsList.ecomorph:type_name -> botany.Ecomorph
	4,  // 5: botany.InputEcomorph.id:type_name -> atlas.rpc.Identifier
	2,  // 6: botany.InputEcomorph.payload:type_name -> botany.InputFormEcomorph
	3,  // 7: botany.EcomorphService.InsertEcomorph:input_type -> botany.InputEcomorph
	6,  // 8: botany.EcomorphService.GetEcomorphById:input_type -> botany.IdRequest
	3,  // 9: botany.EcomorphService.UpdateEcomorph:input_type -> botany.InputEcomorph
	6,  // 10: botany.EcomorphService.DeleteEcomorphById:input_type -> botany.IdRequest
	7,  // 11: botany.EcomorphService.GetListEcomorph:input_type -> botany.EmptyRequest
	0,  // 12: botany.EcomorphService.InsertEcomorph:output_type -> botany.Ecomorph
	0,  // 13: botany.EcomorphService.GetEcomorphById:output_type -> botany.Ecomorph
	0,  // 14: botany.EcomorphService.UpdateEcomorph:output_type -> botany.Ecomorph
	8,  // 15: botany.EcomorphService.DeleteEcomorphById:output_type -> botany.BoolResponse
	1,  // 16: botany.EcomorphService.GetListEcomorph:output_type -> botany.EcomorphsList
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_ecomorphs_proto_init() }
func file_ecomorphs_proto_init() {
	if File_ecomorphs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ecomorphs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ecomorph); i {
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
		file_ecomorphs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcomorphsList); i {
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
		file_ecomorphs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputFormEcomorph); i {
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
		file_ecomorphs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputEcomorph); i {
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
			RawDescriptor: file_ecomorphs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ecomorphs_proto_goTypes,
		DependencyIndexes: file_ecomorphs_proto_depIdxs,
		MessageInfos:      file_ecomorphs_proto_msgTypes,
	}.Build()
	File_ecomorphs_proto = out.File
	file_ecomorphs_proto_rawDesc = nil
	file_ecomorphs_proto_goTypes = nil
	file_ecomorphs_proto_depIdxs = nil
}
