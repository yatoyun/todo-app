package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/yatoyun/todo-app/api/domain/application/repository"
	"github.com/yatoyun/todo-app/api/domain/entity"
	"time"
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

func (i *UserInteractor) Create(ctx context.Context, req CreateUserRequest) (response *UserResponse, err error) {
	existingUser, _ := i.Repository.GetByAuth0ID(ctx, req.Auth0ID)
	if existingUser != nil {
		return nil, entity.ErrConflict
	}
	user := &entity.User{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Email:     req.Email,
		Auth0ID:   req.Auth0ID,
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err = i.Repository.Create(ctx, user)
	response, err = i.convertToUserResponse(user)
	return response, err
}

func (i *UserInteractor) GetList(ctx context.Context) ([]*UserResponse, error) {
	users, err := i.Repository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	userResponses := make([]*UserResponse, 0, len(users))
	for _, user := range users {
		response, err := i.convertToUserResponse(user)
		if err != nil {
			return nil, err
		}
		userResponses = append(userResponses, response)
	}
	return userResponses, nil
}

func (i *UserInteractor) GetByID(ctx context.Context, id string) (*UserResponse, error) {
	user, err := i.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	response, err := i.convertToUserResponse(user)
	return response, err
}

func (i *UserInteractor) GetByAuth0ID(ctx context.Context, auth0ID string) (*UserResponse, error) {
	user, err := i.Repository.GetByAuth0ID(ctx, auth0ID)
	if err != nil {
		return nil, err
	}
	response, err := i.convertToUserResponse(user)
	return response, err
}

func (i *UserInteractor) Update(ctx context.Context, id string, req UpdateUserRequest) (err error) {
	user, err := i.Repository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	setIfNotNil(&user.Name, req.Name)
	setIfNotNil(&user.Email, req.Email)
	setIfNotNil(&user.Role, req.Role)
	user.UpdatedAt = time.Now()
	err = i.Repository.Update(ctx, user)
	return err
}

func (i *UserInteractor) Delete(ctx context.Context, userID string) (err error) {
	err = i.Repository.Delete(ctx, userID)
	return err
}

func (i *UserInteractor) convertToUserResponse(user *entity.User) (*UserResponse, error) {
	outputData := &UserOutputData{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return i.OutputPort.Convert(*outputData)
}

func setIfNotNil[T any](dest *T, src *T) {
	if src != nil {
		*dest = *src
	}
}
