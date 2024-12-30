package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/poolcamacho/jobs-service/internal/domain"
	"github.com/poolcamacho/jobs-service/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetJobs(t *testing.T) {
	// Setup
	mockJobService := new(service.MockJobService)
	jobHandler := NewJobHandler(mockJobService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/jobs", jobHandler.GetJobs)

	// Test data
	mockJobs := []*domain.Job{
		{ID: 1, Title: "Software Engineer", Description: "Develop and maintain software", SalaryRange: "60K-80K"},
		{ID: 2, Title: "Data Scientist", Description: "Analyze data and build models", SalaryRange: "80K-100K"},
	}

	// Mock behavior
	mockJobService.On("GetAllJobs").Return(mockJobs, nil)

	// Prepare HTTP request
	req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	expectedResponse, _ := json.Marshal(mockJobs)
	assert.JSONEq(t, string(expectedResponse), rec.Body.String())
	mockJobService.AssertExpectations(t)
}

func TestCreateJob(t *testing.T) {
	// Setup
	mockJobService := new(service.MockJobService)
	jobHandler := NewJobHandler(mockJobService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/jobs", jobHandler.CreateJob)

	// Test data
	requestBody := domain.Job{
		Title:       "Product Manager",
		Description: "Manage product lifecycle and team collaboration",
		SalaryRange: "90K-110K",
	}

	// Mock behavior
	mockJobService.On("AddJob", mock.AnythingOfType("*domain.Job")).Return(nil)

	// Prepare HTTP request
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"message":"job created successfully"}`, rec.Body.String())
	mockJobService.AssertExpectations(t)
}

func TestCreateJob_BadRequest(t *testing.T) {
	// Setup
	mockJobService := new(service.MockJobService)
	jobHandler := NewJobHandler(mockJobService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/jobs", jobHandler.CreateJob)

	// Test data: Invalid request body (missing required fields)
	invalidRequestBody := `{"title":"Job without description"}` // Description and SalaryRange are missing

	// Prepare HTTP request
	req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewBufferString(invalidRequestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")

	// Verify no interactions with the mock service
	mockJobService.AssertNotCalled(t, "AddJob")
}

func TestGetJobs_InternalServerError(t *testing.T) {
	// Setup
	mockJobService := new(service.MockJobService)
	jobHandler := NewJobHandler(mockJobService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/jobs", jobHandler.GetJobs)

	// Mock behavior
	mockJobService.On("GetAllJobs").Return(nil, assert.AnError)

	// Prepare HTTP request
	req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	rec := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
	mockJobService.AssertExpectations(t)
}
