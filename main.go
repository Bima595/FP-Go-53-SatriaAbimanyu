package main

import (
	"BACKEND/config"
	"BACKEND/docs"
	"BACKEND/routes"
	"BACKEND/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    // for load godotenv
    // for env
    environment := utils.Getenv("ENVIRONMENT", "development")

    if environment == "development" {
      err := godotenv.Load()
      if err != nil {
        log.Fatal("Error loading .env file")
      }
    }
    // programmatically set swagger info
    docs.SwaggerInfo.Title = "Swagger Example API"
    docs.SwaggerInfo.Description = "This is a sample server Movie."
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8081"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

    // database connection
    config.InitDB()
    defer func() {
        if db, err := config.DB.DB(); err == nil {
            db.Close()
        }
    }()

    // router
    r := routes.SetupRouter(config.DB)
    r.Run("localhost:8081")
}
