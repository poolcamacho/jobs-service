package service

import (
	"errors"
	"testing"

	"github.com/poolcamacho/jobs-service/internal/domain"
	"github.com/poolcamacho/jobs-service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetAllJobs(t *testing.T) {
	// Setup
	mockRepo := new(repository.MockJobRepository)
	jobService := NewJobService(mockRepo)

	// Mock data
	jobs := []*domain.Job{
		{
			ID:          1,
			Title:       "Software Engineer",
			Description: "Develop and maintain software.",
			SalaryRange: "4000-6000",
		},
		{
			ID:          2,
			Title:       "Product Manager",
			Description: "Oversee product lifecycle.",
			SalaryRange: "5000-7000",
		},
	}

	// Mock behavior
	mockRepo.On("FindAll").Return(jobs, nil)

	// Execute
	result, err := jobService.GetAllJobs()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, jobs, result)
	mockRepo.AssertExpectations(t)
}

func TestAddJob(t *testing.T) {
	// Setup
	mockRepo := new(repository.MockJobRepository)
	jobService := NewJobService(mockRepo)

	// Mock data
	newJob := &domain.Job{
		Title:       "DevOps Engineer",
		Description: "Manage cloud infrastructure.",
		SalaryRange: "4500-6500",
	}

	// Mock behavior
	mockRepo.On("Create", newJob).Return(nil)

	// Execute
	err := jobService.AddJob(newJob)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddJob_Error(t *testing.T) {
	// Setup
	mockRepo := new(repository.MockJobRepository)
	jobService := NewJobService(mockRepo)

	// Mock data
	newJob := &domain.Job{
		Title:       "QA Engineer",
		Description: "Ensure software quality.",
		SalaryRange: "3500-5000",
	}

	// Mock behavior
	mockRepo.On("Create", newJob).Return(errors.New("database error"))

	// Execute
	err := jobService.AddJob(newJob)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "database error")
	mockRepo.AssertExpectations(t)
}
