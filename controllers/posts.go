package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "sniffle-/config"
    "sniffle-/models"
)

func CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    username, _ := c.Get("username")
    var user models.User
    if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
        return
    }

    post.UserID = user.ID

    if err := config.DB.Create(&post).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {
    var posts []models.Post
    if err := config.DB.Find(&posts).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, posts)
}
