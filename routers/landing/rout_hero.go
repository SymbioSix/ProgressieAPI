package landing

import (
	"github.com/SymbioSix/ProgressieAPI/services/landing"
	"github.com/gofiber/fiber/v3"
)

type LandHeroRouter struct {
	heroService landing.LandHeroService
}

func NewLandHeroRouter(heroService landing.LandHeroService) LandHeroRouter {
	return LandHeroRouter{heroService}
}

func (hr *LandHeroRouter) LandHeroRoutes(rg fiber.Router) {
	router := rg.Group("hero")

	router.Get("/", hr.heroService.GetAllHero)
	router.Post("/", hr.heroService.CreateHeroRequest)
	router.Get("/:id", hr.heroService.GetHeroRequestByID)
	router.Put("/:id", hr.heroService.UpdateHeroRequest)
}
