package config

import (
	"fmt"
	// "os"
	"log"
	"sniffle-/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Host     = "go-postgres.crn397r23dwk.ap-south-1.rds.amazonaws.com"
	User     = "postgres"
	Password = "eKLiipmBjXGqRjCNKwnY"
	Dbname   = "postgres"
	Port     = "5432"
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
