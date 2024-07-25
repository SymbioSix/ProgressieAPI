package landing

import (
    "github.com/gofiber/fiber/v3"
    services "github.com/SymbioSix/ProgressieAPI/services/landing"
)

type LandFaqCategoryRouter struct {
    faqCategoryService services.LandFaqCategoryService
}

func NewLandFaqCategoryRouter(faqCategoryService services.LandFaqCategoryService) LandFaqCategoryRouter {
    return LandFaqCategoryRouter{faqCategoryService}
}

func (router *LandFaqCategoryRouter) LandFaqCategoryRoutes(rg fiber.Router) {
    api := rg.Group("faq-category")

    api.Post("/", router.faqCategoryService.CreateFaqCategoryRequest)
    api.Get("/faqcategory/:id", router.faqCategoryService.GetFaqCategoryRequestByID)
    api.Put("/faqcategory/:id", router.faqCategoryService.UpdateFaqCategoryRequest)
}
