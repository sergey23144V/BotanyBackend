// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: transect.proto

package api

import (
	resource "github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type Transect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *resource.Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Covered         int32                  `protobuf:"varint,3,opt,name=covered,proto3" json:"covered,omitempty"`
	Rating          int32                  `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Square          int32                  `protobuf:"varint,5,opt,name=square,proto3" json:"square,omitempty"`
	SquareTrialSite int32                  `protobuf:"varint,6,opt,name=squareTrialSite,proto3" json:"squareTrialSite,omitempty"`
	CountTypes      int32                  `protobuf:"varint,7,opt,name=count_types,json=countTypes,proto3" json:"count_types,omitempty"`
	TrialSite       []*TrialSite           `protobuf:"bytes,8,rep,name=trialSite,proto3" json:"trialSite,omitempty"`
	Dominant        *TypePlant             `protobuf:"bytes,9,opt,name=dominant,proto3" json:"dominant,omitempty"`
	SubDominant     *TypePlant             `protobuf:"bytes,10,opt,name=subDominant,proto3" json:"subDominant,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Img             *Img                   `protobuf:"bytes,14,opt,name=img,proto3" json:"img,omitempty"`
	UserId          *resource.Identifier   `protobuf:"bytes,15,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Transect) Reset() {
	*x = Transect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transect) ProtoMessage() {}

func (x *Transect) ProtoReflect() protoreflect.Message {
	mi := &file_transect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transect.ProtoReflect.Descriptor instead.
func (*Transect) Descriptor() ([]byte, []int) {
	return file_transect_proto_rawDescGZIP(), []int{0}
}

func (x *Transect) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Transect) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Transect) GetCovered() int32 {
	if x != nil {
		return x.Covered
	}
	return 0
}

func (x *Transect) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Transect) GetSquare() int32 {
	if x != nil {
		return x.Square
	}
	return 0
}

func (x *Transect) GetSquareTrialSite() int32 {
	if x != nil {
		return x.SquareTrialSite
	}
	return 0
}

func (x *Transect) GetCountTypes() int32 {
	if x != nil {
		return x.CountTypes
	}
	return 0
}

func (x *Transect) GetTrialSite() []*TrialSite {
	if x != nil {
		return x.TrialSite
	}
	return nil
}

func (x *Transect) GetDominant() *TypePlant {
	if x != nil {
		return x.Dominant
	}
	return nil
}

func (x *Transect) GetSubDominant() *TypePlant {
	if x != nil {
		return x.SubDominant
	}
	return nil
}

func (x *Transect) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Transect) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Transect) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Transect) GetImg() *Img {
	if x != nil {
		return x.Img
	}
	return nil
}

func (x *Transect) GetUserId() *resource.Identifier {
	if x != nil {
		return x.UserId
	}
	return nil
}

type TransectList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *PagesResponse `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	List []*Transect    `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TransectList) Reset() {
	*x = TransectList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransectList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransectList) ProtoMessage() {}

func (x *TransectList) ProtoReflect() protoreflect.Message {
	mi := &file_transect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransectList.ProtoReflect.Descriptor instead.
func (*TransectList) Descriptor() ([]byte, []int) {
	return file_transect_proto_rawDescGZIP(), []int{1}
}

func (x *TransectList) GetPage() *PagesResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *TransectList) GetList() []*Transect {
	if x != nil {
		return x.List
	}
	return nil
}

type InputFormTransectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title           string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Covered         int32        `protobuf:"varint,3,opt,name=covered,proto3" json:"covered,omitempty"`
	Rating          int32        `protobuf:"varint,10,opt,name=rating,proto3" json:"rating,omitempty"`
	Square          int32        `protobuf:"varint,11,opt,name=square,proto3" json:"square,omitempty"`
	SquareTrialSite int32        `protobuf:"varint,12,opt,name=squareTrialSite,proto3" json:"squareTrialSite,omitempty"`
	CountTypes      int32        `protobuf:"varint,4,opt,name=count_types,json=countTypes,proto3" json:"count_types,omitempty"`
	TrialSite       []*TrialSite `protobuf:"bytes,1,rep,name=trialSite,proto3" json:"trialSite,omitempty"`
	Img             *Img         `protobuf:"bytes,14,opt,name=img,proto3" json:"img,omitempty"`
	Dominant        *TypePlant   `protobuf:"bytes,7,opt,name=dominant,proto3" json:"dominant,omitempty"`
	SubDominant     *TypePlant   `protobuf:"bytes,8,opt,name=subDominant,proto3" json:"subDominant,omitempty"`
}

func (x *InputFormTransectRequest) Reset() {
	*x = InputFormTransectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputFormTransectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputFormTransectRequest) ProtoMessage() {}

func (x *InputFormTransectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputFormTransectRequest.ProtoReflect.Descriptor instead.
func (*InputFormTransectRequest) Descriptor() ([]byte, []int) {
	return file_transect_proto_rawDescGZIP(), []int{2}
}

func (x *InputFormTransectRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InputFormTransectRequest) GetCovered() int32 {
	if x != nil {
		return x.Covered
	}
	return 0
}

func (x *InputFormTransectRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *InputFormTransectRequest) GetSquare() int32 {
	if x != nil {
		return x.Square
	}
	return 0
}

func (x *InputFormTransectRequest) GetSquareTrialSite() int32 {
	if x != nil {
		return x.SquareTrialSite
	}
	return 0
}

func (x *InputFormTransectRequest) GetCountTypes() int32 {
	if x != nil {
		return x.CountTypes
	}
	return 0
}

func (x *InputFormTransectRequest) GetTrialSite() []*TrialSite {
	if x != nil {
		return x.TrialSite
	}
	return nil
}

func (x *InputFormTransectRequest) GetImg() *Img {
	if x != nil {
		return x.Img
	}
	return nil
}

func (x *InputFormTransectRequest) GetDominant() *TypePlant {
	if x != nil {
		return x.Dominant
	}
	return nil
}

func (x *InputFormTransectRequest) GetSubDominant() *TypePlant {
	if x != nil {
		return x.SubDominant
	}
	return nil
}

type InputTransectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *resource.Identifier      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Input *InputFormTransectRequest `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *InputTransectRequest) Reset() {
	*x = InputTransectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputTransectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputTransectRequest) ProtoMessage() {}

