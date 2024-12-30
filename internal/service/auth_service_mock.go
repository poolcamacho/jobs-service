package service

import (
	"github.com/poolcamacho/auth-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockAuthService is a mock implementation of AuthService for testing
type MockAuthService struct {
	mock.Mock
}

// Register mocks the Register method
func (m *MockAuthService) Register(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// Login mocks the Login method
func (m *MockAuthService) Login(email, password string) (string, error) {
	args := m.Called(email, password)
	return args.String(0), args.Error(1)
}
