package handler

import (
	"github.com/gofiber/fiber/v2"
)

// RootHandler godoc
// @Summary Root Endpoint
// @Description This root route returns a simple JSON response
// @Tags Root
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /v1 [get]
func GetRoot(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, Go Ride!",
		"status":  "success",
		"data":    nil,
	})
}

func RootHandler(route fiber.Router) {
	route.Get("/", GetRoot)
}
