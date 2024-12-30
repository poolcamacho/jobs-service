package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poolcamacho/auth-service/internal/repository"
	"github.com/poolcamacho/auth-service/internal/service"
	"github.com/poolcamacho/auth-service/internal/transport"
	"github.com/poolcamacho/auth-service/pkg/config"
	"github.com/poolcamacho/auth-service/pkg/db"
	"github.com/poolcamacho/auth-service/pkg/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/poolcamacho/auth-service/docs" // Import Swagger docs
)

// @title Auth Service API
// @version 1.0
// @description API for user authentication and management.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to the database
	dbConn := db.Connect(cfg.DatabaseURL)

	// Initialize repository
	userRepo := repository.NewUserRepository(dbConn)

	// Initialize password utilities
	passwordUtil := &utils.DefaultPasswordUtils{}

	// Initialize service
	authService := service.NewAuthService(userRepo, cfg.JWTSecretKey, passwordUtil)

	// Initialize Gin and routes
	r := gin.Default()
	authHandler := transport.NewAuthHandler(authService)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	r.POST("/register", authHandler.Register) // Register a new user
	r.POST("/login", authHandler.Login)       // Login and generate token
	r.GET("/health", authHandler.HealthCheck) // Health check

	log.Printf("Auth Service is running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
