package entity

import (
	"time"
)

type User struct {
	ID        string    `json:"id" db:"id"` // UUIDを想定
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Auth0ID   string    `json:"auth0_id" db:"auth0_id"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
