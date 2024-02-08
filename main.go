package main

import (
	"BACKEND/config"
	"BACKEND/controller"
	"BACKEND/docs"
	"BACKEND/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your Project's Name
// @version 1.0
// @description Describe your API here
// @BasePath /api/v1
func main() {
    // Initialize the database
    config.InitDB()

    // Initialize Gin router
    router := gin.Default()

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
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

    // Initialize Swagger documentation
    // Use swaggo to generate Swagger documentation
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Route for registering a new user
    router.POST("/api/v1/register", controller.RegisterUserController)

    // Route for user login
    router.POST("/api/v1/login", controller.LoginUserController)

    // Route for changing user password (requires JWT authentication)
    router.POST("/api/v1/change-password", utils.MiddlewareJWTAuth(controller.ChangePasswordController))

    // Initialize JWT authentication middleware
    // Used to authenticate users on specific routes
    // Use by adding the middleware as a wrapper on handler functions that require JWT authentication
    // Example: router.GET("/api/v1/protected", utils.MiddlewareJWTAuth(protectedHandler))

    // Run the server on port 8081
    router.Run(":8081")
}
