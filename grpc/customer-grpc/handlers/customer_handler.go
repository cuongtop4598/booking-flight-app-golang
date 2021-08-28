package handlers

import (
	"context"
	"database/sql"
	"time"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/customer-grpc/models"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/customer-grpc/repositories"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/customer-grpc/requests"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerHandler struct {
	pb.UnimplementedCustormerServer
	bookingClient      pb.SEOLEBookingClient
	flightClient       pb.FlightClient
	customerRepository repositories.CustomerRepository
}

func NewCustomerHandler(bookingClient pb.SEOLEBookingClient, flightClient pb.FlightClient, customerRepository repositories.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		bookingClient:      bookingClient,
		flightClient:       flightClient,
		customerRepository: customerRepository,
	}, nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	cRequest := &models.Customer{
		Id:          uuid.New(),
		Slut:        in.Slut,
		Name:        in.Name,
		Address:     in.Address,
		LicenseId:   in.LicenseId,
		PhoneNumber: in.PhoneNumber,
		Email:       in.Email,
		Password:    in.Password,
		Active:      true,
	}
	customer, err := h.customerRepository.CreateCustomer(ctx, cRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	cResponse := &pb.CustomerResponse{
		Id:          customer.Id.String(),
		Name:        customer.Name,
		Address:     customer.Address,
		LicenseId:   customer.LicenseId,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		Password:    customer.Password,
		Active:      customer.Active,
		CreatedAt:   timestamppb.New(time.Now()),
		UpdatedAt:   timestamppb.New(time.Now()),
	}
	return cResponse, nil
}

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	customer, err := h.customerRepository.GetCustomerByPhone(ctx, in.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Name != "" {
		customer.Name = in.Name
	}
	if in.Address != "" {
		customer.Address = in.Address
	}
	if in.PhoneNumber != "" {
		customer.PhoneNumber = in.PhoneNumber
	}
	if in.Email != "" {
		customer.Email = in.Email
	}
	if in.LicenseId != "" {
		customer.LicenseId = in.LicenseId
	}
	newCustomer, err := h.customerRepository.UpdateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}
	cResponse := &pb.CustomerResponse{
		Id:          newCustomer.Id.String(),
		Slut:        newCustomer.Slut,
		Name:        newCustomer.Name,
		Address:     newCustomer.Address,
		LicenseId:   newCustomer.LicenseId,
		PhoneNumber: newCustomer.PhoneNumber,
		Email:       newCustomer.Email,
		Password:    newCustomer.Password,
		Active:      newCustomer.Active,
		CreatedAt:   timestamppb.New(newCustomer.CreatedAt),
		UpdatedAt:   timestamppb.New(newCustomer.UpdatedAt),
	}
	return cResponse, nil
}

func (h *CustomerHandler) FindCustomerByPhone(ctx context.Context, in *pb.CustomerAuthen) (*pb.CustomerResponse, error) {
	customer, err := h.customerRepository.GetCustomerByPhone(ctx, in.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	cRes := &pb.CustomerResponse{
		Id:          customer.Id.String(),
		Slut:        customer.Slut,
		Name:        customer.Name,
		Address:     customer.Address,
		LicenseId:   customer.LicenseId,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		Active:      customer.Active,
		CreatedAt:   timestamppb.New(customer.CreatedAt),
		UpdatedAt:   timestamppb.New(customer.UpdatedAt),
	}
	return cRes, nil
}
func (h *CustomerHandler) FindCustomerById(ctx context.Context, in *pb.CustomerAuthen) (*pb.CustomerResponse, error) {
	customer, err := h.customerRepository.GetCustomerById(ctx, in.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	cRes := &pb.CustomerResponse{
		Id:          customer.Id.String(),
		Slut:        customer.Slut,
		Name:        customer.Name,
		Address:     customer.Address,
		LicenseId:   customer.LicenseId,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		Active:      customer.Active,
		CreatedAt:   timestamppb.New(customer.CreatedAt),
		UpdatedAt:   timestamppb.New(customer.UpdatedAt),
	}
	return cRes, nil
}

func (h *CustomerHandler) ListCustomer(ctx context.Context, in *pb.ListCustomerRequest) (*pb.ListCustomerResponse, error) {
	customers, err := h.customerRepository.ListCustomer(ctx, &requests.ListCustomerRequests{Active: in.Active})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	cRes := &pb.ListCustomerResponse{
		ResCustomers: []*pb.CustomerResponse{},
	}
	err = copier.CopyWithOption(&cRes.ResCustomers, &customers, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return nil, err
	}
	return cRes, nil
}

func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.CustomerAuthen) (*pb.Empty, error) {
	customer, err := h.customerRepository.GetCustomerByPhone(ctx, in.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	if in.Password != "" {
		customer.Password = in.Password
	}
	return &pb.Empty{}, nil
}

func (h *CustomerHandler) BookingHistory(ctx context.Context, in *pb.CustomerAuthen) (*pb.HistoryResponse, error) {
	customer, err := h.customerRepository.GetCustomerByPhone(ctx, in.Phone)
	if err != nil {
		if err == nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, err.Error())
			}
			return nil, err
		}
	}
	bookingInfo, err := h.bookingClient.FindBookingByCustomerId(ctx, &pb.CustomeId{Id: customer.Id.String()})
	if err != nil {
		if err == nil {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, "booking not found")
			}
			return nil, err
		}
	}
	historysRes := &pb.HistoryResponse{}

	for _, v := range bookingInfo.Bookings {
		tmp, err := h.flightClient.FindFlight(ctx, &pb.FlightRequest{Id: v.FlightId})
		if err != nil {
			if err == nil {
				if err == sql.ErrNoRows {
					return nil, status.Error(codes.NotFound, "flight not found")
				}
				return nil, err
			}
		}
		historysRes.Historys = append(historysRes.Historys, &pb.History{
			BookingCode: v.Code,
			BookingDate: v.BookedDate.String(),
			Status:      v.Status,
			Flight:      tmp,
		})
	}

	return historysRes, nil
}
