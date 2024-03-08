package controllers

import (
	"errors"
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllVotes(c *gin.Context) {
	var votes []models.Vote
	initializers.DB.Preload("Paslon").Preload("User").Find(&votes)

	c.JSON(200, gin.H{
		"vote": votes,
	})
}

func GetOneVotes(c *gin.Context) {
	id := c.Param("id")

	var votes models.Vote

	err := initializers.DB.Preload("Paslon").Preload("User").First(&votes, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": `data not found`,
		})
		return
	}

	c.JSON(200, gin.H{
		"votes": votes,
	})
}
