// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: type-plant.proto

package type_plant

import (
	context "context"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TypePlantServiceClient is the client API for TypePlantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TypePlantServiceClient interface {
	CreateTypePlant(ctx context.Context, in *InputTypePlantRequest, opts ...grpc.CallOption) (*TypePlant, error)
	// Получение сущности по
	GetTypePlant(ctx context.Context, in *api.IdRequest, opts ...grpc.CallOption) (*TypePlant, error)
	// Обновление сущности по id
	UpdateTypePlant(ctx context.Context, in *InputTypePlantRequest, opts ...grpc.CallOption) (*TypePlant, error)
	// Удаление сущности по заголовку
	DeleteTypePlant(ctx context.Context, in *api.IdRequest, opts ...grpc.CallOption) (*api.BoolResponse, error)
	// Получение списка всех сущностей
	GetAllTypePlant(ctx context.Context, in *api.EmptyRequest, opts ...grpc.CallOption) (*TypePlantList, error)
}

type typePlantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTypePlantServiceClient(cc grpc.ClientConnInterface) TypePlantServiceClient {
	return &typePlantServiceClient{cc}
}

func (c *typePlantServiceClient) CreateTypePlant(ctx context.Context, in *InputTypePlantRequest, opts ...grpc.CallOption) (*TypePlant, error) {
	out := new(TypePlant)
	err := c.cc.Invoke(ctx, "/botany.TypePlantService/CreateTypePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typePlantServiceClient) GetTypePlant(ctx context.Context, in *api.IdRequest, opts ...grpc.CallOption) (*TypePlant, error) {
	out := new(TypePlant)
	err := c.cc.Invoke(ctx, "/botany.TypePlantService/GetTypePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typePlantServiceClient) UpdateTypePlant(ctx context.Context, in *InputTypePlantRequest, opts ...grpc.CallOption) (*TypePlant, error) {
	out := new(TypePlant)
	err := c.cc.Invoke(ctx, "/botany.TypePlantService/UpdateTypePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typePlantServiceClient) DeleteTypePlant(ctx context.Context, in *api.IdRequest, opts ...grpc.CallOption) (*api.BoolResponse, error) {
	out := new(api.BoolResponse)
	err := c.cc.Invoke(ctx, "/botany.TypePlantService/DeleteTypePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typePlantServiceClient) GetAllTypePlant(ctx context.Context, in *api.EmptyRequest, opts ...grpc.CallOption) (*TypePlantList, error) {
	out := new(TypePlantList)
	err := c.cc.Invoke(ctx, "/botany.TypePlantService/GetAllTypePlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TypePlantServiceServer is the server API for TypePlantService service.
// All implementations must embed UnimplementedTypePlantServiceServer
// for forward compatibility
type TypePlantServiceServer interface {
	CreateTypePlant(context.Context, *InputTypePlantRequest) (*TypePlant, error)
	// Получение сущности по
	GetTypePlant(context.Context, *api.IdRequest) (*TypePlant, error)
	// Обновление сущности по id
	UpdateTypePlant(context.Context, *InputTypePlantRequest) (*TypePlant, error)
	// Удаление сущности по заголовку
	DeleteTypePlant(context.Context, *api.IdRequest) (*api.BoolResponse, error)
	// Получение списка всех сущностей
	GetAllTypePlant(context.Context, *api.EmptyRequest) (*TypePlantList, error)
	mustEmbedUnimplementedTypePlantServiceServer()
}

// UnimplementedTypePlantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTypePlantServiceServer struct {
}

func (UnimplementedTypePlantServiceServer) CreateTypePlant(context.Context, *InputTypePlantRequest) (*TypePlant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTypePlant not implemented")
}
func (UnimplementedTypePlantServiceServer) GetTypePlant(context.Context, *api.IdRequest) (*TypePlant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypePlant not implemented")
}
func (UnimplementedTypePlantServiceServer) UpdateTypePlant(context.Context, *InputTypePlantRequest) (*TypePlant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTypePlant not implemented")
}
func (UnimplementedTypePlantServiceServer) DeleteTypePlant(context.Context, *api.IdRequest) (*api.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTypePlant not implemented")
}
func (UnimplementedTypePlantServiceServer) GetAllTypePlant(context.Context, *api.EmptyRequest) (*TypePlantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTypePlant not implemented")
}
func (UnimplementedTypePlantServiceServer) mustEmbedUnimplementedTypePlantServiceServer() {}

// UnsafeTypePlantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TypePlantServiceServer will
// result in compilation errors.
type UnsafeTypePlantServiceServer interface {
	mustEmbedUnimplementedTypePlantServiceServer()
}

func RegisterTypePlantServiceServer(s grpc.ServiceRegistrar, srv TypePlantServiceServer) {
	s.RegisterService(&TypePlantService_ServiceDesc, srv)
}

func _TypePlantService_CreateTypePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputTypePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypePlantServiceServer).CreateTypePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TypePlantService/CreateTypePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypePlantServiceServer).CreateTypePlant(ctx, req.(*InputTypePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypePlantService_GetTypePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypePlantServiceServer).GetTypePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TypePlantService/GetTypePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypePlantServiceServer).GetTypePlant(ctx, req.(*api.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypePlantService_UpdateTypePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputTypePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypePlantServiceServer).UpdateTypePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TypePlantService/UpdateTypePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypePlantServiceServer).UpdateTypePlant(ctx, req.(*InputTypePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypePlantService_DeleteTypePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypePlantServiceServer).DeleteTypePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TypePlantService/DeleteTypePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypePlantServiceServer).DeleteTypePlant(ctx, req.(*api.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypePlantService_GetAllTypePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypePlantServiceServer).GetAllTypePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.TypePlantService/GetAllTypePlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypePlantServiceServer).GetAllTypePlant(ctx, req.(*api.EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TypePlantService_ServiceDesc is the grpc.ServiceDesc for TypePlantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TypePlantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "botany.TypePlantService",
	HandlerType: (*TypePlantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTypePlant",
			Handler:    _TypePlantService_CreateTypePlant_Handler,
		},
		{
			MethodName: "GetTypePlant",
			Handler:    _TypePlantService_GetTypePlant_Handler,
		},
		{
			MethodName: "UpdateTypePlant",
			Handler:    _TypePlantService_UpdateTypePlant_Handler,
		},
		{
			MethodName: "DeleteTypePlant",
			Handler:    _TypePlantService_DeleteTypePlant_Handler,
		},
		{
			MethodName: "GetAllTypePlant",
			Handler:    _TypePlantService_GetAllTypePlant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "type-plant.proto",
}
