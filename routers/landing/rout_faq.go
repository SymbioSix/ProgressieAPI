package landing

import (
    "github.com/gofiber/fiber/v3"
    services "github.com/SymbioSix/ProgressieAPI/services/landing"
)

type LandFaqRouter struct {
    faqService services.LandFaqService
}

func NewLandFaqRouter(faqService services.LandFaqService) LandFaqRouter {
    return LandFaqRouter{faqService}
}

func (router *LandFaqRouter) LandFaqRoutes(rg fiber.Router) {
    api := rg.Group("faq")

    api.Post("/", router.faqService.CreateFaqRequest)
    api.Get("/faq/:id", router.faqService.GetFaqRequestByID)
    api.Put("/faq/:id", router.faqService.UpdateFaqRequest)
}
