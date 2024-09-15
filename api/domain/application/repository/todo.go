package repository

import (
	"context"

	"github.com/yatoyun/todo-app/api/domain/entity"
)

type TodoRepositoryInterface interface {
	Create(ctx context.Context, todo *entity.Todo) error
	GetList(ctx context.Context) ([]*entity.Todo, error)
	GetByID(ctx context.Context, id string) (*entity.Todo, error)
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id string) error
}