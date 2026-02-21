package database

import (
	"log"
	"os"
	"server/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("❌ DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  url,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("❌ failed to connect database")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.Quiz{},
	)
	if err != nil {
		log.Fatal("❌ auto migrate failed:", err)
	}

	DB = db
	log.Println("✅ database connected")
}
