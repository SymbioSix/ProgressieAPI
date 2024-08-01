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

// SignUpWithEmailPassword godoc
//
//	@Summary		Sign Up A User With Email Password
//	@Description	Signing Up A User Using Email and Password. Request Body need email, password, and username.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SignUpRequest	true	"Sign Up Credentials"
//	@Success		200		{object}	object
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/auth/signup-email-password [post]
func (au *AuthController) SignUpWithEmailPassword(c fiber.Ctx) error {
	var signup models.SignUpRequest

	if err := c.Bind().JSON(&signup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var user models.UserModel
	checkUsername := au.DB.Table("usr_user").First(&user, "username = ?", signup.Username)
	if checkUsername.Error != nil {
		if errors.Is(checkUsername.Error, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkUsername.Error.Error()})
		}
	} else if checkUsername.Error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Username has been taken"})
	}

	checkEmail := au.DB.Table("usr_user").First(&user, "email = ?", signup.Email)
	if checkEmail.Error != nil {
		if errors.Is(checkEmail.Error, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkEmail.Error.Error()})
		}
	} else if checkEmail.Error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Email has been taken"})
	}
	data := make(map[string]interface{})
	data["username"] = signup.Username

	result, err := au.API.Auth.Signup(types.SignupRequest{Email: signup.Email, Password: signup.Password, Data: data})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var role models.RoleModel
	if fetchRole := au.DB.Table("usr_role").First(&role, "role_name = ?", "BasicUser"); fetchRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": fetchRole.Error.Error()})
	}

	insertRole := models.InsertUserRole{
		UserID:    result.User.ID,
		RoleID:    role.RoleID,
		CreatedBy: "SignUp System",
	}

	if assignRole := au.DB.Table("usr_roleuser").Create(&insertRole); assignRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": assignRole.Error.Error()})
	}

	var userRoleResponse models.UserRoleResponse
	if getUserRole := au.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", result.User.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	json := fiber.Map{
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

// SignUpForAdmin godoc
//
//	@Summary		Sign Up A Admin With Email Password
//	@Description	Signing Up A Admin Using Email and Password. Request Body need email, password, and username.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SignUpRequest	true	"Sign Up Credentials"
//	@Success		200		{object}	object
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/auth/signup-admin [post]
func (au *AuthController) SignUpForAdmin(c fiber.Ctx) error {
	var signup models.SignUpRequest

	if err := c.Bind().JSON(&signup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var user models.UserModel
	checkUsername := au.DB.Table("usr_user").First(&user, "username = ?", signup.Username)
	if checkUsername.Error != nil {
		if errors.Is(checkUsername.Error, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkUsername.Error.Error()})
		}
	} else if checkUsername.Error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Username has been taken"})
	}

	checkEmail := au.DB.Table("usr_user").First(&user, "email = ?", signup.Email)
	if checkEmail.Error != nil {
		if errors.Is(checkEmail.Error, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkEmail.Error.Error()})
		}
	} else if checkEmail.Error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Email has been taken"})
	}
	data := make(map[string]interface{})
	data["username"] = signup.Username

	result, err := au.API.Auth.Signup(types.SignupRequest{Email: signup.Email, Password: signup.Password, Data: data})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var role models.RoleModel
	if fetchRoleAdmin := au.DB.Table("usr_role").First(&role, "role_name = ?", "Administrator"); fetchRoleAdmin.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": fetchRoleAdmin.Error.Error()})
	}

	insertRole := models.InsertUserRole{
		UserID:    result.User.ID,
		RoleID:    role.RoleID,
		CreatedBy: "SignUp System",
	}

	if assignRole := au.DB.Table("usr_roleuser").Create(&insertRole); assignRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": assignRole.Error.Error()})
	}

	if fetchRoleAdmin := au.DB.Table("usr_role").First(&role, "role_name = ?", "BasicUser"); fetchRoleAdmin.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": fetchRoleAdmin.Error.Error()})
	}

	insertRole = models.InsertUserRole{
		UserID:    result.User.ID,
		RoleID:    role.RoleID,
		CreatedBy: "SignUp System",
	}

	if assignRole := au.DB.Table("usr_roleuser").Create(&insertRole); assignRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": assignRole.Error.Error()})
	}

	var userRoleResponse models.UserRoleResponse
	if getUserRole := au.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", result.User.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	json := fiber.Map{
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

// SignUpForSuperUser godoc
//
//	@Summary		Sign Up A SuperUser With Email Password
//	@Description	Signing Up A SuperUser Using Email and Password. Request Body need email, password, and username.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SignUpRequest	true	"Sign Up Credentials"
//	@Success		200		{object}	object
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/auth/signup-super [post]
func (au *AuthController) SignUpForSuperUser(c fiber.Ctx) error {
	var signup models.SignUpRequest

	if err := c.Bind().JSON(&signup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var user models.UserModel
	checkUsername := au.DB.Table("usr_user").First(&user, "username = ?", signup.Username)
	if checkUsername.Error != nil {
		if errors.Is(checkUsername.Error, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkUsername.Error.Error()})
		}
	} else if checkUsername.Error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Username has been taken"})
	}

	checkEmail := au.DB.Table("usr_user").First(&user, "email = ?", signup.Email)
	if checkEmail.Error != nil {
		if errors.Is(checkEmail.Error, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": checkEmail.Error.Error()})
		}
	} else if checkEmail.Error == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Email has been taken"})
	}
	data := make(map[string]interface{})
	data["username"] = signup.Username

	result, err := au.API.Auth.Signup(types.SignupRequest{Email: signup.Email, Password: signup.Password, Data: data})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var role models.RoleModel
	if fetchRoleSuper := au.DB.Table("usr_role").First(&role, "role_name = ?", "SuperUser"); fetchRoleSuper.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": fetchRoleSuper.Error.Error()})
	}

	insertRole := models.InsertUserRole{
		UserID:    result.User.ID,
		RoleID:    role.RoleID,
		CreatedBy: "SignUp System",
	}

	if assignRole := au.DB.Table("usr_roleuser").Create(&insertRole); assignRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": assignRole.Error.Error()})
	}

	var userRoleResponse models.UserRoleResponse
	if getUserRole := au.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", result.User.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	json := fiber.Map{
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

