package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/utils"
)

func CreateBudget(c *gin.Context) {
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	budget.UserId = user.(models.User).ID

	if err := utils.DB.Create(&budget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budget})
}

func GetBudgets(c *gin.Context) {
	var budgets []models.Budget
	user, _ := c.Get("user")

	if err := utils.DB.Where("user_id = ?", user.(models.User).ID).Find(&budgets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgets})
}

func UpdateBudget(c *gin.Context) {
	var budget models.Budget
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&budget).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Budget not found"})
		return
	}

	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.DB.Save(&budget)
	c.JSON(http.StatusOK, gin.H{"data": budget})
}

func DeleteBudget(c *gin.Context) {
	var budget models.Budget
	if err := utils.DB.Where("id = ?", c.Param("id")).First(&budget).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Budget not found"})
		return
	}

	utils.DB.Delete(&budget)
	c.JSON(http.StatusOK, gin.H{"data": "Budget deleted"})
}
