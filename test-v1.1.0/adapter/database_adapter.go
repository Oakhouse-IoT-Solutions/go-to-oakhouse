package adapter

import (
	"fmt"
	"log"

	"test-v1.1.0/config"
	"test-v1.1.0/adapter/postgres"

	"gorm.io/gorm"
)

type DatabaseAdapter struct {
	DB *gorm.DB
}

func NewDatabaseAdapter(cfg *config.Config) (*DatabaseAdapter, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate models
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate: %w", err)
	}

	log.Println("✅ Database connected successfully")

	return &DatabaseAdapter{DB: db}, nil
}

func autoMigrate(db *gorm.DB) error {
	// Add your models here for auto-migration
	// return db.AutoMigrate(
	//     &entity.User{},
	//     &entity.Product{},
	// )
	return nil
}

func (d *DatabaseAdapter) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
