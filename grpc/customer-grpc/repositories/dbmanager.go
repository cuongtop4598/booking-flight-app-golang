package repositories

import (
	"context"

	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/database"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/customer-grpc/models"
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/grpc/customer-grpc/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomerByEmail(ctx context.Context, email string) (*models.Customer, error)
	GetCustomerByPhone(ctx context.Context, phone string) (*models.Customer, error)
	GetCustomerById(ctx context.Context, id string) (*models.Customer, error)
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, model *models.Customer) error
	ListCustomer(ctx context.Context, req *requests.ListCustomerRequests) ([]*models.Customer, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Customer{},
	)

	if err != nil {
		return nil, err
	}
	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) GetCustomerByEmail(ctx context.Context, email string) (*models.Customer, error) {
	customer := models.Customer{}
	if err := m.Where(&models.Customer{Email: email}).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (m *dbmanager) GetCustomerByPhone(ctx context.Context, phone string) (*models.Customer, error) {
	customer := models.Customer{}
	if err := m.Where(&models.Customer{PhoneNumber: phone}).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
func (m *dbmanager) GetCustomerById(ctx context.Context, id string) (*models.Customer, error) {
	customer := models.Customer{}
	if err := m.Where(&models.Customer{Id: uuid.MustParse(id)}).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
func (m *dbmanager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := m.Where(&models.Customer{Id: model.Id}).Updates(&models.Customer{Name: model.Name, Email: model.Email, Active: model.Active, Address: model.Address, LicenseId: model.LicenseId, Password: model.Password, PhoneNumber: model.PhoneNumber}).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) DeleteCustomer(ctx context.Context, model *models.Customer) error {
	if err := m.Where(&models.Customer{Id: model.Id}).Delete(&models.Customer{Name: model.Name}).Error; err != nil {
		return err
	}
	return nil
}

func (m *dbmanager) ListCustomer(ctx context.Context, req *requests.ListCustomerRequests) ([]*models.Customer, error) {
	Customers := []*models.Customer{}

	if err := m.Where(&models.Customer{Active: req.Active}).Find(&Customers).Error; err != nil {
		return nil, err
	}
	return Customers, nil
}
