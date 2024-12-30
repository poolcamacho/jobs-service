package utils

import "golang.org/x/crypto/bcrypt"

// PasswordUtils defines methods for password operations
// @Description Interface that provides methods for hashing and verifying passwords.
type PasswordUtils interface {
	// HashPassword hashes a plain-text password using bcrypt
	// @Description Hashes a plain-text password using the bcrypt algorithm.
	// @Param password string The plain-text password to hash.
	// @Return string The hashed password as a string.
	// @Return error An error if the hashing process fails.
	HashPassword(password string) (string, error)

	// CheckPassword compares a plain-text password with its hashed counterpart
	// @Description Validates if a plain-text password matches its hashed counterpart using bcrypt.
	// @Param hashedPassword string The hashed password.
	// @Param plainPassword string The plain-text password to compare.
	// @Return error An error if the passwords do not match or if the comparison fails.
	CheckPassword(hashedPassword, plainPassword string) error
}

// DefaultPasswordUtils is the default implementation of PasswordUtils
// @Description Default implementation of the PasswordUtils interface using bcrypt.
type DefaultPasswordUtils struct{}

// HashPassword hashes a plain-text password using bcrypt
// @Description Hashes a plain-text password using the bcrypt algorithm with the default cost.
// @Param password string The plain-text password to hash.
// @Return string The hashed password as a string.
// @Return error An error if the hashing process fails.
func (d *DefaultPasswordUtils) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a plain-text password with its hashed counterpart
// @Description Validates if a plain-text password matches its hashed counterpart using bcrypt.
// @Param hashedPassword string The hashed password.
// @Param plainPassword string The plain-text password to compare.
// @Return error An error if the passwords do not match or if the comparison fails.
func (d *DefaultPasswordUtils) CheckPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
