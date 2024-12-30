package domain

import "time"

// User represents a user in the system
// @Description Represents a user entity in the system with all associated details.
type User struct {
	ID           int       `json:"id"`            // User ID
	Username     string    `json:"username"`      // User's username
	Email        string    `json:"email"`         // User's email
	PasswordHash string    `json:"password_hash"` // Hashed password of the user
	Role         string    `json:"role"`          // User's role (e.g., admin, user)
	CreatedAt    time.Time `json:"created_at"`    // Timestamp when the user was created
	UpdatedAt    time.Time `json:"updated_at"`    // Timestamp when the user was last updated
}
