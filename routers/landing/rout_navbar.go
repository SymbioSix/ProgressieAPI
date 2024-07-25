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
    router.Get("/navbar/:id", nr.navbarService.GetNavbarRequestByID)
    router.Put("/navbar/:id", nr.navbarService.UpdateNavbarRequest)
}
