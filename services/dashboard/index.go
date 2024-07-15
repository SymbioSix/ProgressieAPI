package dashboard

import (
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
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
	return c.JSON("")
}
