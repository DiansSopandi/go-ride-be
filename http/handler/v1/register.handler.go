package handler

import (
	"database/sql"
	"fmt"

	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/middlewares"
	model "github.com/DiansSopandi/goride_be/models"
	"github.com/DiansSopandi/goride_be/pkg"
	helper "github.com/DiansSopandi/goride_be/pkg/helper"
	"github.com/DiansSopandi/goride_be/repository"
	service "github.com/DiansSopandi/goride_be/services"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	RegisterService *service.UserService
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{}
}

func RegisterRoutes(route fiber.Router) {
	handler := NewRegisterHandler()
	route.Post("/register", middlewares.WithTransaction(RegisterUserHandler(handler)))
}

func RegisterUserHandler(handler *RegisterHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var registerDto dto.UserRegisterRequest
		if errDto := c.BodyParser(&registerDto); errDto != nil {
			return pkg.ResponseApiErrorBadRequest(c, fmt.Sprintf("Failed to parse request body: %v", errDto))
		}

		if err := helper.ValidateRegisterUserRequest(&registerDto); err != nil {
			return pkg.ResponseApiErrorBadRequest(c, fmt.Sprintf("Validation failed: %v", err))
		}

		res, err := handler.RegisterUser(c, registerDto)
		if err != nil {
			// return pkg.ResponseApiError(c, fiber.StatusInternalServerError, "Failed to register user", err)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return pkg.ResponseApiOK(c, "User registered successfully", res)
	}
}

// RegisterUser handles user registration, including role assignment.
// @summary Register a new user with roles
// @description Register a new user and assign roles if provided.
// @tags Auth
// @accept json
// @produce json
// @param registerDto body dto.UserRegisterRequest true "User registration data"
// @success 200 {object} model.User "User registration successful"
// @failure 400 {object} map[string]interface{} "Bad request, validation errors"
// @failure 500 {object} map[string]interface{} "Internal server error, database or service errors"
// @router /v1/auth/register [post]
func (h *RegisterHandler) RegisterUser(c *fiber.Ctx, regDto dto.UserRegisterRequest) (model.User, error) {
	tx := c.Locals(middlewares.TxContextKey).(*sql.Tx)
	userRepo, err := repository.NewUserRepository(tx)
	if err != nil {
		return model.User{}, err
	}
	userServiceWithTx := service.NewUserService(userRepo)

	var roleIDs []int64
	if len(regDto.Roles) > 0 {
		roleIDs, err = userServiceWithTx.ValidateRolesExist(tx, regDto.Roles)

		if err != nil {
			return model.User{}, fmt.Errorf("role validation failed: %w", err)
		}
	}

	registerDto := dto.UserCreateRequest{
		Username: regDto.Username,
		Email:    regDto.Email,
		Password: regDto.Password,
		Roles:    []string{"user"},
	}

	res, err := userServiceWithTx.CreateUser(tx, &registerDto)
	if err != nil {
		return model.User{}, err
	}

	if len(roleIDs) > 0 {
		err = userServiceWithTx.AssignRolesToUserWithTx(tx, uint(res.ID), roleIDs)
		if err != nil {
			return model.User{}, err
		}
	}

	return res, nil
}
