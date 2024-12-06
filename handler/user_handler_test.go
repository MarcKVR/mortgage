package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MarcKVR/mortgage/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(id string) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserService) Create(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestUserHandler_Get(t *testing.T) {
	// Create a new instance of our mock service
	mockUserService := new(MockUserService)

	// Create a new Fiber app
	app := fiber.New()

	// Create a new UserHandler with the mock service
	userHandler := &UserHandler{
		service: mockUserService,
	}

	// Register the handler
	app.Get("/user/:id", userHandler.Get)

	// Define the test cases
	tests := []struct {
		name           string
		userID         string
		mockUser       *domain.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "User found",
			userID:         "1",
			mockUser:       &domain.User{ID: "1", Name: "John Doe", Email: "jonh@gmail.com"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":200,"data":{"email":"jonh@gmail.com","id":"1","name":"John Doe"}}`,
		},
		{
			name:           "User not found",
			userID:         "2",
			mockUser:       nil,
			mockError:      fiber.ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"status":404,"error":"Not Found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserService.ExpectedCalls = nil

			// Set up the expected calls and return values for the mock service
			mockUserService.On("Get", tt.userID).Return(tt.mockUser, tt.mockError)

			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodGet, "/user/"+tt.userID, nil)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Requesf failed: %v", err)
			}
			defer resp.Body.Close()

			// Assert the response status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Assert the response body
			body, err := io.ReadAll(resp.Body)

			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}
			assert.JSONEq(t, tt.expectedBody, string(body))

			// Assert that the expectations were met
			mockUserService.AssertExpectations(t)
		})
	}
}
