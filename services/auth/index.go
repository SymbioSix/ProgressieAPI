package auth

import (
	models "github.com/SymbioSix/ProgressieAPI/models/auth"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/supabase-community/gotrue-go/types"
	"gorm.io/gorm"
)

type AuthController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewAuthController(DB *gorm.DB, API *utils.Client) AuthController {
	return AuthController{DB, API}
}

func (au *AuthController) VerifySignUp(c fiber.Ctx) error {
	// TODO: Find out how to invoke send email confirmation signup!
	confirm := new(models.ConfirmationSignup)
	if err := c.Bind().Query(confirm); err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	_, err := au.API.Auth.Verify(types.VerifyRequest{Type: types.VerificationType(confirm.Type), Token: confirm.TokenHash, RedirectTo: confirm.RedirectTo})
	if err != nil {
		c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		return c.Redirect().Status(fiber.StatusNotAcceptable).To("/v1/auth-failed")
	}
	return c.Redirect().Status(fiber.StatusOK).To(confirm.RedirectTo)
}

func (au *AuthController) SignUpWithEmailPassword(c fiber.Ctx) error {
	var payload *models.SignUpRequest

	if err := c.Bind().JSON(&payload); err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	data := make(map[string]interface{})
	data["username"] = payload.Username

	result, err := au.API.Auth.Signup(types.SignupRequest{Email: payload.Email, Password: payload.Password, Data: data})
	if err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.JSON(fiber.Map{"status": "success", "result": result})
}

func (au *AuthController) SignInWithEmailPassword(c fiber.Ctx) error {
	//TODO: Do Something
	return c.JSON("")
}

func (au *AuthController) SignOut(c fiber.Ctx) error {
	//TODO: Do Something
	return nil
}
