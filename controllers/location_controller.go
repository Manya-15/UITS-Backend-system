package controllers

import (
    "net/http"
    "webtechproject/models"
    "github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
    locations, err := models.FetchAllLocations()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, locations)
}
