package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBProduct() *gorm.DB {
	LoadEnv()
	GORM_DB_PRODUCT := os.Getenv("GORM_DB_PRODUCT")

	db, err := gorm.Open(postgres.Open(GORM_DB_PRODUCT), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}
	return db
}
