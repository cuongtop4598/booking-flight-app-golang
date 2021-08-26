package handlers

import (
	"net/http"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/flight-api/requests"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/flight-api/responses"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	FindFlight(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.FlightClient
}

func NewFlightHandler(flightClient pb.FlightClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := requests.FlightRequest{}

	if err := c.ShouldBind(&req); err != nil {

		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Kind().String())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}

	fReq := &pb.FlightRequest{
		Slut: req.Slut,
		Name: req.Name,
		From: req.From,
		To:   req.To,
		Date: timestamppb.New(req.Date),
	}

	fRes, err := h.flightClient.CreateFlight(c.Request.Context(), fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.FlightReponse{
		Id:            fRes.Id,
		Slut:          fRes.Slut,
		Name:          fRes.Name,
		From:          fRes.From,
		To:            fRes.To,
		Date:          fRes.Date.AsTime(),
		Status:        fRes.Status,
		AvailableSlot: fRes.AvailableSlot,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *flightHandler) FindFlight(c *gin.Context) {
	req := requests.FlightRequest{}

	if err := c.ShouldBind(&req); err != nil {

		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Kind().String())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})
		return
	}

	fReq := &pb.FlightRequest{
		Slut: req.Slut,
	}

	fRes, err := h.flightClient.FindFlight(c.Request.Context(), fReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.FlightReponse{
		Id:            fRes.Id,
		Slut:          fRes.Slut,
		Name:          fRes.Name,
		From:          fRes.From,
		To:            fRes.To,
		Date:          fRes.Date.AsTime(),
		Status:        fRes.Status,
		AvailableSlot: fRes.AvailableSlot,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
