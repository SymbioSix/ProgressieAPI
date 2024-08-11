package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/landing"
	"github.com/gofiber/fiber/v3"
)

type LandNavbarRouter struct {
	navbarService services.LandNavbarService
}

func NewLandNavbarRouter(navbarService services.LandNavbarService) LandNavbarRouter {
	return LandNavbarRouter{navbarService}
}

func (nr *LandNavbarRouter) LandNavbarRoutes(rg fiber.Router) {
	router := rg.Group("navbar")

	router.Get("/", nr.navbarService.GetAllNavbar)
	router.Post("/", nr.navbarService.CreateNavbarRequest)
	router.Get("/:id", nr.navbarService.GetNavbarRequestByID)
	router.Put("/:id", nr.navbarService.UpdateNavbarRequest)
}
