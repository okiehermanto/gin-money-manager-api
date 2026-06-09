package dto

type CreateTodoRequest struct {
	Todo string `json:"todo" binding:"required,min=3,max=100"`
}
