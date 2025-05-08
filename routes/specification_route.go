// routes/specification_routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "webtechproject/controllers"
    "webtechproject/middlewares"
)

func RegisterSpecificationRoutes(r *gin.Engine) {
    api := r.Group("/api")
    api.Use(middlewares.AuthMiddleware())

    spec := api.Group("/specifications")
    spec.Use(middlewares.RoleMiddleware("admin", "DOE"))

    spec.GET("/templates/:device_id", controllers.GetDeviceSpecificationTemplates)
    spec.POST("/add-template", controllers.AddSpecificationTemplate)
    spec.GET("/values", controllers.GetSpecificationValues)
    spec.POST("/add-value", controllers.AddSpecificationValue)
    spec.POST("/add", controllers.AddDeviceSpecifications)
}
