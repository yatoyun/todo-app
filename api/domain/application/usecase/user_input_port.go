package usecase

import (
	"context"
)

type CreateUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Auth0ID string `json:"auth0_id" binding:"required"`
	Role    string `json:"role" binding:"required,oneof=admin user"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`
	Email *string `json:"email" binding:"email"`
	Role  *string `json:"role"`
}

type UserInputPortInterface interface {
	Create(ctx context.Context, req CreateUserRequest) (UserResponse, error)
	GetByID(ctx context.Context, id string) (UserResponse, error)
	GetList(ctx context.Context) ([]UserResponse, error)
	GetByAuth0ID(ctx context.Context, auth0ID string) (UserResponse, error)
	Update(ctx context.Context, req UpdateUserRequest) error
	Delete(ctx context.Context, id string) error
}
