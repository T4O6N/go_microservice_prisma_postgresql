package database

import (
	"context"
	"fmt"

	"github.com/tonpcst/go-microservice-prisma-postgresql/prisma/db"
)

type PrismaDB struct {
	Client  *db.PrismaClient
	Context context.Context
}

var PClient = &PrismaDB{}

func ConnectDB() (*PrismaDB, error) {
	fmt.Println("Connecting to database...")

	client := db.NewClient()
	// Attempt connection with detailed error logging
	if err := client.Prisma.Connect(); err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Database connected successfully")

	PClient.Client = client
	PClient.Context = context.Background()
	return PClient, nil
}
