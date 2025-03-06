package common

import (
	"veo/pkg/errors"

	"github.com/gin-gonic/gin"
)

// ParseRequest parses the incoming JSON request body into the specified struct.
// Returns true if binding is successful, otherwise sends an error response and returns false.
func ParseRequest(c *gin.Context, out interface{}) bool {
	if err := c.ShouldBindJSON(out); err != nil {
		RespondError(c, errors.NewInvalidParams("Invalid request format"))
		return false
	}
	return true
}

// ParseQuery parses the incoming query parameters into the specified struct.
// Returns true if binding is successful, otherwise sends an error response and returns false.
func ParseQuery(c *gin.Context, out interface{}) bool {
	if err := c.ShouldBindQuery(out); err != nil {
		RespondError(c, errors.NewInvalidParams("Invalid query parameters"))
		return false
	}
	return true
}

// ParseForm parses the incoming form data into the specified struct.
// Returns true if binding is successful, otherwise sends an error response and returns false.
func ParseForm(c *gin.Context, out interface{}) bool {
	if err := c.ShouldBind(out); err != nil {
		RespondError(c, errors.NewInvalidParams("Invalid form data"))
		return false
	}
	return true
}
