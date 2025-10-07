package config

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all the configuration settings for the backend
// This struct centralizes all our application settings in one place
// Each field corresponds to a specific aspect of our application
type Config struct {
	JWTSecret  string // Secret key for signing JWT tokens (should be kept secure)
	Port       string // HTTP server port (e.g., "8080")
	MaxRetries int    // Maximum number of retry attempts for failed operations
	DebugMode  bool   // Whether to run in debug mode (default: true for development)
}

// Load reads configuration from environment variables and returns a Config struct
// This function is responsible for:
// 1. Reading environment variables
// 2. Providing sensible defaults if variables aren't set
// 3. Converting string values to appropriate types
// 4. Centralizing all configuration logic
func Load() *Config {
	return &Config{
		JWTSecret:  getEnv("JWT_SECRET", "supersecret"),
		Port:       getEnv("PORT", "8080"),
		MaxRetries: getIntEnv("MAX_RETRIES", 3),
		DebugMode:  getBoolEnv("DEBUG_MODE", true),
	}
}

// getEnv reads an environment variable and returns its value
// If the environment variable is not set, it returns the default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}


// getIntEnv reads an environment variable and converts it to an integer
func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getBoolEnv reads an environment variable and converts it to a boolean
// Valid values: "true", "false", "1", "0", "yes", "no" (case insensitive)
func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		value = strings.ToLower(value)
		if value == "true" || value == "1" || value == "yes" {
			return true
		}
		if value == "false" || value == "0" || value == "no" {
			return false
		}
	}
	return defaultValue
}
