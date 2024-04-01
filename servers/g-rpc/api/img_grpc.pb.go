// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: img.proto

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

// ImgServiceClient is the client API for ImgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImgServiceClient interface {
	GetImgByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Img, error)
	GetListImg(ctx context.Context, in *PagesRequest, opts ...grpc.CallOption) (*ImgList, error)
}

type imgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImgServiceClient(cc grpc.ClientConnInterface) ImgServiceClient {
	return &imgServiceClient{cc}
}

func (c *imgServiceClient) GetImgByID(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Img, error) {
	out := new(Img)
	err := c.cc.Invoke(ctx, "/botany.ImgService/GetImgByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imgServiceClient) GetListImg(ctx context.Context, in *PagesRequest, opts ...grpc.CallOption) (*ImgList, error) {
	out := new(ImgList)
	err := c.cc.Invoke(ctx, "/botany.ImgService/GetListImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImgServiceServer is the server API for ImgService service.
// All implementations must embed UnimplementedImgServiceServer
// for forward compatibility
type ImgServiceServer interface {
	GetImgByID(context.Context, *IdRequest) (*Img, error)
	GetListImg(context.Context, *PagesRequest) (*ImgList, error)
	MustEmbedUnimplementedImgServiceServer()
}

// UnimplementedImgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImgServiceServer struct {
}

func (UnimplementedImgServiceServer) GetImgByID(context.Context, *IdRequest) (*Img, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImgByID not implemented")
}
func (UnimplementedImgServiceServer) GetListImg(context.Context, *PagesRequest) (*ImgList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListImg not implemented")
}
func (UnimplementedImgServiceServer) MustEmbedUnimplementedImgServiceServer() {}

// UnsafeImgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImgServiceServer will
// result in compilation errors.
type UnsafeImgServiceServer interface {
	MustEmbedUnimplementedImgServiceServer()
}

func RegisterImgServiceServer(s grpc.ServiceRegistrar, srv ImgServiceServer) {
	s.RegisterService(&ImgService_ServiceDesc, srv)
}

func _ImgService_GetImgByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImgServiceServer).GetImgByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.ImgService/GetImgByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImgServiceServer).GetImgByID(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImgService_GetListImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImgServiceServer).GetListImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/botany.ImgService/GetListImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImgServiceServer).GetListImg(ctx, req.(*PagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImgService_ServiceDesc is the grpc.ServiceDesc for ImgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "botany.ImgService",
	HandlerType: (*ImgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImgByID",
			Handler:    _ImgService_GetImgByID_Handler,
		},
		{
			MethodName: "GetListImg",
			Handler:    _ImgService_GetListImg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "img.proto",
}