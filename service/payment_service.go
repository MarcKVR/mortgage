package service

import (
	"log"
	"time"

	"github.com/MarcKVR/mortgage/domain"
	"github.com/MarcKVR/mortgage/repository"
)

type (
	PaymentService interface {
		Create(monthlyPayment, damageInsurance, lifeInsurance, interests, capital float64,
			rate float32,
			paymentNumber int32,
			datePayment string) (*domain.Payment, error)
	}

	service struct {
		log  *log.Logger
		repo repository.PaymentRepository
	}
)

func NewService(log *log.Logger, repo repository.PaymentRepository) PaymentService {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s *service) Create(monthlyPayment, damageInsurance, lifeInsurance, interests, capital float64,
	rate float32,
	paymentNumber int32,
	datePayment string) (*domain.Payment, error) {

	datePaymentParsed, err := time.Parse("2006-01-02", datePayment)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	payment := &domain.Payment{
		MonthlyPayment:  monthlyPayment,
		DamageInsurance: damageInsurance,
		LifeInsurance:   lifeInsurance,
		Interests:       interests,
		Capital:         capital,
		Rate:            rate,
		PaymentNumber:   paymentNumber,
		DatePayment:     datePaymentParsed,
	}

	if err := s.repo.Create(payment); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return payment, nil
}
