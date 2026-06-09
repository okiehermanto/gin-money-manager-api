package repository

import (
	baseRepository "gin-money-manager-api/modules/shared/repository"
	"gin-money-manager-api/modules/user/entity"

	"gorm.io/gorm"
)

type roleRepository struct {
	baseRepository.BaseRepository[entity.Role]
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		BaseRepository: baseRepository.NewBaseRepository[entity.Role](db),
	}
}
