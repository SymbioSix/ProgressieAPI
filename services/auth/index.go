package auth

import (
	"errors"

	models "github.com/SymbioSix/ProgressieAPI/models/auth"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/supabase-community/gotrue-go/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var confirm *models.ConfirmationSignup
	if err := c.Bind().Query(&confirm); err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	_, err := au.API.Auth.Verify(types.VerifyRequest{Type: types.VerificationType(confirm.Type), Token: confirm.TokenHash, RedirectTo: confirm.RedirectTo})
	if err != nil {
		c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		return c.Redirect().Status(fiber.StatusConflict).To("/v1/auth-failed")
	}
	return c.Redirect().Status(fiber.StatusOK).To(confirm.RedirectTo)
}

func (au *AuthController) SignUpWithEmailPassword(c fiber.Ctx) error {
	var signup *models.SignUpRequest

	if err := c.Bind().JSON(&signup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var user *models.UserModel
	checkUsername := au.DB.Table("usr_user").First(&user, "username = ?", signup.Username)
	if checkUsername.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Username has been taken"})
	} else if !errors.Is(checkUsername.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkUsername.Error.Error()})
	}

	checkEmail := au.DB.Table("usr_user").First(&user, "email = ?", signup.Email)
	if checkEmail.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Email has been used"})
	} else if !errors.Is(checkEmail.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkEmail.Error.Error()})
	}

	data := make(map[string]interface{})
	data["username"] = signup.Username

	result, err := au.API.Auth.Signup(types.SignupRequest{Email: signup.Email, Password: signup.Password, Data: data})
	if err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var role *models.RoleModel
	if fetchRole := au.DB.Table("usr_role").First(&role, "role_name = ?", "BasicUser"); fetchRole.Error != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": fetchRole.Error.Error()})
	}

	insertRole := models.InsertUserRole{
		UserID:    result.User.ID,
		RoleID:    role.ID,
		CreatedBy: "SignUp System",
	}

	if assignRole := au.DB.Table("usr_roleuser").Create(&insertRole); assignRole.Error != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": assignRole.Error.Error()})
	}

	var userRoleResponse *models.UserRoleResponse
	if getUserRole := au.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", result.User.ID); getUserRole.Error != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	json := fiber.Map{
		"status":        "success",
		"access_token":  result.AccessToken,
		"refresh_token": result.RefreshToken,
		"expired_at":    result.ExpiresAt,
		"expires_in":    result.ExpiresIn,
		"token_type":    result.TokenType,
		"data":          userRoleResponse,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": json})
}

func (au *AuthController) SignInWithEmailPassword(c fiber.Ctx) error {
	var signIn *models.SignInRequest

	if err := c.Bind().JSON(&signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	result, err := au.API.Auth.SignInWithEmailPassword(signIn.Email, signIn.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var userRoleResponse *models.UserRoleResponse
	if getUserRole := au.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", result.User.ID); getUserRole.Error != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	json := fiber.Map{
		"status":        "success",
		"access_token":  result.AccessToken,
		"refresh_token": result.RefreshToken,
		"expired_at":    result.ExpiresAt,
		"expires_in":    result.ExpiresIn,
		"token_type":    result.TokenType,
		"data":          userRoleResponse,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": json})
}

func (au *AuthController) SignOut(c fiber.Ctx) error {
	//TODO: Do Something
	return nil
}
