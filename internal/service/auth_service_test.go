package service

import (
	"errors"
	"testing"

	"github.com/poolcamacho/auth-service/internal/domain"
	"github.com/poolcamacho/auth-service/internal/repository"
	_ "github.com/poolcamacho/auth-service/pkg/utils"
	"github.com/stretchr/testify/assert"
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

func TestRegister_Success(t *testing.T) {
	// Arrange
	mockRepo := new(repository.MockUserRepository)
	mockPasswordUtil := new(MockPasswordUtils)
	authService := NewAuthService(mockRepo, "mock-secret", mockPasswordUtil)

	user := &domain.User{
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "plaintextpassword",
		Role:         "user",
	}

	mockPasswordUtil.On("HashPassword", "plaintextpassword").Return("hashedpassword", nil)
	mockRepo.On("Create", mock.Anything).Return(nil)

	// Act
	err := authService.Register(user)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "hashedpassword", user.PasswordHash)
	mockPasswordUtil.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestRegister_ErrorHashingPassword(t *testing.T) {
	// Arrange
	mockRepo := new(repository.MockUserRepository)
	mockPasswordUtil := new(MockPasswordUtils)
	authService := NewAuthService(mockRepo, "mock-secret", mockPasswordUtil)

	mockPasswordUtil.On("HashPassword", "plaintextpassword").Return("", errors.New("hashing error"))

	user := &domain.User{
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "plaintextpassword",
		Role:         "user",
	}

	// Act
	err := authService.Register(user)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "hashing error", err.Error())
	mockRepo.AssertNotCalled(t, "Create", mock.Anything)
	mockPasswordUtil.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	// Arrange
	mockRepo := new(repository.MockUserRepository)
	mockPasswordUtil := new(MockPasswordUtils)
	authService := NewAuthService(mockRepo, "mock-secret", mockPasswordUtil)

	email := "test@example.com"
	password := "plaintextpassword"
	hashedPassword := "hashedpassword"

	user := &domain.User{
		ID:           1,
		Username:     "testuser",
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         "user",
	}

	mockRepo.On("FindByEmail", email).Return(user, nil)
	mockPasswordUtil.On("CheckPassword", hashedPassword, password).Return(nil)

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
	mockPasswordUtil.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Arrange
	mockRepo := new(repository.MockUserRepository)
	mockPasswordUtil := new(MockPasswordUtils)
	authService := NewAuthService(mockRepo, "mock-secret", mockPasswordUtil)

	email := "test@example.com"
	password := "wrongpassword"

	mockRepo.On("FindByEmail", email).Return(nil, errors.New("not found"))

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	mockPasswordUtil.AssertExpectations(t)
}

func TestLogin_InvalidPassword(t *testing.T) {
	// Arrange
	mockRepo := new(repository.MockUserRepository)
	mockPasswordUtil := new(MockPasswordUtils)
	authService := NewAuthService(mockRepo, "mock-secret", mockPasswordUtil)

	email := "test@example.com"
	password := "wrongpassword"
	hashedPassword := "hashedpassword"

	user := &domain.User{
		ID:           1,
		Username:     "testuser",
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         "user",
	}

	mockRepo.On("FindByEmail", email).Return(user, nil)
	mockPasswordUtil.On("CheckPassword", hashedPassword, password).Return(errors.New("invalid password"))

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidCredentials, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
	mockPasswordUtil.AssertExpectations(t)
}
