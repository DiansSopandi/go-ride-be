package middlewares

import (
	"fmt"
	"strings"

	"github.com/DiansSopandi/goride_be/errors"
	"github.com/DiansSopandi/goride_be/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthGuard(c *fiber.Ctx) error {
	var tokenString string

	publicRoutes := []string{
		"/",
		"/v1",
		"/swagger/*",
		"/swagger/index.html",
		"/swagger/doc.json",
		"/v1/auth/login",
		"/v1/auth/register",
	}

	for _, r := range publicRoutes {
		if c.Path() == r || strings.HasPrefix(c.Path(), "/swagger/") {
			return c.Next()
		}
	}

	// Ambil token dari Authorization header
	authHeader := c.Get("Authorization")
	// if authHeader == "" {
	// 	return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	// }
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	}

	if tokenString == "" {
		tokenString = c.Cookies("jwt_at")
	}

	if tokenString == "" {
		// return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
		return errors.Unauthorized("unauthorized access")
	}

	// tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	// if tokenString == authHeader {
	// 	return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
	// }

	// Parse & validasi JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma sesuai
		// SigningMethodHS256
		// if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			return nil, errors.Unauthorized("Unexpected signing method")
		}
		return []byte(pkg.Cfg.Application.JwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return errors.Unauthorized(fmt.Sprintf("Invalid token: %v", err))
	}

	// if !token.Valid {
	// 	return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	// }

	// Simpan user info ke context (opsional)
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user", claims)

	return c.Next()
}
