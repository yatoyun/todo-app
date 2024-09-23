package usecase

import "time"

type UserOutputData struct {
	ID        string
	Name      string
	Email     string
	Auth0ID   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserOutputPortInterface interface {
	Convert(UserOutputData) (*UserResponse, error)
}
