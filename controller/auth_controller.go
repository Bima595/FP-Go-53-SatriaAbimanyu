// controller/user_controller.go
package controller

import (
	"net/http"

	"BACKEND/config"
	"BACKEND/models"
	"BACKEND/utils"

	"github.com/gin-gonic/gin"
)


func RegisterUserController(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password sebelum disimpan di database
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
    }
    user.Password = hashedPassword

    // Simpan user ke dalam database menggunakan model
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

// LoginUserController menghandle proses login pengguna
func LoginUserController(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Temukan user berdasarkan username
    var storedUser models.User
    if err := config.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // Periksa apakah password cocok
    if !utils.CheckPasswordHash(user.Password, storedUser.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }

    // Generate token JWT
    token, err := utils.GenerateToken(storedUser.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// ChangePasswordController menghandle proses pergantian password pengguna
func ChangePasswordController(c *gin.Context) {
    userID, _ := c.Get("userID")
    userIDUint := userID.(uint)
    var newPassword struct {
        NewPassword string `json:"new_password"`
    }
    if err := c.ShouldBindJSON(&newPassword); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password baru
    hashedPassword, err := utils.HashPassword(newPassword.NewPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
    }

    // Update password di database
    if err := config.DB.Model(&models.User{}).Where("id = ?", userIDUint).Update("password", hashedPassword).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}
