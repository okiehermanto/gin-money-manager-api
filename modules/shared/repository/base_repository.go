package repository

import (
	"gin-money-manager-api/modules/shared/repository/options"
	"strings"

	"gorm.io/gorm"
)

type baseRepository[T any] struct {
	db *gorm.DB
}

func (b *baseRepository[T]) Create(entity T) (T, error) {
	result := b.db.Create(&entity)

	return entity, result.Error
}

func (b *baseRepository[T]) Delete(id string) error {
	var entity T

	result := b.db.Delete(&entity, "id = ?", id)

	return result.Error
}

func (b *baseRepository[T]) Find(id string, options *options.FindOptions) (T, error) {
	var entity T
	query := b.db

	if options != nil {
		for _, relationship := range options.Relationships {
			query = query.Preload(relationship)
		}
	}

	result := query.First(&entity, "id = ?", id)

	return entity, result.Error
}

func (b *baseRepository[T]) FindAll(options *options.FindAllOptions) ([]T, error) {
	var entities []T
	query := b.db

	if options != nil {
		for _, relationship := range options.Relationships {
			query = query.Preload(relationship)
		}

		for _, filter := range options.Filters {
			switch filter.Operator {

			case "IN":
				query = query.Where(
					filter.Field+" IN ?",
					filter.Value,
				)

			default:
				query = query.Where(
					filter.Field+" "+filter.Operator+" ?",
					filter.Value,
				)
			}
		}

		if options.Search != nil &&
			options.Search.Keyword != "" &&
			len(options.Search.Fields) > 0 {
			var conditions []string
			var values []any

			for _, field := range options.Search.Fields {
				conditions = append(
					conditions,
					field+" ILIKE ?",
				)

				values = append(
					values,
					"%"+options.Search.Keyword+"%",
				)
			}

			query = query.Where(
				strings.Join(conditions, " OR "),
				values...,
			)
		}

		if options.OrderBy != "" {
			query = query.Order(options.OrderBy)
		}
	}

	result := query.Find(&entities)

	return entities, result.Error
}

func (b *baseRepository[T]) Update(entity T) (T, error) {
	result := b.db.Save(&entity)

	return entity, result.Error
}

func (b *baseRepository[T]) FindBy(where map[string]any) ([]T, error) {
	var entities []T
	result := b.db.Where(where).Find(&entities)
	return entities, result.Error
}

func (b *baseRepository[T]) DB() *gorm.DB {
	return b.db
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

// user, err := userRepository.FindOneBy( map[string]any{ "email": "test@gmail.com", }, )
