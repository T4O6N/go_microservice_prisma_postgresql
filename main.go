package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tonpcst/go-microservice-prisma-postgresql/database"
	_ "github.com/tonpcst/go-microservice-prisma-postgresql/docs"
	"github.com/tonpcst/go-microservice-prisma-postgresql/router"
	"go.uber.org/zap"
)

// @title           Event Service API Docs
// @version         1.0.0
// @description     A microservice for managing Event Service with Prisma and PostgreSQL
// @host           localhost:7000
// @BasePath       /api

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Logger *zap.Logger
}

func (app *Application) Serve() error {
	app.Logger.Info("Serving server...", zap.String("port", app.Config.Port))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: router.Routers(),
	}

	return srv.ListenAndServe()
}

func main() {
	// Initialize zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Failed to initialize zap logger: %v", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("Loading environment variables...")
	err = godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}

	logger.Info("Connecting to database...")
	db, err := database.ConnectDB()
	if err != nil {
		logger.Fatal("Cannot connect to database", zap.Error(err))
	}
	defer db.Client.Disconnect()

	// Load config
	config := Config{
		Port: os.Getenv("PORT"),
	}

	// Create application instance
	app := &Application{
		Config: config,
		Logger: logger,
	}

	// Start server
	err = app.Serve()
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
