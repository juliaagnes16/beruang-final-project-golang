package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/middleware"
)

func BudgetRoutes(r *gin.Engine) {
	r.GET("/budgets", middleware.AuthMiddleware(), controllers.GetBudgets)
	r.POST("/budgets", middleware.AuthMiddleware(), controllers.CreateBudget)
	r.PUT("/budgets/:id", middleware.AuthMiddleware(), controllers.UpdateBudget)
	r.DELETE("/budgets/:id", middleware.AuthMiddleware(), controllers.DeleteBudget)
}
