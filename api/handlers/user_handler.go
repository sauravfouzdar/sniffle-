package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sauravfouzdar/project/models"
)



// CreateUser adds a new user to the database
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}


// GetUsers returns all users from the database
func GetUsers(c *gin.Context) {
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}


// GetUser retrieves a user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, user)
}





// Other handler functions for user CRUD operations
