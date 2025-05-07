package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        for _, allowed := range allowedRoles {
            if role == allowed {
                c.Next()
                return
            }
        }
        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
    }
}
