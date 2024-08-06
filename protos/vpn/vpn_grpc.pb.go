// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: vpn.proto

package vpn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	VPN_Connect_FullMethodName         = "/VPN/Connect"
	VPN_Disconnect_FullMethodName      = "/VPN/Disconnect"
	VPN_RegisterRoute_FullMethodName   = "/VPN/RegisterRoute"
	VPN_UnregisterRoute_FullMethodName = "/VPN/UnregisterRoute"
	VPN_TransferData_FullMethodName    = "/VPN/TransferData"
)

// VPNClient is the client API for VPN service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VPNClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	RegisterRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	UnregisterRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	TransferData(ctx context.Context, opts ...grpc.CallOption) (VPN_TransferDataClient, error)
}

type vPNClient struct {
	cc grpc.ClientConnInterface
}

func NewVPNClient(cc grpc.ClientConnInterface) VPNClient {
	return &vPNClient{cc}
}

func (c *vPNClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, VPN_Connect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, VPN_Disconnect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNClient) RegisterRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, VPN_RegisterRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNClient) UnregisterRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, VPN_UnregisterRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPNClient) TransferData(ctx context.Context, opts ...grpc.CallOption) (VPN_TransferDataClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VPN_ServiceDesc.Streams[0], VPN_TransferData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &vPNTransferDataClient{ClientStream: stream}
	return x, nil
}

type VPN_TransferDataClient interface {
	Send(*DataPacket) error
	Recv() (*DataPacket, error)
	grpc.ClientStream
}

type vPNTransferDataClient struct {
	grpc.ClientStream
}

func (x *vPNTransferDataClient) Send(m *DataPacket) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vPNTransferDataClient) Recv() (*DataPacket, error) {
	m := new(DataPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VPNServer is the server API for VPN service.
// All implementations must embed UnimplementedVPNServer
// for forward compatibility
type VPNServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	RegisterRoute(context.Context, *RouteRequest) (*RouteResponse, error)
	UnregisterRoute(context.Context, *RouteRequest) (*RouteResponse, error)
	TransferData(VPN_TransferDataServer) error
	mustEmbedUnimplementedVPNServer()
}

// UnimplementedVPNServer must be embedded to have forward compatible implementations.
type UnimplementedVPNServer struct {
}

func (UnimplementedVPNServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedVPNServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedVPNServer) RegisterRoute(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRoute not implemented")
}
func (UnimplementedVPNServer) UnregisterRoute(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterRoute not implemented")
}
func (UnimplementedVPNServer) TransferData(VPN_TransferDataServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferData not implemented")
}
func (UnimplementedVPNServer) mustEmbedUnimplementedVPNServer() {}

// UnsafeVPNServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VPNServer will
// result in compilation errors.
type UnsafeVPNServer interface {
	mustEmbedUnimplementedVPNServer()
}

func RegisterVPNServer(s grpc.ServiceRegistrar, srv VPNServer) {
	s.RegisterService(&VPN_ServiceDesc, srv)
}

func _VPN_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VPN_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPN_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VPN_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPN_RegisterRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServer).RegisterRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VPN_RegisterRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServer).RegisterRoute(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPN_UnregisterRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPNServer).UnregisterRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VPN_UnregisterRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPNServer).UnregisterRoute(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPN_TransferData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VPNServer).TransferData(&vPNTransferDataServer{ServerStream: stream})
}

type VPN_TransferDataServer interface {
	Send(*DataPacket) error
	Recv() (*DataPacket, error)
	grpc.ServerStream
}

type vPNTransferDataServer struct {
	grpc.ServerStream
}

func (x *vPNTransferDataServer) Send(m *DataPacket) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vPNTransferDataServer) Recv() (*DataPacket, error) {
	m := new(DataPacket)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VPN_ServiceDesc is the grpc.ServiceDesc for VPN service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VPN_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VPN",
	HandlerType: (*VPNServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _VPN_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _VPN_Disconnect_Handler,
		},
		{
			MethodName: "RegisterRoute",
			Handler:    _VPN_RegisterRoute_Handler,
		},
		{
			MethodName: "UnregisterRoute",
			Handler:    _VPN_UnregisterRoute_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransferData",
			Handler:       _VPN_TransferData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "vpn.proto",
}
