package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `gorm:"not null" json:"name"`
	Email string    `json:"email"`
}
