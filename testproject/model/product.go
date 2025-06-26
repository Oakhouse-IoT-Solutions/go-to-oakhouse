package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Price string `gorm:"column:price" json:"price"`
	Description string `gorm:"column:description" json:"description"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Product) TableName() string {
	return "products"
}
