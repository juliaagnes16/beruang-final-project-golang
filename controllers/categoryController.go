package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/utils"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := utils.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.DB.Save(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	utils.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"data": "Category deleted"})
}
