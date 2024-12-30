package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/poolcamacho/auth-service/internal/domain"
	"github.com/poolcamacho/auth-service/internal/service"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	// Setup
	mockAuthService := new(service.MockAuthService)
	authHandler := NewAuthHandler(mockAuthService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", authHandler.Register)

	// Test data
	requestBody := domain.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	user := &domain.User{
		Username:     requestBody.Username,
		Email:        requestBody.Email,
		Role:         "user",
		PasswordHash: requestBody.Password,
	}

	// Mock behavior
	mockAuthService.On("Register", user).Return(nil)

	// Prepare HTTP request
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"message":"user registered successfully"}`, rec.Body.String())
	mockAuthService.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	// Setup
	mockAuthService := new(service.MockAuthService)
	authHandler := NewAuthHandler(mockAuthService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", authHandler.Login)

	// Test data
	requestBody := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	expectedToken := "mocked_jwt_token"

	// Mock behavior
	mockAuthService.On("Login", "test@example.com", "password123").Return(expectedToken, nil)

	// Prepare HTTP request
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message":"login successful", "token":"mocked_jwt_token"}`, rec.Body.String())
	mockAuthService.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Setup
	mockAuthService := new(service.MockAuthService)
	authHandler := NewAuthHandler(mockAuthService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", authHandler.Login)

	// Test data
	requestBody := map[string]string{
		"email":    "wrong@example.com",
		"password": "wrongpassword",
	}

	// Mock behavior
	mockAuthService.On("Login", "wrong@example.com", "wrongpassword").Return("", service.ErrInvalidCredentials)

	// Prepare HTTP request
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, `{"error":"invalid credentials"}`, rec.Body.String())
	mockAuthService.AssertExpectations(t)
}
