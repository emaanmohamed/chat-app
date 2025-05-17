package db

import (
	"fmt"
	"github.com/emaanmohamed/chat-app/configs"
	"github.com/emaanmohamed/chat-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {

	cfg := configs.Envs
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHOST, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)

	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Cannot get DB from gorm.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)

	DB = db
	fmt.Println("Connected to the database successfully")

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	err = DB.AutoMigrate(&models.User{}, &models.Message{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
