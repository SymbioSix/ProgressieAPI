package dashboard

import (
	auth "github.com/SymbioSix/ProgressieAPI/models/auth"
	dashb "github.com/SymbioSix/ProgressieAPI/models/dashboard"
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

func (dash *DashboardController) SidebarMapper(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(roleSidebarResponse), "data": roleSidebarResponse})
}

func (dash *DashboardController) GetUserProfile(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var userRoleResponse *auth.UserRoleResponse
	if getUserRole := dash.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", getUser.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "result": userRoleResponse})
}
