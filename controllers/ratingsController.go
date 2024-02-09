package controllers

import (
    "BACKEND/models"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type RatingInput struct {
    ReviewID uint `json:"review_id" binding:"required"`
    Rating   int  `json:"rating" binding:"required"`
    UserID   uint `json:"user_id" binding:"required"`
}

// CreateRating godoc
// @Summary Create a new rating.
// @Description Create a new rating for a specific review.
// @Tags Ratings
// @Param Authorization header string true "JWT access token"
// @Param Body body RatingInput true "Rating details"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Router /ratings [post]
func CreateRating(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input RatingInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    rating := models.Rating{
        ReviewID: input.ReviewID,
        Rating:   input.Rating,
        UserID:   input.UserID,
    }

    if err := db.Create(&rating).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create rating"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "rating created successfully", "rating": rating})
}

// GetRatings godoc
// @Summary Get all ratings.
// @Description Get all ratings from the database.
// @Tags Ratings
// @Produce json
// @Success 200 {array} models.Rating
// @Failure 500 {object} ErrorResponse
// @Router /ratings [get]
func GetRatings(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var ratings []models.Rating
    if err := db.Find(&ratings).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch ratings"})
        return
    }

    c.JSON(http.StatusOK, ratings)
}


// GetRating godoc
// @Summary Get a rating by ID.
// @Description Get a rating by its ID.
// @Tags Ratings
// @Param id path int true "Rating ID"
// @Produce json
// @Success 200 {object} models.Rating
// @Failure 404 {object} ErrorResponse
// @Router /ratings/{id} [get]
func GetRating(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    id := c.Param("id")

    var rating models.Rating
    if err := db.First(&rating, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "rating not found"})
        return
    }

    c.JSON(http.StatusOK, rating)
}


// UpdateRating godoc
// @Summary Update a rating.
// @Description Update an existing rating with the provided details.
// @Tags Ratings
// @Param Authorization header string true "JWT access token"
// @Param id path int true "Rating ID"
// @Param Body body RatingInput true "Rating details"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Router /ratings/{id} [put]
func UpdateRating(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input RatingInput

    id := c.Param("id")

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var rating models.Rating
    if err := db.First(&rating, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "rating not found"})
        return
    }

    rating.ReviewID = input.ReviewID
    rating.Rating = input.Rating
    rating.UserID = input.UserID

    if err := db.Save(&rating).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update rating"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "rating updated successfully", "rating": rating})
}

// DeleteRating godoc
// @Summary Delete a rating.
// @Description Delete an existing rating.
// @Tags Ratings
// @Param Authorization header string true "JWT access token"
// @Param id path int true "Rating ID"
// @Produce json
// @Success 200 {string} string "rating deleted successfully"
// @Failure 404 {object} ErrorResponse
// @Router /ratings/{id} [delete]
func DeleteRating(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    id := c.Param("id")

    var rating models.Rating
    if err := db.First(&rating, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "rating not found"})
        return
    }

    if err := db.Delete(&rating).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete rating"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "rating deleted successfully"})
}
