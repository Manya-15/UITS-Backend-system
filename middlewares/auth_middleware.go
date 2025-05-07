package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "webtechproject/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, role, err := utils.ValidateToken(c.GetHeader("Authorization"))
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Set("user_id", userID)
        c.Set("role", role)
        c.Next()
    }
}
