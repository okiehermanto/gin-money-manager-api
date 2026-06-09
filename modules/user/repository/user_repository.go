package repository

import (
	baseRepository "gin-money-manager-api/modules/shared/repository"
	"gin-money-manager-api/modules/user/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	baseRepository.BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: baseRepository.NewBaseRepository[entity.User](db),
	}
}

func (repository *userRepository) FindByUsername(
	username string,
) (*entity.User, error) {

	var user entity.User

	err := repository.DB().
		Where("username = ?", username).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
