package handler

import (
	"database/sql"
	"fmt"

	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/errors"
	"github.com/DiansSopandi/goride_be/middlewares"
	"github.com/DiansSopandi/goride_be/pkg"
	helper "github.com/DiansSopandi/goride_be/pkg/helper"
	"github.com/DiansSopandi/goride_be/repository"
	service "github.com/DiansSopandi/goride_be/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService *service.UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func AuthRoutes(route fiber.Router) {
	handler := NewAuthHandler()
	route.Post("/register", middlewares.WithTransaction(RegisterUserHandler(handler)))
	route.Post("/login", middlewares.WithTransaction(LoginUserHandler(handler)))
}

func RegisterUserHandler(handler *AuthHandler) fiber.Handler {
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
			// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			return err
		}

		return pkg.ResponseApiOK(c, "User registered successfully...", res)
	}
}

func LoginUserHandler(handler *AuthHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var loginDto dto.UserLoginRequest
		if err := c.BodyParser(&loginDto); err != nil {
			return pkg.ResponseApiErrorBadRequest(c, fmt.Sprintf("Failed to parse request body: %v", err))
		}

		res, err := handler.LoginUser(c, loginDto)
		if err != nil {
			// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			return err
		}

		return pkg.ResponseApiOK(c, "User logged in successfully...", res)
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
func (h *AuthHandler) RegisterUser(c *fiber.Ctx, regDto dto.UserRegisterRequest) (dto.UserResponse, error) {
	tx := c.Locals(middlewares.TxContextKey).(*sql.Tx)
	userRepo, _ := repository.NewUserRepository(tx)
	roleRepo, _ := repository.NewRoleRepository(tx)

	userServiceWithTx := service.NewUserService(userRepo, roleRepo)

	var (
		roleIDs []int64
		err     error
	)
	if len(regDto.Roles) > 0 {
		roleIDs, err = userServiceWithTx.ValidateRolesExist(tx, regDto.Roles)

		if err != nil {
			return dto.UserResponse{}, errors.RoleValidationFailed(fmt.Sprintf("role validation failed: %v", err))
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
		return dto.UserResponse{}, err
	}

	if len(roleIDs) > 0 {
		err = userServiceWithTx.AssignRolesToUserWithTx(tx, uint(res.ID), roleIDs)
		if err != nil {
			return dto.UserResponse{}, errors.InternalError(fmt.Sprintf("failed to assign roles to user: %v", err))
		}
	}

	return dto.UserResponse{
		ID:       uint(res.ID),
		Username: &registerDto.Username,
		Email:    res.Email,
		Roles:    registerDto.Roles,
	}, nil
}

// LoginUser handles user login and returns user details and token.
// @summary Login a user
// @description Login a user and return user details and token.
// @tags Auth
// @accept json
// @produce json
// @param loginDto body dto.UserLoginRequest true "User login data"
// @success 200 {object} dto.UserLoginResponse "User login successful"
// @failure 400 {object} map[string]interface{} "Bad request, validation errors"
// @failure 500 {object} map[string]interface{} "Internal server error, database or service errors"
// @router /v1/auth/login [post]
func (h *AuthHandler) LoginUser(c *fiber.Ctx, loginDto dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	tx := c.Locals(middlewares.TxContextKey).(*sql.Tx)
	userRepo, _ := repository.NewUserRepository(tx)
	roleRepo, _ := repository.NewRoleRepository(tx)

	userServiceWithTx := service.NewUserService(userRepo, roleRepo)

	res, err := userServiceWithTx.LoginUser(tx, loginDto)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return res, nil
}
