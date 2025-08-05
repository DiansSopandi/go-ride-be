package handler

import (
	"database/sql"

	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/middlewares"
	"github.com/DiansSopandi/goride_be/models"
	model "github.com/DiansSopandi/goride_be/models"
	"github.com/DiansSopandi/goride_be/repository"
	service "github.com/DiansSopandi/goride_be/services"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	service *service.RoleService
}

func NewRoleHandler() *RoleHandler {
	var tx *sql.Tx
	// roleRepo, _ := repository.NewRoleRepository(false)
	roleRepo, _ := repository.NewRoleRepository(tx)
	roleService := service.NewRoleService(roleRepo)
	return &RoleHandler{
		service: roleService,
	}
}

func RolesRoutes(route fiber.Router) {
	handler := NewRoleHandler()
	route.Get("/roles", GetAllRolesHandler(handler))
	route.Post("/roles", middlewares.WithTransaction(CreateRoleHandler(handler)))
}

func GetAllRolesHandler(handler *RoleHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, err := handler.GetAllRoles()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"status":  fiber.StatusInternalServerError,
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Role fetch successfuly...",
			"status":  fiber.StatusOK,
			"success": true,
			"data":    roles,
		})
	}
}

func CreateRoleHandler(handler *RoleHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createRoledto dto.RoleCreateRequest

		// if err := c.BodyParser(&createRoledto); err != nil {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"message": err.Error(),
		// 		"status":  fiber.StatusBadRequest,
		// 		"data":    nil,
		// 	})
		// }
		if err := c.BodyParser(&createRoledto); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		res, err := handler.CreateRole(c, &createRoledto)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Role created successfully",
			"status":  fiber.StatusOK,
			"success": true,
			"data":    res,
		})
	}
}

// GetAllRoles
// @Summary GetAllRoles
// @Description Get all roles
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {array} models.Role
// @Router /v1/roles [get]
func (h *RoleHandler) GetAllRoles() ([]models.Role, error) {
	return h.service.GetAllRoles()
}

// CreateRole
// @Summary CreateRole
// @Description Create a new role
// @Tags Role
// @Accept json
// @Produce json
// @Param roleDto body dto.RoleCreateRequest true "Create Role Request"
// @Success 200 {object} models.Role
// @Router /v1/roles [post]
func (h *RoleHandler) CreateRole(c *fiber.Ctx, roleDto *dto.RoleCreateRequest) (models.Role, error) {
	// Ambil TX dari context
	tx := c.Locals(middlewares.TxContextKey).(*sql.Tx)
	// isTx := tx != nil
	// roleRepo, err := repository.NewRoleRepository(isTx)
	roleRepo, err := repository.NewRoleRepository(tx)
	if err != nil {
		return model.Role{}, err
	}

	role := models.Role{
		Name:        roleDto.Name,
		Description: roleDto.Description,
	}

	roleServiceWithTx := service.NewRoleService(roleRepo)

	return roleServiceWithTx.CreateRoles(tx, &role)
}
