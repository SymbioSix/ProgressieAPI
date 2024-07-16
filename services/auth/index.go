package auth

import (
	models "github.com/SymbioSix/ProgressieAPI/models/auth"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AuthController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewAuthController(DB *gorm.DB, API *utils.Client) AuthController {
	return AuthController{DB, API}
}

func (au *AuthController) SignUpWithEmailPassword(c fiber.Ctx) error {
	var payload *models.SignUpRequest

	if err := c.Bind().JSON(&payload); err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	data := make(map[string]interface{})
	data["username"] = payload.Username
	//TODO: Implement PKCE Flow
	// result, err := au.API.Auth.Signup(types.SignupRequest{Email: payload.Email, Password: payload.Password, Data: data})
	return c.JSON("")
}

func (au *AuthController) SignInWithEmailPassword(c fiber.Ctx) error {
	//TODO: Do Something
	return c.JSON("")
}

func (au *AuthController) SignOut(c fiber.Ctx) error {
	//TODO: Do Something
	return nil
}
