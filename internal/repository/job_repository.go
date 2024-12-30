package repository

import (
	"database/sql"
	"github.com/poolcamacho/jobs-service/internal/domain"
	"time"
)

// JobRepository defines methods for accessing the jobs table
// This interface abstracts database operations for the jobs table.
type JobRepository interface {
	// FindAll retrieves all jobs from the database
	// @return []*domain.Job - A slice of jobs
	// @return error - An error if the query fails
	FindAll() ([]*domain.Job, error)

	// Create inserts a new job into the database
	// @param job *domain.Job - The job data to be inserted
	// @return error - An error if the query fails
	Create(job *domain.Job) error
}

type jobRepositoryImpl struct {
	db *sql.DB // Database connection instance
}

// NewJobRepository creates a new JobRepository instance
// @param db *sql.DB - The database connection to be used for queries
// @return JobRepository - The implementation of the repository
func NewJobRepository(db *sql.DB) JobRepository {
	return &jobRepositoryImpl{db: db}
}

// FindAll retrieves all jobs from the database
// Executes a SELECT query on the jobs table and maps the results to a slice of Job structs.
// @return []*domain.Job - A slice of jobs
// @return error - An error if the query fails
func (r *jobRepositoryImpl) FindAll() ([]*domain.Job, error) {
	query := "SELECT id, title, description, salary_range, created_at, updated_at FROM jobs"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err // Return error if the query fails
	}
	defer rows.Close() // Ensure rows are closed after processing

	var jobs []*domain.Job
	for rows.Next() {
		var job domain.Job
		var createdAt, updatedAt []uint8 // Temporary variables to handle MySQL DATETIME/TIMESTAMP as []uint8
		// Scan values into variables
		if err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.SalaryRange, &createdAt, &updatedAt); err != nil {
			return nil, err // Return error if scanning fails
		}
		// Convert []uint8 to time.Time
		job.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
		job.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		jobs = append(jobs, &job)
	}
	return jobs, nil
}

// Create inserts a new job into the database
// Executes an INSERT query to add a new job record to the jobs table.
// @param job *domain.Job - The job data to be inserted
// @return error - An error if the query fails
func (r *jobRepositoryImpl) Create(job *domain.Job) error {
	query := "INSERT INTO jobs (title, description, salary_range) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, job.Title, job.Description, job.SalaryRange)
	return err // Return error if the query fails
}
