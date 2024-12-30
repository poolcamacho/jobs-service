package domain

// RegisterRequest represents the payload for a user registration request
// @Description The request body for registering a new user.
// @Param username query string true "Username of the new user"
// @Param email query string true "Email address of the new user"
// @Param password query string true "Password of the new user"
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`       // Username of the new user
	Email    string `json:"email" binding:"required,email"`    // Email address of the new user
	Password string `json:"password" binding:"required,min=8"` // Password for the new user
}

// LoginRequest represents the payload for user login
// @Description The request body for logging in a user.
// @Param email query string true "User's email address"
// @Param password query string true "User's password"
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"` // Email address of the user
	Password string `json:"password" binding:"required"`    // Password of the user
}
