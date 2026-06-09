package query

import (
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/todo/repository"

	"github.com/gin-gonic/gin"
)

type TodoFindCompleted struct {
	todoRepository repository.TodoRepository
}

func NewTodoFindCompleted(todoRepository repository.TodoRepository) *TodoFindCompleted {
	return &TodoFindCompleted{
		todoRepository: todoRepository,
	}
}

func (h *TodoFindCompleted) Handler(c *gin.Context) {
	todos, err := h.todoRepository.FindCompleted()

	response.Success(c, todos, err)
}
