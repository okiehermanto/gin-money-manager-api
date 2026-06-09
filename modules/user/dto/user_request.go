package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Name     string      `json:"name" binding:"required,min=3,max=100"`
	Username string      `json:"username" binding:"required,min=3,max=100"`
	Email    string      `json:"email" binding:"required,min=3,max=100"`
	Password string      `json:"password" binding:"min=3,max=100"`
	Roles    []uuid.UUID `json:"roles"`
}
