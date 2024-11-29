package repository

import (
	"log"

	"github.com/MarcKVR/mortgage/domain"
	"gorm.io/gorm"
)

type (
	PaymentRepository interface {
		Create(payment *domain.Payment) error
		Get(id string) (*domain.Payment, error)
	}

	paymentRepository struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepository(db *gorm.DB, log *log.Logger) PaymentRepository {
	return &paymentRepository{
		db:  db,
		log: log,
	}
}

func (repo *paymentRepository) Create(payment *domain.Payment) error {
	if err := repo.db.Create(payment).Error; err != nil {
		repo.log.Printf("Error: %v", err)
	}
	repo.log.Println("Payment was with id: received successfully", payment.ID)

	return nil
}

func (repo *paymentRepository) Get(id string) (*domain.Payment, error) {
	payment := domain.Payment{ID: id}

	if err := repo.db.First(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}
