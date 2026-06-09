package repository

import (
	"gin-money-manager-api/modules/shared/repository"
	"gin-money-manager-api/modules/user/entity"
)

type RoleRepository interface {
	repository.BaseRepository[entity.Role]
}
