package v1

import (
	"log"
	"net/http"

	"veo/internal/service"

	"github.com/gin-gonic/gin"
)

// UserAPI provides API endpoints for user-related operations.
type UserAPI struct {
	userService service.UserService
}

// NewUserAPI initializes a new UserAPI instance.
func NewUserAPI(userService service.UserService) *UserAPI {
	return &UserAPI{userService: userService}
}

// SetupUserRouter sets up the user-related routes in the Gin engine.
func SetupUserRouter(router *gin.Engine, api *UserAPI) {
	protected := router.Group("/api")

	// Apply JWT authentication middleware
	protected.Use(service.AuthMiddleware())
	{
		protected.GET("/getUserInfo", api.GetUserInfo) // Route for retrieving user information
	}
}

// GetUserInfo handles requests to fetch user information based on JWT authentication.
func (api *UserAPI) GetUserInfo(c *gin.Context) {
	// Retrieve username from the JWT claims stored in the context
	username := c.MustGet("username").(string)
	log.Printf("Fetching user info for: %s", username)

	// Get user details from the service layer
	user, err := api.userService.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Respond with sanitized user information
	c.JSON(http.StatusOK, user.Sanitize())
}
