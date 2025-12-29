package main

import (
	"log"

	"github.com/Badankamon/gochat_backend/internal/config"
	"github.com/Badankamon/gochat_backend/internal/platform/database"
	"github.com/Badankamon/gochat_backend/internal/shared/logger"
)

func main() {
	// 1. Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Logger
	logger.Init(cfg.App.Env)

	// 3. Connect
	database.Connect(cfg.Database)

	// 4. Drop Tables
	log.Println("Dropping tables...")
	if err := database.DB.Migrator().DropTable("sessions"); err != nil {
		log.Printf("Failed to drop sessions: %v", err)
	}
	if err := database.DB.Migrator().DropTable("users"); err != nil { // creating cascade effect if needed, but GORM handles order usually if explicit
		log.Printf("Failed to drop users: %v", err)
	}
	if err := database.DB.Migrator().DropTable("profiles"); err != nil { // Cleanup old table
		log.Printf("Failed to drop profiles: %v", err)
	}

	log.Println("Database reset complete.")
}
