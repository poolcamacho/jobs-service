package repository

import (
	"database/sql"
	"github.com/poolcamacho/auth-service/internal/domain"
	"time"
)

// UserRepository defines methods for interacting with the users table
// @Description Interface that defines methods to interact with the users table in the database.
type UserRepository interface {
	// FindByEmail retrieves a user by their email address
	// @Description Retrieves a user by their email address from the database.
	// @Param email string The email address of the user to retrieve.
	// @Return *domain.User The user object if found, or an error if not.
	FindByEmail(email string) (*domain.User, error)

	// Create adds a new user to the database
	// @Description Inserts a new user record into the database.
	// @Param user *domain.User The user object to create.
	// @Return error An error if the operation fails.
	Create(user *domain.User) error
}

// userRepositoryImpl is the MySQL implementation of UserRepository
type userRepositoryImpl struct {
	db *sql.DB // Database connection
}

// NewUserRepository creates a new UserRepository instance
// @Description Creates a new instance of the UserRepository for interacting with the users table.
// @Param db *sql.DB The database connection object.
// @Return UserRepository The new repository instance.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// FindByEmail retrieves a user by their email address
// @Description Fetches a user from the database using their email address.
// @Param email string The email address of the user to retrieve.
// @Return *domain.User The user object if found, or an error if not.
func (r *userRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	query := "SELECT id, username, email, password_hash, role, created_at, updated_at FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	var user domain.User
	var createdAt, updatedAt []uint8 // Temporary variables to handle MySQL DATETIME/TIMESTAMP as []uint8

	// Scan values into variables
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Convert []uint8 to time.Time
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return nil, err
	}

	user.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Create adds a new user to the database
// @Description Inserts a new user record into the database.
// @Param user *domain.User The user object containing details to be added.
// @Return error An error if the operation fails.
func (r *userRepositoryImpl) Create(user *domain.User) error {
	query := "INSERT INTO users (username, email, password_hash, role) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.Role)
	return err
}
