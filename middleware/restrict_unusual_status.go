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
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "unauthorized", "message": "Request performed unauthorized"})
		}

		var user *models.UserModel
		if result := s.DB.Table("usr_user").First(&user, "user_id = ?", res.ID); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
		}

		if strings.Contains(strings.ToLower(user.Status), "banned") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": "Request performed by banned user"})
		} else if strings.Contains(strings.ToLower(user.Status), "deactivated") || strings.Contains(strings.ToLower(user.Status), "inactive") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": "Request performed by inactive/deactivated user"})
		} else if strings.Contains(strings.ToLower(user.Status), "locked") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "forbidden", "message": "Request performed by locked user"})
		} else {
			return c.Next()
		}
	}
}
