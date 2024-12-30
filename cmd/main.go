package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/poolcamacho/jobs-service/internal/repository"
	"github.com/poolcamacho/jobs-service/internal/service"
	"github.com/poolcamacho/jobs-service/internal/transport"
	"github.com/poolcamacho/jobs-service/pkg/config"
	"github.com/poolcamacho/jobs-service/pkg/db"
	jwtUtil "github.com/poolcamacho/jobs-service/pkg/jwt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/poolcamacho/jobs-service/docs" // Import Swagger docs
)

// @title Jobs Service API
// @version 1.0
// @description API for managing jobs in the system.
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
	// Load the application configuration from environment variables
	cfg := config.Load()

	// Connect to the database
	// Establish a connection to the database using the configuration
	dbConn := db.Connect(cfg.DatabaseURL)

	// Initialize repository
	// Create a new instance of JobRepository to interact with the database
	candidateRepo := repository.NewJobRepository(dbConn)

	// Initialize service
	// Create a new instance of JobService to manage business logic
	candidateService := service.NewJobService(candidateRepo)

	// Initialize Gin and routes
	// Setup the Gin HTTP router
	r := gin.Default()
	candidateHandler := transport.NewJobHandler(candidateService)

	// Swagger route
	// Serve Swagger documentation at /swagger/*any
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	// Protected routes requiring authentication
	r.GET("/jobs", jwtUtil.AuthMiddleware(cfg.JWTSecretKey), candidateHandler.GetJobs)    // Get all jobs
	r.POST("/jobs", jwtUtil.AuthMiddleware(cfg.JWTSecretKey), candidateHandler.CreateJob) // Add a new candidate

	// Public route
	// Health check endpoint to verify if the service is running
	r.GET("/health", candidateHandler.HealthCheck)

	// Start server
	// Start the Gin HTTP server on the configured port
	log.Printf("Jobs Service is running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
