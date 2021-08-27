package main

import (
	"fmt"
	"net"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/booking-grpc/handlers"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/booking-grpc/repositories"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/helper"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func main() {

	flightConn, err := grpc.Dial(":2223",
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithCodes(codes.DeadlineExceeded, codes.Internal),
				grpc_retry.WithMax(2)),
		)),
		grpc.WithChainStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_retry.StreamClientInterceptor(
				grpc_retry.WithCodes(codes.DeadlineExceeded, codes.Internal),
				grpc_retry.WithMax(2)),
		)),
	)
	if err != nil {
		panic(err)
	}
	customerConn, err := grpc.Dial(":2224",
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithCodes(codes.DeadlineExceeded, codes.Internal),
				grpc_retry.WithMax(2)),
		)),
		grpc.WithChainStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_retry.StreamClientInterceptor(
				grpc_retry.WithCodes(codes.DeadlineExceeded, codes.Internal),
				grpc_retry.WithMax(2)),
		)))
	if err != nil {
		panic(err)
	}
	flightClient := pb.NewFlightClient(flightConn)
	customerClient := pb.NewCustormerClient(customerConn)

	err = helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", ":2225")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	bookingRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}
	h := handlers.NewBookingHandler(flightClient, customerClient, bookingRepository)
	reflection.Register(s)
	pb.RegisterSEOLEBookingServer(s, h)
	fmt.Println("Listen at port: 2225")
	s.Serve(listen)
}
