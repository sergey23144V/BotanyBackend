// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: analysis.proto

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

type TypeAnalysis int32

const (
	TypeAnalysis_TypeAnalysisPlant    TypeAnalysis = 0
	TypeAnalysis_TypeAnalysisTransect TypeAnalysis = 1
)

// Enum value maps for TypeAnalysis.
var (
	TypeAnalysis_name = map[int32]string{
		0: "TypeAnalysisPlant",
		1: "TypeAnalysisTransect",
	}
	TypeAnalysis_value = map[string]int32{
		"TypeAnalysisPlant":    0,
		"TypeAnalysisTransect": 1,
	}
)

func (x TypeAnalysis) Enum() *TypeAnalysis {
	p := new(TypeAnalysis)
	*p = x
	return p
}

func (x TypeAnalysis) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeAnalysis) Descriptor() protoreflect.EnumDescriptor {
	return file_analysis_proto_enumTypes[0].Descriptor()
}

func (TypeAnalysis) Type() protoreflect.EnumType {
	return &file_analysis_proto_enumTypes[0]
}

func (x TypeAnalysis) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeAnalysis.Descriptor instead.
func (TypeAnalysis) EnumDescriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{0}
}

type Analysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *resource.Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	TypeAnalysis TypeAnalysis           `protobuf:"varint,5,opt,name=typeAnalysis,proto3,enum=botany.TypeAnalysis" json:"typeAnalysis,omitempty"`
	Title        string                 `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Transect     *Transect              `protobuf:"bytes,7,opt,name=transect,proto3" json:"transect,omitempty"`
	Path         string                 `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	UserId       *resource.Identifier   `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Analysis) Reset() {
	*x = Analysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Analysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Analysis) ProtoMessage() {}

func (x *Analysis) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Analysis.ProtoReflect.Descriptor instead.
func (*Analysis) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *Analysis) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Analysis) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Analysis) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Analysis) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Analysis) GetTypeAnalysis() TypeAnalysis {
	if x != nil {
		return x.TypeAnalysis
	}
	return TypeAnalysis_TypeAnalysisPlant
}

func (x *Analysis) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Analysis) GetTransect() *Transect {
	if x != nil {
		return x.Transect
	}
	return nil
}

func (x *Analysis) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Analysis) GetUserId() *resource.Identifier {
	if x != nil {
		return x.UserId
	}
	return nil
}

type AnalysisList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *PagesResponse `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	List []*Analysis    `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AnalysisList) Reset() {
	*x = AnalysisList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisList) ProtoMessage() {}

func (x *AnalysisList) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisList.ProtoReflect.Descriptor instead.
func (*AnalysisList) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *AnalysisList) GetPage() *PagesResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *AnalysisList) GetList() []*Analysis {
	if x != nil {
		return x.List
	}
	return nil
}

type InputCreateAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string       `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	TypeAnalysis TypeAnalysis `protobuf:"varint,2,opt,name=typeAnalysis,proto3,enum=botany.TypeAnalysis" json:"typeAnalysis,omitempty"`
	Transect     *Transect    `protobuf:"bytes,3,opt,name=transect,proto3" json:"transect,omitempty"`
	Ecomorph     []*Ecomorph  `protobuf:"bytes,4,rep,name=ecomorph,proto3" json:"ecomorph,omitempty"`
}

func (x *InputCreateAnalysis) Reset() {
	*x = InputCreateAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputCreateAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputCreateAnalysis) ProtoMessage() {}

func (x *InputCreateAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputCreateAnalysis.ProtoReflect.Descriptor instead.
func (*InputCreateAnalysis) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{2}
}

func (x *InputCreateAnalysis) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InputCreateAnalysis) GetTypeAnalysis() TypeAnalysis {
	if x != nil {
		return x.TypeAnalysis
	}
	return TypeAnalysis_TypeAnalysisPlant
}

func (x *InputCreateAnalysis) GetTransect() *Transect {
	if x != nil {
		return x.Transect
	}
	return nil
}

func (x *InputCreateAnalysis) GetEcomorph() []*Ecomorph {
	if x != nil {
		return x.Ecomorph
	}
	return nil
}

type InputUpdateAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Ecomorph []*Ecomorph          `protobuf:"bytes,3,rep,name=ecomorph,proto3" json:"ecomorph,omitempty"`
}

func (x *InputUpdateAnalysis) Reset() {
	*x = InputUpdateAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputUpdateAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputUpdateAnalysis) ProtoMessage() {}

func (x *InputUpdateAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputUpdateAnalysis.ProtoReflect.Descriptor instead.
func (*InputUpdateAnalysis) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{3}
}

func (x *InputUpdateAnalysis) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *InputUpdateAnalysis) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InputUpdateAnalysis) GetEcomorph() []*Ecomorph {
	if x != nil {
		return x.Ecomorph
	}
	return nil
}

type FilterAnalysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           []*resource.Identifier `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	SearchTitle  string                 `protobuf:"bytes,2,opt,name=searchTitle,proto3" json:"searchTitle,omitempty"`
	TypeAnalysis TypeAnalysis           `protobuf:"varint,3,opt,name=typeAnalysis,proto3,enum=botany.TypeAnalysis" json:"typeAnalysis,omitempty"`
	Transect     []*Transect            `protobuf:"bytes,4,rep,name=transect,proto3" json:"transect,omitempty"`
	Ecomorph     []*Ecomorph            `protobuf:"bytes,5,rep,name=ecomorph,proto3" json:"ecomorph,omitempty"`
}

func (x *FilterAnalysis) Reset() {
	*x = FilterAnalysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterAnalysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterAnalysis) ProtoMessage() {}

func (x *FilterAnalysis) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterAnalysis.ProtoReflect.Descriptor instead.
func (*FilterAnalysis) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{4}
}

func (x *FilterAnalysis) GetId() []*resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *FilterAnalysis) GetSearchTitle() string {
	if x != nil {
		return x.SearchTitle
	}
	return ""
}

func (x *FilterAnalysis) GetTypeAnalysis() TypeAnalysis {
	if x != nil {
		return x.TypeAnalysis
	}
	return TypeAnalysis_TypeAnalysisPlant
}

func (x *FilterAnalysis) GetTransect() []*Transect {
	if x != nil {
		return x.Transect
	}
	return nil
}

func (x *FilterAnalysis) GetEcomorph() []*Ecomorph {
	if x != nil {
		return x.Ecomorph
	}
	return nil
}

type AnalysisListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *PagesRequest   `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Filter *FilterAnalysis `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *AnalysisListRequest) Reset() {
	*x = AnalysisListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analysis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisListRequest) ProtoMessage() {}

func (x *AnalysisListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analysis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisListRequest.ProtoReflect.Descriptor instead.
func (*AnalysisListRequest) Descriptor() ([]byte, []int) {
	return file_analysis_proto_rawDescGZIP(), []int{5}
}

func (x *AnalysisListRequest) GetPage() *PagesRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *AnalysisListRequest) GetFilter() *FilterAnalysis {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_analysis_proto protoreflect.FileDescriptor

var file_analysis_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x1a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72,
	0x6d, 0x40, 0x76, 0x31, 0x2e, 0x31, 0x2e, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x40, 0x76, 0x31, 0x2e, 0x31, 0x2e,
	0x34, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72,
	0x79, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x65, 0x63,
	0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x03,
	0x0a, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x35, 0x0a, 0x02, 0x69, 0x64,
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
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e,
	0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x42, 0x0c, 0xba, 0xb9, 0x19, 0x08, 0x22, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x38, 0x01, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x47, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x17, 0xba, 0xb9, 0x19, 0x13,
	0x0a, 0x11, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x7a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x08, 0x01, 0x22, 0x5f, 0x0a, 0x0c, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62,
	0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0xc1, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e,
	0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x2c, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74,
	0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x65, 0x63,
	0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62,
	0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x52, 0x08,
	0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x22, 0x80, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x12, 0x25, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2c, 0x0a,
	0x08, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70,
	0x68, 0x52, 0x08, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x22, 0xef, 0x01, 0x0a, 0x0e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x25,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x12, 0x2c, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x65, 0x63, 0x74, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x12,
	0x2c, 0x0a, 0x08, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x45, 0x63, 0x6f, 0x6d, 0x6f,
	0x72, 0x70, 0x68, 0x52, 0x08, 0x65, 0x63, 0x6f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x22, 0x6f, 0x0a,
	0x13, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2e,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2a, 0x3f,
	0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x15,
	0x0a, 0x11, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x65, 0x63, 0x74, 0x10, 0x01, 0x32,
	0xe1, 0x02, 0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x74, 0x61,
	0x6e, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61,
	0x6e, 0x79, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62,
	0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x74,
	0x61, 0x6e, 0x79, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x65, 0x72, 0x67, 0x65, 0x79, 0x32, 0x33, 0x31, 0x34, 0x34, 0x56, 0x2f, 0x42,
	0x6f, 0x74, 0x61, 0x6e, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analysis_proto_rawDescOnce sync.Once
	file_analysis_proto_rawDescData = file_analysis_proto_rawDesc
)

func file_analysis_proto_rawDescGZIP() []byte {
	file_analysis_proto_rawDescOnce.Do(func() {
		file_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_analysis_proto_rawDescData)
	})
	return file_analysis_proto_rawDescData
}

var file_analysis_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_analysis_proto_goTypes = []interface{}{
	(TypeAnalysis)(0),             // 0: botany.TypeAnalysis
	(*Analysis)(nil),              // 1: botany.Analysis
	(*AnalysisList)(nil),          // 2: botany.AnalysisList
	(*InputCreateAnalysis)(nil),   // 3: botany.InputCreateAnalysis
	(*InputUpdateAnalysis)(nil),   // 4: botany.InputUpdateAnalysis
	(*FilterAnalysis)(nil),        // 5: botany.FilterAnalysis
	(*AnalysisListRequest)(nil),   // 6: botany.AnalysisListRequest
	(*resource.Identifier)(nil),   // 7: atlas.rpc.Identifier
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Transect)(nil),              // 9: botany.Transect
	(*PagesResponse)(nil),         // 10: botany.PagesResponse
	(*Ecomorph)(nil),              // 11: botany.Ecomorph
	(*PagesRequest)(nil),          // 12: botany.PagesRequest
	(*IdRequest)(nil),             // 13: botany.IdRequest
	(*BoolResponse)(nil),          // 14: botany.BoolResponse
}
var file_analysis_proto_depIdxs = []int32{
	7,  // 0: botany.Analysis.id:type_name -> atlas.rpc.Identifier
	8,  // 1: botany.Analysis.created_at:type_name -> google.protobuf.Timestamp
	8,  // 2: botany.Analysis.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 3: botany.Analysis.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 4: botany.Analysis.typeAnalysis:type_name -> botany.TypeAnalysis
	9,  // 5: botany.Analysis.transect:type_name -> botany.Transect
	7,  // 6: botany.Analysis.user_id:type_name -> atlas.rpc.Identifier
	10, // 7: botany.AnalysisList.page:type_name -> botany.PagesResponse
	1,  // 8: botany.AnalysisList.list:type_name -> botany.Analysis
	0,  // 9: botany.InputCreateAnalysis.typeAnalysis:type_name -> botany.TypeAnalysis
	9,  // 10: botany.InputCreateAnalysis.transect:type_name -> botany.Transect
	11, // 11: botany.InputCreateAnalysis.ecomorph:type_name -> botany.Ecomorph
	7,  // 12: botany.InputUpdateAnalysis.id:type_name -> atlas.rpc.Identifier
	11, // 13: botany.InputUpdateAnalysis.ecomorph:type_name -> botany.Ecomorph
	7,  // 14: botany.FilterAnalysis.id:type_name -> atlas.rpc.Identifier
	0,  // 15: botany.FilterAnalysis.typeAnalysis:type_name -> botany.TypeAnalysis
	9,  // 16: botany.FilterAnalysis.transect:type_name -> botany.Transect
	11, // 17: botany.FilterAnalysis.ecomorph:type_name -> botany.Ecomorph
	12, // 18: botany.AnalysisListRequest.page:type_name -> botany.PagesRequest
	5,  // 19: botany.AnalysisListRequest.filter:type_name -> botany.FilterAnalysis
	3,  // 20: botany.AnalysisService.CreatAnalysis:input_type -> botany.InputCreateAnalysis
	4,  // 21: botany.AnalysisService.RepeatedAnalysis:input_type -> botany.InputUpdateAnalysis
	13, // 22: botany.AnalysisService.GetAnalysis:input_type -> botany.IdRequest
	6,  // 23: botany.AnalysisService.GetListAnalysis:input_type -> botany.AnalysisListRequest
	13, // 24: botany.AnalysisService.DeleteAnalysis:input_type -> botany.IdRequest
	1,  // 25: botany.AnalysisService.CreatAnalysis:output_type -> botany.Analysis
	1,  // 26: botany.AnalysisService.RepeatedAnalysis:output_type -> botany.Analysis
	1,  // 27: botany.AnalysisService.GetAnalysis:output_type -> botany.Analysis
	2,  // 28: botany.AnalysisService.GetListAnalysis:output_type -> botany.AnalysisList
	14, // 29: botany.AnalysisService.DeleteAnalysis:output_type -> botany.BoolResponse
	25, // [25:30] is the sub-list for method output_type
	20, // [20:25] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_analysis_proto_init() }
func file_analysis_proto_init() {
	if File_analysis_proto != nil {
		return
	}
	file_transect_proto_init()
	file_elementary_type_proto_init()
	file_ecomorphs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Analysis); i {
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
		file_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisList); i {
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
		file_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputCreateAnalysis); i {
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
		file_analysis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputUpdateAnalysis); i {
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
		file_analysis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterAnalysis); i {
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
		file_analysis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisListRequest); i {
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
			RawDescriptor: file_analysis_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analysis_proto_goTypes,
		DependencyIndexes: file_analysis_proto_depIdxs,
		EnumInfos:         file_analysis_proto_enumTypes,
		MessageInfos:      file_analysis_proto_msgTypes,
	}.Build()
	File_analysis_proto = out.File
	file_analysis_proto_rawDesc = nil
	file_analysis_proto_goTypes = nil
	file_analysis_proto_depIdxs = nil
}
