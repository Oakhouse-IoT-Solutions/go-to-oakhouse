//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"oakhouse-release-v1.11.0/config"
	"oakhouse-release-v1.11.0/adapter"

	"github.com/google/wire"
)

// ProviderSet is a Wire provider set that includes all the dependencies
var ProviderSet = wire.NewSet(
	config.Load,
	NewDatabaseAdapter,
	NewAppServer,
)

// NewDatabaseAdapter creates a new database adapter with error handling
func NewDatabaseAdapter(cfg *config.Config) (*adapter.DatabaseAdapter, func(), error) {
	db, err := adapter.NewDatabaseAdapter(cfg)
	if err != nil {
		// Return nil database but don't fail - server can run without DB
		return nil, func() {}, nil
	}
	
	cleanup := func() {
		if db != nil {
			db.Close()
		}
	}
	
	return db, cleanup, nil
}

// InitializeApp initializes the application with all dependencies using Wire
func InitializeApp() (*AppServer, func(), error) {
	wire.Build(ProviderSet)
	return &AppServer{}, func() {}, nil
}