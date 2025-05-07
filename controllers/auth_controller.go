package controllers

import (
    "net/http"
    "webtechproject/config"
    "webtechproject/models"
    "webtechproject/utils"
    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
        Role     string `json:"role"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := utils.HashPassword(input.Password)

    user := models.User{Email: input.Email, Password: hashedPassword, Role: input.Role}
    result := config.DB.Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    token, _ := utils.GenerateToken(user.ID, user.Role)
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

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, _ := utils.GenerateToken(user.ID, user.Role)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
