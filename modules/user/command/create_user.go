package command

import (
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/dto"
	"gin-money-manager-api/modules/user/entity"
	"gin-money-manager-api/modules/user/repository"

	"github.com/gin-gonic/gin"
)

type CreateUser struct {
	userRepository repository.UserRepository
	roleRepository repository.RoleRepository
}

func NewCreateUser(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
) *CreateUser {
	return &CreateUser{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (h *CreateUser) Handler(c *gin.Context) {
	var body dto.CreateUserRequest
	var roles []entity.Role

	if !helper.BindAndValidate(c, &body) {
		return
	}

	h.roleRepository.DB().Where("id IN ?", body.Roles).Find(&roles)

	hashedPassword, err := helper.HashPassword(body.Password)

	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	newUser := entity.User{
		Name:     body.Name,
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
		Roles:    roles,
	}

	user, err := h.userRepository.Create(newUser)

	response.Created(c, user, err)
}
