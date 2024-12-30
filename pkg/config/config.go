package config

import (
	"os"
)

// Config holds application configuration values
// @Description Contains all configuration values required by the application,
// such as database connection details, JWT secret key, and server port.
type Config struct {
	DatabaseURL  string // URL for the database connection
	JWTSecretKey string // Secret key used for JWT token generation
	Port         string // Port on which the server will run
}

// Load reads configuration from environment variables
// @Description Loads configuration values from environment variables.
// If an environment variable is not set, it uses a default fallback value.
// @Return *Config A pointer to the loaded Config structure.
func Load() *Config {
	return &Config{
		DatabaseURL:  getEnv("DATABASE_URL", "admin_db:dadgic-qafkuh-Hipto0@tcp(talent-management-db.cne4yyyawn11.us-east-1.rds.amazonaws.com:3306)/talent_management_db"),
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "d18aa05bbce170dc073b548f721170fee6e8085e8f10b10548854a489b93afb8"),
		Port:         getEnv("PORT", "3000"),
	}
}

// getEnv retrieves the value of the environment variable named by the key
// @Description Retrieves the value of an environment variable. If the variable
// is not set, it returns a default fallback value.
// @Param key string The name of the environment variable to retrieve.
// @Param fallback string The default value to return if the variable is not set.
// @Return string The value of the environment variable or the fallback value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
