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

    spec.GET("/templates/:type_id", controllers.GetDeviceSpecificationTemplates)
    spec.POST("/add-template", controllers.AddSpecificationTemplate)
    spec.POST("/values", controllers.GetSpecificationValues) //template id passed in this
    spec.POST("/add-value", controllers.AddSpecificationValue) //yemplate id passes in this as well
    spec.POST("/add", controllers.AddDeviceSpecifications)
}
