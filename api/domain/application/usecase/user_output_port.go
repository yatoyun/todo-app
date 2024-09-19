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
	ID        string
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserOutputPortInterface interface {
	Convert(UserOutputData) (*UserResponse, error)
}
