package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/customer-api/handlers"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//create grpc client connect
	clientConn, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	customerClient := pb.NewCustormerClient(clientConn)
	// handler for gin
	h := handlers.NewCustomerHandler(customerClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(gin.Logger(), gin.Recovery())
	gr := g.Group("/v1/api")
	gr.GET("/history", h.BookingHistory)
	gr.POST("/create", h.CreateCustom)
	fmt.Println("Listen and serve at port: 3335")
	http.ListenAndServe(":3335", g)
}
