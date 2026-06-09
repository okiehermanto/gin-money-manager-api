package dto

type UpdateTodoRequest struct {
	Todo      string `json:"todo" binding:"required,min=3,max=100"`
	Completed bool   `json:"completed"`
}
