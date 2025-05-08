package routes

import (
    "github.com/gin-gonic/gin"
    "webtechproject/controllers"
    "webtechproject/middlewares"
)

func RegisterDeviceRoutes(r *gin.Engine) {
    api := r.Group("/api")
    api.Use(middlewares.AuthMiddleware())

    admin := api.Group("/device")
    admin.Use(middlewares.RoleMiddleware("admin", "DOE"))

    admin.GET("/categories", controllers.GetDeviceCategories)
    admin.GET("/types/:category_id", controllers.GetDeviceTypesByCategory)
    admin.POST("/add-category", controllers.AddDeviceCategory)
    admin.POST("/add-type", controllers.AddDeviceType)
    admin.POST("/add", controllers.AddDevice)
	admin.POST("/view", controllers.ViewFilteredDevices)

	// for assigning the ownership 
	


}
