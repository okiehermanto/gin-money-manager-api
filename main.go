package main

import (
	"gin-money-manager-api/database"
	"gin-money-manager-api/modules/auth"
	"gin-money-manager-api/modules/todo"
	"gin-money-manager-api/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	todo.Register(r, database.DB)
	user.Register(r, database.DB)
	auth.Register(r, database.DB)

	r.Run(":8080")
}
