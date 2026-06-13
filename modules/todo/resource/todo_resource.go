package resource

import (
	"gin-money-manager-api/modules/shared/resource"
	"gin-money-manager-api/modules/todo/dto"
	"gin-money-manager-api/modules/todo/entity"
	"gin-money-manager-api/modules/todo/repository"
)

type TodoResource struct {
	*resource.BaseResource[
		entity.Todo,
		dto.CreateTodoRequest,
		dto.UpdateTodoRequest,
	]
}

func NewTodoResource(
	repo repository.TodoRepository,
) *TodoResource {
	return &TodoResource{
		BaseResource: &resource.BaseResource[
			entity.Todo,
			dto.CreateTodoRequest,
			dto.UpdateTodoRequest,
		]{
			Repository: repo,
			SearchFields: []string{
				"todo",
			},
		},
	}
}
