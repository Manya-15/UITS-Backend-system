package controllers

import (
    "net/http"
    "webtechproject/models"
    "webtechproject/utils"
    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var input struct {
        Email       string `json:"email"`
        Password    string `json:"password"`
        Role        string `json:"role"`
        FullName    string `json:"full_name"`
        Designation string `json:"designation"`
        Username    string `json:"username"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // Build user model
    user := models.User{
        Email:       input.Email,
        Password:    hashedPassword,
        Role:        input.Role,
        FullName:    input.FullName,
        Designation: input.Designation,
        Username:    input.Username,
        Status:      1, // default status
    }

    // Save user
    if err := models.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Generate token
    token, err := utils.GenerateToken(user.ID, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Retrieve user
    user, err := models.GetUserByEmail(input.Email)
    if err != nil || !utils.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Generate token
    token, err := utils.GenerateToken(user.ID, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
