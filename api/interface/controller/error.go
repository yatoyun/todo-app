package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yatoyun/todo-app/api/domain/entity"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func handleError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(getStatusCode(err), ResponseError{Message: err.Error()})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case entity.ErrInternalServerError:
		return http.StatusInternalServerError
	case entity.ErrNotFound:
		return http.StatusNotFound
	case entity.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}