func (x *InputTransectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputTransectRequest.ProtoReflect.Descriptor instead.
func (*InputTransectRequest) Descriptor() ([]byte, []int) {
	return file_transect_proto_rawDescGZIP(), []int{3}
}

func (x *InputTransectRequest) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *InputTransectRequest) GetInput() *InputFormTransectRequest {
	if x != nil {
		return x.Input
	}
	return nil
}

type FilterTransect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []*resource.Identifier `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	SearchTitle string                 `protobuf:"bytes,2,opt,name=searchTitle,proto3" json:"searchTitle,omitempty"`
	Dominant    *TypePlant             `protobuf:"bytes,7,opt,name=dominant,proto3" json:"dominant,omitempty"`
	SubDominant *TypePlant             `protobuf:"bytes,8,opt,name=subDominant,proto3" json:"subDominant,omitempty"`
}

func (x *FilterTransect) Reset() {
	*x = FilterTransect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transect_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterTransect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterTransect) ProtoMessage() {}

func (x *FilterTransect) ProtoReflect() protoreflect.Message {
	mi := &file_transect_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterTransect.ProtoReflect.Descriptor instead.
func (*FilterTransect) Descriptor() ([]byte, []int) {
	return file_transect_proto_rawDescGZIP(), []int{4}
}

func (x *FilterTransect) GetId() []*resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *FilterTransect) GetSearchTitle() string {
	if x != nil {
		return x.SearchTitle
	}
	return ""
}

func (x *FilterTransect) GetDominant() *TypePlant {
	if x != nil {
		return x.Dominant
	}
	return nil
}

func (x *FilterTransect) GetSubDominant() *TypePlant {
	if x != nil {
		return x.SubDominant
	}
	return nil
}

type TransectListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *PagesRequest   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Filter *FilterTransect `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *TransectListRequest) Reset() {
	*x = TransectListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transect_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransectListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransectListRequest) ProtoMessage() {}

