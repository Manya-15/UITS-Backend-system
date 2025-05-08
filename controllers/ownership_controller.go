package controllers

import (
    "net/http"
    "webtechproject/models"
    "github.com/gin-gonic/gin"
)

func AssignOwnershipBulk(c *gin.Context) {
    var input models.OwnershipAssignment
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
        return
    }

    if input.UserID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
        return
    }

    if input.LocationID == nil && input.TypeID == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "At least one filter (location_id or type_id) must be provided"})
        return
    }

    if err := models.AssignOwnershipBulk(input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Assignment failed: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Ownership assigned successfully"})
}
