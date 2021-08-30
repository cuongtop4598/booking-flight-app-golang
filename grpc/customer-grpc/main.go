package main

import (
	"fmt"
	"net"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/customer-grpc/handlers"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/customer-grpc/repositories"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/helper"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func main() {

	bookingCnn, err := grpc.Dial(":2225", grpc.WithInsecure(),
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
	bookingClient := pb.NewSEOLEBookingClient(bookingCnn)

	flightCnn, err := grpc.Dial(":2223", grpc.WithInsecure(),
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
	flightClient := pb.NewFlightClient(flightCnn)

	err = helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", viper.GetInt("socket.port")))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	customertRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}
	h, err := handlers.NewCustomerHandler(bookingClient, flightClient, customertRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterCustormerServer(s, h)
	fmt.Printf("Listen at port: %v", viper.GetInt("socket.port"))
	s.Serve(listen)
}
