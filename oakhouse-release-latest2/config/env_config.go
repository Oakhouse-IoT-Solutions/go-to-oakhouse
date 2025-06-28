// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package config

import "os"

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	Port       string
	Env        string
	JWTSecret  string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "oakhouse-release-latest2_db"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
		Port:       getEnv("PORT", "8080"),
		Env:        getEnv("ENV", "development"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
