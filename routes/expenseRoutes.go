package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/middleware"
)

func ExpenseRoutes(r *gin.Engine) {
	r.GET("/expenses", middleware.AuthMiddleware(), controllers.GetExpenses)
	r.POST("/expenses", middleware.AuthMiddleware(), controllers.CreateExpense)
	r.PUT("/expenses/:id", middleware.AuthMiddleware(), controllers.UpdateExpense)
	r.DELETE("/expenses/:id", middleware.AuthMiddleware(), controllers.DeleteExpense)
}
