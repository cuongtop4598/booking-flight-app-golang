package main

import (
	"net/http"
	"os"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/flight-api/handlers"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// create grpc client connect
	flightConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	flightClient := pb.NewFlightClient(flightConn)

	//Handler for Gin Gonic
	h := handlers.NewFlightHandler(flightClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(gin.Logger(), gin.Recovery())
	//Create routes
	gr := g.Group("/v1/api")
	gr.POST("/create", h.CreateFlight)
	gr.GET("/find", h.FindFlight)
	//Listen and serve
	http.ListenAndServe(":3333", g)
}
