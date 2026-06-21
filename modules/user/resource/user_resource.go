package resource

import (
	"gin-money-manager-api/modules/shared/dto"
	"gin-money-manager-api/modules/shared/resource"
	"gin-money-manager-api/modules/user/entity"
	"gin-money-manager-api/modules/user/repository"
)

type UserResource struct {
	*resource.BaseResource[
		entity.User,
		dto.EmptyDto,
		dto.EmptyDto,
	]
}

func NewUserResource(repository repository.UserRepository) *UserResource {
	return &UserResource{
		BaseResource: &resource.BaseResource[entity.User, dto.EmptyDto, dto.EmptyDto]{
			Repository: repository,
			RouteParam: entity.User{}.EntityName(),
			Relationships: []string{
				"Roles",
			},
			SearchFields: []string{
				"name",
				"username",
				"email",
			},
		},
	}
}
