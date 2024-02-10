package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "BACKEND/controllers"
    // "BACKEND/middleware"

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Game routes
    r.POST("/games", controllers.CreateGame)
    r.GET("/games", controllers.GetGames)
    r.PUT("/games/:id", controllers.UpdateGame)
    r.DELETE("/games/:id", controllers.DeleteGame)

    // Review routes
    r.POST("/reviews", controllers.CreateReview) // Tambahkan route untuk membuat review

 // Routes untuk GameType
 r.POST("/game_types", controllers.CreateGameType)
 r.GET("/game_types/:id", controllers.GetGameType)
 r.PUT("/game_types/:id", controllers.UpdateGameType)
 r.DELETE("/game_types/:id", controllers.DeleteGameType)

    
    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
