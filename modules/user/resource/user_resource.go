package resource

import (
	"gin-money-manager-api/modules/shared/repository/options"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/repository"

	"github.com/gin-gonic/gin"
)

type UserResource struct {
	userRepository repository.UserRepository
}

func NewUserResource(
	userRepository repository.UserRepository,
) *UserResource {
	return &UserResource{
		userRepository: userRepository,
	}
}

func (r *UserResource) Index(c *gin.Context) {
	var search = c.Query("search")

	users, err := r.userRepository.FindAll(
		&options.FindAllOptions{
			Relationships: []string{
				"Roles",
			},
			Search: &options.SearchOptions{
				Keyword: search,
				Fields: []string{
					"name",
					"username",
					"email",
				},
			},
		},
	)

	response.Success(c, users, err)
}
