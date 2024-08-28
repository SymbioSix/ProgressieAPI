package dashboard

import (
	"time"

	auth "github.com/SymbioSix/ProgressieAPI/models/auth"
	dashb "github.com/SymbioSix/ProgressieAPI/models/dashboard"
	profile "github.com/SymbioSix/ProgressieAPI/models/profile"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DashboardController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewDashboardController(DB *gorm.DB, API *utils.Client) DashboardController {
	return DashboardController{DB, API}
}

// TODO: Add Another Feature/Service Relatable to Dashboard Below this Line

// SidebarMapper godoc
//
//	@Summary		Get sidebar mapping for the user
//	@Description	Get sidebar mapping for the authenticated user based on their roles
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dashb.RoleSidebarResponse
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/dashboard/sidebar [get]
func (dash *DashboardController) SidebarMapperForAuthenticatedUser(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}
	var userRoleResponse []auth.UserRoleResponse
	if getUserRole := dash.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", getUser.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	var roleSidebarResponse []dashb.RoleSidebarResponse
	for i := 0; i < len(userRoleResponse); i++ {
		if getRoleSidebarFromAuthenticatedUser := dash.DB.Table("usr_rolesidebar").Preload(clause.Associations).Find(&roleSidebarResponse, "role_id = ?", userRoleResponse[i].RoleData); getRoleSidebarFromAuthenticatedUser != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getRoleSidebarFromAuthenticatedUser.Error.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(roleSidebarResponse)
}

// GetUserProfile godoc
//
//	@Summary		Get user profile
//	@Description	Get the profile of the authenticated user
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	auth.UserRoleResponse
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/dashboard/profile [get]
func (dash *DashboardController) GetUserProfile(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}
	var userRoleResponse *auth.UserRoleResponse
	if getUserRole := dash.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", getUser.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(userRoleResponse)
}

// UpdateUserProfile godoc
//
//	@Summary		Update user profile
//	@Description	Update the profile of the authenticated user
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		auth.UpdateUserProfileRequest	true	"Update User Profile Request"
//	@Success		200		{object}	status.StatusModel
//	@Failure		401		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/dashboard/profile [put]
func (dash *DashboardController) UpdateUserProfile(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}

	var request auth.UpdateUserProfileRequest
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var updateUser auth.UserModel
	if res := dash.DB.Where(auth.UserModel{UserID: getUser.ID}).First(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	if request.FirstName != "" {
		updateUser.FirstName = request.FirstName
	}
	if request.LastName != "" {
		updateUser.LastName = request.LastName
	}
	if request.Email != "" {
		updateUser.Email = request.Email
	}
	if request.PhotoProfile != "" {
		updateUser.PhotoProfile = request.PhotoProfile
	}
	if request.TitleProfile != "" {
		updateUser.TitleProfile = request.TitleProfile
	}
	if request.PhoneNumber != "" {
		updateUser.PhoneNumber = request.PhoneNumber
	}
	if request.Description != "" {
		updateUser.Description = request.Description
	}
	if request.Gender != "" {
		updateUser.Gender = request.Gender
	}

	if res := dash.DB.Save(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Profile Updated Successfully!"})
}

// UpdateUserSkill godoc
//
//	@Summary		Update user skill
//	@Description	Update the skill of the authenticated user
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		profile.UpdateUserTitleSkillRequest	true	"Update User Skill Request"
//	@Success		200		{object}	status.StatusModel
//	@Failure		401		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/dashboard/skill [put]
func (dash *DashboardController) CreateOrUpdateUserSkill(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}

	var request profile.UpdateUserTitleSkillRequest
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var updateUser profile.UserTitleSkill
	if res := dash.DB.Where(auth.UserModel{UserID: getUser.ID}).First(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	if request.TitleSkill != "" {
		updateUser.TitleSkill = request.TitleSkill
	}
	if request.Subtitle != "" {
		updateUser.Subtitle = request.Subtitle
	}
	if request.UpdatedAt.String() != "" {
		updateUser.UpdatedAt = request.UpdatedAt
	}

	if res := dash.DB.Save(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Skills Updated Successfully!"})
}

// SoftDeleteUser godoc
//
//	@Summary		Soft Delete User
//	@Description	Soft Delete User (Can only be performed by user with at least Administrator role)
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of the user who wants to be deleted"
//	@Success		200	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/dashboard/{id}/soft [delete]
func (dash *DashboardController) SoftDeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	var now time.Time = time.Now()
	_, _, err := dash.API.Rest.From("auth.users").Update("deleted_at = "+now.String(), "", "exact").Eq("id", id).Execute()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	var user auth.UserModel
	if res := dash.DB.Table("usr_user").Where("user_id = ?", id).First(&user); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	user.Status = "Deactivated"

	if res := dash.DB.Table("usr_user").Save(&user); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Is Successfully Soft Deleted!"})
}
