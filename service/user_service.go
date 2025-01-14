package service

import (
	"errors"
	"log"

	"github.com/MarcKVR/mortgage/domain"
	"github.com/MarcKVR/mortgage/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	Filters struct {
		Name  string
		Email string
	}

	UserService interface {
		Create(user *domain.User) (*domain.User, error)
		Get(id string) (*domain.User, error)
		GetUsers(filters Filters, limit, offset int) ([]domain.User, error)
		Count(filters Filters) (int, error)
		Update(id string, user *domain.User) error
	}

	userService struct {
		repo repository.UserRepository
		log  *log.Logger
	}
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func NewUserService(repo repository.UserRepository, log *log.Logger) UserService {
	return &userService{
		repo: repo,
		log:  log,
	}
}

func (s *userService) Create(user *domain.User) (*domain.User, error) {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	if err := s.repo.Create(user); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return user, nil
}

func (s *userService) Get(id string) (*domain.User, error) {
	return s.repo.Get(id)
}

func (s *userService) GetUsers(filters Filters, limit, offset int) ([]domain.User, error) {
	repoFilters := repository.Filters{
		Name:  filters.Name,
		Email: filters.Email,
	}
	users, err := s.repo.GetUsers(repoFilters, limit, offset)
	return users, err
}

func (s *userService) Count(filters Filters) (int, error) {
	repoFilters := repository.Filters{
		Name:  filters.Name,
		Email: filters.Email,
	}
	return s.repo.Count(repoFilters)
}

func (s *userService) Update(id string, user *domain.User) error {

	existentUser, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	if existentUser == nil {
		return errors.New("user not found")
	}

	existentUser.Name = user.Name
	existentUser.Email = user.Email
	if user.Password != "" {
		existentUser.Password = user.Password
	}

	return s.repo.Update(existentUser)
}
