package controllers

import (
    "net/http"
    "webtechproject/models"
    "github.com/gin-gonic/gin"
)

func ViewFilteredDevices(c *gin.Context) {
    var filter models.DeviceFilter
    if err := c.ShouldBindJSON(&filter); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    devices, err := models.FetchFilteredDevices(filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"devices": devices})
}
