package resp

import (
	"errors"
	"net/http"
	CustomError "veo/pkg/errors"

	"github.com/gin-gonic/gin"
)

// Response represents the standard API response structure.
type Response struct {
	Code    int         `json:"code"`    // Response status code
	Message string      `json:"message"` // Response message
	Data    interface{} `json:"data"`    // Response payload
}

// JSONResponse sends a JSON response to the client based on the provided data and error.
// If an error is provided, it checks whether the error is a custom error type (CustomError.CustomError).
// - If it is a custom error, the response includes the specific error code and message.
// - If it is a general error, it responds with HTTP 500 (Internal Server Error) and a generic error message.
// If there is no error, the function returns the data with an HTTP 200 (OK) status.
func JSONResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		var be *CustomError.CustomError
		if errors.As(err, &be) {
			// Handle custom error and return its details
			c.JSON(http.StatusOK, Response{
				Code:    int(be.Code),
				Message: be.Message,
				Data:    data,
			})
		} else {
			// Handle generic errors with a 500 status
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: "service fail",
			})
		}
		return
	}

	// Return successful response with data
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: data,
	})
}
