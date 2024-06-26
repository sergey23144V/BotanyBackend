// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: transect.proto

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

// TransectServiceClient is the client API for TransectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransectServiceClient interface {
	CreateTransect(ctx context.Context, in *InputTransectRequest, opts ...grpc.CallOption) (*Transect, error)
	// Получение сущности по
	GetTransect(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Transect, error)
	// Обновление сущности по id
	UpTransect(ctx context.Context, in *InputTransectRequest, opts ...grpc.CallOption) (*Transect, error)
	// Обновление сущности по id
	AddTrialSiteToTransect(ctx context.Context, in *InputTransectRequest, opts ...grpc.CallOption) (*Transect, error)
	// Удаление сущности по заголовку
	DeleteTransect(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*BoolResponse, error)
	// Получение списка всех сущностей
	GetAllTransect(ctx context.Context, in *TransectListRequest, opts ...grpc.CallOption) (*TransectList, error)
}

type transectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransectServiceClient(cc grpc.ClientConnInterface) TransectServiceClient {
	return &transectServiceClient{cc}
}

func (c *transectServiceClient) CreateTransect(ctx context.Context, in *InputTransectRequest, opts ...grpc.CallOption) (*Transect, error) {
	out := new(Transect)
	err := c.cc.Invoke(ctx, "/botany.TransectService/CreateTransect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transectServiceClient) GetTransect(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Transect, error) {
	out := new(Transect)
	err := c.cc.Invoke(ctx, "/botany.TransectService/GetTransect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transectServiceClient) UpTransect(ctx context.Context, in *InputTransectRequest, opts ...grpc.CallOption) (*Transect, error) {
	out := new(Transect)
	err := c.cc.Invoke(ctx, "/botany.TransectService/UpTransect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transectServiceClient) AddTrialSiteToTransect(ctx context.Context, in *InputTransectRequest, opts ...grpc.CallOption) (*Transect, error) {
	out := new(Transect)
	err := c.cc.Invoke(ctx, "/botany.TransectService/AddTrialSiteToTransect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transectServiceClient) DeleteTransect(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := c.cc.Invoke(ctx, "/botany.TransectService/DeleteTransect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transectServiceClient) GetAllTransect(ctx context.Context, in *TransectListRequest, opts ...grpc.CallOption) (*TransectList, error) {
	out := new(TransectList)
	err := c.cc.Invoke(ctx, "/botany.TransectService/GetAllTransect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransectServiceServer is the server API for TransectService service.
// All implementations must embed UnimplementedTransectServiceServer
// for forward compatibility
type TransectServiceServer interface {
	CreateTransect(context.Context, *InputTransectRequest) (*Transect, error)
	// Получение сущности по
	GetTransect(context.Context, *IdRequest) (*Transect, error)
	// Обновление сущности по id
	UpTransect(context.Context, *InputTransectRequest) (*Transect, error)
	// Обновление сущности по id
	AddTrialSiteToTransect(context.Context, *InputTransectRequest) (*Transect, error)
	// Удаление сущности по заголовку
	DeleteTransect(context.Context, *IdRequest) (*BoolResponse, error)
	// Получение списка всех сущностей
	GetAllTransect(context.Context, *TransectListRequest) (*TransectList, error)
	MustEmbedUnimplementedTransectServiceServer()
}

// UnimplementedTransectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransectServiceServer struct {
}

func (UnimplementedTransectServiceServer) CreateTransect(context.Context, *InputTransectRequest) (*Transect, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransect not implemented")
}
func (UnimplementedTransectServiceServer) GetTransect(context.Context, *IdRequest) (*Transect, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransect not implemented")
}
func (UnimplementedTransectServiceServer) UpTransect(context.Context, *InputTransectRequest) (*Transect, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpTransect not implemented")
}
func (UnimplementedTransectServiceServer) AddTrialSiteToTransect(context.Context, *InputTransectRequest) (*Transect, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrialSiteToTransect not implemented")
}
func (UnimplementedTransectServiceServer) DeleteTransect(context.Context, *IdRequest) (*BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransect not implemented")
}
func (UnimplementedTransectServiceServer) GetAllTransect(context.Context, *TransectListRequest) (*TransectList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTransect not implemented")
}
func (UnimplementedTransectServiceServer) MustEmbedUnimplementedTransectServiceServer() {}

// UnsafeTransectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransectServiceServer will
// result in compilation errors.
type UnsafeTransectServiceServer interface {
	MustEmbedUnimplementedTransectServiceServer()
}

func RegisterTransectServiceServer(s grpc.ServiceRegistrar, srv TransectServiceServer) {
	s.RegisterService(&TransectService_ServiceDesc, srv)
}

func _TransectService_CreateTransect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputTransectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransectServiceServer).CreateTransect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TransectService/CreateTransect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransectServiceServer).CreateTransect(ctx, req.(*InputTransectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransectService_GetTransect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransectServiceServer).GetTransect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TransectService/GetTransect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransectServiceServer).GetTransect(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransectService_UpTransect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputTransectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransectServiceServer).UpTransect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TransectService/UpTransect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransectServiceServer).UpTransect(ctx, req.(*InputTransectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransectService_AddTrialSiteToTransect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputTransectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransectServiceServer).AddTrialSiteToTransect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TransectService/AddTrialSiteToTransect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransectServiceServer).AddTrialSiteToTransect(ctx, req.(*InputTransectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransectService_DeleteTransect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransectServiceServer).DeleteTransect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TransectService/DeleteTransect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransectServiceServer).DeleteTransect(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransectService_GetAllTransect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransectServiceServer).GetAllTransect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TransectService/GetAllTransect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransectServiceServer).GetAllTransect(ctx, req.(*TransectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransectService_ServiceDesc is the grpc.ServiceDesc for TransectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "botany.TransectService",
	HandlerType: (*TransectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransect",
			Handler:    _TransectService_CreateTransect_Handler,
		},
		{
			MethodName: "GetTransect",
			Handler:    _TransectService_GetTransect_Handler,
		},
		{
			MethodName: "UpTransect",
			Handler:    _TransectService_UpTransect_Handler,
		},
		{
			MethodName: "AddTrialSiteToTransect",
			Handler:    _TransectService_AddTrialSiteToTransect_Handler,
		},
		{
			MethodName: "DeleteTransect",
			Handler:    _TransectService_DeleteTransect_Handler,
		},
		{
			MethodName: "GetAllTransect",
			Handler:    _TransectService_GetAllTransect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transect.proto",
}
