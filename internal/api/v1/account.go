package v1

import (
	"veo/internal/api/common"
	"veo/internal/service"
	"veo/pkg/errors"

	"github.com/gin-gonic/gin"
)

// 导入 common 包中的函数到当前包
var (
	ParseRequest   = common.ParseRequest
	ParseQuery     = common.ParseQuery
	ParseForm      = common.ParseForm
	AbortIfError   = common.AbortIfError
	RespondData    = common.RespondData
	RespondMessage = common.RespondMessage
	GenerateJWT    = common.GenerateJWT
	AuthMiddleware = common.AuthMiddleware
	NewUserExists  = errors.NewUserExists
	NewAuthFailed  = errors.NewAuthFailed
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
	protected.Use(AuthMiddleware())
	{
		protected.POST("/updatePassword", api.UpdatePassword)
	}
}

// Register handles user registration
func (api *AccountAPI) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if !ParseRequest(c, &req) {
		return
	}

	user, err := api.userService.Register(req.Username, req.Password)
	if AbortIfError(c, err) {
		return
	}

	token, err := GenerateJWT(user.ID, user.Username)
	if AbortIfError(c, err) {
		return
	}

	RespondData(c, token)
}

// Login handles user authentication
func (api *AccountAPI) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if !ParseRequest(c, &req) {
		return
	}

	user, err := api.userService.Login(req.Username, req.Password)
	if AbortIfError(c, err) {
		return
	}

	token, err := GenerateJWT(user.ID, user.Username)
	if AbortIfError(c, err) {
		return
	}

	RespondData(c, token)
}

// UpdatePassword allows users to change their password
func (api *AccountAPI) UpdatePassword(c *gin.Context) {
	var req struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	if !ParseRequest(c, &req) {
		return
	}

	id := c.MustGet("userId").(int)
	if AbortIfError(c, api.userService.UpdatePassword(id, req.OldPassword, req.NewPassword)) {
		return
	}

	RespondMessage(c, "Password updated successfully")
}
