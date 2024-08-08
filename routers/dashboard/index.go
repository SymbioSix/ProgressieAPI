package routers

import (
	"github.com/SymbioSix/ProgressieAPI/services/dashboard"
	"github.com/gofiber/fiber/v3"
)

type DashboardRouter struct {
	dashboardRouter dashboard.DashboardController
}

func NewRouteAuthController(dashboardRouter dashboard.DashboardController) DashboardRouter {
	return DashboardRouter{dashboardRouter}
}

func (dr *DashboardRouter) DashboardRoutes(rg fiber.Router) {
	router := rg.Group("dashboard")

	router.Get("/sidebar", dr.dashboardRouter.SidebarMapper)
	router.Get("/profile", dr.dashboardRouter.GetUserProfile)
}
