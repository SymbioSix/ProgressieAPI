package landing

import (
    "github.com/gofiber/fiber/v3"
    "github.com/SymbioSix/ProgressieAPI/services/landing"
)

type LandHeroRouter struct {
    heroService landing.LandHeroService
}

func NewLandHeroRouter(heroService landing.LandHeroService) LandHeroRouter {
    return LandHeroRouter{heroService}
}

func (hr *LandHeroRouter) LandHeroRoutes(rg fiber.Router) {
    router := rg.Group("hero")

    router.Post("/", hr.heroService.CreateHeroRequest)
    router.Get("/hero/:id", hr.heroService.GetHeroRequestByID)
    router.Put("/hero/:id", hr.heroService.UpdateHeroRequest)
}
