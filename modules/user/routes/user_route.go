package routes

import (
	"gin-money-manager-api/modules/auth/middleware"
	"gin-money-manager-api/modules/user/command"
	"gin-money-manager-api/modules/user/container"
	"gin-money-manager-api/modules/user/resource"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(r *gin.Engine, container *container.UserContainer) {
	auth := middleware.Auth(container.UserRepository)

	createUser := command.NewCreateUser(&container.UserService)
	userResource := resource.NewUserResource(container.UserRepository)

	users := r.Group("/users")
	users.POST("", createUser.Handler)

	users.GET(
		"",
		auth,
		middleware.HasRoles("user"),
		userResource.Index,
	)

	users.GET(
		"/:user",
		auth,
		middleware.HasRoles("user"),
		userResource.Show,
	)
}
