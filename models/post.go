package models

import (
	"gorm.io/gorm"

)


type Post struct {

	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	Description string `json:"description"`
	UserID uint `json:"user_id"`
	User User `json:"user"`
}

