package middleware

import (
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/entity"

	"github.com/gin-gonic/gin"
)

func HasRoles(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]bool)

	allowed["admin"] = true

	for _, role := range roles {
		allowed[role] = true
	}

	return func(c *gin.Context) {
		user := c.MustGet("user").(entity.User)

		for _, role := range user.Roles {
			if allowed[role.Name] {
				c.Next()
				return
			}
		}
		response.Forbidden(c)
	}
}
