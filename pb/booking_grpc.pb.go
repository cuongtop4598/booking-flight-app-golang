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

// SEOLEBookingClient is the client API for SEOLEBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SEOLEBookingClient interface {
	CreateBooking(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error)
	FindBooking(ctx context.Context, in *BookingCode, opts ...grpc.CallOption) (*BookingInfo, error)
	CancelBooking(ctx context.Context, in *BookingCode, opts ...grpc.CallOption) (*BookingInfo, error)
	BookingHistory(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*ListBookingReponse, error)
}

type sEOLEBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewSEOLEBookingClient(cc grpc.ClientConnInterface) SEOLEBookingClient {
	return &sEOLEBookingClient{cc}
}

func (c *sEOLEBookingClient) CreateBooking(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, "/booking.SEOLEBooking/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sEOLEBookingClient) FindBooking(ctx context.Context, in *BookingCode, opts ...grpc.CallOption) (*BookingInfo, error) {
	out := new(BookingInfo)
	err := c.cc.Invoke(ctx, "/booking.SEOLEBooking/FindBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sEOLEBookingClient) CancelBooking(ctx context.Context, in *BookingCode, opts ...grpc.CallOption) (*BookingInfo, error) {
	out := new(BookingInfo)
	err := c.cc.Invoke(ctx, "/booking.SEOLEBooking/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sEOLEBookingClient) BookingHistory(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*ListBookingReponse, error) {
	out := new(ListBookingReponse)
	err := c.cc.Invoke(ctx, "/booking.SEOLEBooking/BookingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SEOLEBookingServer is the server API for SEOLEBooking service.
// All implementations must embed UnimplementedSEOLEBookingServer
// for forward compatibility
type SEOLEBookingServer interface {
	CreateBooking(context.Context, *BookingRequest) (*BookingResponse, error)
	FindBooking(context.Context, *BookingCode) (*BookingInfo, error)
	CancelBooking(context.Context, *BookingCode) (*BookingInfo, error)
	BookingHistory(context.Context, *CustomerID) (*ListBookingReponse, error)
	mustEmbedUnimplementedSEOLEBookingServer()
}

// UnimplementedSEOLEBookingServer must be embedded to have forward compatible implementations.
type UnimplementedSEOLEBookingServer struct {
}

func (UnimplementedSEOLEBookingServer) CreateBooking(context.Context, *BookingRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedSEOLEBookingServer) FindBooking(context.Context, *BookingCode) (*BookingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBooking not implemented")
}
func (UnimplementedSEOLEBookingServer) CancelBooking(context.Context, *BookingCode) (*BookingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedSEOLEBookingServer) BookingHistory(context.Context, *CustomerID) (*ListBookingReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookingHistory not implemented")
}
func (UnimplementedSEOLEBookingServer) mustEmbedUnimplementedSEOLEBookingServer() {}

// UnsafeSEOLEBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SEOLEBookingServer will
// result in compilation errors.
type UnsafeSEOLEBookingServer interface {
	mustEmbedUnimplementedSEOLEBookingServer()
}

func RegisterSEOLEBookingServer(s grpc.ServiceRegistrar, srv SEOLEBookingServer) {
	s.RegisterService(&SEOLEBooking_ServiceDesc, srv)
}

func _SEOLEBooking_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOLEBookingServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.SEOLEBooking/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOLEBookingServer).CreateBooking(ctx, req.(*BookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SEOLEBooking_FindBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOLEBookingServer).FindBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.SEOLEBooking/FindBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOLEBookingServer).FindBooking(ctx, req.(*BookingCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _SEOLEBooking_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOLEBookingServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.SEOLEBooking/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOLEBookingServer).CancelBooking(ctx, req.(*BookingCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _SEOLEBooking_BookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SEOLEBookingServer).BookingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.SEOLEBooking/BookingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SEOLEBookingServer).BookingHistory(ctx, req.(*CustomerID))
	}
	return interceptor(ctx, in, info, handler)
}

// SEOLEBooking_ServiceDesc is the grpc.ServiceDesc for SEOLEBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SEOLEBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.SEOLEBooking",
	HandlerType: (*SEOLEBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _SEOLEBooking_CreateBooking_Handler,
		},
		{
			MethodName: "FindBooking",
			Handler:    _SEOLEBooking_FindBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _SEOLEBooking_CancelBooking_Handler,
		},
		{
			MethodName: "BookingHistory",
			Handler:    _SEOLEBooking_BookingHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}