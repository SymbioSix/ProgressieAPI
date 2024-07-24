package landing

import (
    "github.com/gofiber/fiber/v3"
    services "github.com/SymbioSix/ProgressieAPI/services/landing"
)

type LandNavbarRouter struct {
    navbarService services.LandNavbarService
}

func NewLandNavbarRouter(navbarService services.LandNavbarService) LandNavbarRouter {
    return LandNavbarRouter{navbarService}
}

func (nr *LandNavbarRouter) LandNavbarRoutes(rg fiber.Router) {
    router := rg.Group("navbar")

    router.Post("/", nr.navbarService.CreateNavbarRequest)
    router.Get("/nama-endpoint/:id", nr.navbarService.GetNavbarRequestByID)
    router.Put("/nama-endpoint/:id", nr.navbarService.UpdateNavbarRequest)
}
