package user

import (
	"gin-money-manager-api/modules/auth/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, container *UserContainer) {
	roles := r.Group("/roles")
	{
		roles.GET("", container.RoleResource.Index)
		roles.GET("/:id", container.RoleResource.Show)
	}

	users := r.Group("/users")
	users.Use(
		middleware.Auth(container.UserRepository),
		middleware.HasRoles("user"),
	)
	{
		users.POST("", container.CreateUser.Handler)
		users.GET("", container.UserResource.Index)
	}
}
