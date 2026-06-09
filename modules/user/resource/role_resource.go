package resource

import (
	"gin-money-manager-api/modules/shared/repository/options"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/repository"

	"github.com/gin-gonic/gin"
)

type RoleResource struct {
	repository repository.RoleRepository
}

func NewRoleResource(repository repository.RoleRepository) *RoleResource {
	return &RoleResource{
		repository: repository,
	}
}

func (r *RoleResource) Index(c *gin.Context) {
	var name = c.Query("name")

	roles, err := r.repository.FindAll(
		&options.FindAllOptions{
			Search: &options.SearchOptions{
				Keyword: name,
				Fields: []string{
					"name",
				},
			},
		},
	)

	response.Success(c, roles, err)
}

func (r *RoleResource) Show(c *gin.Context) {
	var id = c.Param("id")

	role, err := r.repository.Find(id, nil)

	response.Success(c, role, err)
}
