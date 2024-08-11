package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"main.go/routes"
	"main.go/utils"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	r := gin.Default()

	utils.ConnectDatabase()

	routes.AuthRoutes(r)
	routes.ExpenseRoutes(r)
	routes.CategoryRoutes(r)
	routes.BudgetRoutes(r)

	r.Run(":8080")

}
