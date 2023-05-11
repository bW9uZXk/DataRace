package command

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

// RoutingServiceClient is the client API for RoutingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutingServiceClient interface {
	SubscribeRoutingStats(ctx context.Context, in *SubscribeRoutingStatsRequest, opts ...grpc.CallOption) (RoutingService_SubscribeRoutingStatsClient, error)
	TestRoute(ctx context.Context, in *TestRouteRequest, opts ...grpc.CallOption) (*RoutingContext, error)
	GetBalancerInfo(ctx context.Context, in *GetBalancerInfoRequest, opts ...grpc.CallOption) (*GetBalancerInfoResponse, error)
	OverrideBalancerTarget(ctx context.Context, in *OverrideBalancerTargetRequest, opts ...grpc.CallOption) (*OverrideBalancerTargetResponse, error)
}

type routingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingServiceClient(cc grpc.ClientConnInterface) RoutingServiceClient {
	return &routingServiceClient{cc}
}

func (c *routingServiceClient) SubscribeRoutingStats(ctx context.Context, in *SubscribeRoutingStatsRequest, opts ...grpc.CallOption) (RoutingService_SubscribeRoutingStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoutingService_ServiceDesc.Streams[0], "/v2ray.core.app.router.command.RoutingService/SubscribeRoutingStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingServiceSubscribeRoutingStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingService_SubscribeRoutingStatsClient interface {
	Recv() (*RoutingContext, error)
	grpc.ClientStream
}

type routingServiceSubscribeRoutingStatsClient struct {
	grpc.ClientStream
}

func (x *routingServiceSubscribeRoutingStatsClient) Recv() (*RoutingContext, error) {
	m := new(RoutingContext)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routingServiceClient) TestRoute(ctx context.Context, in *TestRouteRequest, opts ...grpc.CallOption) (*RoutingContext, error) {
	out := new(RoutingContext)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.router.command.RoutingService/TestRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingServiceClient) GetBalancerInfo(ctx context.Context, in *GetBalancerInfoRequest, opts ...grpc.CallOption) (*GetBalancerInfoResponse, error) {
	out := new(GetBalancerInfoResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.router.command.RoutingService/GetBalancerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingServiceClient) OverrideBalancerTarget(ctx context.Context, in *OverrideBalancerTargetRequest, opts ...grpc.CallOption) (*OverrideBalancerTargetResponse, error) {
	out := new(OverrideBalancerTargetResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.router.command.RoutingService/OverrideBalancerTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutingServiceServer is the server API for RoutingService service.
// All implementations must embed UnimplementedRoutingServiceServer
// for forward compatibility
type RoutingServiceServer interface {
	SubscribeRoutingStats(*SubscribeRoutingStatsRequest, RoutingService_SubscribeRoutingStatsServer) error
	TestRoute(context.Context, *TestRouteRequest) (*RoutingContext, error)
	GetBalancerInfo(context.Context, *GetBalancerInfoRequest) (*GetBalancerInfoResponse, error)
	OverrideBalancerTarget(context.Context, *OverrideBalancerTargetRequest) (*OverrideBalancerTargetResponse, error)
	mustEmbedUnimplementedRoutingServiceServer()
}

// UnimplementedRoutingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutingServiceServer struct {
}

func (UnimplementedRoutingServiceServer) SubscribeRoutingStats(*SubscribeRoutingStatsRequest, RoutingService_SubscribeRoutingStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeRoutingStats not implemented")
}
func (UnimplementedRoutingServiceServer) TestRoute(context.Context, *TestRouteRequest) (*RoutingContext, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRoute not implemented")
}
func (UnimplementedRoutingServiceServer) GetBalancerInfo(context.Context, *GetBalancerInfoRequest) (*GetBalancerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalancerInfo not implemented")
}
func (UnimplementedRoutingServiceServer) OverrideBalancerTarget(context.Context, *OverrideBalancerTargetRequest) (*OverrideBalancerTargetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverrideBalancerTarget not implemented")
}
func (UnimplementedRoutingServiceServer) mustEmbedUnimplementedRoutingServiceServer() {}

// UnsafeRoutingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingServiceServer will
// result in compilation errors.
type UnsafeRoutingServiceServer interface {
	mustEmbedUnimplementedRoutingServiceServer()
}

func RegisterRoutingServiceServer(s grpc.ServiceRegistrar, srv RoutingServiceServer) {
	s.RegisterService(&RoutingService_ServiceDesc, srv)
}

func _RoutingService_SubscribeRoutingStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRoutingStatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingServiceServer).SubscribeRoutingStats(m, &routingServiceSubscribeRoutingStatsServer{stream})
}

type RoutingService_SubscribeRoutingStatsServer interface {
	Send(*RoutingContext) error
	grpc.ServerStream
}

type routingServiceSubscribeRoutingStatsServer struct {
	grpc.ServerStream
}

func (x *routingServiceSubscribeRoutingStatsServer) Send(m *RoutingContext) error {
	return x.ServerStream.SendMsg(m)
}

func _RoutingService_TestRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingServiceServer).TestRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.router.command.RoutingService/TestRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingServiceServer).TestRoute(ctx, req.(*TestRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingService_GetBalancerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalancerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingServiceServer).GetBalancerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.router.command.RoutingService/GetBalancerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingServiceServer).GetBalancerInfo(ctx, req.(*GetBalancerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingService_OverrideBalancerTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverrideBalancerTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingServiceServer).OverrideBalancerTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.router.command.RoutingService/OverrideBalancerTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingServiceServer).OverrideBalancerTarget(ctx, req.(*OverrideBalancerTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoutingService_ServiceDesc is the grpc.ServiceDesc for RoutingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2ray.core.app.router.command.RoutingService",
	HandlerType: (*RoutingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestRoute",
			Handler:    _RoutingService_TestRoute_Handler,
		},
		{
			MethodName: "GetBalancerInfo",
			Handler:    _RoutingService_GetBalancerInfo_Handler,
		},
		{
			MethodName: "OverrideBalancerTarget",
			Handler:    _RoutingService_OverrideBalancerTarget_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeRoutingStats",
			Handler:       _RoutingService_SubscribeRoutingStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "app/router/command/command.proto",
}
