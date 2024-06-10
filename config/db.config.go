package config

import (

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sniffle-/models"
)

var DB *gorm.DB

func InitDatabase() {
    var err error

	//dbURL := "postgres://pg:postgres@localhost:5432/?sslmode=disable"
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    DB.AutoMigrate(&models.User{}, &models.Post{})
}


