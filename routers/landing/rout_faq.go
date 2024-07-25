package landing

import (
	services "github.com/SymbioSix/ProgressieAPI/services/landing"
	"github.com/gofiber/fiber/v3"
)

type LandFaqRouter struct {
	faqService services.LandFaqService
}

func NewLandFaqRouter(faqService services.LandFaqService) LandFaqRouter {
	return LandFaqRouter{faqService}
}

func (router *LandFaqRouter) LandFaqRoutes(rg fiber.Router) {
	api := rg.Group("faq")

	api.Get("/", router.faqService.GetAllFaq)
	api.Post("/", router.faqService.CreateFaqRequest)
	api.Get("/:id", router.faqService.GetFaqRequestByID)
	api.Put("/:id", router.faqService.UpdateFaqRequest)
}
