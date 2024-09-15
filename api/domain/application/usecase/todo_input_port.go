package usecase

import (
	"context"

	"github.com/yatoyun/todo-app/api/domain/entity"
)

// input Boundary <I>
type TodoInputPortInterface interface {
	Create(ctx context.Context, todo *entity.Todo) error
	GetList(ctx context.Context) ([]TodoResponse, error)
	GetByID(ctx context.Context, id string) (TodoResponse, error)
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id string) error
}
