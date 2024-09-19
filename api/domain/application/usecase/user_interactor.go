package usecase

import (
	"context"
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
	existingUser, _ := i.Repository.GetByAuth0ID(ctx, user.Auth0ID)
	if existingUser != nil {
		return entity.ErrConflict
	}
	err = i.Repository.Create(ctx, user)
	return err
}

func (i *UserInteractor) GetList(ctx context.Context) ([]UserResponse, error) {
	users, err := i.Repository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	userResponses := make([]UserResponse, 0, len(users))
	for _, user := range users {
		outputData := &UserOutputData{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Auth0ID:   user.Auth0ID,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		response, err := i.OutputPort.Convert(*outputData)
		if err != nil {
			return nil, err
		}
		userResponses = append(userResponses, *response)
	}
	return userResponses, nil
}

func (i *UserInteractor) GetByID(ctx context.Context, id string) (UserResponse, error) {
	user, err := i.Repository.GetByID(ctx, id)
	if err != nil {
		return UserResponse{}, err
	}
	outputData := &UserOutputData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Auth0ID:   user.Auth0ID,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	response, err := i.OutputPort.Convert(*outputData)
	return *response, err
}

func (i *UserInteractor) GetByAuth0ID(ctx context.Context, auth0ID string) (UserResponse, error) {
	user, err := i.Repository.GetByAuth0ID(ctx, auth0ID)
	if err != nil {
		return UserResponse{}, err
	}
	outputData := &UserOutputData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Auth0ID:   user.Auth0ID,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	response, err := i.OutputPort.Convert(*outputData)
	return *response, err
}

func (i *UserInteractor) Update(ctx context.Context, user *entity.User) (err error) {
	err = i.Repository.Update(ctx, user)
	return err
}

func (i *UserInteractor) Delete(ctx context.Context, userID string) (err error) {
	err = i.Repository.Delete(ctx, userID)
	return err
}
