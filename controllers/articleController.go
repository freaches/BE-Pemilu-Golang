package controllers

import (
	"errors"
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	initializers.DB.Preload("User").Find(&articles)

	c.JSON(200, gin.H{
		"articles": articles,
	})
}

func GetOneArticles(c *gin.Context) {
	id := c.Param("id")

	var articles models.Article

	err := initializers.DB.Preload("User").First(&articles, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": `data not found`,
		})
		return
	}

	c.JSON(200, gin.H{
		"articles": articles,
	})
}