func (x *TransectListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transect_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransectListRequest.ProtoReflect.Descriptor instead.
func (*TransectListRequest) Descriptor() ([]byte, []int) {
	return file_transect_proto_rawDescGZIP(), []int{5}
}

func (x *TransectListRequest) GetPage() *PagesRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *TransectListRequest) GetFilter() *FilterTransect {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_transect_proto protoreflect.FileDescriptor

var file_transect_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x1a, 0x09, 0x69, 0x6d, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x40, 0x76, 0x31, 0x2e,
	0x31, 0x2e, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x61, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f,
	0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x40, 0x76, 0x31, 0x2e, 0x31, 0x2e, 0x34, 0x2f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x2d, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x72, 0x69, 0x61, 0x6c,
	0x2d, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x05, 0x0a, 0x08,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x0e, 0xba, 0xb9, 0x19,
	0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x69,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x09, 0x74, 0x72,
	0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65,
	0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x2a, 0x08, 0x30, 0x01, 0x38, 0x01, 0x40, 0x01, 0x48, 0x01,
	0x52, 0x09, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x64,
	0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x42, 0x0c, 0xba, 0xb9, 0x19, 0x08, 0x22, 0x06, 0x20, 0x01, 0x28, 0x01, 0x38, 0x01, 0x52, 0x08,
	0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x44,
	0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x42, 0x0c, 0xba, 0xb9, 0x19, 0x08, 0x22, 0x06, 0x20, 0x01, 0x28, 0x01, 0x38, 0x01, 0x52, 0x0b,
	0x73, 0x75, 0x62, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x03,
	0x69, 0x6d, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x6f, 0x74, 0x61,
	0x6e, 0x79, 0x2e, 0x49, 0x6d, 0x67, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x22, 0x02, 0x38, 0x01,
	0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x47, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x17, 0xba,
	0xb9, 0x19, 0x13, 0x0a, 0x11, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x7a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x06,
	0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0x5f, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63,
	0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xf9, 0x02, 0x0a, 0x18, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x54, 0x72,
	0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x2f, 0x0a, 0x09, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x69, 0x61,
	0x6c, 0x53, 0x69, 0x74, 0x65, 0x52, 0x09, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6d, 0x67, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12,
	0x2d, 0x0a, 0x08, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x33,
	0x0a, 0x0b, 0x73, 0x75, 0x62, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x44, 0x6f, 0x6d, 0x69, 0x6e,
	0x61, 0x6e, 0x74, 0x22, 0x75, 0x0a, 0x14, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xbd, 0x01, 0x0a, 0x0e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x25, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e,
	0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x69, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x44, 0x6f, 0x6d, 0x69,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x74,
	0x61, 0x6e, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x6f, 0x0a, 0x13, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6f,
	0x74, 0x61, 0x6e, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x65, 0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x32, 0x9f, 0x03, 0x0a, 0x0f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63,
	0x74, 0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63,
	0x74, 0x12, 0x32, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74,
	0x12, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x55, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x65, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x65, 0x63, 0x74, 0x12, 0x48, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x53,
	0x69, 0x74, 0x65, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x2e,
	0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f,
	0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x49, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12,
	0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x74,
	0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x67,
	0x65, 0x79, 0x32, 0x33, 0x31, 0x34, 0x34, 0x56, 0x2f, 0x42, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x2d, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transect_proto_rawDescOnce sync.Once
	file_transect_proto_rawDescData = file_transect_proto_rawDesc
)

func file_transect_proto_rawDescGZIP() []byte {
	file_transect_proto_rawDescOnce.Do(func() {
		file_transect_proto_rawDescData = protoimpl.X.CompressGZIP(file_transect_proto_rawDescData)
	})
	return file_transect_proto_rawDescData
}

