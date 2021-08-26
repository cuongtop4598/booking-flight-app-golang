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

// CustormerClient is the client API for Custormer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustormerClient interface {
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
	UpdateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
	FindCustomer(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*CustomerResponse, error)
	ListCustomer(ctx context.Context, in *ListCustomerRequest, opts ...grpc.CallOption) (*ListCustomerResponse, error)
	ChangePassword(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*Empty, error)
	BookingHistory(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*History, error)
}

type custormerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustormerClient(cc grpc.ClientConnInterface) CustormerClient {
	return &custormerClient{cc}
}

func (c *custormerClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/booking.Custormer/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *custormerClient) UpdateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/booking.Custormer/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *custormerClient) FindCustomer(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/booking.Custormer/FindCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *custormerClient) ListCustomer(ctx context.Context, in *ListCustomerRequest, opts ...grpc.CallOption) (*ListCustomerResponse, error) {
	out := new(ListCustomerResponse)
	err := c.cc.Invoke(ctx, "/booking.Custormer/ListCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *custormerClient) ChangePassword(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/booking.Custormer/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *custormerClient) BookingHistory(ctx context.Context, in *CustomerID, opts ...grpc.CallOption) (*History, error) {
	out := new(History)
	err := c.cc.Invoke(ctx, "/booking.Custormer/BookingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustormerServer is the server API for Custormer service.
// All implementations must embed UnimplementedCustormerServer
// for forward compatibility
type CustormerServer interface {
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
	UpdateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
	FindCustomer(context.Context, *CustomerID) (*CustomerResponse, error)
	ListCustomer(context.Context, *ListCustomerRequest) (*ListCustomerResponse, error)
	ChangePassword(context.Context, *CustomerID) (*Empty, error)
	BookingHistory(context.Context, *CustomerID) (*History, error)
	mustEmbedUnimplementedCustormerServer()
}

// UnimplementedCustormerServer must be embedded to have forward compatible implementations.
type UnimplementedCustormerServer struct {
}

func (UnimplementedCustormerServer) CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustormerServer) UpdateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustormerServer) FindCustomer(context.Context, *CustomerID) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCustomer not implemented")
}
func (UnimplementedCustormerServer) ListCustomer(context.Context, *ListCustomerRequest) (*ListCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCustomer not implemented")
}
func (UnimplementedCustormerServer) ChangePassword(context.Context, *CustomerID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedCustormerServer) BookingHistory(context.Context, *CustomerID) (*History, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookingHistory not implemented")
}
func (UnimplementedCustormerServer) mustEmbedUnimplementedCustormerServer() {}

// UnsafeCustormerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustormerServer will
// result in compilation errors.
type UnsafeCustormerServer interface {
	mustEmbedUnimplementedCustormerServer()
}

func RegisterCustormerServer(s grpc.ServiceRegistrar, srv CustormerServer) {
	s.RegisterService(&Custormer_ServiceDesc, srv)
}

func _Custormer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustormerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Custormer/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustormerServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Custormer_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustormerServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Custormer/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustormerServer).UpdateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Custormer_FindCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustormerServer).FindCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Custormer/FindCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustormerServer).FindCustomer(ctx, req.(*CustomerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Custormer_ListCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustormerServer).ListCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Custormer/ListCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustormerServer).ListCustomer(ctx, req.(*ListCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Custormer_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustormerServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Custormer/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustormerServer).ChangePassword(ctx, req.(*CustomerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Custormer_BookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustormerServer).BookingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Custormer/BookingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustormerServer).BookingHistory(ctx, req.(*CustomerID))
	}
	return interceptor(ctx, in, info, handler)
}

// Custormer_ServiceDesc is the grpc.ServiceDesc for Custormer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Custormer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.Custormer",
	HandlerType: (*CustormerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Custormer_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _Custormer_UpdateCustomer_Handler,
		},
		{
			MethodName: "FindCustomer",
			Handler:    _Custormer_FindCustomer_Handler,
		},
		{
			MethodName: "ListCustomer",
			Handler:    _Custormer_ListCustomer_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Custormer_ChangePassword_Handler,
		},
		{
			MethodName: "BookingHistory",
			Handler:    _Custormer_BookingHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
