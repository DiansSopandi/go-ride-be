package handler

import (
	"database/sql"
	"fmt"

	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/errors"
	"github.com/DiansSopandi/goride_be/middlewares"
	model "github.com/DiansSopandi/goride_be/models"
	"github.com/DiansSopandi/goride_be/pkg"
	helper "github.com/DiansSopandi/goride_be/pkg/helper"
	"github.com/DiansSopandi/goride_be/repository"
	service "github.com/DiansSopandi/goride_be/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler() *UserHandler {
	var tx *sql.Tx
	// userRepo, _ := repository.NewUserRepository(tx, false)
	userRepo, _ := repository.NewUserRepository(tx)
	roleRepo, _ := repository.NewRoleRepository(tx)
	userService := service.NewUserService(userRepo, roleRepo)

	return &UserHandler{
		UserService: userService, // service.NewUserService(),
	}
}

func UserRoutes(route fiber.Router) {
	handler := NewUserHandler()

	route.Get("/users", GetUserHandler(handler))
	route.Post("/users", middlewares.WithTransaction(CreateUserHandler(handler)))
}

func CreateUserHandler(handler *UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createUserDto dto.UserCreateRequest

		if err := c.BodyParser(&createUserDto); err != nil {
			return pkg.ResponseApiErrorBadRequest(c, fmt.Sprintf("Failed to parse request body: %v", err))
		}

		if err := helper.ValidateCreateUserRequest(&createUserDto); err != nil {
			return pkg.ResponseApiErrorBadRequest(c, fmt.Sprintf("Validation failed: %v", err))
		}

		res, err := handler.CreateUser(c, &createUserDto)

		if err != nil {
			// return pkg.ResponseApiErrorInternalServer(c, fmt.Sprintf("Failed to create user: %v", err))
			// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			return err
			// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// 	"message": err.Error(),
			// 	"status":  fiber.StatusInternalServerError,
			// 	"data":    nil,
			// })
		}

		return pkg.ResponseApiOK(c, "User created successfully", res)
	}
}

func GetUserHandler(handler *UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := handler.GetUser()

		if err != nil {
			return errors.InternalError(fmt.Sprintf("Failed to fetch users: %v", err))
			// return pkg.ResponseApiErrorInternalServer(c, fmt.Sprintf("Failed to fetch users: %v", err))
		}

		return pkg.ResponseApiOK(c, "User fetch successfully...", res)
	}
}

// PostUser godoc
// @Summary User Endpoint
// @Description This user route returns a simple JSON response
// @Tags User
// @Produce json
// @Param createUserDto body dto.UserCreateRequest true "Create User Request"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /v1/users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx, createUserDto *dto.UserCreateRequest) (model.User, error) {
	// Ambil TX dari context
	tx := c.Locals(middlewares.TxContextKey).(*sql.Tx)
	// isTx := tx != nil
	// userRepo, err := repository.NewUserRepository(tx, isTx)
	userRepo, _ := repository.NewUserRepository(tx)
	roleRepo, _ := repository.NewRoleRepository(tx)

	userServiceWithTx := service.NewUserService(userRepo, roleRepo)

	// userJson, err := json.MarshalIndent(createUserDto, "", "  ")
	// fmt.Println(string(userJson))

	var (
		roleIDs []int64
		err     error
	)
	if len(createUserDto.Roles) > 0 {
		// roleIDs, err = h.UserService.ValidateRolesExist(tx, createUserDto.Roles)
		roleIDs, err = userServiceWithTx.ValidateRolesExist(tx, createUserDto.Roles)
		// roleIDs, err = svc.ValidateRolesExist(createUserDto.Roles)

		if err != nil {
			return model.User{}, errors.RoleValidationFailed(fmt.Sprintf("role validation failed: %v", err))
		}
	}

	// Ambil service dari context (TX sudah aktif) di middleware
	// svc := c.Locals(middlewares.UserServiceCtxKey).(*service.UserService)

	// res, err := h.UserService.CreateUser(tx, createUserDto)
	res, err := userServiceWithTx.CreateUser(tx, createUserDto)
	// res, err := svc.CreateUser(createUserDto)
	if err != nil {
		return model.User{}, err
	}

	if len(roleIDs) > 0 {
		// err = h.UserService.AssignRolesToUserWithTx(tx, uint(res.ID), roleIDs)
		err = userServiceWithTx.AssignRolesToUserWithTx(tx, uint(res.ID), roleIDs)
		// err = svc.AssignRolesToUserWithTx(uint(res.ID), roleIDs)
		if err != nil {
			return model.User{}, errors.InternalError(fmt.Sprintf("failed to assign roles: %v", err))
		}
	}

	return res, nil
}

// UserHandler godoc
// @Summary User Endpoint
// @Description This user route returns a simple JSON response
// @Tags User
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /v1/users [get]
func (h *UserHandler) GetUser() ([]dto.UserResponse, error) {

	res, err := h.UserService.GetAllUsers()

	if err != nil {
		return []dto.UserResponse{}, errors.InternalError(fmt.Sprintf("Failed to fetch users: %v", err))
	}

	return res, nil
}
