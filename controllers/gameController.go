package controllers

import (
    "BACKEND/models"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type ErrorResponse struct {
    Message string `json:"message"`
}

type GameInput struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
}

// CreateGame godoc
// @Summary Create a new game.
// @Description Create a new game with the provided details.
// @Tags Games
// @Param Authorization header string true "JWT access token"
// @Param Body body GameInput true "Game details"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Router /games [post]
func CreateGame(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input GameInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    game := models.Game{
        Name:        input.Name,
        Description: input.Description,
    }

    if err := db.Create(&game).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create game"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "game created successfully", "game": game})
}

// Getgames godoc
// @Summary Get all games
// @Description Get all games from the database
// @Tags Games
// @Produce json
// @Success 200 {array} models.Game
// @Failure 500 {object} ErrorResponse
// @Router /games [get]
func GetGames(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var games []models.Game
    if err := db.Find(&games).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch games", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, games)
}

// UpdateGame godoc
// @Summary Update a game.
// @Description Update an existing game with the provided details.
// @Tags Games
// @Param Authorization header string true "JWT access token"
// @Param id path int true "Game ID"
// @Param Body body GameInput true "Game details"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Router /games/{id} [put]
func UpdateGame(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input GameInput

    id := c.Param("id")

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    var game models.Game
    if err := db.First(&game, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
        return
    }

    game.Name = input.Name
    game.Description = input.Description

    if err := db.Save(&game).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update game"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "game updated successfully", "game": game})
}

// DeleteGame godoc
// @Summary Delete a game.
// @Description Delete an existing game.
// @Tags Games
// @Param Authorization header string true "JWT access token"
// @Param id path int true "Game ID"
// @Produce json
// @Success 200 {string} string "game deleted successfully"
// @Failure 404 {object} ErrorResponse
// @Router /games/{id} [delete]
func DeleteGame(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    id := c.Param("id")

    var game models.Game
    if err := db.First(&game, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
        return
    }

    if err := db.Delete(&game).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete game"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "game deleted successfully"})
}
