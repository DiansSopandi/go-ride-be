package handler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/middlewares"
	model "github.com/DiansSopandi/goride_be/models"
	helper "github.com/DiansSopandi/goride_be/pkg/helper"
	"github.com/DiansSopandi/goride_be/pkg/utils"
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
	userService := service.NewUserService(userRepo)
	return &UserHandler{
		UserService: userService, // service.NewUserService(),
	}
}

func UserRoutes(route fiber.Router) {
	handler := NewUserHandler()

	route.Post("/users", middlewares.WithTransaction(CreateUserHandler(handler)))
	route.Get("/users", GetUserHandler(handler))
}

func CreateUserHandler(handler *UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createUserDto dto.UserCreateRequest

		if err := c.BodyParser(&createUserDto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
				"status":  fiber.StatusBadRequest,
				"data":    nil,
			})
		}

		if err := helper.ValidateCreateUserRequest(&createUserDto); err != nil {
			log.Printf("Validation error: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Validation failed",
				"status":  fiber.StatusBadRequest,
				"data":    nil,
				"error":   err.Error(),
			})
		}

		res, err := handler.CreateUser(c, &createUserDto)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// 	"message": err.Error(),
			// 	"status":  fiber.StatusInternalServerError,
			// 	"data":    nil,
			// })
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User created...",
			"status":  fiber.StatusOK,
			"success": true,
			"data":    res,
		})
	}
}

func GetUserHandler(handler *UserHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := handler.GetUser()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"status":  fiber.StatusInternalServerError,
				"success": false,
				"data":    nil,
			})
		}
		return c.JSON(fiber.Map{
			"message": "User fetch successfuly...",
			"status":  fiber.StatusOK,
			"success": true,
			"data":    res,
		})
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
	userRepo, err := repository.NewUserRepository(tx)
	if err != nil {
		return model.User{}, err
	}

	// userJson, err := json.MarshalIndent(createUserDto, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(userJson))

	password, _ := utils.HashPassword(createUserDto.Password)
	userDto := model.User{
		Username: createUserDto.Username,
		Email:    createUserDto.Email,
		Password: password,
	}

	// tx, err := h.UserService.BeginTransaction() // sudah dihandle di middleware
	// if err != nil {
	// 	return model.User{}, fmt.Errorf("failed to begin transaction: %w", err)
	// }

	// defer func() {
	// 	if p := recover(); p != nil {
	// 		tx.Rollback()
	// 		panic(p)
	// 	} else if err != nil {
	// 		tx.Rollback()
	// 	}
	// }()

	// defer db.RollbackOnError(tx, err) // sudah dihandle di middleware
	userServiceWithTx := service.NewUserService(userRepo)
	// Ambil service dari context (TX sudah aktif) di middleware
	// svc := c.Locals(middlewares.UserServiceCtxKey).(*service.UserService)

	var roleIDs []int64
	if len(createUserDto.Roles) > 0 {
		// roleIDs, err = h.UserService.ValidateRolesExist(tx, createUserDto.Roles)
		roleIDs, err = userServiceWithTx.ValidateRolesExist(tx, createUserDto.Roles)
		// roleIDs, err = svc.ValidateRolesExist(createUserDto.Roles)

		if err != nil {
			return model.User{}, fmt.Errorf("role validation failed: %w", err)
		}
	}

	// res, err := h.UserService.CreateUser(tx, &userDto)
	res, err := userServiceWithTx.CreateUser(tx, &userDto)
	// res, err := svc.CreateUser(&userDto)
	if err != nil {
		return model.User{}, err
	}

	if len(roleIDs) > 0 {
		// err = h.UserService.AssignRolesToUserWithTx(tx, uint(res.ID), roleIDs)
		err = userServiceWithTx.AssignRolesToUserWithTx(tx, uint(res.ID), roleIDs)
		// err = svc.AssignRolesToUserWithTx(uint(res.ID), roleIDs)
		if err != nil {
			return model.User{}, err
		}
	}

	// Commit transaction
	// err = tx.Commit() // sudah  dihandle di middleware
	// if err != nil {
	// 	return model.User{}, fmt.Errorf("failed to commit transaction: %w", err)
	// }

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
		return []dto.UserResponse{}, err
	}

	return res, nil
}
