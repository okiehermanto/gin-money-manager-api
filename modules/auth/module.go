package auth

import (
	"gin-money-manager-api/database"
	"gin-money-manager-api/modules/auth/command"
	"gin-money-manager-api/modules/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(
	r *gin.Engine,
	db *gorm.DB,
) {
	userRepository := repository.NewUserRepository(database.DB)

	authCommand := command.NewAuthCommand(userRepository)

	auths := r.Group("/auth")
	{
		auths.POST("", authCommand.Handler)
	}
}
