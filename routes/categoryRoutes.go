package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/middleware"
)

func CategoryRoutes(r *gin.Engine) {
	r.GET("/categories", middleware.AuthMiddleware(), controllers.GetCategories)
	r.POST("/categories", middleware.AuthMiddleware(), controllers.CreateCategory)
	r.PUT("/categories/:id", middleware.AuthMiddleware(), controllers.UpdateCategory)
	r.DELETE("/categories/:id", middleware.AuthMiddleware(), controllers.DeleteCategory)
}
