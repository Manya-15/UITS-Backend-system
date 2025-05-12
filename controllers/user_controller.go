package controllers

import (
    "net/http"
    "webtechproject/models"
    "github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
    userID := c.GetUint("user_id")

    // Retrieve user profile from the database
    user, err := models.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Return user profile information
    c.JSON(http.StatusOK, gin.H{
        "email":       user.Email,
        "role":        user.Role,
        "full_name":   user.FullName,
        "designation": user.Designation,
        "username":    user.Username,
        "status":      user.Status,
    })
}

func GetUsers(c *gin.Context) {
    users, err := models.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"users": users})
}
