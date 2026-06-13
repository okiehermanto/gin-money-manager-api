package routes

import (
	"gin-money-manager-api/modules/auth/middleware"
	"gin-money-manager-api/modules/user/container"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(r *gin.Engine, container *container.UserContainer) {
	users := r.Group("/users")
	users.POST("", container.CreateUser.Handler)

	users.Use(
		middleware.Auth(container.UserRepository),
		middleware.HasRoles("user"),
	)
	{
		users.GET("", container.UserResource.Index)
	}

}
