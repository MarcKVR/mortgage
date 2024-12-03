package auth_test

import (
	"testing"

	"github.com/MarcKVR/mortgage/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) ValidateUser(email, password string) bool {
	args := m.Called(email, password)
	return args.Bool(0)
}

func Login(repo *MockAuthRepository, email, password string) (string, error) {
	if repo.ValidateUser(email, password) {
		return auth.GenerateToken(email)
	}

	return "", nil
}

func TestLogin(t *testing.T) {
	repo := new(MockAuthRepository)

	// Success Case
	repo.On("ValidateUser", "admin", "password").Return(true)
	token, err := Login(repo, "admin", "password")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Failure Case
	repo.On("ValidateUser", "user", "wrongpassword").Return(false)
	token, err = Login(repo, "user", "wrongpassword")
	assert.NoError(t, err)
	assert.Empty(t, token) // El token debe estar vacío
}
