package entity

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Todo      string    `json:"todo"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
