package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"

	"github.com/yatoyun/todo-app/api/domain/entity"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func handleError(ctx *gin.Context, err error) {
	// Log the error message with more details
	slog.Error("An error occurred", slog.String("error", err.Error()))

	// Get the appropriate HTTP status code for the error
	statusCode := getStatusCode(err)

	// Send the error response with the status code and message
	ctx.AbortWithStatusJSON(statusCode, ResponseError{Message: err.Error(), Code: statusCode})
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