// SignInWithEmailPassword godoc
//
//	@Summary		Sign In A User With Email Password
//	@Description	Signing In A User Using Email and Password. Request Body need email and password.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.SignInRequest	true	"Sign In Credentials"
//	@Success		200		{object}	object
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/auth/signin-email-password [post]
func (au *AuthController) SignInWithEmailPassword(c fiber.Ctx) error {
	var signIn models.SignInRequest

	if err := c.Bind().JSON(&signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	result, err := au.API.Auth.SignInWithEmailPassword(signIn.Email, signIn.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var userRoleResponse models.UserRoleResponse
	if getUserRole := au.DB.Table("usr_user").Table("usr_role").Table("usr_roleuser").Preload("UserData").Preload("RoleData").Find(&userRoleResponse, "user_id = ?", result.User.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	json := fiber.Map{
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

// SignOut godoc
//
//	@Summary		Sign Out A User
//	@Description	Signing Out A User. Required Authenticated User.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	object
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/auth/signout [post]
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

// SendForgotPasswordEmail godoc
//
//	@Summary		Send Email For Forgot Password Feature
//	@Description	Request Body need email.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.ForgotPasswordRequest	true	"Forgot Password Requirement"
//	@Success		200		{object}	object
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/auth/send-forgot-password-email [post]
func (au *AuthController) SendForgotPasswordEmail(c fiber.Ctx) error {
	var forgotRequest models.ForgotPasswordRequest
	if err := c.Bind().JSON(&forgotRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := au.API.Auth.Recover(types.RecoverRequest{Email: forgotRequest.Email}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Email Sent!"})
}

// VerifyPasswordRecovery godoc
//
//	@Summary		Verifying Forgotten Password
//	@Description	Request Body need email, password, and username. Required SendForgotPasswordEmail.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			type		query		string	true	"The Request Type From Confirmation Email"
//	@Param			token_hash	query		string	true	"Secret Token Hashed From Confirmation Email"
//	@Param			redirect_to	query		string	true	"Redirect API URL From Confirmation Email"
//	@Success		200			{object}	object
//	@Failure		400			{object}	object
//	@Failure		404			{object}	object
//	@Failure		500			{object}	object
//	@Router			/auth/verify-password-recovery [get]
func (au *AuthController) VerifyPasswordRecovery(c fiber.Ctx) error {
	var verify models.ConfirmationSignup
	if err := c.Bind().Query(&verify); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	_, err := au.API.Auth.Verify(types.VerifyRequest{Type: types.VerificationType(verify.Type), Token: verify.TokenHash, RedirectTo: verify.RedirectTo})
	if err != nil {
		c.JSON(fiber.Map{"status": "fail", "message": err.Error()})
		return c.Redirect().Status(fiber.StatusConflict).To("/v1/auth/failed?type=" + verify.Type)
	}
	return c.Redirect().Status(fiber.StatusOK).To(verify.RedirectTo)
}

// UpdateUserPassword godoc
//
//	@Summary		Update The User Password (After Confirming From Verify Password Recovery)
//	@Description	Request Body need new_password.
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.UpdatePasswordAfterForgotPassword	true	"Update The Password"
//	@Success		200		{object}	object
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/auth/update-user-password [put]
func (au *AuthController) UpdateUserPassword(c fiber.Ctx) error {
	var password models.UpdatePasswordAfterForgotPassword
	if err := c.Bind().JSON(&password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	_, err := au.API.Auth.UpdateUser(types.UpdateUserRequest{Password: &password.NewPassword})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Password has been changed!"})
}

// FailedAuthService godoc
//
//	@Summary		Failed Auth Service Messages and Status Responser
//	@Description	When Redirected To This Service, It Will Return JSON Response With Various Error Message
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			type	query		string	false	"Get The type Value When Being Redirected from Forgot Password Recovery or Verify Signup"
//	@Success		200		{object}	object
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/auth/failed [get]
func (au *AuthController) FailedAuthService(c fiber.Ctx) error {
	var failedType models.FailedAuth
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

// / [ THE CONFIRMATION EMAIL SEND SYSTEM WILL BE AUTOMATICALLY INVOKED WHEN CONFIRM EMAIL IN SUPABASE DASHBOARD CONFIGURATION IS ENABLED ]
// VerifySignup godoc
//
//	@Summary		Verify New Signed Up User
//	@Description	Not Very Useful At The Moment
//	@Tags			Auth Service
//	@Accept			json
//	@Produce		json
//	@Param			type		query		string	true	"The Request Type From Confirmation Email"
//	@Param			token_hash	query		string	true	"Secret Token Hashed From Confirmation Email"
//	@Param			redirect_to	query		string	true	"Redirect API URL From Confirmation Email"
//	@Success		200			{object}	object
//	@Failure		400			{object}	object
//	@Failure		404			{object}	object
//	@Failure		500			{object}	object
//	@Router			/auth/verify-signup [get]
func (au *AuthController) VerifySignUp(c fiber.Ctx) error {
	var confirm models.ConfirmationSignup
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
