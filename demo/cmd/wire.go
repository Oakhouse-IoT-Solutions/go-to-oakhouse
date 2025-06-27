// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"demo/config"
	"demo/adapter"

	"github.com/google/wire"
)

// ProviderSet is a Wire provider set that includes all the dependencies
var ProviderSet = wire.NewSet(
	config.Load,
	ProvideDatabase,
	NewAppServer,
)

// ProvideDatabase creates a new database adapter with error handling
func ProvideDatabase(cfg *config.Config) *adapter.DatabaseAdapter {
	db, err := adapter.NewDatabaseAdapter(cfg)
	if err != nil {
		// Return nil database but don't fail - server can run without DB
		return nil
	}
	return db
}

// InitializeApp initializes the application with all dependencies using Wire
func InitializeApp() (*AppServer, func(), error) {
	wire.Build(ProviderSet)
	return &AppServer{}, func() {}, nil
}
