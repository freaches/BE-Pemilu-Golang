package controllers

import (
	"errors"
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPaslons(c *gin.Context) {
	var paslons []models.Paslon
	initializers.DB.Preload("Partai").Find(&paslons)

	c.JSON(200, gin.H{
		"paslon": paslons,
	})
}

func GetOnePaslons(c *gin.Context) {
	id := c.Param("id")

	var paslons models.Paslon

	err := initializers.DB.Preload("Partai").First(&paslons, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": `data not found`,
		})
		return
	}

	c.JSON(200, gin.H{
		"paslons": paslons,
	})
}
