package config

import (
	"fmt"
	"os"
	"log"
	"sniffle-/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	// println(key, os.Getenv(key))
	return os.Getenv(key)
}

var (
	Host     = goDotEnvVariable("DB_HOST")
	User     = goDotEnvVariable("DB_USERNAME")
	Password = goDotEnvVariable("DB_PASSWORD")
	Dbname   = goDotEnvVariable("DB_NAME")
	Port     = goDotEnvVariable("DB_PORT")
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
