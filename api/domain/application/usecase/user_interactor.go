package usecase

import (
	"context"
	"time"

	"github.com/yatoyun/todo-app/api/domain/application/repository"
	"github.com/yatoyun/todo-app/api/domain/entity"
)

type UserInteractor struct {
	OutputPort UserOutputPortInterface
	Repository repository.UserRepositoryInterface
}

func NewUserInteractor(
	outputPort UserOutputPortInterface,
	repository repository.UserRepositoryInterface,
) UserInputPortInterface {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (i *UserInteractor) Create(ctx context.Context, user *entity.User) (err error) {
	existingUser, err := i.Repository.GetByAuth0ID(ctx, user.Auth0ID)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return entity.ErrConflict
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user, err = i.Repository.Create(ctx, user)
	return err
}

func (i *UserInteractor) GetList(ctx context.Context) ([]UserResponse, error) {
	users, err := i.Repository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	userResponses := make([]UserResponse, 0, len(users))
	for _, user := range users {
		outputData := UserOutputData{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			Auth0ID: user.Auth0ID,
			Role: user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		response, err := i.OutputPort.Convert(outputData)
		if err != nil {
			return nil, err
		}
		userResponses = append(userResponses, *response)
	}
	return userResponses, nil
}

func (i *UserInteractor) GetByID(ctx context.Context, id string) (user *entity.User, err error) {
	user, err = i.Repository.GetByID(ctx, id)
	return user, err
}

func (i *UserInteractor) GetByAuth0ID(ctx context.Context, auth0ID string) (user *entity.User, err error) {
	user, err = i.Repository.GetByAuth0ID(ctx, auth0ID)
	return user, err
}

func (i *UserInteractor) Update(ctx context.Context, user *entity.User) (err error) {
	user.UpdatedAt = time.Now()
	user, err = i.Repository.Update(ctx, user)
	return err
}

func (i *UserInteractor) Delete(ctx context.Context, user *entity.User) (err error) {
	err = i.Repository.Delete(ctx, user)
	return err
}