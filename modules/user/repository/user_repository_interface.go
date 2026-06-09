package repository

import (
	"gin-money-manager-api/modules/shared/repository"
	"gin-money-manager-api/modules/user/entity"
)

type UserRepository interface {
	repository.BaseRepository[entity.User]

	FindByUsername(username string) (*entity.User, error)
}
