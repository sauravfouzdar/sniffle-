package config

import (
	"fmt"
	"os"
	"log"
	"sniffle-/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/joho/godotenv"
)

var (
	Host     = os.Getenv("DB_HOST")
	User     = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	Dbname   = os.Getenv("DB_NAME")
	Port     = os.Getenv("DB_PORT")
)

func BuildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		Host, User, Password, Dbname, Port)
}

var DB *gorm.DB

func InitDatabase() {

	var err error
	dsn := BuildDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")

	DB.AutoMigrate(&models.User{}, &models.Post{})
}
