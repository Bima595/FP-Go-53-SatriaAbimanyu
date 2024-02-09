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

    // Review routes
// Ratings routes
r.POST("/ratings", controllers.CreateRating) // Route untuk membuat rating
r.GET("/ratings", controllers.GetRatings)    // Route untuk membaca semua ratings
r.GET("/ratings/:id", controllers.GetRating) // Route untuk membaca satu rating
r.PUT("/ratings/:id", controllers.UpdateRating) // Route untuk memperbarui rating
r.DELETE("/ratings/:id", controllers.DeleteRating) // Route untuk menghapus rating


    
    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
