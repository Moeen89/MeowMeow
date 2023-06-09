// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: GRPC/connect.proto

package go_connection_grpc

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	ReqPq(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*ReplyMsg, error)
	Req_DHParam(ctx context.Context, in *NewMsg, opts ...grpc.CallOption) (*NewReplyMsg, error)
	CheckKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) ReqPq(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*ReplyMsg, error) {
	out := new(ReplyMsg)
	err := c.cc.Invoke(ctx, "/GRPC.authService/req_pq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Req_DHParam(ctx context.Context, in *NewMsg, opts ...grpc.CallOption) (*NewReplyMsg, error) {
	out := new(NewReplyMsg)
	err := c.cc.Invoke(ctx, "/GRPC.authService/req_DH_param", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error) {
	out := new(Val)
	err := c.cc.Invoke(ctx, "/GRPC.authService/check_key", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	ReqPq(context.Context, *Msg) (*ReplyMsg, error)
	Req_DHParam(context.Context, *NewMsg) (*NewReplyMsg, error)
	CheckKey(context.Context, *Key) (*Val, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) ReqPq(context.Context, *Msg) (*ReplyMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReqPq not implemented")
}
func (UnimplementedAuthServiceServer) Req_DHParam(context.Context, *NewMsg) (*NewReplyMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Req_DHParam not implemented")
}
func (UnimplementedAuthServiceServer) CheckKey(context.Context, *Key) (*Val, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckKey not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_ReqPq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ReqPq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GRPC.authService/req_pq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ReqPq(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Req_DHParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Req_DHParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GRPC.authService/req_DH_param",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Req_DHParam(ctx, req.(*NewMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GRPC.authService/check_key",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckKey(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GRPC.authService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "req_pq",
			Handler:    _AuthService_ReqPq_Handler,
		},
		{
			MethodName: "req_DH_param",
			Handler:    _AuthService_Req_DHParam_Handler,
		},
		{
			MethodName: "check_key",
			Handler:    _AuthService_CheckKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GRPC/connect.proto",
}
