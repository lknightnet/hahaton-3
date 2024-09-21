package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    bool      `json:"status"`
	UUID      string    `json:"uuid"`
	TypeUser  string    `json:"type_user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	Student    = "student"
	University = "university"
	Enterprise = "enterprise"
)
