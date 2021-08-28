package handlers

import (
	"context"
	"database/sql"
	"time"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/booking-grpc/models"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/booking-grpc/repositories"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/booking-grpc/requests"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/helper"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingHandler struct {
	flightClient        pb.FlightClient
	customersClient     pb.CustormerClient
	bookingRepositories repositories.BookingRepositories
	pb.UnimplementedSEOLEBookingServer
}

func NewBookingHandler(flightClient pb.FlightClient, customersClient pb.CustormerClient,
	bookingRepositories repositories.BookingRepositories) *BookingHandler {
	return &BookingHandler{
		flightClient:        flightClient,
		customersClient:     customersClient,
		bookingRepositories: bookingRepositories,
	}
}

func (h *BookingHandler) CreateBooking(ctx context.Context, in *pb.BookingRequest) (*pb.BookingResponse, error) {
	if in.CustomerId == "" {
		return nil, status.Error(codes.InvalidArgument, "customer_id is required")
	}
	if in.FlightId == "" {
		return nil, status.Error(codes.InvalidArgument, "flight_id is required")
	}
	flight, err := h.flightClient.FindFlight(ctx, &pb.FlightRequest{
		Id: in.FlightId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "flight_id is not found")
			}
		} else {
			return nil, err
		}
	}
	if helper.CaculateDistanceOfTwoTime(flight.Date.AsTime(), time.Now()) > 12 {
		return nil, status.Error(codes.Internal, "flight date exceeded")
	}
	if flight.AvailableSlot == 100 {
		return nil, status.Error(codes.Internal, "out of slot")
	}
	_, err = h.customersClient.FindCustomerById(ctx, &pb.CustomerAuthen{
		Id: in.CustomerId,
	})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.NotFound {
				return nil, status.Error(codes.NotFound, "customer_id is not found")
			}
		} else {
			return nil, err
		}
	}

	booking := &requests.BookingRequest{
		Id:         uuid.New(),
		Slut:       helper.Random_generate_string(10),
		CustomerId: uuid.MustParse(in.CustomerId),
		FlightId:   uuid.MustParse(in.FlightId),
	}
	res, err := h.bookingRepositories.CreateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	bRes := &pb.BookingResponse{
		Id:         res.Id.String(),
		CustomerId: res.CustomerId.String(),
		FlightId:   res.FlightId.String(),
		Code:       res.Code,
		Status:     res.Status,
		BookedDate: timestamppb.New(res.Booked_date),
	}

	if err != nil {
		return nil, err
	}
	bRes.BookedDate = timestamppb.New(res.Booked_date)
	return bRes, nil
}

func (h *BookingHandler) FindBooking(ctx context.Context, in *pb.BookingCode) (*pb.BookingInfo, error) {
	bReq, err := h.bookingRepositories.FindBooking(ctx, in.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	cutomerInfo, err := h.customersClient.FindCustomerById(ctx, &pb.CustomerAuthen{Id: bReq.CustomerId.String()})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "customer_id not found")
		}
		return nil, err
	}
	flightInfo, err := h.flightClient.FindFlight(ctx, &pb.FlightRequest{Id: bReq.FlightId.String()})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "flight_id not found")
		}
		return nil, err
	}
	return &pb.BookingInfo{
		Id:         bReq.Id.String(),
		Customer:   cutomerInfo,
		Flight:     flightInfo,
		Code:       bReq.Code,
		Status:     bReq.Status,
		BookedDate: timestamppb.New(bReq.Booked_date),
		CreatedAt:  timestamppb.New(bReq.CreatedAt),
		UpdatedAt:  timestamppb.New(bReq.UpdatedAt),
	}, nil
}

func (h *BookingHandler) FindBookingByCustomerId(ctx context.Context, in *pb.CustomeId) (*pb.ListMyBookingResponse, error) {
	bReq, err := h.bookingRepositories.FindBookingByCustomerId(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	listBooking := []*pb.BookingResponse{}
	for _, v := range bReq {
		listBooking = append(listBooking, &pb.BookingResponse{
			Id:         v.Id.String(),
			Status:     v.Status,
			Code:       v.Code,
			CustomerId: v.CustomerId.String(),
			FlightId:   v.FlightId.String(),
			BookedDate: timestamppb.New(v.Booked_date),
		})
	}
	return &pb.ListMyBookingResponse{Bookings: listBooking}, nil
}

func (h *BookingHandler) CancelBooking(ctx context.Context, in *pb.BookingCode) (*pb.BookingResponse, error) {
	bReq, err := h.bookingRepositories.FindBooking(ctx, in.Code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	_, err = h.bookingRepositories.UpdateBooking(ctx, &models.Booking{Status: "Canceled"})
	if err != nil {
		return nil, err
	}
	return &pb.BookingResponse{
		Id:         bReq.Id.String(),
		CustomerId: bReq.CustomerId.String(),
		FlightId:   bReq.FlightId.String(),
		Code:       bReq.Code,
		Status:     "Canceled",
		BookedDate: timestamppb.New(bReq.Booked_date),
	}, nil
}
