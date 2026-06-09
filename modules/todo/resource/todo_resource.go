package resource

import (
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/todo/dto"
	"gin-money-manager-api/modules/todo/entity"
	"gin-money-manager-api/modules/todo/repository"

	"github.com/gin-gonic/gin"
)

type TodoResource struct {
	todoRepository repository.TodoRepository
}

func NewTodoResource(todoRepository repository.TodoRepository) *TodoResource {
	return &TodoResource{
		todoRepository: todoRepository,
	}
}

func (h *TodoResource) CreateTodo(c *gin.Context) {
	var body dto.CreateTodoRequest

	if !helper.BindAndValidate(c, &body) {
		return
	}

	newTodo := entity.Todo{
		Todo:      body.Todo,
		Completed: false,
	}

	todo, err := h.todoRepository.Create(newTodo)

	response.Created(c, todo, err)
}

func (h *TodoResource) GetTodos(c *gin.Context) {
	todos, err := h.todoRepository.FindAll(nil)

	response.Success(c, todos, err)
}

func (h *TodoResource) GetTodo(c *gin.Context) {
	var id = c.Param("id")

	todo, err := h.todoRepository.Find(id, nil)

	response.Success(c, todo, err)
}

func (h *TodoResource) UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var body dto.UpdateTodoRequest

	if !helper.BindAndValidate(c, &body) {
		return
	}

	todo, err := h.todoRepository.Find(id, nil)

	if err != nil {
		response.NotFound(c)
		return
	}

	todo.Todo = body.Todo
	todo.Completed = body.Completed

	todo, err = h.todoRepository.Update(todo)

	response.Success(c, todo, err)
}

func (h *TodoResource) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := h.todoRepository.Find(id, nil)

	err = h.todoRepository.Delete(id)

	response.Success(c, todo, err)
}
