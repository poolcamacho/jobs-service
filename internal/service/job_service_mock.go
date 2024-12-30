package service

import (
	"github.com/poolcamacho/jobs-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockJobService is a mock implementation of JobService for testing
type MockJobService struct {
	mock.Mock
}

// GetAllJobs mocks the GetAllJobs method
// @return []*domain.Job - A slice of jobs
// @return error - An error if the operation fails
func (m *MockJobService) GetAllJobs() ([]*domain.Job, error) {
	args := m.Called()
	if jobs, ok := args.Get(0).([]*domain.Job); ok {
		return jobs, args.Error(1)
	}
	return nil, args.Error(1)
}

// AddJob mocks the AddJob method
// @param job *domain.Job - The job data to be added
// @return error - An error if the operation fails
func (m *MockJobService) AddJob(job *domain.Job) error {
	args := m.Called(job)
	return args.Error(0)
}
