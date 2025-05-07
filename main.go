package main

import (
    "webtechproject/config"
    "webtechproject/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time" // <-- Import the time package
)

func main() {
    config.ConnectDB()
    r := gin.Default()

    // Enable CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Allow the frontend's origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Allow HTTP methods
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow headers
        AllowCredentials: true,
        MaxAge: 12 * time.Hour, // Use the time package to define MaxAge
    }))
    
    routes.SetupRoutes(r)
    r.Run(":8080") // Backend on port 8080
}
