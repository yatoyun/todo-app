package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/yatoyun/todo-app/api/domain/application/usecase"
	"github.com/yatoyun/todo-app/api/interface/controller"
	"github.com/yatoyun/todo-app/api/interface/presenter"
	"github.com/yatoyun/todo-app/api/interface/repository"
)

func InitializeUserController(conn *sqlx.DB) *controller.UserController {
	userPresenter := presenter.NewUserPresenter()
	userRepository := repository.NewUserRepository(conn)
	userInteractor := usecase.NewUserInteractor(userPresenter, userRepository)
	return controller.NewUserController(userInteractor)
}
