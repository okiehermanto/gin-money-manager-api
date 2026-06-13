package container

import (
	"gin-money-manager-api/modules/user/command"
	"gin-money-manager-api/modules/user/repository"
	"gin-money-manager-api/modules/user/resource"
)

type UserContainer struct {
	RoleRepository repository.RoleRepository
	UserRepository repository.UserRepository

	RoleResource *resource.RoleResource
	UserResource *resource.UserResource

	CreateUser *command.CreateUser
}
