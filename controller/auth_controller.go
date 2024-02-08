package controller

import (
	"net/http"
	"time"

	"BACKEND/config"
	"BACKEND/models"
	"BACKEND/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ErrorResponse adalah struktur data untuk menangani kesalahan dalam respon HTTP
type ErrorResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
    Token string `json:"token"`
}


// @Summary Register new user
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {string} string "user created successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
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

// @Summary User login
// @Description Log in as an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
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

// @Summary Change user password
// @Description Change password for the authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param password body struct { NewPassword string } true "New password details"
// @Success 200 {string} string "password updated successfully"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /change-password [post]

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

// @Summary Generate JWT token
// @Description Generate JWT token with the provided userID
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userID body int true "User ID"
// @Success 200 {object} TokenResponse
// @Failure 500 {object} ErrorResponse
// @Router /generate-token [post]
func GenerateToken(userID uint) (models.TokenResponse, error) {
    // Buat payload token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["userID"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token berlaku selama 24 jam

    // Sign token dengan secret key
    tokenString, err := token.SignedString(utils.JWTKey)
    if err != nil {
        return models.TokenResponse{}, err
    }
    return models.TokenResponse{Token: tokenString}, nil
}

