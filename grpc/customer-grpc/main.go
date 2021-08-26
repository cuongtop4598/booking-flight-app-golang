package main

import (
	"fmt"
	"net"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/flight-grpc/handlers"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/flight-grpc/repositories"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/helper"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := helper.AutoBindConfig("config.yml")
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", viper.GetInt("socket.port")))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	flightRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}
	h, err := handlers.NewFlighHandler(flightRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterFlightServer(s, h)
	fmt.Printf("Listen at port: %v", viper.GetInt("socket.port"))
	s.Serve(listen)
}
