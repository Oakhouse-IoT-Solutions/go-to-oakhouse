// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package adapter

import (
	"fmt"
	"log"

	"momo/config"
	"momo/adapter/postgres"
	"gorm.io/gorm"
)

// InitializeDatabase initializes the database connection
// Returns nil, error if connection fails but allows application to continue
func InitializeDatabase(cfg *config.Config) (*gorm.DB, error) {
	// Check if database configuration is provided
	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBName == "" {
		log.Println("⚠️ Database configuration not provided or incomplete")
		return nil, fmt.Errorf("database configuration incomplete")
	}

	return postgres.NewGormDB(cfg)
}

// GetDSN returns the database connection string
func GetDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)
}
