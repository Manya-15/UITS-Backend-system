package controllers

import (
    "net/http"
    "webtechproject/config"
    "webtechproject/models"
    "github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
    userID := c.GetUint("user_id")

    var user models.User
    if err := config.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"email": user.Email, "role": user.Role})
}
