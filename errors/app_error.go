package errors

import (
	"net/http"

	"github.com/DiansSopandi/goride_be/pkg"
	"github.com/gofiber/fiber/v2"
)

// type AppError struct {
// 	Code    string `json:"code"`
// 	Message string `json:"message"`
// 	Status  int    `json:"status"`
// }

type DetailResponse struct {
	Path       string `json:"path" example:"/api/v1/path"`
	Param      string `json:"param" example:"?page=1&limit=10"`
	StatusCode int    `json:"status_code" example:"200"`
	Method     string `json:"method" example:"GET"`
	Status     string `json:"status" example:"success_ok"`
} // @name	DetailResponse

type AppErrorResponse struct {
	Details    DetailResponse `json:"details"`
	Success    bool           `json:"success" example:"true"`
	Data       any            `json:"data" swaggertype:"array,object"`
	Errors     any            `json:"errors" swaggertype:"array,object"`
	Message    interface{}    `json:"message" example:"API Message"`
	LogMessage string         `json:"log_message" example:"Log message for internal user"`
} // @name ResponseApi

func (e *AppErrorResponse) Error() string {
	return e.Message.(string)
}

func NewAppErrorResponse(c *fiber.Ctx, code string, statusCode int, logMessage string, status string) *AppErrorResponse {
	message, ok := ErrorCodeMap[code]

	if !ok {
		message = "Unknown error"
	}

	return &AppErrorResponse{
		Details: DetailResponse{
			Path:       c.Path(),
			Param:      string(c.Request().URI().QueryString()),
			StatusCode: statusCode,
			Method:     c.Method(),
			Status:     status,
		},
		Success:    false,
		Data:       nil,
		Errors:     nil,
		Message:    message,
		LogMessage: logMessage,
	}
}

func UserNotFound(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "USER_NOT_FOUND", http.StatusNotFound, logMessage, string(pkg.ApiStatusErrorNotFound))
}

func InvalidCredential(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "INVALID_CREDENTIAL", http.StatusUnauthorized, logMessage, string(pkg.ApiStatusErrorUnauthorized))
}

func DatabaseError(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "DB_ERROR", http.StatusInternalServerError, logMessage, string(pkg.ApiStatusErrorInternalServerError))
}

func InternalError(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "INTERNAL_ERROR", http.StatusInternalServerError, logMessage, string(pkg.ApiStatusErrorInternalServerError))
}

func RoleNotFound(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "ROLE_NOT_FOUND", http.StatusNotFound, logMessage, string(pkg.ApiStatusErrorNotFound))
}

func PermissionDenied(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "PERMISSION_DENIED", http.StatusForbidden, logMessage, string(pkg.ApiStatusErrorForbidden))
}

func InvalidInput(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "INVALID_INPUT", http.StatusBadRequest, logMessage, string(pkg.ApiStatusErrorBadRequest))
}

func EmailAlreadyExists(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "EMAIL_ALREADY_EXISTS", http.StatusConflict, logMessage, string(pkg.ApiStatusErrorConflict))
}

func PhoneAlreadyExists(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "PHONE_ALREADY_EXISTS", http.StatusConflict, logMessage, string(pkg.ApiStatusErrorConflict))
}

func Unauthorized(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "UNAUTHORIZED", http.StatusUnauthorized, logMessage, string(pkg.ApiStatusErrorUnauthorized))
}

func ResourceNotFound(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "RESOURCE_NOT_FOUND", http.StatusNotFound, logMessage, string(pkg.ApiStatusErrorNotFound))
}

func InvalidToken(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "INVALID_TOKEN", http.StatusUnauthorized, logMessage, string(pkg.ApiStatusErrorUnauthorized))
}

func PasswordMismatch(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "PASSWORD_MISMATCH", http.StatusBadRequest, logMessage, string(pkg.ApiStatusErrorBadRequest))
}

func UserAlreadyExists(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "USER_ALREADY_EXISTS", http.StatusConflict, logMessage, string(pkg.ApiStatusErrorConflict))
}

func RoleAlreadyExists(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "ROLE_ALREADY_EXISTS", http.StatusConflict, logMessage, string(pkg.ApiStatusErrorConflict))
}

func EmailNotVerified(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "EMAIL_NOT_VERIFIED", http.StatusForbidden, logMessage, string(pkg.ApiStatusErrorForbidden))
}

func PhoneNotVerified(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "PHONE_NOT_VERIFIED", http.StatusForbidden, logMessage, string(pkg.ApiStatusErrorForbidden))
}

func AccountLocked(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "ACCOUNT_LOCKED", http.StatusForbidden, logMessage, string(pkg.ApiStatusErrorForbidden))
}

func TooManyRequests(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "TOO_MANY_REQUESTS", http.StatusTooManyRequests, logMessage, string(pkg.ApiStatusErrorTooManyRequests))
}

func InvalidFileType(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "INVALID_FILE_TYPE", http.StatusBadRequest, logMessage, string(pkg.ApiStatusErrorBadRequest))
}

func FileTooLarge(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "FILE_TOO_LARGE", http.StatusBadRequest, logMessage, string(pkg.ApiStatusErrorBadRequest))
}

func OperationNotAllowed(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "OPERATION_NOT_ALLOWED", http.StatusMethodNotAllowed, logMessage, string(pkg.ApiStatusErrorMethodNotAllowed))
}

func ResourceConflict(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "RESOURCE_CONFLICT", http.StatusConflict, logMessage, string(pkg.ApiStatusErrorConflict))
}
func InvalidRequest(c *fiber.Ctx, logMessage string) *AppErrorResponse {
	return NewAppErrorResponse(c, "INVALID_REQUEST", http.StatusBadRequest, logMessage, string(pkg.ApiStatusErrorBadRequest))
}
func InvalidRequestWithMessage(c *fiber.Ctx, message string) *AppErrorResponse {
	return &AppErrorResponse{
		Details: DetailResponse{
			Path:       "",
			Param:      "",
			StatusCode: http.StatusBadRequest,
			Method:     "",
			Status:     "",
		},
		Success: false,
		Data:    nil,
		Errors:  nil,
		Message: message,
	}
}

func InvalidRequestWithCode(c *fiber.Ctx, code string) *AppErrorResponse {
	return &AppErrorResponse{
		Details: DetailResponse{
			Path:       "",
			Param:      "",
			StatusCode: http.StatusBadRequest,
			Method:     "",
			Status:     "",
		},
		Success: false,
		Data:    nil,
		Errors:  nil,
		Message: "Invalid request",
	}
}

func InvalidRequestWithCodeAndMessage(c *fiber.Ctx, code, message string) *AppErrorResponse {
	return &AppErrorResponse{
		Details: DetailResponse{
			Path:       "",
			Param:      "",
			StatusCode: http.StatusBadRequest,
			Method:     "",
			Status:     "",
		},
		Success: false,
		Data:    nil,
		Errors:  nil,
		Message: message,
	}
}

func InvalidRequestWithCodeMessageAndStatus(c *fiber.Ctx, code, message string, status int) *AppErrorResponse {
	return &AppErrorResponse{
		Details: DetailResponse{
			Path:       "",
			Param:      "",
			StatusCode: status,
			Method:     "",
			Status:     "",
		},
		Success: false,
		Data:    nil,
		Errors:  nil,
		Message: message,
	}
}
