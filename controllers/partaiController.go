package controllers

import (
	"errors"
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPartais(c *gin.Context) {
	var partais []models.Partai
	initializers.DB.Preload("Paslon").Find(&partais)

	c.JSON(200, gin.H{
		"partai": partais,
	})
}

func GetOnePartais(c *gin.Context) {
	id := c.Param("id")

	var partais models.Partai

	err := initializers.DB.Preload("Paslon").First(&partais, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": `data not found`,
		})
		return
	}

	c.JSON(200, gin.H{
		"partais": partais,
	})
}
