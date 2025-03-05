package v1

import (
	"veo/internal/api/resp"
	"veo/internal/service"
	CustomError "veo/pkg/errors"

	"github.com/gin-gonic/gin"
)

// AccountAPI handles user authentication and account management
type AccountAPI struct {
	userService service.UserService
}

// NewAccountAPI creates a new instance of AccountAPI
func NewAccountAPI(userService service.UserService) *AccountAPI {
	return &AccountAPI{userService: userService}
}

// SetupAccountRouter configures account-related routes
func SetupAccountRouter(router *gin.Engine, api *AccountAPI) {
	protected := router.Group("/api")

	// Public endpoints (No authentication required)
	protected.POST("/register", api.Register)
	protected.POST("/login", api.Login)

	// Protected endpoints (Require JWT authentication)
	protected.Use(service.AuthMiddleware())
	{
		protected.POST("/updatePassword", api.UpdatePassword)
	}
}

// Register handles user registration
func (api *AccountAPI) Register(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Validate JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		resp.JSONResponse(c, nil, CustomError.NewError(CustomError.CodeInvalidParams, "Invalid request format"))
		return
	}

	// Register the user
	err := api.userService.Register(request.Username, request.Password)
	if err != nil {
		resp.JSONResponse(c, nil, err)
		return
	}

	// Retrieve the newly registered user
	user, err := api.userService.GetUserByUsername(request.Username)
	if err != nil {
		resp.JSONResponse(c, nil, err)
		return
	}

	// Generate JWT token for authentication
	token, err := service.GenerateJWT(user.ID, user.Username)
	if err != nil {
		resp.JSONResponse(c, nil, CustomError.NewError(CustomError.CodeError, "Failed to generate authentication token"))
		return
	}

	resp.JSONResponse(c, token, nil)
}

// Login handles user authentication
func (api *AccountAPI) Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Validate JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		resp.JSONResponse(c, nil, CustomError.NewError(CustomError.CodeInvalidParams, "Invalid request format"))
		return
	}

	// Authenticate user
	user, err := api.userService.Login(request.Username, request.Password)
	if err != nil {
		resp.JSONResponse(c, nil, err)
		return
	}

	// Generate JWT token
	token, err := service.GenerateJWT(user.ID, user.Username)
	if err != nil {
		resp.JSONResponse(c, nil, CustomError.NewError(CustomError.CodeError, "Failed to generate authentication token"))
		return
	}

	resp.JSONResponse(c, token, nil)
}

// UpdatePassword allows users to change their password
func (api *AccountAPI) UpdatePassword(c *gin.Context) {
	var request struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	// Validate JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		resp.JSONResponse(c, nil, CustomError.NewError(CustomError.CodeInvalidParams, "Invalid request format"))
		return
	}

	// Get user ID from JWT claims
	id := c.MustGet("userId").(int)

	// Fetch user details
	user, err := api.userService.GetUserByID(id)
	if err != nil {
		resp.JSONResponse(c, nil, err)
		return
	}

	// Attempt password update
	err = api.userService.UpdatePassword(user.ID, request.OldPassword, request.NewPassword)
	if err != nil {
		resp.JSONResponse(c, nil, err)
		return
	}

	resp.JSONResponse(c, "Password updated successfully", nil)
}
