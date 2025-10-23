package main

import (
	"os"
	"shorter-url/internal/link"
	"shorter-url/internal/stat"
	"shorter-url/internal/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&link.Link{}, &user.User{}, stat.Stat{})
	if err != nil {
		panic(err)
	}
}
