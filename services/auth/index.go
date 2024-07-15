package auth

import (
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

func (au *AuthController) SignInWithEmailPassword(c fiber.Ctx) error {
	//TODO: Do Something
	return c.JSON("")
}

func (au *AuthController) SignUpWithEmailPassword(c fiber.Ctx) error {
	//TODO: Do Something
	return c.JSON("")
}

func (au *AuthController) SignInWithGoogle(c fiber.Ctx) error {
	//TODO: Do Something
	return c.JSON("")
}

func (au *AuthController) SignUpWithGoogle(c fiber.Ctx) error {
	//TODO: Do Something
	return c.JSON("")
}

func (au *AuthController) SignOut(c fiber.Ctx) error {
	//TODO: Do Something
	return nil
}
