package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/yatoyun/todo-app/api/domain/application/repository"
	"github.com/yatoyun/todo-app/api/domain/entity"
	"time"
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
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (i *TodoInteractor) Create(ctx context.Context, req CreateTodoRequest) (response *TodoResponse, err error) {
	todo := &entity.Todo{
		ID:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = i.Repository.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	response, err = i.convertToTodoResponse(todo)
	return response, err
}

func (i *TodoInteractor) GetList(ctx context.Context) ([]*TodoResponse, error) {
	todos, err := i.Repository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	todoResponses := make([]*TodoResponse, 0, len(todos))
	for _, todo := range todos {
		response, err := i.convertToTodoResponse(todo)
		if err != nil {
			return nil, err
		}
		todoResponses = append(todoResponses, response)
	}
	return todoResponses, nil
}

func (i *TodoInteractor) GetByID(ctx context.Context, id string) (response *TodoResponse, err error) {
	todo, err := i.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	response, err = i.convertToTodoResponse(todo)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (i *TodoInteractor) Update(ctx context.Context, id string, req UpdateTodoRequest) (err error) {
	todo, err := i.Repository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	setIfNotNil(&todo.Title, req.Title)
	setIfNotNil(&todo.Description, req.Description)
	setIfNotNil(&todo.Completed, req.Completed)
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

func (i *TodoInteractor) convertToTodoResponse(todo *entity.Todo) (*TodoResponse, error) {
	outputData := &TodoOutputData{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
	return i.OutputPort.Convert(*outputData)
}
