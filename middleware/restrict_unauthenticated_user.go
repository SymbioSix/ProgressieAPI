package middleware

import (
	s "github.com/SymbioSix/ProgressieAPI/setup"
	"github.com/gofiber/fiber/v3"
)

func RestrictUnauthenticatedUser() fiber.Handler {
	return func(c fiber.Ctx) error {
		_, err := s.Client.Auth.GetUser()
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "unauthorized", "message": "Request performed unauthorized"})
		}
		return c.Next()
	}
}
