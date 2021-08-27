package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/booking-api/handlers"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	bookingConn, err := grpc.Dial(":2225", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	bookingClient := pb.NewSEOLEBookingClient(bookingConn)

	//Handler for gin
	h := handlers.NewBookingHandler(bookingClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(gin.Logger(), gin.Recovery())
	gr := g.Group("/v1/api")
	gr.POST("/create", h.Booking)
	gr.GET("/find", h.ViewBooking)
	gr.GET("/cancel", h.CancelBooking)
	//Listen and serve
	fmt.Println("Listen and serve at oort: 3334")
	http.ListenAndServe(":3334", g)
}
