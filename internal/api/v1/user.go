package v1

import (
	"veo/internal/service"
	"veo/internal/utils"

	"github.com/gin-gonic/gin"
)

var logger = utils.GetLogger()

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
	protected.Use(AuthMiddleware())
	{
		protected.GET("/getUserInfo", api.GetUserInfo) // Route for retrieving user information
	}
}

// GetUserInfo handles requests to fetch user information based on JWT authentication.
func (api *UserAPI) GetUserInfo(c *gin.Context) {
	// Retrieve username from the JWT claims stored in the context
	username := c.MustGet("username").(string)
	logger.Infof("Fetching user info for: %s", username)

	// Get user details from the service layer
	user, err := api.userService.GetUserByUsername(username)
	if AbortIfError(c, err) {
		return
	}

	// Respond with sanitized user information
	RespondData(c, user.Sanitize())
}
