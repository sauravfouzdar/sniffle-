package middlewares

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "sniffle-/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
            c.Abort()
            return
        }

        tokenStr := strings.Split(authHeader, "Bearer ")[1]
        claims, err := utils.ParseJWT(tokenStr)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Set("username", claims.Username)
        c.Next()
    }
}



