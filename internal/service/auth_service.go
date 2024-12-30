package service

import (
	"errors"
	"time"

	"github.com/poolcamacho/auth-service/internal/domain"
	"github.com/poolcamacho/auth-service/internal/repository"
	jwtUtil "github.com/poolcamacho/auth-service/pkg/jwt"
	"github.com/poolcamacho/auth-service/pkg/utils"
)

// ErrInvalidCredentials is returned when the provided credentials are invalid
var ErrInvalidCredentials = errors.New("invalid credentials")

// AuthService defines methods for user registration and authentication operations
type AuthService interface {
	Register(user *domain.User) error             // Handles user registration
	Login(email, password string) (string, error) // Authenticates user and returns a JWT token
}

// authServiceImpl is the implementation of AuthService
type authServiceImpl struct {
	repo         repository.UserRepository // User repository interface
	jwtSecretKey string                    // JWT secret key for token generation
	passwordUtil utils.PasswordUtils
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(repo repository.UserRepository, jwtSecretKey string, passwordUtil utils.PasswordUtils) AuthService {
	return &authServiceImpl{
		repo:         repo,
		jwtSecretKey: jwtSecretKey,
		passwordUtil: passwordUtil,
	}
}

// Register hashes the user's password and saves the user in the repository
func (s *authServiceImpl) Register(user *domain.User) error {
	// Hash the plain-text password
	hashedPassword, err := s.passwordUtil.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}

	// Assign the hashed password to the user
	user.PasswordHash = hashedPassword

	// Save the user in the repository
	return s.repo.Create(user)
}

// Login authenticates the user by validating credentials and generates a JWT token
func (s *authServiceImpl) Login(email, password string) (string, error) {
	// Fetch user by email from the repository
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Validate the provided password against the stored hash
	err = s.passwordUtil.CheckPassword(user.PasswordHash, password)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Prepare JWT claims
	claims := map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expiration time (24 hours)
	}

	// Generate the JWT token
	token, err := jwtUtil.GenerateToken(s.jwtSecretKey, claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
