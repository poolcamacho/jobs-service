package domain

import "time"

// Job represents a job in the system
type Job struct {
	ID          int       `json:"id"`           // Job ID
	Title       string    `json:"title"`        // Job title
	Description string    `json:"description"`  // Job description
	SalaryRange string    `json:"salary_range"` // Salary range
	CreatedAt   time.Time `json:"created_at"`   // Creation timestamp
	UpdatedAt   time.Time `json:"updated_at"`   // Last update timestamp
}
