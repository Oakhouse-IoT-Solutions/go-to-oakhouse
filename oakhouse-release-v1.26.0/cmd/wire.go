//go:build wireinject
// +build wireinject

// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
package main

import (
	"oakhouse-release-v1.26.0/config"
	"oakhouse-release-v1.26.0/adapter"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Wire providers
var AppSet = wire.NewSet(
	config.LoadConfig,
	adapter.InitializeDatabase,
	NewAppServer,
)

// InitializeApp initializes the application with dependency injection
func InitializeApp() (*AppServer, error) {
	wire.Build(AppSet)
	return &AppServer{}, nil
}
