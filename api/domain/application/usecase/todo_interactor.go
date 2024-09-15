package usecase

import (
	"context"
	"time"

	"github.com/yatoyun/todo-app/api/domain/application/repository"
	"github.com/yatoyun/todo-app/api/domain/entity"
)


type TodoInteractor struct {
	Repository repository.TodoRepositoryInterface
	OutputPort TodoOutputPortInterface // 追加
}

func NewTodoInteractor(
	outputPort TodoOutputPortInterface,
	repository repository.TodoRepositoryInterface,
) TodoInputPortInterface {
	return &TodoInteractor{
		OutputPort: outputPort, // 追加
		Repository: repository,
	}
}

func (i *TodoInteractor) Create(ctx context.Context, todo *entity.Todo) (err error) {
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	err = i.Repository.Create(ctx, todo)
	return err
}

func (i *TodoInteractor) GetList(ctx context.Context) ([]TodoResponse, error) {
	todos, err := i.Repository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	todoResponses := make([]TodoResponse, 0, len(todos))
	for _, todo := range todos {
		outputData := TodoOutputData{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}
		response, err := i.OutputPort.Convert(outputData)
		if err != nil {
			return nil, err
		}
		todoResponses = append(todoResponses, *response)
	}
	return todoResponses, nil
}

func (i *TodoInteractor) GetByID(ctx context.Context, id string) (todo *entity.Todo, err error) {
	todo, err = i.Repository.GetByID(ctx, id)
	return todo, err
}

func (i *TodoInteractor) Update(ctx context.Context, todo *entity.Todo) (err error) {
	todo.UpdatedAt = time.Now()
	err = i.Repository.Update(ctx, todo)
	return err
}

func (i *TodoInteractor) Delete(ctx context.Context, id string) (err error) {
	existedTodo, err := i.Repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if existedTodo == nil {
		return entity.ErrNotFound
	}
	return i.Repository.Delete(ctx, id)
}