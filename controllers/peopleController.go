package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yosefalemu/go-crud/models"
	"gorm.io/gorm"
)

func CreatePerson(c *gin.Context, db *gorm.DB) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&person).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"data": person})
}

func GetPerson(c *gin.Context, db *gorm.DB) {
	fmt.Println("GetPerson")
	var people []models.Person
	if err := db.Find(&people).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": people})
}
