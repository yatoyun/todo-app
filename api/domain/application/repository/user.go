package repository

import (
	"context"

	"github.com/yatoyun/todo-app/api/domain/entity"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *entity.User) error
	GetList(ctx context.Context) ([]*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByAuth0ID(ctx context.Context, auth0ID string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}