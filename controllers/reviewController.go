package controllers

import (
	"BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewInput struct {
	GameID  uint   `json:"game_id" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
	Rating  int    `json:"rating" binding:"required"`
	Comment string `json:"comment"`
}

// CreateReview godoc
// @Summary Create a new review.
// @Description Create a new review for a specific game.
// @Tags Reviews
// @Param Authorization header string true "JWT access token"
// @Param Body body ReviewInput true "Review details"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Router /reviews [post]
func CreateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input ReviewInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Review{
		GameID:  input.GameID,
		UserID:  input.UserID,
		Rating:  input.Rating,
		Comment: input.Comment,
	}

	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create review"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "review created successfully", "review": review})
}
