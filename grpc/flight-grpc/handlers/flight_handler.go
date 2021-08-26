package handlers

import (
	"context"
	"database/sql"
	"time"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/flight-grpc/models"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/flight-grpc/repositories"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/flight-grpc/requests"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler struct {
	pb.UnimplementedFlightServer
	flightRepository repositories.FlightRepository
}

func NewFlighHandler(flightRepository repositories.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepository,
	}, nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, in *pb.FlightRequest) (*pb.FlightResponse, error) {
	fRequest := &models.Flight{
		Id:            uuid.New(),
		Slut:          in.Slut,
		Name:          in.Name,
		From:          in.From,
		To:            in.To,
		Status:        in.Status,
		Date:          time.Now(),
		AvailableSlot: in.AvailableSlot,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	flight, err := h.flightRepository.CreateFlight(ctx, fRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	fResponse := &pb.FlightResponse{
		Id:     flight.Id.String(),
		Slut:   flight.Slut,
		Name:   flight.Name,
		From:   flight.From,
		To:     flight.To,
		Date:   timestamppb.New(flight.Date),
		Status: flight.Status,
	}
	return fResponse, nil
}

func (h *FlightHandler) UpdateFlight(ctx context.Context, in *pb.FlightRequest) (*pb.FlightResponse, error) {
	flight, err := h.flightRepository.GetFlightById(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Slut != "" {
		flight.Slut = in.Slut
	}
	if in.Status != "" {
		flight.Name = in.Name
	}
	if in.From != "" {
		flight.From = in.From
	}
	if in.To != "" {
		flight.To = in.To
	}

	if in.Date.AsTime().After(time.Now()) {
		flight.Date = in.Date.AsTime()
	}
	if in.Status != "" {
		flight.Status = in.Status
	}
	if in.AvailableSlot >= 0 {
		flight.AvailableSlot = in.AvailableSlot
	}
	newFlight, err := h.flightRepository.UpdateFlight(ctx, flight)
	if err != nil {
		return nil, err
	}
	fResponse := &pb.FlightResponse{
		Id:            newFlight.Id.String(),
		Slut:          newFlight.Slut,
		Name:          newFlight.Name,
		From:          newFlight.From,
		To:            newFlight.To,
		Date:          timestamppb.New(newFlight.Date),
		Status:        "",
		AvailableSlot: 0,
		CreatedAt:     timestamppb.New(newFlight.CreatedAt),
		UpdatedAt:     timestamppb.New(time.Now()),
	}
	return fResponse, nil
}

func (h *FlightHandler) FindFlight(ctx context.Context, in *pb.FlightRequest) (*pb.FlightResponse, error) {
	flight, err := h.flightRepository.GetFlightBySlut(ctx, in.Slut)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	fres := &pb.FlightResponse{
		Id:            flight.Id.String(),
		Slut:          flight.Slut,
		Name:          flight.Name,
		From:          flight.From,
		To:            flight.To,
		Date:          timestamppb.New(flight.Date),
		Status:        flight.Status,
		AvailableSlot: flight.AvailableSlot,
		CreatedAt:     timestamppb.New(flight.CreatedAt),
		UpdatedAt:     timestamppb.New(flight.UpdatedAt),
	}
	return fres, nil
}

func (h *FlightHandler) FindListFlight(ctx context.Context, in *pb.FlightRequest) (*pb.ListFlightResponse, error) {
	flights, err := h.flightRepository.ListFlight(ctx, &requests.ListFlightRequest{
		Name: in.Name,
		From: in.From,
		To:   in.To,
		Date: in.Date.AsTime(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	fres := &pb.ListFlightResponse{
		Flights: []*pb.FlightResponse{},
	}

	err = copier.CopyWithOption(&fres.Flights, &flights, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return nil, err
	}
	return fres, nil
}
