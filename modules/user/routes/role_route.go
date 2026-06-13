package routes

import (
	"gin-money-manager-api/modules/user/container"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoute(r *gin.Engine, container *container.UserContainer) {
	roles := r.Group("/roles")
	{
		roles.GET("", container.RoleResource.Index)
		roles.GET("/:id", container.RoleResource.Show)
	}
}
