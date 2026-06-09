package middleware

import (
	"gin-money-manager-api/modules/auth/service"
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/repository/options"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/repository"

	"github.com/gin-gonic/gin"
)

func Auth(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := helper.ExtractToken(c)

		if err != nil {
			response.Unauthorized(c, err)
			return
		}

		claims, err := service.ParseToken(tokenString)

		if err != nil {
			response.Unauthorized(c, err)
			return
		}
		userId := claims["sub"].(string)

		user, err := userRepository.Find(
			userId,
			&options.FindOptions{
				Relationships: []string{
					"Roles",
				},
			},
		)

		if err != nil {
			response.Unauthorized(c, err)
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
