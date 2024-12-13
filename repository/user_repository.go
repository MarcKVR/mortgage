package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/MarcKVR/mortgage/domain"
	"gorm.io/gorm"
)

type (
	Filters struct {
		Name  string
		Email string
	}

	UserRepository interface {
		Create(user *domain.User) error
		Get(id string) (*domain.User, error)
		GetUsers(filters Filters, limit, offset int) ([]domain.User, error)
		Count(filters Filters) (int, error)
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

	return nil
}

func (repo *userRepository) Get(id string) (*domain.User, error) {
	user := domain.User{ID: id}

	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetUsers(filters Filters, limit, offset int) ([]domain.User, error) {
	var users []domain.User

	query := repo.db.Model((&users))
	query = applyFilters(query, filters)
	query = query.Offset(offset).Limit(limit)
	result := query.Order("id desc").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repo *userRepository) Count(filters Filters) (int, error) {
	var total int64
	user := repo.db.Model(domain.User{})
	user = applyFilters(user, filters)

	if err := user.Count(&total).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}

func applyFilters(user *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		user = user.Where("LOWER(name) LIKE ?", filters.Name)
	}

	if filters.Email != "" {
		filters.Email = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Email))
		user = user.Where("LOWER(email) LIKE ?", filters.Email)
	}

	return user
}
