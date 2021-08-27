// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// FlightClient is the client API for Flight service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightClient interface {
	CreateFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*FlightResponse, error)
	UpdateFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*FlightResponse, error)
	FindFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*FlightResponse, error)
	FindListFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*ListFlightResponse, error)
}

type flightClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightClient(cc grpc.ClientConnInterface) FlightClient {
	return &flightClient{cc}
}

func (c *flightClient) CreateFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*FlightResponse, error) {
	out := new(FlightResponse)
	err := c.cc.Invoke(ctx, "/booking.Flight/CreateFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightClient) UpdateFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*FlightResponse, error) {
	out := new(FlightResponse)
	err := c.cc.Invoke(ctx, "/booking.Flight/UpdateFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightClient) FindFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*FlightResponse, error) {
	out := new(FlightResponse)
	err := c.cc.Invoke(ctx, "/booking.Flight/FindFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightClient) FindListFlight(ctx context.Context, in *FlightRequest, opts ...grpc.CallOption) (*ListFlightResponse, error) {
	out := new(ListFlightResponse)
	err := c.cc.Invoke(ctx, "/booking.Flight/FindListFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightServer is the server API for Flight service.
// All implementations must embed UnimplementedFlightServer
// for forward compatibility
type FlightServer interface {
	CreateFlight(context.Context, *FlightRequest) (*FlightResponse, error)
	UpdateFlight(context.Context, *FlightRequest) (*FlightResponse, error)
	FindFlight(context.Context, *FlightRequest) (*FlightResponse, error)
	FindListFlight(context.Context, *FlightRequest) (*ListFlightResponse, error)
	mustEmbedUnimplementedFlightServer()
}

// UnimplementedFlightServer must be embedded to have forward compatible implementations.
type UnimplementedFlightServer struct {
}

func (UnimplementedFlightServer) CreateFlight(context.Context, *FlightRequest) (*FlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlight not implemented")
}
func (UnimplementedFlightServer) UpdateFlight(context.Context, *FlightRequest) (*FlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlight not implemented")
}
func (UnimplementedFlightServer) FindFlight(context.Context, *FlightRequest) (*FlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFlight not implemented")
}
func (UnimplementedFlightServer) FindListFlight(context.Context, *FlightRequest) (*ListFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListFlight not implemented")
}
func (UnimplementedFlightServer) mustEmbedUnimplementedFlightServer() {}

// UnsafeFlightServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightServer will
// result in compilation errors.
type UnsafeFlightServer interface {
	mustEmbedUnimplementedFlightServer()
}

func RegisterFlightServer(s grpc.ServiceRegistrar, srv FlightServer) {
	s.RegisterService(&Flight_ServiceDesc, srv)
}

func _Flight_CreateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServer).CreateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Flight/CreateFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServer).CreateFlight(ctx, req.(*FlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flight_UpdateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServer).UpdateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Flight/UpdateFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServer).UpdateFlight(ctx, req.(*FlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flight_FindFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServer).FindFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Flight/FindFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServer).FindFlight(ctx, req.(*FlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Flight_FindListFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServer).FindListFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Flight/FindListFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServer).FindListFlight(ctx, req.(*FlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Flight_ServiceDesc is the grpc.ServiceDesc for Flight service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Flight_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.Flight",
	HandlerType: (*FlightServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFlight",
			Handler:    _Flight_CreateFlight_Handler,
		},
		{
			MethodName: "UpdateFlight",
			Handler:    _Flight_UpdateFlight_Handler,
		},
		{
			MethodName: "FindFlight",
			Handler:    _Flight_FindFlight_Handler,
		},
		{
			MethodName: "FindListFlight",
			Handler:    _Flight_FindListFlight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flight.proto",
}
