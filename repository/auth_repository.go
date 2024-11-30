package repository

import (
	"log"

	"github.com/MarcKVR/mortgage/domain"
	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		Login(email, password string) (bool, error)
	}

	authRepository struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewAuthRepository(db *gorm.DB, log *log.Logger) AuthRepository {
	return &authRepository{
		db:  db,
		log: log,
	}
}

func (repo *authRepository) Login(email, password string) (bool, error) {
	var user domain.User

	err := repo.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		repo.log.Printf("Error: %v", err)
		return false, err
	}

	return true, nil
}
