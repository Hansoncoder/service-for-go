package errors

import (
	"veo/internal/utils"
)

// ErrorCode represents the type of error
type ErrorCode int

const (
	// 通用状态码
	CodeSuccess ErrorCode = 200
	CodeError   ErrorCode = 500

	// 业务错误码（从 1000 开始，避免和 HTTP 状态码冲突）
	CodeInvalidParams    ErrorCode = 1000 + iota // 参数错误
	CodeUserNotFound                             // 用户不存在
	CodeUserExists                               // 用户已注册
	CodeAuthFailed                               // 认证失败
	CodeTokenExpired                             // Token 过期
	CodePermissionDenied                         // 权限不足
)

var logger = utils.GetLogger()

// Error represents a custom error with code and message
type Error struct {
	Code    ErrorCode
	Message string
}

// Error implements the error interface
func (e *Error) Error() string {
	return e.Message
}

// Error implements the error interface
func (e *Error) GetCode() int {
	return int(e.Code)
}

// New creates a new error with the given code and message
func New(code ErrorCode, message string) error {
	logger.Error(message)
	return &Error{Code: code, Message: message}
}

// NewInvalidParams creates a new invalid parameters error
func NewInvalidParams(message string) error {
	return New(CodeInvalidParams, message)
}

// NewUserNotFound creates a new user not found error
func NewUserNotFound(message string) error {
	return New(CodeUserNotFound, message)
}

// NewAuthFailed creates a new authentication failed error
func NewAuthFailed(message string) error {
	return New(CodeAuthFailed, message)
}

// NewTokenExpired creates a new token expired error
func NewTokenExpired(message string) error {
	return New(CodeTokenExpired, message)
}

// NewPermissionDenied creates a new permission denied error
func NewPermissionDenied(message string) error {
	return New(CodePermissionDenied, message)
}

// NewUserExists creates a new permission denied error
func NewUserExists(message string) error {
	return New(CodeUserExists, message)
}
