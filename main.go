package main

import (
	"gin-money-manager-api/database"
	"gin-money-manager-api/modules/auth"
	expensereport "gin-money-manager-api/modules/expense-report"
	"gin-money-manager-api/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	user.Register(r, database.DB)
	auth.Register(r, database.DB)
	expensereport.Register(r, database.DB)

	r.Run(":8080")
}
