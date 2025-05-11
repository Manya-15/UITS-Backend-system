package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "webtechproject/models"
)

func GetDeviceSpecificationTemplates(c *gin.Context) {
    typeID, _ := strconv.Atoi(c.Param("type_id"))
    templates, err := models.FetchTemplatesByTypeID(typeID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch templates"})
        return
    }
    c.JSON(http.StatusOK, templates)
}

func AddSpecificationTemplate(c *gin.Context) {
    var req struct {
        TypeID    int    `json:"type_id"`
        SpecName  string `json:"spec_name"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    err := models.InsertSpecificationTemplate(req.TypeID, req.SpecName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert template"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Template added successfully"})
}

func GetSpecificationValues(c *gin.Context) {
    var req struct {
        TemplateID    int    `json:"template_id"`
    }
    values, err := models.FetchSpecificationValues(req.TemplateID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch values"})
        return
    }
    c.JSON(http.StatusOK, values)
}

func AddSpecificationValue(c *gin.Context) {
    var req struct {
        SpecValue string `json:"spec_value"`
        TemplateID int `json:"template_id"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    err := models.InsertSpecificationValue(req.SpecValue, req.TemplateID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert value"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Value added successfully"})
}

func AddDeviceSpecifications(c *gin.Context) {
    var specs []models.DeviceSpecificationInput
    if err := c.BindJSON(&specs); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    err := models.InsertDeviceSpecifications(specs)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add specifications"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Specifications added successfully"})
}
