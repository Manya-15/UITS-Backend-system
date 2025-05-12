package routes

import (
    "github.com/gin-gonic/gin"
    "webtechproject/controllers"
    "webtechproject/middlewares"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    api := r.Group("/api")
    api.Use(middlewares.AuthMiddleware())
    api.GET("/profile", controllers.GetProfile)

	// r.POST("/add-device", controllers.AddDevice)

    admin := api.Group("/admin")
    admin.Use(middlewares.RoleMiddleware("admin","DOE"))
    admin.GET("/dashboard", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Welcome, admin!"})
    })
    admin.POST("/ownership/assign", controllers.AssignOwnershipBulk)
    admin.GET("/users", controllers.GetUsers)

	// admin.POST("/add-device", controllers.AddDevice)
}
