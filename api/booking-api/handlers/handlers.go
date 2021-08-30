package handlers

import (
	"net/http"
	"time"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/booking-api/requests"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/api/booking-api/responses"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/helper"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type BookingHandler interface {
	Booking(c *gin.Context)
	ViewBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.SEOLEBookingClient
}

func NewBookingHandler(bookingClient pb.SEOLEBookingClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}

func (h *bookingHandler) Booking(c *gin.Context) {
	req := requests.BookingRequest{}

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
	bReq := &pb.BookingRequest{
		Id:         req.Id,
		Slut:       req.Slut,
		CustomerId: req.CustomerId,
		FlightId:   req.FlightId,
	}
	bRes, err := h.bookingClient.CreateBooking(c.Request.Context(), bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	dto := &responses.BookingResponse{
		Id:          bRes.Id,
		CustomerId:  bReq.CustomerId,
		FlightId:    bReq.FlightId,
		Code:        bRes.Code,
		Status:      bRes.Status,
		Booked_date: bRes.BookedDate.AsTime().String(),
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusCreated,
		"payload": dto,
	})
}

func (h *bookingHandler) ViewBooking(c *gin.Context) {
	req := requests.BookingCodeRequest{}

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
	bReq := &pb.BookingCode{
		Code: req.Code,
	}
	bRes, err := h.bookingClient.FindBooking(c.Request.Context(), bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	customerRes := &responses.CustomerResponse{
		Id:      bRes.Customer.Id,
		Name:    bRes.Customer.Name,
		Address: bRes.Customer.Address,
		Email:   bRes.Customer.Email,
		Phone:   bRes.Customer.PhoneNumber,
	}
	flightRes := &responses.FlightResponse{
		Id:         bRes.Flight.Id,
		Name:       bRes.Flight.Name,
		From:       bRes.Flight.From,
		To:         bRes.Flight.To,
		Status:     bRes.Flight.Status,
		Slot:       helper.Random_generate_string(3),
		DepartDate: helper.ConvertTimeToDate(time.Now().AddDate(0, 0, 3)),
		DepartTime: time.Now().Format("15:04:05"),
	}
	dto := &responses.BookingViewResponse{
		Code:        bRes.Code,
		Booked_date: helper.ConvertTimeToDate(bRes.BookedDate.AsTime()),
		Customer:    *customerRes,
		Flight:      *flightRes,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusCreated,
		"payload": dto,
	})
}

func (h *bookingHandler) CancelBooking(c *gin.Context) {
	req := requests.BookingCodeRequest{}

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
	bReq := &pb.BookingCode{
		Code: req.Code,
	}
	bRes, err := h.bookingClient.CancelBooking(c.Request.Context(), bReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingCancelResponse{
		Id:         bRes.Id,
		CustomerId: bRes.CustomerId,
		FlightId:   bRes.FlightId,
		Code:       bRes.Code,
		Status:     bRes.Status,
		CancelDate: time.Now().String(),
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusCreated,
		"payload": dto,
	})
}
