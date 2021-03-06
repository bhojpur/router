// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// RouterUIClient is the client API for RouterUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouterUIClient interface {
	// ListEngineSpecs returns a list of Routing Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (RouterUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type routerUIClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterUIClient(cc grpc.ClientConnInterface) RouterUIClient {
	return &routerUIClient{cc}
}

func (c *routerUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (RouterUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouterUI_ServiceDesc.Streams[0], "/v1.RouterUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouterUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type routerUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *routerUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.RouterUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterUIServer is the server API for RouterUI service.
// All implementations must embed UnimplementedRouterUIServer
// for forward compatibility
type RouterUIServer interface {
	// ListEngineSpecs returns a list of Routing Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, RouterUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedRouterUIServer()
}

// UnimplementedRouterUIServer must be embedded to have forward compatible implementations.
type UnimplementedRouterUIServer struct {
}

func (UnimplementedRouterUIServer) ListEngineSpecs(*ListEngineSpecsRequest, RouterUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedRouterUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedRouterUIServer) mustEmbedUnimplementedRouterUIServer() {}

// UnsafeRouterUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterUIServer will
// result in compilation errors.
type UnsafeRouterUIServer interface {
	mustEmbedUnimplementedRouterUIServer()
}

func RegisterRouterUIServer(s grpc.ServiceRegistrar, srv RouterUIServer) {
	s.RegisterService(&RouterUI_ServiceDesc, srv)
}

func _RouterUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterUIServer).ListEngineSpecs(m, &routerUIListEngineSpecsServer{stream})
}

type RouterUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type routerUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *routerUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RouterUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RouterUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouterUI_ServiceDesc is the grpc.ServiceDesc for RouterUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouterUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RouterUI",
	HandlerType: (*RouterUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _RouterUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _RouterUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "router-ui.proto",
}
