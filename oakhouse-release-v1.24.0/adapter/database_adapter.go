// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package adapter

import (
	"fmt"

	"oakhouse-release-v1.24.0/config"
	"oakhouse-release-v1.24.0/adapter/postgres"
	"gorm.io/gorm"
)

// InitializeDatabase initializes the database connection
func InitializeDatabase(cfg *config.Config) (*gorm.DB, error) {
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
