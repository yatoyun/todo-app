package usecase

import (
	"context"
)

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      string `json:"user_id" validate:"required"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

// input Boundary <I>
type TodoInputPortInterface interface {
	Create(ctx context.Context, req CreateTodoRequest) (*TodoResponse, error)
	GetList(ctx context.Context) ([]*TodoResponse, error)
	GetByID(ctx context.Context, id string) (*TodoResponse, error)
	Update(ctx context.Context, id string, req UpdateTodoRequest) error
	Delete(ctx context.Context, id string) error
}