var file_transect_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transect_proto_goTypes = []interface{}{
	(*Transect)(nil),                 // 0: botany.Transect
	(*TransectList)(nil),             // 1: botany.TransectList
	(*InputFormTransectRequest)(nil), // 2: botany.InputFormTransectRequest
	(*InputTransectRequest)(nil),     // 3: botany.InputTransectRequest
	(*FilterTransect)(nil),           // 4: botany.FilterTransect
	(*TransectListRequest)(nil),      // 5: botany.TransectListRequest
	(*resource.Identifier)(nil),      // 6: atlas.rpc.Identifier
	(*TrialSite)(nil),                // 7: botany.TrialSite
	(*TypePlant)(nil),                // 8: botany.TypePlant
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(*Img)(nil),                      // 10: botany.Img
	(*PagesResponse)(nil),            // 11: botany.PagesResponse
	(*PagesRequest)(nil),             // 12: botany.PagesRequest
	(*IdRequest)(nil),                // 13: botany.IdRequest
	(*BoolResponse)(nil),             // 14: botany.BoolResponse
}
var file_transect_proto_depIdxs = []int32{
	6,  // 0: botany.Transect.id:type_name -> atlas.rpc.Identifier
	7,  // 1: botany.Transect.trialSite:type_name -> botany.TrialSite
	8,  // 2: botany.Transect.dominant:type_name -> botany.TypePlant
	8,  // 3: botany.Transect.subDominant:type_name -> botany.TypePlant
	9,  // 4: botany.Transect.created_at:type_name -> google.protobuf.Timestamp
	9,  // 5: botany.Transect.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 6: botany.Transect.deleted_at:type_name -> google.protobuf.Timestamp
	10, // 7: botany.Transect.img:type_name -> botany.Img
	6,  // 8: botany.Transect.user_id:type_name -> atlas.rpc.Identifier
	11, // 9: botany.TransectList.page:type_name -> botany.PagesResponse
	0,  // 10: botany.TransectList.list:type_name -> botany.Transect
	7,  // 11: botany.InputFormTransectRequest.trialSite:type_name -> botany.TrialSite
	10, // 12: botany.InputFormTransectRequest.img:type_name -> botany.Img
	8,  // 13: botany.InputFormTransectRequest.dominant:type_name -> botany.TypePlant
	8,  // 14: botany.InputFormTransectRequest.subDominant:type_name -> botany.TypePlant
	6,  // 15: botany.InputTransectRequest.id:type_name -> atlas.rpc.Identifier
	2,  // 16: botany.InputTransectRequest.input:type_name -> botany.InputFormTransectRequest
	6,  // 17: botany.FilterTransect.id:type_name -> atlas.rpc.Identifier
	8,  // 18: botany.FilterTransect.dominant:type_name -> botany.TypePlant
	8,  // 19: botany.FilterTransect.subDominant:type_name -> botany.TypePlant
	12, // 20: botany.TransectListRequest.page:type_name -> botany.PagesRequest
	4,  // 21: botany.TransectListRequest.filter:type_name -> botany.FilterTransect
	3,  // 22: botany.TransectService.CreateTransect:input_type -> botany.InputTransectRequest
	13, // 23: botany.TransectService.GetTransect:input_type -> botany.IdRequest
	3,  // 24: botany.TransectService.UpTransect:input_type -> botany.InputTransectRequest
	3,  // 25: botany.TransectService.AddTrialSiteToTransect:input_type -> botany.InputTransectRequest
	13, // 26: botany.TransectService.DeleteTransect:input_type -> botany.IdRequest
	5,  // 27: botany.TransectService.GetAllTransect:input_type -> botany.TransectListRequest
	0,  // 28: botany.TransectService.CreateTransect:output_type -> botany.Transect
	0,  // 29: botany.TransectService.GetTransect:output_type -> botany.Transect
	0,  // 30: botany.TransectService.UpTransect:output_type -> botany.Transect
	0,  // 31: botany.TransectService.AddTrialSiteToTransect:output_type -> botany.Transect
	14, // 32: botany.TransectService.DeleteTransect:output_type -> botany.BoolResponse
	1,  // 33: botany.TransectService.GetAllTransect:output_type -> botany.TransectList
	28, // [28:34] is the sub-list for method output_type
	22, // [22:28] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_transect_proto_init() }
func file_transect_proto_init() {
	if File_transect_proto != nil {
		return
	}
	file_img_proto_init()
	file_elementary_type_proto_init()
	file_type_plant_proto_init()
	file_trial_site_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transect); i {
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
		file_transect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransectList); i {
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
		file_transect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputFormTransectRequest); i {
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
		file_transect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputTransectRequest); i {
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
		file_transect_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterTransect); i {
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
		file_transect_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransectListRequest); i {
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
			RawDescriptor: file_transect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transect_proto_goTypes,
		DependencyIndexes: file_transect_proto_depIdxs,
		MessageInfos:      file_transect_proto_msgTypes,
	}.Build()
	File_transect_proto = out.File
	file_transect_proto_rawDesc = nil
	file_transect_proto_goTypes = nil
	file_transect_proto_depIdxs = nil
}
