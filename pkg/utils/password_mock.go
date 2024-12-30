package utils

import (
	"github.com/stretchr/testify/mock"
)

// MockPasswordUtils is a mock implementation of PasswordUtils
type MockPasswordUtils struct {
	mock.Mock
}

func (m *MockPasswordUtils) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordUtils) CheckPassword(hashedPassword, plainPassword string) error {
	args := m.Called(hashedPassword, plainPassword)
	return args.Error(0)
}
