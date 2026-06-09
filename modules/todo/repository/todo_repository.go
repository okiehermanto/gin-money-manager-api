package repository

import (
	baseRepository "gin-money-manager-api/modules/shared/repository"
	"gin-money-manager-api/modules/todo/entity"

	"gorm.io/gorm"
)

type todoRepository struct {
	baseRepository.BaseRepository[entity.Todo]
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		BaseRepository: baseRepository.NewBaseRepository[entity.Todo](db),
	}
}

func (r *todoRepository) FindCompleted() ([]entity.Todo, error) {

	var todos []entity.Todo

	result := r.DB().
		Where("completed = ?", true).
		Find(&todos)

	// result := r.db.
	// 	Where("completed = ?", true).
	// 	Find(&todos)

	return todos, result.Error
}
