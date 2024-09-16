package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/yatoyun/todo-app/api/domain/application/usecase"
	"github.com/yatoyun/todo-app/api/interface/controller"
	"github.com/yatoyun/todo-app/api/interface/presenter"
	"github.com/yatoyun/todo-app/api/interface/repository"
)

func InitializeTodoController(conn *sqlx.DB) *controller.TodoController {
	todoPresenter := presenter.NewTodoPresenter()
	todoRespository := repository.NewTodoRepository(conn)
	todoInteractor := usecase.NewTodoInteractor(todoPresenter, todoRespository)
	return controller.NewTodoController(todoInteractor)
}
