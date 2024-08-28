package routers

import (
	"github.com/SymbioSix/ProgressieAPI/middleware"
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

	router.Get("/sidebar", dr.dashboardRouter.SidebarMapperForAuthenticatedUser, middleware.RestrictUnauthenticatedUser())
	router.Get("/profile", dr.dashboardRouter.GetUserProfile, middleware.RestrictUnauthenticatedUser())
	router.Put("/profile", dr.dashboardRouter.UpdateUserProfile, middleware.RestrictUnauthenticatedUser())
	router.Put("/skill", dr.dashboardRouter.CreateOrUpdateUserSkill, middleware.RestrictUnauthenticatedUser())
	router.Delete("/:id/soft", dr.dashboardRouter.SoftDeleteUser, middleware.RestrictNonAdmin())
}
