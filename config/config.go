package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connects Go to MySQL database
func ConnectDB() *gorm.DB {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DSN")
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect MySQL database")
	}

	return db
}

// DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
