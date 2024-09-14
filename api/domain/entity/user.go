package entity

import (
	"time"
)

type User struct {
	ID string `json:"id"` // UUIDを想定
	Name string `json:"name"`
	Email string `json:"email"`
	Auth0ID string `json:"auth0_id"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}