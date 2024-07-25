package landing

import (
	services "github.com/SymbioSix/ProgressieAPI/services/landing"
	"github.com/gofiber/fiber/v3"
)

type LandFooterRouter struct {
	footerService services.FooterService
}

func NewLandFooterRouter(footerService services.FooterService) LandFooterRouter {
	return LandFooterRouter{footerService}
}

func (fo *LandFooterRouter) LandFooterRoutes(rg fiber.Router) {
	router := rg.Group("footer")

	router.Get("/", fo.footerService.GetAllFooterHandler)
	// router.Post("/", fo.footerService.CreateFooterHandler)
	// router.Get("/:id", fo.footerService.GetFooterRequestByIDHandler)
	// router.Put("/:id", fo.footerService.UpdateFooterRequestHandler)
	// router.Delete("/:id", fo.footerService.DeleteFooterHandler)
}
