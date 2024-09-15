package presenter

import (
	"github.com/yatoyun/todo-app/api/domain/application/usecase"
)

type TodoPresenter struct {
}

func NewTodoPresenter() usecase.TodoOutputPortInterface {
	return &TodoPresenter{}
}

func (p *TodoPresenter) Convert(outputData usecase.TodoOutputData) (*usecase.TodoResponse, error) {
	return &usecase.TodoResponse{
		ID:          outputData.ID,
		Title:       outputData.Title,
		Description: outputData.Description,
		Completed:   outputData.Completed,
		CreatedAt:   outputData.CreatedAt,
		UpdatedAt:   outputData.UpdatedAt,
	}, nil
}