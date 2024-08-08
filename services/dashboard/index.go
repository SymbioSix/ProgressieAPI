package dashboard

import (
	auth "github.com/SymbioSix/ProgressieAPI/models/auth"
	dashb "github.com/SymbioSix/ProgressieAPI/models/dashboard"
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
func (dash *DashboardController) SidebarMapper(c fiber.Ctx) error {
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
