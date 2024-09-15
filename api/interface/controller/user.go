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

func (uc *UserController) Create(c *gin.Context) {
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

func (uc *UserController) GetUsers(c *gin.Context) {
	context := c.Request.Context()
	users, err := uc.InputPort.GetList(context)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

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