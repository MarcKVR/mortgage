package service

import (
	"log"

	"github.com/MarcKVR/mortgage/repository"
)

type (
	AuthService interface {
		Login(email, password string) (bool, error)
	}

	authService struct {
		repo repository.AuthRepository
		log  *log.Logger
	}
)

func NewAuthService(repo repository.AuthRepository, log *log.Logger) AuthService {
	return &authService{
		repo: repo,
		log:  log,
	}
}

func (s *authService) Login(email, password string) (bool, error) {
	return s.repo.Login(email, password)
}
