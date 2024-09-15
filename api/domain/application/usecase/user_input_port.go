package usecase

import (
	"context"

	"github.com/yatoyun/todo-app/api/domain/entity"
)

type UserInputPortInterface interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (UserResponse, error)
	GetList(ctx context.Context) ([]UserResponse, error)
	GetByAuth0ID(ctx context.Context, auth0ID string) (UserResponse, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}