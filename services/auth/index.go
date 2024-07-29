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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var role *models.RoleModel
	if fetchRole := au.DB.Table("usr_role").First(&role, "role_name = ?", "BasicUser"); fetchRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": fetchRole.Error.Error()})
	}

	insertRole := models.InsertUserRole{
		UserID:    result.User.ID,
		RoleID:    role.ID,
		CreatedBy: "SignUp System",
	}

	if assignRole := au.DB.Table("usr_roleuser").Create(&insertRole); assignRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": assignRole.Error.Error()})
	}

	var userRoleResponse *models.UserRoleResponse
	if getUserRole := au.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", result.User.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
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
	au.API.UpdateAuthSession(result.Session)
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
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
	au.API.UpdateAuthSession(result.Session)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": json})
}

func (au *AuthController) SignOut(c fiber.Ctx) error {
	result, err := au.API.Auth.GetUser()
	if result == nil || err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "User was not logged in at first"})
	}
	if err := au.API.Auth.Logout(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Logged Out Successfully"})
}

func (au *AuthController) SendForgotPasswordEmail(c fiber.Ctx) error {
	var forgotRequest *models.ForgotPasswordRequest
	if err := c.Bind().JSON(&forgotRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := au.API.Auth.Recover(types.RecoverRequest{Email: forgotRequest.Email}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.JSON(fiber.Map{})
}

func (au *AuthController) VerifyPasswordRecovery(c fiber.Ctx) error {
	var verify *models.ConfirmationSignup
	if err := c.Bind().Query(&verify); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	_, err := au.API.Auth.Verify(types.VerifyRequest{Type: types.VerificationType(verify.Type), Token: verify.TokenHash, RedirectTo: verify.RedirectTo})
	if err != nil {
		c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		return c.Redirect().Status(fiber.StatusConflict).To("/v1/auth/failed?type=" + verify.Type)
	}
	return c.Status(fiber.StatusOK).Redirect().To(verify.RedirectTo)
}

func (au *AuthController) UpdateUserPassword(c fiber.Ctx) error {
	var password *models.UpdatePasswordAfterForgotPassword
	if err := c.Bind().JSON(&password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	_, err := au.API.Auth.UpdateUser(types.UpdateUserRequest{Password: &password.NewPassword})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Password has been changed!"})
}

func (au *AuthController) FailedAuthService(c fiber.Ctx) error {
	var failedType *models.FailedAuth
	if err := c.Bind().Query(&failedType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	if failedType.Type == "signup" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "conflict", "message": "Error When Verifying SignUp. Please Confirm Again Later!"})
	} else if failedType.Type == "recovery" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "conflict", "message": "Error When Verifying Reset Password Request. Please Try Again Later!"})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Unknown error"})
	}
}

/// [ THE CONFIRMATION EMAIL SEND SYSTEM WILL BE AUTOMATICALLY INVOKED WHEN CONFIRM EMAIL IN SUPABASE DASHBOARD CONFIGURATION IS ENABLED ]

func (au *AuthController) VerifySignUp(c fiber.Ctx) error {
	var confirm *models.ConfirmationSignup
	if err := c.Bind().Query(&confirm); err != nil {
		return c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	_, err := au.API.Auth.Verify(types.VerifyRequest{Type: types.VerificationType(confirm.Type), Token: confirm.TokenHash, RedirectTo: confirm.RedirectTo})
	if err != nil {
		c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		return c.Redirect().Status(fiber.StatusConflict).To("/v1/auth/failed?type=" + confirm.Type)
	}
	return c.Redirect().Status(fiber.StatusOK).To(confirm.RedirectTo)
}
