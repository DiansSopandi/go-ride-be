package pkg

import "github.com/gofiber/fiber/v2"

type HttpStatusCode int

const (
	HttpStatusOK                  HttpStatusCode = 200
	HttpStatusCreated             HttpStatusCode = 201
	HttpStatusNoContent           HttpStatusCode = 204
	HttpStatusBadRequest          HttpStatusCode = 400
	HttpStatusUnauthorized        HttpStatusCode = 401
	HttpStatusForbidden           HttpStatusCode = 403
	HttpStatusNotFound            HttpStatusCode = 404
	HttpStatusInternalServerError HttpStatusCode = 500
	HttpStatusUnprocessableEntity HttpStatusCode = 422
	HttpStatusTooManyRequests     HttpStatusCode = 429
	HttpStatusGone                HttpStatusCode = 410
	HttpStatusServiceUnavailable  HttpStatusCode = 503
)

type ApiStatusOK string

const (
	ApiStatusSuccessOk      ApiStatusOK = "success_ok"
	ApiStatusSuccessCreated ApiStatusOK = "success_created"
	ApiStatusSuccessUpdated ApiStatusOK = "success_updated"
	ApiStatusSuccessDeleted ApiStatusOK = "success_deleted"
)

type ApiStatusError string

const (
	ApiStatusErrorBadRequest          ApiStatusError = "error_bad_request"
	ApiStatusErrorUnauthorized        ApiStatusError = "error_unauthorized"
	ApiStatusErrorForbidden           ApiStatusError = "error_forbidden"
	ApiStatusErrorNotFound            ApiStatusError = "error_not_found"
	ApiStatusErrorInternalServerError ApiStatusError = "error_internal_server_error"
	ApiStatusErrorUnprocessableEntity ApiStatusError = "error_unprocessable_entity"
	ApiStatusErrorTooManyRequests     ApiStatusError = "error_too_many_requests"
	ApiStatusErrorGone                ApiStatusError = "error_gone"
	ApiStatusErrorServiceUnavailable  ApiStatusError = "error_service_unavailable"
	ApiErrorUnprocessAble             ApiStatusError = "error_unprocessable"
	ApiErrorLimitReached              ApiStatusError = "error_limit_reached"
)

type DetailResponse struct {
	Path       string `json:"path" example:"/api/v1/path"`
	Param      string `json:"param" example:"?page=1&limit=10"`
	StatusCode int    `json:"status_code" example:"200"`
	Method     string `json:"method" example:"GET"`
	Status     string `json:"status" example:"success_ok"`
} // @name	DetailResponse

type ResponseApi struct {
	Details DetailResponse `json:"details"`
	Success bool           `json:"success" example:"true"`
	Data    any            `json:"data" swaggertype:"array,object"`
	Errors  any            `json:"errors" swaggertype:"array,object"`
	Message interface{}    `json:"message" example:"API Message"`
} // @name ResponseApi

type ResponseApiError struct {
	// HttpStatusCode *HttpStatusCode `json:"http_status_code,omitempty" example:"400"`
	HttpStatusCode HttpStatusCode `json:"http_status_code" example:"400"`
	Status         ApiStatusError `json:"status"`
	Message        interface{}    `json:"message"`
}

func ResponseApiWrapper(ctx *fiber.Ctx, msg string, status string, statusCode int, data any, errors any) error {
	details := DetailResponse{
		StatusCode: statusCode,
		Path:       ctx.Request().URI().String(),
		Method:     string(ctx.Request().Header.Method()),
		Status:     status,
	}

	if statusCode >= 400 {
		CreateAccessLog(ctx, "[ACCESS:API][ERROR]", statusCode, errors)

		return ctx.Status(statusCode).JSON(ResponseApi{
			Success: false,
			Message: msg,
			Data:    data,
			Errors:  errors,
			Details: details,
		})
	}

	CreateAccessLog(ctx, "[ACCESS:API][SUCCESS]", statusCode, data)

	return ctx.Status(statusCode).JSON(ResponseApi{
		Success: true,
		Message: msg,
		Data:    data,
		Errors:  errors,
		Details: details,
	})
}

func ResponseApiOK(ctx *fiber.Ctx, msg string, data any) error {
	return ResponseApiWrapper(ctx, msg, string(ApiStatusSuccessOk), int(HttpStatusOK), data, nil)
}

func ResponseApiCreated(ctx *fiber.Ctx, msg string, data any) error {
	return ResponseApiWrapper(ctx, msg, string(ApiStatusSuccessCreated), int(HttpStatusCreated), data, nil)
}

func ResponseApiUpdated(ctx *fiber.Ctx, msg string, data any) error {
	return ResponseApiWrapper(ctx, msg, string(ApiStatusSuccessUpdated), int(HttpStatusOK), data, nil)
}

func ResponseApiDeleted(ctx *fiber.Ctx, msg string) error {
	return ResponseApiWrapper(ctx, msg, string(ApiStatusSuccessDeleted), int(HttpStatusNoContent), nil, nil)
}

