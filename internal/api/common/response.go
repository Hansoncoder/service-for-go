package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized API response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// RespondError sends an error response
func RespondError(c *gin.Context, err error) {
	if customErr, ok := err.(interface{ GetCode() int }); ok {
		c.JSON(http.StatusOK, Response{
			Code:    customErr.GetCode(),
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	})
}

// RespondData sends a success response with data
func RespondData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

// RespondMessage sends a success response with a custom message
func RespondMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
	})
}

// AbortIfError checks if an error occurred, responds with the error, and aborts the current request processing.
// Returns true if an error was encountered and the request was aborted, otherwise false.
func AbortIfError(c *gin.Context, err error) bool {
	if err != nil {
		RespondError(c, err)
		c.Abort()
		return true
	}
	return false
}
