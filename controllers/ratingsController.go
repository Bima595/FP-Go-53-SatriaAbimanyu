package controllers

// Import model yang sesuai
import (
	"BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Gunakan struct yang sesuai untuk input
type GameTypeInput struct {
    GameID   uint   `json:"game_id" binding:"required"`
    Theme    string `json:"theme" binding:"required"`
}

// Perbarui fungsi CreateGameType untuk membuat GameType baru
func CreateGameType(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input GameTypeInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    gameType := models.GameType{
        GameID:   input.GameID,
        Theme:    input.Theme,
    }

    if err := db.Create(&gameType).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create game type"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "game type created successfully", "game_type": gameType})
}

// Perbarui fungsi GetGameType untuk mendapatkan GameType berdasarkan ID
func GetGameType(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    id := c.Param("id")

    var gameType models.GameType
    if err := db.First(&gameType, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "game type not found"})
        return
    }

    c.JSON(http.StatusOK, gameType)
}

// Perbarui fungsi UpdateGameType untuk mengupdate GameType
func UpdateGameType(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input GameTypeInput

    id := c.Param("id")

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var gameType models.GameType
    if err := db.First(&gameType, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "game type not found"})
        return
    }

    gameType.GameID = input.GameID
    gameType.Theme = input.Theme

    if err := db.Save(&gameType).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update game type"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "game type updated successfully", "game_type": gameType})
}

// Perbarui fungsi DeleteGameType untuk menghapus GameType
func DeleteGameType(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    id := c.Param("id")

    var gameType models.GameType
    if err := db.First(&gameType, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "game type not found"})
        return
    }

    if err := db.Delete(&gameType).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete game type"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "game type deleted successfully"})
}
