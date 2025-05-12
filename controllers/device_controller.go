package controllers

import (
    "net/http"
    "strconv"
    "webtechproject/models"
    "github.com/gin-gonic/gin"
)

func GetDeviceCategories(c *gin.Context) {
    categories, err := models.FetchDeviceCategories()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}

func GetDeviceTypesByCategory(c *gin.Context) {
    categoryID, _ := strconv.Atoi(c.Param("category_id"))
    types, err := models.FetchDeviceTypesByCategory(categoryID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, types)
}

func AddDeviceCategory(c *gin.Context) {
    var input struct {
        CategoryName string `json:"category_name"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := models.InsertDeviceCategory(input.CategoryName); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Category added"})
}

func AddDeviceType(c *gin.Context) {
    var input struct {
        CategoryID int    `json:"category_id"`
        TypeName   string `json:"type_name"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := models.InsertDeviceType(input.CategoryID, input.TypeName); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Device type added"})
}

func AddDevice(c *gin.Context) {
    var input models.DeviceInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    addedBy := c.GetUint("user_id")
    result, err := models.InsertDevice(input, int(addedBy))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":   "Device added successfully",
        "deviceId":  result.DeviceID,
    })
}


