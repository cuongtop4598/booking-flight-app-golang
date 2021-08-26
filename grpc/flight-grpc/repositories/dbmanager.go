package repositories

import (
	"context"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/database"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/flight-grpc/models"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/flight-grpc/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlightRepository interface {
	GetFlightById(ctx context.Context, Id uuid.UUID) (*models.Flight, error)
	GetFlightBySlut(ctx context.Context, slut string) (*models.Flight, error)
	CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error)
	DeleteFlight(ctx context.Context, model *models.Flight) error
	ListFlight(ctx context.Context, req *requests.ListFlightRequest) ([]*models.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Flight{},
	)

	if err != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) GetFlightById(ctx context.Context, Id uuid.UUID) (*models.Flight, error) {
	flight := models.Flight{}
	if err := m.Where(&models.Flight{Id: Id}).First(&flight).Error; err != nil {
		return nil, err
	}
	return &flight, nil
}

func (m *dbmanager) GetFlightBySlut(ctx context.Context, slut string) (*models.Flight, error) {
	flight := models.Flight{}
	if err := m.Where(&models.Flight{Slut: slut}).First(&flight).Error; err != nil {
		return nil, err
	}
	return &flight, nil
}

func (m *dbmanager) CreateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) UpdateFlight(ctx context.Context, model *models.Flight) (*models.Flight, error) {
	if err := m.Where(&models.Flight{Id: model.Id}).Updates(&models.Flight{Name: model.Name, From: model.From, To: model.To}).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) DeleteFlight(ctx context.Context, model *models.Flight) error {
	if err := m.Where(&models.Flight{Id: model.Id}).Delete(&models.Flight{Name: model.Name}).Error; err != nil {
		return err
	}
	return nil
}

func (m *dbmanager) ListFlight(ctx context.Context, req *requests.ListFlightRequest) ([]*models.Flight, error) {
	flights := []*models.Flight{}

	if err := m.Where(&models.Flight{Name: req.Name, From: req.From, To: req.To, Date: req.Date}).Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}
