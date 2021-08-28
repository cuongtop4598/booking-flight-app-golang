package repositories

import (
	"context"
	"time"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/database"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/booking-grpc/models"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/booking-grpc/requests"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingRepositories interface {
	CreateBooking(ctx context.Context, req *requests.BookingRequest) (*models.Booking, error)
	FindBooking(ctx context.Context, bookingCode string) (*models.Booking, error)
	UpdateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error)
	FindBookingByCustomerId(ctx context.Context, customerId string) ([]*models.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&models.Booking{},
	)
	if err != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreateBooking(ctx context.Context, req *requests.BookingRequest) (*models.Booking, error) {

	booking := &models.Booking{
		Id:          uuid.New(),
		Slut:        req.Slut,
		CustomerId:  req.CustomerId,
		FlightId:    req.FlightId,
		Code:        helper.Random_generate_string(10),
		Status:      "OK",
		Booked_date: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := m.Create(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (m *dbmanager) FindBooking(ctx context.Context, bookingCode string) (*models.Booking, error) {
	booking := models.Booking{}
	if err := m.Where(&models.Booking{Code: bookingCode}).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
func (m *dbmanager) FindBookingByCustomerId(ctx context.Context, customerId string) ([]*models.Booking, error) {
	booking := []*models.Booking{}
	if err := m.Where(&models.Booking{CustomerId: uuid.MustParse(customerId)}).Find(&booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (m *dbmanager) UpdateBooking(ctx context.Context, model *models.Booking) (*models.Booking, error) {
	if err := m.Where(&models.Booking{Code: model.Code}).Updates(&models.Booking{Status: model.Status, Booked_date: model.Booked_date, UpdatedAt: model.UpdatedAt}).Error; err != nil {
		return nil, err
	}
	return model, nil
}
