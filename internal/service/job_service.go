package service

import (
	"github.com/poolcamacho/jobs-service/internal/domain"
	"github.com/poolcamacho/jobs-service/internal/repository"
)

// JobService defines methods for job-related operations
// This interface abstracts the business logic for managing jobs.
type JobService interface {
	GetAllJobs() ([]*domain.Job, error) // Retrieves all jobs
	AddJob(job *domain.Job) error       // Adds a new job
}

type jobServiceImpl struct {
	repo repository.JobRepository // Dependency on the JobRepository
}

// NewJobService creates a new JobService instance
// @param repo repository.JobRepository - The repository to interact with the database
// @return JobService - The implementation of the service interface
func NewJobService(repo repository.JobRepository) JobService {
	return &jobServiceImpl{repo: repo}
}

// GetAllJobs retrieves all jobs from the repository
// Delegates the operation to the repository's FindAll method.
// @return []*domain.Job - A slice of all jobs
// @return error - An error if the retrieval fails
func (s *jobServiceImpl) GetAllJobs() ([]*domain.Job, error) {
	return s.repo.FindAll() // Call repository method to fetch all jobs
}

// AddJob adds a new job to the repository
// Delegates the operation to the repository's Create method.
// @param job *domain.Job - The job data to be added
// @return error - An error if the creation fails
func (s *jobServiceImpl) AddJob(job *domain.Job) error {
	return s.repo.Create(job) // Call repository method to add a new job
}
