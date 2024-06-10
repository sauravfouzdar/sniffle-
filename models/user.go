package models


import (

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}



