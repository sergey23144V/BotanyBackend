// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: ecomorphs-entity.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EcomorphEntityServiceClient is the client API for EcomorphEntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcomorphEntityServiceClient interface {
	// Создание новой сущности
	InsertEcomorphEntity(ctx context.Context, in *InputEcomorphsEntity, opts ...grpc.CallOption) (*EcomorphsEntity, error)
	// Получение сущности по заголовку
	GetEcomorphEntityByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EcomorphsEntity, error)
	// Обновление сущности по заголовку
	UpdateEcomorphEntity(ctx context.Context, in *InputEcomorphsEntity, opts ...grpc.CallOption) (*EcomorphsEntity, error)
	// Удаление сущности по заголовку
	DeleteEcomorphEntityByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	// Получение списка всех сущностей
	GetAllEcomorphEntity(ctx context.Context, in *EcomorphsEntityListRequest, opts ...grpc.CallOption) (*EcomorphsEntityList, error)
}

type ecomorphEntityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEcomorphEntityServiceClient(cc grpc.ClientConnInterface) EcomorphEntityServiceClient {
	return &ecomorphEntityServiceClient{cc}
}

func (c *ecomorphEntityServiceClient) InsertEcomorphEntity(ctx context.Context, in *InputEcomorphsEntity, opts ...grpc.CallOption) (*EcomorphsEntity, error) {
	out := new(EcomorphsEntity)
	err := c.cc.Invoke(ctx, "/botany.EcomorphEntityService/InsertEcomorphEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomorphEntityServiceClient) GetEcomorphEntityByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EcomorphsEntity, error) {
	out := new(EcomorphsEntity)
	err := c.cc.Invoke(ctx, "/botany.EcomorphEntityService/GetEcomorphEntityByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomorphEntityServiceClient) UpdateEcomorphEntity(ctx context.Context, in *InputEcomorphsEntity, opts ...grpc.CallOption) (*EcomorphsEntity, error) {
	out := new(EcomorphsEntity)
	err := c.cc.Invoke(ctx, "/botany.EcomorphEntityService/UpdateEcomorphEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomorphEntityServiceClient) DeleteEcomorphEntityByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/botany.EcomorphEntityService/DeleteEcomorphEntityByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecomorphEntityServiceClient) GetAllEcomorphEntity(ctx context.Context, in *EcomorphsEntityListRequest, opts ...grpc.CallOption) (*EcomorphsEntityList, error) {
	out := new(EcomorphsEntityList)
	err := c.cc.Invoke(ctx, "/botany.EcomorphEntityService/GetAllEcomorphEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcomorphEntityServiceServer is the server API for EcomorphEntityService service.
// All implementations must embed UnimplementedEcomorphEntityServiceServer
// for forward compatibility
type EcomorphEntityServiceServer interface {
	// Создание новой сущности
	InsertEcomorphEntity(context.Context, *InputEcomorphsEntity) (*EcomorphsEntity, error)
	// Получение сущности по заголовку
	GetEcomorphEntityByID(context.Context, *IdRequest) (*EcomorphsEntity, error)
	// Обновление сущности по заголовку
	UpdateEcomorphEntity(context.Context, *InputEcomorphsEntity) (*EcomorphsEntity, error)
	// Удаление сущности по заголовку
	DeleteEcomorphEntityByID(context.Context, *IdRequest) (*BoolResponse, error)
	// Получение списка всех сущностей
	GetAllEcomorphEntity(context.Context, *EcomorphsEntityListRequest) (*EcomorphsEntityList, error)
	MustEmbedUnimplementedEcomorphEntityServiceServer()
}

// UnimplementedEcomorphEntityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEcomorphEntityServiceServer struct {
}

func (UnimplementedEcomorphEntityServiceServer) InsertEcomorphEntity(context.Context, *InputEcomorphsEntity) (*EcomorphsEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEcomorphEntity not implemented")
}
func (UnimplementedEcomorphEntityServiceServer) GetEcomorphEntityByID(context.Context, *IdRequest) (*EcomorphsEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcomorphEntityByID not implemented")
}
func (UnimplementedEcomorphEntityServiceServer) UpdateEcomorphEntity(context.Context, *InputEcomorphsEntity) (*EcomorphsEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEcomorphEntity not implemented")
}
func (UnimplementedEcomorphEntityServiceServer) DeleteEcomorphEntityByID(context.Context, *IdRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEcomorphEntityByID not implemented")
}
func (UnimplementedEcomorphEntityServiceServer) GetAllEcomorphEntity(context.Context, *EcomorphsEntityListRequest) (*EcomorphsEntityList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEcomorphEntity not implemented")
}
func (UnimplementedEcomorphEntityServiceServer) MustEmbedUnimplementedEcomorphEntityServiceServer() {}

// UnsafeEcomorphEntityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcomorphEntityServiceServer will
// result in compilation errors.
type UnsafeEcomorphEntityServiceServer interface {
	MustEmbedUnimplementedEcomorphEntityServiceServer()
}

func RegisterEcomorphEntityServiceServer(s grpc.ServiceRegistrar, srv EcomorphEntityServiceServer) {
	s.RegisterService(&EcomorphEntityService_ServiceDesc, srv)
}

func _EcomorphEntityService_InsertEcomorphEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputEcomorphsEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomorphEntityServiceServer).InsertEcomorphEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.EcomorphEntityService/InsertEcomorphEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomorphEntityServiceServer).InsertEcomorphEntity(ctx, req.(*InputEcomorphsEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcomorphEntityService_GetEcomorphEntityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomorphEntityServiceServer).GetEcomorphEntityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.EcomorphEntityService/GetEcomorphEntityByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomorphEntityServiceServer).GetEcomorphEntityByID(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcomorphEntityService_UpdateEcomorphEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputEcomorphsEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomorphEntityServiceServer).UpdateEcomorphEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.EcomorphEntityService/UpdateEcomorphEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomorphEntityServiceServer).UpdateEcomorphEntity(ctx, req.(*InputEcomorphsEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcomorphEntityService_DeleteEcomorphEntityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomorphEntityServiceServer).DeleteEcomorphEntityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.EcomorphEntityService/DeleteEcomorphEntityByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomorphEntityServiceServer).DeleteEcomorphEntityByID(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcomorphEntityService_GetAllEcomorphEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EcomorphsEntityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcomorphEntityServiceServer).GetAllEcomorphEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.EcomorphEntityService/GetAllEcomorphEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcomorphEntityServiceServer).GetAllEcomorphEntity(ctx, req.(*EcomorphsEntityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EcomorphEntityService_ServiceDesc is the grpc.ServiceDesc for EcomorphEntityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EcomorphEntityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "botany.EcomorphEntityService",
	HandlerType: (*EcomorphEntityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertEcomorphEntity",
			Handler:    _EcomorphEntityService_InsertEcomorphEntity_Handler,
		},
		{
			MethodName: "GetEcomorphEntityByID",
			Handler:    _EcomorphEntityService_GetEcomorphEntityByID_Handler,
		},
		{
			MethodName: "UpdateEcomorphEntity",
			Handler:    _EcomorphEntityService_UpdateEcomorphEntity_Handler,
		},
		{
			MethodName: "DeleteEcomorphEntityByID",
			Handler:    _EcomorphEntityService_DeleteEcomorphEntityByID_Handler,
		},
		{
			MethodName: "GetAllEcomorphEntity",
			Handler:    _EcomorphEntityService_GetAllEcomorphEntity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecomorphs-entity.proto",
}
