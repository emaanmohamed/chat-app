package db

import (
	"fmt"
	"github.com/emaanmohamed/chat-app/configs"
	"github.com/emaanmohamed/chat-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {

	cfg := configs.Envs
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHOST, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	err = db.AutoMigrate(&models.User{}, &models.Message{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
