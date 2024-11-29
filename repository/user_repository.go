package repository

import (
	"log"

	"github.com/MarcKVR/mortgage/domain"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(user *domain.User) error
		Get(id string) (*domain.User, error)
	}

	userRepository struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewUserRepository(db *gorm.DB, log *log.Logger) UserRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}

func (repo *userRepository) Create(user *domain.User) error {

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Printf("Error: %v", err)
	}
	// repo.log.Println("User was with id: received successfully", user.ID)

	return nil
}

func (repo *userRepository) Get(id string) (*domain.User, error) {
	user := domain.User{ID: id}

	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
