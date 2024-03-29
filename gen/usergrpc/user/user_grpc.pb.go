// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user/user.proto

package usergrpc

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

// UsererClient is the client API for Userer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsererClient interface {
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error)
}

type usererClient struct {
	cc grpc.ClientConnInterface
}

func NewUsererClient(cc grpc.ClientConnInterface) UsererClient {
	return &usererClient{cc}
}

func (c *usererClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/usergrpc.Userer/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usererClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/usergrpc.Userer/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usererClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/usergrpc.Userer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usererClient) GetByID(ctx context.Context, in *GetByIDRequest, opts ...grpc.CallOption) (*GetByIDResponse, error) {
	out := new(GetByIDResponse)
	err := c.cc.Invoke(ctx, "/usergrpc.Userer/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsererServer is the server API for Userer service.
// All implementations must embed UnimplementedUsererServer
// for forward compatibility
type UsererServer interface {
	Profile(context.Context, *ProfileRequest) (*ProfileResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error)
	mustEmbedUnimplementedUsererServer()
}

// UnimplementedUsererServer must be embedded to have forward compatible implementations.
type UnimplementedUsererServer struct {
}

func (UnimplementedUsererServer) Profile(context.Context, *ProfileRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedUsererServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUsererServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUsererServer) GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUsererServer) mustEmbedUnimplementedUsererServer() {}

// UnsafeUsererServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsererServer will
// result in compilation errors.
type UnsafeUsererServer interface {
	mustEmbedUnimplementedUsererServer()
}

func RegisterUsererServer(s grpc.ServiceRegistrar, srv UsererServer) {
	s.RegisterService(&Userer_ServiceDesc, srv)
}

func _Userer_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsererServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usergrpc.Userer/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsererServer).Profile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsererServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usergrpc.Userer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsererServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsererServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usergrpc.Userer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsererServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userer_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsererServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usergrpc.Userer/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsererServer).GetByID(ctx, req.(*GetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Userer_ServiceDesc is the grpc.ServiceDesc for Userer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Userer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usergrpc.Userer",
	HandlerType: (*UsererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Profile",
			Handler:    _Userer_Profile_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Userer_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Userer_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Userer_GetByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
