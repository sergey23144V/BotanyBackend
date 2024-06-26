// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: img.proto

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

type Img struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *resource.Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path      string                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	UserId    *resource.Identifier   `protobuf:"bytes,12,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Img) Reset() {
	*x = Img{}
	if protoimpl.UnsafeEnabled {
		mi := &file_img_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Img) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Img) ProtoMessage() {}

func (x *Img) ProtoReflect() protoreflect.Message {
	mi := &file_img_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Img.ProtoReflect.Descriptor instead.
func (*Img) Descriptor() ([]byte, []int) {
	return file_img_proto_rawDescGZIP(), []int{0}
}

func (x *Img) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Img) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Img) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Img) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Img) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Img) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Img) GetUserId() *resource.Identifier {
	if x != nil {
		return x.UserId
	}
	return nil
}

type ImgList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *PagesResponse `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	List []*Img         `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ImgList) Reset() {
	*x = ImgList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_img_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgList) ProtoMessage() {}

func (x *ImgList) ProtoReflect() protoreflect.Message {
	mi := &file_img_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgList.ProtoReflect.Descriptor instead.
func (*ImgList) Descriptor() ([]byte, []int) {
	return file_img_proto_rawDescGZIP(), []int{1}
}

func (x *ImgList) GetPage() *PagesResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ImgList) GetList() []*Img {
	if x != nil {
		return x.List
	}
	return nil
}

var File_img_proto protoreflect.FileDescriptor

var file_img_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x6d, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x6f, 0x74,
	0x61, 0x6e, 0x79, 0x1a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
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
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02, 0x0a, 0x03, 0x49, 0x6d, 0x67, 0x12,
	0x35, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x4e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x13, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x52, 0x0b, 0x73, 0x6f,
	0x66, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x17, 0xba, 0xb9,
	0x19, 0x13, 0x0a, 0x11, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x7a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x06, 0xba,
	0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0x55, 0x0a, 0x07, 0x49, 0x6d, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x6f, 0x74, 0x61,
	0x6e, 0x79, 0x2e, 0x49, 0x6d, 0x67, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x73, 0x0a, 0x0a,
	0x49, 0x6d, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x67, 0x42, 0x79, 0x49, 0x44, 0x12, 0x11, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e,
	0x79, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x62, 0x6f,
	0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6d, 0x67, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x67, 0x12, 0x14, 0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e,
	0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x62, 0x6f, 0x74, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6d, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x65, 0x72, 0x67, 0x65, 0x79, 0x32, 0x33, 0x31, 0x34, 0x34, 0x56, 0x2f, 0x42, 0x6f, 0x74,
	0x61, 0x6e, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x2f, 0x67, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_img_proto_rawDescOnce sync.Once
	file_img_proto_rawDescData = file_img_proto_rawDesc
)

func file_img_proto_rawDescGZIP() []byte {
	file_img_proto_rawDescOnce.Do(func() {
		file_img_proto_rawDescData = protoimpl.X.CompressGZIP(file_img_proto_rawDescData)
	})
	return file_img_proto_rawDescData
}

var file_img_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_img_proto_goTypes = []interface{}{
	(*Img)(nil),                   // 0: botany.Img
	(*ImgList)(nil),               // 1: botany.ImgList
	(*resource.Identifier)(nil),   // 2: atlas.rpc.Identifier
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*PagesResponse)(nil),         // 4: botany.PagesResponse
	(*IdRequest)(nil),             // 5: botany.IdRequest
	(*PagesRequest)(nil),          // 6: botany.PagesRequest
}
var file_img_proto_depIdxs = []int32{
	2, // 0: botany.Img.id:type_name -> atlas.rpc.Identifier
	3, // 1: botany.Img.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: botany.Img.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: botany.Img.deleted_at:type_name -> google.protobuf.Timestamp
	2, // 4: botany.Img.user_id:type_name -> atlas.rpc.Identifier
	4, // 5: botany.ImgList.page:type_name -> botany.PagesResponse
	0, // 6: botany.ImgList.list:type_name -> botany.Img
	5, // 7: botany.ImgService.GetImgByID:input_type -> botany.IdRequest
	6, // 8: botany.ImgService.GetListImg:input_type -> botany.PagesRequest
	0, // 9: botany.ImgService.GetImgByID:output_type -> botany.Img
	1, // 10: botany.ImgService.GetListImg:output_type -> botany.ImgList
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_img_proto_init() }
func file_img_proto_init() {
	if File_img_proto != nil {
		return
	}
	file_elementary_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_img_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Img); i {
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
		file_img_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgList); i {
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
			RawDescriptor: file_img_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_img_proto_goTypes,
		DependencyIndexes: file_img_proto_depIdxs,
		MessageInfos:      file_img_proto_msgTypes,
	}.Build()
	File_img_proto = out.File
	file_img_proto_rawDesc = nil
	file_img_proto_goTypes = nil
	file_img_proto_depIdxs = nil
}
