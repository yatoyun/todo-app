package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/yatoyun/todo-app/api/domain/application/usecase"
	"github.com/yatoyun/todo-app/api/domain/entity"
)

// Controller
type TodoController struct {
	InputPort usecase.TodoInputPortInterface
	Validator *validator.Validate
}

func NewTodoController(inputPort usecase.TodoInputPortInterface) *TodoController {
	return &TodoController{InputPort: inputPort}
}

// isRequestValidTodo リクエストバリデーション
func (c *TodoController) isRequestValidTodo(m *entity.Todo) (bool, error) {
	err := c.Validator.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo entity.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		handleError(ctx, err)
		return
	}

	if ok, err := c.isRequestValidTodo(&todo); !ok {
		handleError(ctx, err)
		return
	}

	context := ctx.Request.Context()
	err := c.InputPort.Create(context, &todo)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, todo)
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	context := ctx.Request.Context()
	todos, err := c.InputPort.GetList(context)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	context := ctx.Request.Context()
	todo, err := c.InputPort.GetByID(context, id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	var todo entity.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		handleError(ctx, err)
		return
	}

	if ok, err := c.isRequestValidTodo(&todo); !ok {
		handleError(ctx, err)
		return
	}
	context := ctx.Request.Context()
	err := c.InputPort.Update(context, &todo)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	context := ctx.Request.Context()
	err := c.InputPort.Delete(context, id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, ResponseError{Message: "Deleted"})
}