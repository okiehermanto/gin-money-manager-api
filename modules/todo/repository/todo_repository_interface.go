package repository

import (
	baseRepository "gin-money-manager-api/modules/shared/repository"

	"gin-money-manager-api/modules/todo/entity"
)

type TodoRepository interface {
	baseRepository.BaseRepository[entity.Todo]

	FindCompleted() ([]entity.Todo, error)
}
