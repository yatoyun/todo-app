package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/yatoyun/todo-app/api/domain/application/usecase"
)

// Controller
type TodoController struct {
	InputPort usecase.TodoInputPortInterface
	Validator *validator.Validate
}

func NewTodoController(inputPort usecase.TodoInputPortInterface) *TodoController {
	return &TodoController{InputPort: inputPort, Validator: validator.New()}
}

// CreateTodo godoc
// @Summary todoを作成
// @Tags CreateTodo
// @Accept json
// @Produce json
// @Param request body usecase.CreateTodoRequest true "Todo"
// @Success 201 {object} usecase.TodoResponse
// @Router /todos [post]
func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var req usecase.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, err)
		return
	}

	if ok, err := c.isRequestValidCreateTodo(&req); !ok {
		handleError(ctx, err)
		return
	}

	context := ctx.Request.Context()
	todo, err := c.InputPort.Create(context, req)
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
// @Success 200 {array} usecase.TodoResponse
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
// @Router /todos/{id} [get]
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
// @Param id path string true "Todo ID"
// @Param request body usecase.UpdateTodoRequest true "Todo"
// @Success 200 {object} usecase.TodoResponse
// @Router /todos/{id} [put]
func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	var req usecase.UpdateTodoRequest
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, err)
		return
	}

	if ok, err := c.isRequestValidUpdateTodo(&req); !ok {
		handleError(ctx, err)
		return
	}
	context := ctx.Request.Context()
	err := c.InputPort.Update(context, id, req)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
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
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Todo deleted successfully"})
}

func (c *TodoController) isRequestValidCreateTodo(req *usecase.CreateTodoRequest) (bool, error) {
	err := c.Validator.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *TodoController) isRequestValidUpdateTodo(req *usecase.UpdateTodoRequest) (bool, error) {
	err := c.Validator.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}
