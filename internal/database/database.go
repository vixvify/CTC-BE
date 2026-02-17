package database

import (
	"log"
	"server/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgresql://neondb_owner:npg_glXvR2t6uVcy@ep-purple-resonance-a142qjbz-pooler.ap-southeast-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
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