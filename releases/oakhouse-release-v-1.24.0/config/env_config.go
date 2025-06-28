// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package config

import (
	"os"
	"strconv"
)

type Config struct {
	// Application
	AppName string
	AppPort string
	AppEnv  string
	AppDebug bool

	// Database
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
	DBTimezone string

	// JWT
	JWTSecret    string
	JWTExpiresIn string

	// Redis
	RedisURL      string
	RedisPassword string
	RedisDB       int

	// CORS
	CorsAllowedOrigins string
	CorsAllowedMethods string
	CorsAllowedHeaders string

	// Rate Limiting
	RateLimitRequests int
	RateLimitDuration string
}

func Load() *Config {
	return &Config{
		// Application
		AppName:  getEnv("APP_NAME", "go-to-oakhouse"),
		AppPort:  getEnv("APP_PORT", "8080"),
		AppEnv:   getEnv("APP_ENV", "development"),
		AppDebug: getEnvBool("APP_DEBUG", true),

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "oakhouse_db"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
		DBTimezone: getEnv("DB_TIMEZONE", "UTC"),

		// JWT
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),

		// Redis
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		// CORS
		CorsAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
		CorsAllowedMethods: getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS"),
		CorsAllowedHeaders: getEnv("CORS_ALLOWED_HEADERS", "*"),

		// Rate Limiting
		RateLimitRequests: getEnvInt("RATE_LIMIT_REQUESTS", 100),
		RateLimitDuration: getEnv("RATE_LIMIT_DURATION", "1m"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
