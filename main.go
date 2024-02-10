package main

import (
	"BACKEND/config"
	"BACKEND/docs"
	"BACKEND/routes"
	"BACKEND/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
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
    docs.SwaggerInfo.Host = ":3000"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

 // database connection
    db := config.ConnectDataBase()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()

    // router
    var port = envPortOr("3000")
    r := routes.SetupRouter(db)
    r.Run("0.0.0.0" + port)
    
}

func envPortOr(port string) string {
  // If `PORT` variable in environment exists, return it
  if envPort := os.Getenv("PORT"); envPort != "" {
    return ":" + envPort
  }
  // Otherwise, return the value of `port` variable from function argument
  return ":" + port
}