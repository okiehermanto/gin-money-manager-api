package user

import (
	"gin-money-manager-api/modules/user/command"
	"gin-money-manager-api/modules/user/repository"
	"gin-money-manager-api/modules/user/resource"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	roleResource := resource.NewRoleResource(roleRepository)

	CreateUser := command.NewCreateUser(userRepository, roleRepository)
	UserResource := resource.NewUserResource(userRepository)

	userContainer := &UserContainer{
		RoleRepository: roleRepository,
		UserRepository: userRepository,

		RoleResource: roleResource,
		CreateUser:   CreateUser,
		UserResource: UserResource,
	}

	RegisterRoute(r, userContainer)
}
