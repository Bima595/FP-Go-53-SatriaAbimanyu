package routes

import (
	"BACKEND/controllers"
	middlewares "BACKEND/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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
    r.GET("/games", controllers.GetGames)
    gameMiddlewareRoute := r.Group("/games")
    gameMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    gameMiddlewareRoute.POST("/", controllers.CreateGame)
    gameMiddlewareRoute.PATCH("/:id", controllers.UpdateGame)
    gameMiddlewareRoute.DELETE("/:id", controllers.DeleteGame)

    // Review routes
    r.POST("/reviews", controllers.CreateReview)

    // Routes untuk GameType
    r.GET("/game_types/:id", controllers.GetGameType)
    gameTypeMiddlewareRoute := r.Group("/game_types")
    gameTypeMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    gameTypeMiddlewareRoute.POST("/", controllers.CreateGameType)
    gameTypeMiddlewareRoute.PATCH("/:id", controllers.UpdateGameType)
    gameTypeMiddlewareRoute.DELETE("/:id", controllers.DeleteGameType)

    
    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
