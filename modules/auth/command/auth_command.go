package command

import (
	"gin-money-manager-api/modules/auth/dto"
	"gin-money-manager-api/modules/auth/service"
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/repository"

	"github.com/gin-gonic/gin"
)

type AuthCommand struct {
	userRepository repository.UserRepository
}

func NewAuthCommand(
	userRepository repository.UserRepository,
) *AuthCommand {
	return &AuthCommand{
		userRepository: userRepository,
	}
}

func (command *AuthCommand) Handler(
	c *gin.Context,
) {
	var body dto.LoginRequest

	if !helper.BindAndValidate(c, &body) {
		return
	}

	user, err := command.userRepository.FindByUsername(
		body.Username,
	)

	if err != nil {
		response.ServerError(c, err.Error())
	}

	if !helper.CheckPassword(
		body.Password,
		user.Password,
	) {
		response.ServerError(c, err.Error())
	}

	token, err := service.GenerateToken(
		user.ID.String(),
		user.Username,
	)

	if err != nil {
		response.ServerError(c, err.Error())
	}

	response.Success(
		c,
		map[string]string{
			"token": token,
		},
		err,
	)
}
