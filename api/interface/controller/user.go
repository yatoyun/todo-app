package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yatoyun/todo-app/api/domain/application/usecase"
	"github.com/yatoyun/todo-app/api/domain/entity"
)

type UserController struct {
	InputPort usecase.UserInputPortInterface
	Validator *validator.Validate
}

func NewUserController(inputPort usecase.UserInputPortInterface) *UserController {
	return &UserController{InputPort: inputPort, Validator: validator.New()}
}

func (c *UserController) isRequestValidUser(m *entity.User) (bool, error) {
	err := c.Validator.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateUser godoc
// @Summary userを作成する
// @tags CreateUser
// @Description userを作成する
// @Accept json
// @Produce json
// @Param user body entity.User true "user"
// @Success 201 {object} entity.User
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	if ok, err := uc.isRequestValidUser(&user); !ok {
		handleError(c, err)
		return
	}

	context := c.Request.Context()
	err := uc.InputPort.Create(context, &user)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUsers godoc
// @Summary user一覧を取得する
// @tags GetUsers
// @Description user一覧を取得する
// @Produce json
// @Success 200 {array} entity.User
// @Router /users [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	context := c.Request.Context()
	users, err := uc.InputPort.GetList(context)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary userをIDで取得する
// @tags GetUserByID
// @Description userをIDで取得する
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} entity.User
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	context := c.Request.Context()
	user, err := uc.InputPort.GetByID(context, id)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUserByAuth0ID godoc
// @Summary userをAuth0IDで取得する
// @tags GetUserByAuth0ID
// @Description userをAuth0IDで取得する
// @Produce json
// @Param auth0_id path string true "Auth0ID"
// @Success 200 {object} entity.User
// @Router /users/auth0/{auth0_id} [get]
func (uc *UserController) GetUserByAuth0ID(c *gin.Context) {
	auth0ID := c.Param("auth0_id")

	context := c.Request.Context()
	user, err := uc.InputPort.GetByAuth0ID(context, auth0ID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary userを更新する
// @tags UpdateUser
// @Description userを更新する
// @Accept json
// @Produce json
// @Param user body entity.User true "user"
// @Router /users/update/ [post]
func (uc *UserController) UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	if ok, err := uc.isRequestValidUser(&user); !ok {
		handleError(c, err)
		return
	}

	context := c.Request.Context()
	err := uc.InputPort.Update(context, &user)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary userを削除する
// @tags DeleteUser
// @Description userを削除する
// @Param id path string true "ID"
// @Success 200 {string} string "User deleted successfully"
// @Router /users/delete/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	context := c.Request.Context()
	err := uc.InputPort.Delete(context, id)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
