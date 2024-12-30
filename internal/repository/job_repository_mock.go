package repository

import (
	"github.com/poolcamacho/jobs-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

// MockJobRepository is a mock implementation of JobRepository for testing
type MockJobRepository struct {
	mock.Mock
}

// FindAll mocks the FindAll method
// Simulates the retrieval of all jobs from the database
func (m *MockJobRepository) FindAll() ([]*domain.Job, error) {
	args := m.Called()
	if jobs, ok := args.Get(0).([]*domain.Job); ok {
		return jobs, args.Error(1)
	}
	return nil, args.Error(1)
}

// Create mocks the Create method
// Simulates the insertion of a new job into the database
func (m *MockJobRepository) Create(job *domain.Job) error {
	args := m.Called(job)
	return args.Error(0)
}
