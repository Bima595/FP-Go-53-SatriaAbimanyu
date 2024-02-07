package main

import (
	"BACKEND/config"     // Sesuaikan dengan path yang sesuai
	"BACKEND/controller" // Sesuaikan dengan path yang sesuai
	"BACKEND/utils"      // Sesuaikan dengan path yang sesuai

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @title Your Project's Name
// @version 1.0
// @description Describe your API here
// @BasePath /api/v1
func main() {
    // Inisialisasi database
    config.InitDB()

    // Inisialisasi router Gin
    router := gin.Default()

    // Inisialisasi dokumentasi Swagger
    // Menggunakan swaggo untuk meng-generate dokumentasi Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Rute untuk registrasi pengguna baru
    router.POST("/api/v1/register", controller.RegisterUserController)

    // Rute untuk login pengguna
    router.POST("/api/v1/login", controller.LoginUserController)

    // Rute untuk mengubah password pengguna (memerlukan autentikasi JWT)
    router.POST("/api/v1/change-password", utils.MiddlewareJWTAuth(controller.ChangePasswordController))


    // Inisialisasi middleware JWT auth
    // Digunakan untuk mengautentikasi pengguna pada rute tertentu
    // Gunakan dengan cara menambahkan middleware tersebut sebagai wrapper pada handler fungsi yang memerlukan autentikasi JWT
    // Contoh: router.GET("/api/v1/protected", utils.MiddlewareJWTAuth(protectedHandler))

    // Jalankan server di port 8080
    router.Run(":8081")
}
