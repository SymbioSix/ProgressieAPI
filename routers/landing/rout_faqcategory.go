package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/landing"
	"github.com/gofiber/fiber/v3"
)

type LandFaqCategoryRouter struct {
	faqCategoryService services.LandFaqCategoryService
}

func NewLandFaqCategoryRouter(faqCategoryService services.LandFaqCategoryService) LandFaqCategoryRouter {
	return LandFaqCategoryRouter{faqCategoryService}
}

func (router *LandFaqCategoryRouter) LandFaqCategoryRoutes(rg fiber.Router) {
	api := rg.Group("faqcategory")

	api.Get("/", router.faqCategoryService.GetAllFaqCategory)
	api.Post("/", router.faqCategoryService.CreateFaqCategoryRequest)
	api.Get("/:id", router.faqCategoryService.GetFaqCategoryRequestByID)
	api.Put("/:id", router.faqCategoryService.UpdateFaqCategoryRequest)
}
