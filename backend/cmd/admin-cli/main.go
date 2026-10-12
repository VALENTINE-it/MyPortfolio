package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
)

func main() {
	username := flag.String("username", "", "Administrator username (3-32 chars)")
	password := flag.String("password", "", "Administrator password (minimum 10 chars)")
	dbPath := flag.String("db", "./portfolio.db", "Path to SQLite database file")
	flag.Parse()

	if *username == "" || *password == "" {
		fmt.Println("Usage: go run ./cmd/admin-cli -username <admin_username> -password <secure_password> [-db ./portfolio.db]")
		os.Exit(1)
	}

	db, err := database.InitDB(*dbPath)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	adminRepo := repositories.NewAdminRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)
	authService := services.NewAuthService(adminRepo, sessionRepo, 24*time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	admin, err := authService.CreateAdmin(ctx, *username, *password)
	if err != nil {
		fmt.Printf("Failed to provision administrator: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Success! Administrator provisioned successfully:\n  ID: %d\n  Username: %s\n  Created: %s\n",
		admin.ID, admin.Username, admin.CreatedAt.Format(time.RFC3339))
}
