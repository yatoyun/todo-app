package presenter

import "github.com/yatoyun/todo-app/api/domain/application/usecase"

type UserPresenter struct {
}

func NewUserPresenter() usecase.UserOutputPortInterface {
	return &UserPresenter{}
}

func (p *UserPresenter) Convert(outputData usecase.UserOutputData) (*usecase.UserResponse, error) {
	return &usecase.UserResponse{
		ID:        outputData.ID,
		Name:      outputData.Name,
		Email:     outputData.Email,
		Role:      outputData.Role,
		CreatedAt: outputData.CreatedAt,
		UpdatedAt: outputData.UpdatedAt,
	}, nil
}
