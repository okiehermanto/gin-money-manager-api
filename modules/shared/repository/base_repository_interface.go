package repository

import (
	"gin-money-manager-api/modules/shared/repository/options"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	FindAll(options *options.FindAllOptions) ([]T, error)
	Find(id string, options *options.FindOptions) (T, error)
	FindBy(where map[string]any) ([]T, error)
	Create(entity T) (T, error)
	Update(entity T) (T, error)
	Delete(id string) error
	DB() *gorm.DB
}
