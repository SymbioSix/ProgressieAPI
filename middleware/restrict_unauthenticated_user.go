package middleware

import (
	s "github.com/SymbioSix/ProgressieAPI/setup"
	"github.com/gofiber/fiber/v3"
)

func RestrictUnauthenticatedUser() fiber.Handler {
	return func(c fiber.Ctx) error {
		_, err := s.Client.Auth.GetUser()
		if err != nil {
			return c.Redirect().Status(fiber.StatusUnauthorized).To("/v1/unauthorized")
		}
		return c.Next()
	}
}
