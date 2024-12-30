package transport

import (
	"github.com/poolcamacho/auth-service/internal/domain"
	"github.com/poolcamacho/auth-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles HTTP requests for authentication
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// HealthCheck provides a simple health status of the service
// @Summary Check service health
// @Description Returns the health status of the authentication service
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "Service is healthy"
// @Router /health [get]
func (h *AuthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

// Register handles the registration of a new user
// @Summary Register a new user
// @Description Register a new user by providing username, email, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body domain.RegisterRequest true "User Registration Request"
// @Success 201 {object} map[string]string "User registered successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map RegisterRequest to User
	user := &domain.User{
		Username:     req.Username,
		Email:        req.Email,
		Role:         "user",       // Default role
		PasswordHash: req.Password, // Pass the plain-text password for hashing in the service
	}

	// Delegate the registration process to the service
	if err := h.authService.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

// Login handles user authentication and token generation
// @Summary Login a user
// @Description Authenticate a user using email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body domain.LoginRequest true "User Login Request"
// @Success 200 {object} map[string]string "Login successful with JWT token"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var credentials domain.LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": user})
}
