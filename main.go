package main

import (
	"github.com/mamxalf/eniqilo-store/config"
	"github.com/mamxalf/eniqilo-store/shared/logger"
)

//go:generate go run github.com/swaggo/swag/cmd/swag init
//go:generate go run github.com/google/wire/cmd/wire

var cfg *config.Config

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
func main() {
	logger.InitLogger()
	// Initialize config
	cfg = config.Get()

	// Set desired log level
	logger.SetLogLevel(cfg)

	http := InitializeService()
	http.SetupAndServe()
}
