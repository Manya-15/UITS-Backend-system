package config

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "webtechproject/models" // Replace with your actual module name
)

var DB *gorm.DB
var JWT_SECRET []byte

func init() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Load JWT secret from the environment
    JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

    // Connect to the database
    ConnectDB()
}

func ConnectDB() {
    // Fetch database credentials from environment variables
    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Form the Data Source Name (DSN)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUsername, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    // Auto migrate your models
    DB.AutoMigrate(&models.User{})
}
