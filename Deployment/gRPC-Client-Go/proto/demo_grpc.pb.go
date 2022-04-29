// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package demo

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

// ServicioNodejsClient is the client API for ServicioNodejs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicioNodejsClient interface {
	IniciarJuego(ctx context.Context, in *PlayerGameRequest, opts ...grpc.CallOption) (*PlayerGameReply, error)
}

type servicioNodejsClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioNodejsClient(cc grpc.ClientConnInterface) ServicioNodejsClient {
	return &servicioNodejsClient{cc}
}

func (c *servicioNodejsClient) IniciarJuego(ctx context.Context, in *PlayerGameRequest, opts ...grpc.CallOption) (*PlayerGameReply, error) {
	out := new(PlayerGameReply)
	err := c.cc.Invoke(ctx, "/ServicioNodejs/IniciarJuego", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioNodejsServer is the server API for ServicioNodejs service.
// All implementations must embed UnimplementedServicioNodejsServer
// for forward compatibility
type ServicioNodejsServer interface {
	IniciarJuego(context.Context, *PlayerGameRequest) (*PlayerGameReply, error)
	mustEmbedUnimplementedServicioNodejsServer()
}

// UnimplementedServicioNodejsServer must be embedded to have forward compatible implementations.
type UnimplementedServicioNodejsServer struct {
}

func (UnimplementedServicioNodejsServer) IniciarJuego(context.Context, *PlayerGameRequest) (*PlayerGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IniciarJuego not implemented")
}
func (UnimplementedServicioNodejsServer) mustEmbedUnimplementedServicioNodejsServer() {}

// UnsafeServicioNodejsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicioNodejsServer will
// result in compilation errors.
type UnsafeServicioNodejsServer interface {
	mustEmbedUnimplementedServicioNodejsServer()
}

func RegisterServicioNodejsServer(s grpc.ServiceRegistrar, srv ServicioNodejsServer) {
	s.RegisterService(&ServicioNodejs_ServiceDesc, srv)
}

func _ServicioNodejs_IniciarJuego_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioNodejsServer).IniciarJuego(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServicioNodejs/IniciarJuego",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioNodejsServer).IniciarJuego(ctx, req.(*PlayerGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicioNodejs_ServiceDesc is the grpc.ServiceDesc for ServicioNodejs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicioNodejs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServicioNodejs",
	HandlerType: (*ServicioNodejsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IniciarJuego",
			Handler:    _ServicioNodejs_IniciarJuego_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/demo.proto",
}