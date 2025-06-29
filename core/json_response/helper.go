package json_response

import (
	"github.com/gin-gonic/gin"
	"github.com/vfyuliawan/my-golang-sdk/core/base_response"
	"net/http"
)

// Success wraps c.JSON for a success response and formats it as needed.
func Success[T any](c *gin.Context, result T) {
	c.JSON(http.StatusOK, base_response.Success(result))
}

// SuccessWithMessage wraps c.JSON for a success response and formats it as needed.
func SuccessWithMessage[T any](c *gin.Context, result T, message string) {
	c.JSON(http.StatusOK, base_response.SuccessWithMessage(result, message))
}

// Error wraps c.JSON for an error response and formats it as needed.
func Error(codeSystem, code, message string, httpStatus int) {
	panic(base_response.NewError(codeSystem, code, message, httpStatus))
}

// ErrorWithResult wraps c.JSON for an error response and formats it as needed.
func ErrorWithResult(codeSystem, code, message string, httpStatus int, result any) {
	panic(base_response.NewErrorWithResult(codeSystem, code, message, httpStatus, result))
}
