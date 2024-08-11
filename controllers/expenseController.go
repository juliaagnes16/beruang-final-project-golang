package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/utils"
)

func CreateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	expense.UserID = user.(models.User).ID

	if err := utils.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": expense})
}

func GetExpenses(c *gin.Context) {
	var expenses []models.Expense
	user, _ := c.Get("user")

	if err := utils.DB.Where("user_id = ?", user.(models.User).ID).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": expenses})
}

func UpdateExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenseID := c.Param("id")
	var existingExpense models.Expense

	if err := utils.DB.Where("id = ?", expenseID).First(&existingExpense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	user, _ := c.Get("user")
	if existingExpense.UserID != user.(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to update this expense"})
		return
	}

	if err := utils.DB.Model(&existingExpense).Updates(expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": existingExpense})
}

func DeleteExpense(c *gin.Context) {
	expenseID := c.Param("id")
	var expense models.Expense

	if err := utils.DB.Where("id = ?", expenseID).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	user, _ := c.Get("user")
	if expense.UserID != user.(models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to delete this expense"})
		return
	}

	if err := utils.DB.Delete(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}
