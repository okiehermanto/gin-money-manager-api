package entity

import "github.com/google/uuid"

type Role struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name  string    `json:"name"`
	Label string    `json:"label"`
}
