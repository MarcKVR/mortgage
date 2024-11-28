package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID              string         `json:"id" gorm:"type:char(36);not null;primaryKey;unique_index"`
	PaymentNumber   int32          `json:"payment_number" gorm:"type:int;not null"`
	DatePayment     time.Time      `json:"date_payment"`
	Rate            float32        `json:"rate" gorm:"type:decimal(10,2);not null"`
	MonthlyPayment  float64        `json:"monthly_payment" gorm:"type:decimal(10,2);not null"`
	DamageInsurance float64        `json:"damage_insurance" gorm:"type:decimal(10,2);not null"`
	LifeInsurance   float64        `json:"life_insurance" gorm:"type:decimal(10,2);not null"`
	Interests       float64        `json:"interests" gorm:"type:decimal(10,2);not null"`
	Capital         float64        `json:"capital" gorm:"type:decimal(10,2);not null"`
	CreatedAt       *time.Time     `json:"-"`
	UpdatedAt       *time.Time     `json:"-"`
	Deleted         gorm.DeletedAt `json:"-"`
}

func (c *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}

	return
}
