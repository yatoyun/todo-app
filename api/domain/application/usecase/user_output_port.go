package usecase

import "time"

type UserOutputData struct {
	Id string
	Name string
	Email string
	Auth0ID string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	Id string
	Name string
	Email string
	Auth0ID string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserOutputPortInterface interface {
	Convert(UserOutputData) (*UserResponse, error)
}