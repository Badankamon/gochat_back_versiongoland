package response

import (
	"net/http"

	"github.com/Badankamon/gochat_backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Error   *ErrorDetails `json:"error,omitempty"`
}

type ErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	var appErr *errors.AppError
	if e, ok := err.(*errors.AppError); ok {
		appErr = e
	} else {
		appErr = errors.ErrInternalServer
	}

	c.JSON(appErr.HttpCode, Response{
		Success: false,
		Message: appErr.Message,
		Error: &ErrorDetails{
			Code:    appErr.Code,
			Message: appErr.Message,
		},
	})
}
