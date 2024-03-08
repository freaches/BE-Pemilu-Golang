package controllers

import (
	"errors"
	"micro-feature-pemilu/initializers"
	"micro-feature-pemilu/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUser(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"user": users,
	})
}

func GetOneUser(c *gin.Context) {
	id := c.Param("id")

	var users models.User

	err := initializers.DB.First(&users, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": `data not found`,
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
