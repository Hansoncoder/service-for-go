package CustomError

import "fmt"

type ErrorCode int

const (
	// 通用状态码
	CodeSuccess ErrorCode = 200
	CodeError   ErrorCode = 500

	// 业务错误码（从 1000 开始，避免和 HTTP 状态码冲突）
	CodeInvalidParams    ErrorCode = 1000 + iota // 参数错误
	CodeUserNotFound                             // 用户不存在
	CodeAuthFailed                               // 认证失败
	CodeTokenExpired                             // Token 过期
	CodePermissionDenied                         // 权限不足
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewError(code ErrorCode, message string) error {
	return &CustomError{Code: code, Message: message}
}
