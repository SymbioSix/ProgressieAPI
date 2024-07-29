package middleware

import (
	"strings"

	models "github.com/SymbioSix/ProgressieAPI/models/auth"
	s "github.com/SymbioSix/ProgressieAPI/setup"
	"github.com/gofiber/fiber/v3"
)

func RestrictUserWithUnusualStatus() fiber.Handler {
	return func(c fiber.Ctx) error {
		res, err := s.Client.Auth.GetUser()
		if err != nil {
			return c.Redirect().Status(fiber.StatusUnauthorized).To("/v1/unauthorized")
		}

		var user *models.UserModel
		if result := s.DB.Table("usr_user").First(&user, "user_id = ?", res.ID); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
		}

		if strings.Contains(strings.ToLower(user.Status), "banned") {
			return c.Redirect().Status(fiber.StatusForbidden).To("/v1/banned")
		} else if strings.Contains(strings.ToLower(user.Status), "deactivated") || strings.Contains(strings.ToLower(user.Status), "inactive") {
			return c.Redirect().Status(fiber.StatusForbidden).To("/v1/inactive")
		} else if strings.Contains(strings.ToLower(user.Status), "locked") {
			return c.Redirect().Status(fiber.StatusForbidden).To("/v1/locked")
		} else {
			return c.Next()
		}
	}
}
