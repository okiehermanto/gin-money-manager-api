package resource

import (
	"gin-money-manager-api/modules/shared/dto"
	"gin-money-manager-api/modules/shared/resource"
	"gin-money-manager-api/modules/user/entity"
	"gin-money-manager-api/modules/user/repository"
)

type RoleResource struct {
	*resource.BaseResource[
		entity.Role,
		dto.EmptyDto,
		dto.EmptyDto,
	]
}

func NewRoleResource(repository repository.RoleRepository) *RoleResource {
	return &RoleResource{
		BaseResource: &resource.BaseResource[entity.Role, dto.EmptyDto, dto.EmptyDto]{
			Repository: repository,
			SearchFields: []string{
				"name",
			},
		},
	}
}
