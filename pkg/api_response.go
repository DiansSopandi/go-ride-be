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
	Valid   bool           `json:"valid" example:"true"`
	Data    any            `json:"data" swaggertype:"array,object"`
	Errors  any            `json:"errors" swaggertype:"array,object"`
	Message string         `json:"message" example:"API Message"`
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
			Valid:   false,
			Message: msg,
			Data:    data,
			Errors:  errors,
			Details: details,
		})
	}

	CreateAccessLog(ctx, "[ACCESS:API][SUCCESS]", statusCode, data)

	return ctx.Status(statusCode).JSON(ResponseApi{
		Valid:   true,
		Message: msg,
		Data:    data,
		Errors:  errors,
		Details: details,
	})
}
