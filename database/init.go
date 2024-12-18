package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func DbInstance() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("no .env file found", err)
	}

	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		log.Fatal("database environment not set")
	}

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect", err)
	}

	fmt.Println("database connected sucessfully")

	DB = db
}
