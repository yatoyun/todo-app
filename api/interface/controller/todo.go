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
	return &TodoController{InputPort: inputPort, Validator: validator.New()}
}

// isRequestValidTodo リクエストバリデーション
func (c *TodoController) isRequestValidTodo(m *entity.Todo) (bool, error) {
	err := c.Validator.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateTodo godoc
// @Summary todoを作成
// @Tags CreateTodo
// @Accept json
// @Produce json
// @Param request body entity.Todo true "Todo"
// @Success 201 {object} entity.Todo
// @Router /todos [post]
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

// GetTodos godoc
// @Summary todo一覧を取得
// @Tags GetTodos
// @Produce json
// @Success 200 {array} entity.Todo
// @Router /todos [get]
func (c *TodoController) GetTodos(ctx *gin.Context) {
	context := ctx.Request.Context()
	todos, err := c.InputPort.GetList(context)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
// @Summary todoを取得
// @Tags GetTodoByID
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} entity.Todo
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

// UpdateTodo godoc
// @Summary todoを更新
// @Tags UpdateTodo
// @Accept json
// @Produce json
// @Param request body entity.Todo true "Todo"
// @Success 200 {object} entity.Todo
// @Router /todos/update [post]
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

// DeleteTodo godoc
// @Summary todoを削除
// @Tags DeleteTodo
// @Param id path string true "Todo ID"
// @Success 204
// @Router /todos/{id} [delete]
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