func ResponseApiErrorWrapper(ctx *fiber.Ctx, status ApiStatusError, statusCode HttpStatusCode, message interface{}) error {
	details := DetailResponse{
		StatusCode: int(statusCode),
		Path:       ctx.Request().URI().String(),
		Method:     string(ctx.Request().Header.Method()),
		Status:     string(status),
	}

	CreateAccessLog(ctx, "[ACCESS:API][ERROR]", int(statusCode), message)

	return ctx.Status(int(statusCode)).JSON(ResponseApi{
		Success: false,
		Message: message,
		Errors:  nil,
		Data:    nil,
		Details: details,
	})
}
func ResponseApiErrorCustom(ctx *fiber.Ctx, status ApiStatusError, statusCode HttpStatusCode, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, status, statusCode, message)
}
func ResponseApiErrorBadRequest(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorBadRequest, HttpStatusBadRequest, message)
}

func ResponseApiErrorUnauthorized(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorUnauthorized, HttpStatusUnauthorized, message)
}

func ResponseApiErrorForbidden(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorForbidden, HttpStatusForbidden, message)
}

func ResponseApiErrorNotFound(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorNotFound, HttpStatusNotFound, message)
}

func ResponseApiErrorInternalServer(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorInternalServerError, HttpStatusInternalServerError, message)
}

func ResponseApiErrorUnprocessableEntity(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorUnprocessableEntity, HttpStatusUnprocessableEntity, message)
}

func ResponseApiErrorTooManyRequests(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorTooManyRequests, HttpStatusTooManyRequests, message)
}

func ResponseApiErrorGone(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorGone, HttpStatusGone, message)
}

func ResponseApiErrorServiceUnavailable(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorServiceUnavailable, HttpStatusServiceUnavailable, message)
}

func ResponseApiErrorUnprocessAble(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiErrorUnprocessAble, HttpStatusUnprocessableEntity, message)
}

func ResponseApiErrorLimitReached(ctx *fiber.Ctx, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiErrorLimitReached, HttpStatusTooManyRequests, message)
}

func ResponseApiErrorWithStatus(ctx *fiber.Ctx, status ApiStatusError, message interface{}) error {
	switch status {
	case ApiStatusErrorBadRequest:
		return ResponseApiErrorBadRequest(ctx, message)
	case ApiStatusErrorUnauthorized:
		return ResponseApiErrorUnauthorized(ctx, message)
	case ApiStatusErrorForbidden:
		return ResponseApiErrorForbidden(ctx, message)
	case ApiStatusErrorNotFound:
		return ResponseApiErrorNotFound(ctx, message)
	case ApiStatusErrorInternalServerError:
		return ResponseApiErrorInternalServer(ctx, message)
	case ApiStatusErrorUnprocessableEntity:
		return ResponseApiErrorUnprocessableEntity(ctx, message)
	case ApiStatusErrorTooManyRequests:
		return ResponseApiErrorTooManyRequests(ctx, message)
	case ApiStatusErrorGone:
		return ResponseApiErrorGone(ctx, message)
	case ApiStatusErrorServiceUnavailable:
		return ResponseApiErrorServiceUnavailable(ctx, message)
	default:
		return ResponseApiErrorCustom(ctx, ApiStatusErrorInternalServerError, HttpStatusInternalServerError, message)
	}
}

func ResponseApiErrorWithStatusCode(ctx *fiber.Ctx, statusCode HttpStatusCode, message interface{}) error {
	switch statusCode {
	case HttpStatusBadRequest:
		return ResponseApiErrorBadRequest(ctx, message)
	case HttpStatusUnauthorized:
		return ResponseApiErrorUnauthorized(ctx, message)
	case HttpStatusForbidden:
		return ResponseApiErrorForbidden(ctx, message)
	case HttpStatusNotFound:
		return ResponseApiErrorNotFound(ctx, message)
	case HttpStatusInternalServerError:
		return ResponseApiErrorInternalServer(ctx, message)
	case HttpStatusUnprocessableEntity:
		return ResponseApiErrorUnprocessableEntity(ctx, message)
	case HttpStatusTooManyRequests:
		return ResponseApiErrorTooManyRequests(ctx, message)
	case HttpStatusGone:
		return ResponseApiErrorGone(ctx, message)
	case HttpStatusServiceUnavailable:
		return ResponseApiErrorServiceUnavailable(ctx, message)
	default:
		return ResponseApiErrorCustom(ctx, ApiStatusErrorInternalServerError, statusCode, message)
	}
}

func ResponseApiErrorWithStatusAndCode(ctx *fiber.Ctx, status ApiStatusError, statusCode HttpStatusCode, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, status, statusCode, message)
}
func ResponseApiErrorWithStatusAndMessage(ctx *fiber.Ctx, status ApiStatusError, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, status, HttpStatusInternalServerError, message)
}

func ResponseApiErrorWithMessage(ctx *fiber.Ctx, statusCode HttpStatusCode, message interface{}) error {
	return ResponseApiErrorWrapper(ctx, ApiStatusErrorInternalServerError, statusCode, message)
}

func ResponseApiErrorWithStatusAndData(ctx *fiber.Ctx, status ApiStatusError, data any) error {
	details := DetailResponse{
		StatusCode: int(HttpStatusInternalServerError),
		Path:       ctx.Request().URI().String(),
		Method:     string(ctx.Request().Header.Method()),
		Status:     string(status),
	}

	return ctx.Status(int(HttpStatusInternalServerError)).JSON(ResponseApi{
		Success: false,
		Message: "",
		Errors:  nil,
		Data:    nil,
		Details: details,
	})
}
