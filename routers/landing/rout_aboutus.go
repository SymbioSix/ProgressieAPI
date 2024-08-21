package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/landing"
	"github.com/gofiber/fiber/v3"
)

type LandAboutUsRouter struct {
	aboutUsService services.AboutUsService
}

func NewLandAboutUsRouter(aboutUsService services.AboutUsService) LandAboutUsRouter {
	return LandAboutUsRouter{aboutUsService}
}

func (ab *LandAboutUsRouter) LandAboutUsRoutes(rg fiber.Router) {
	router := rg.Group("aboutus")

	router.Get("/", ab.aboutUsService.GetAllAboutUsHandler)
	router.Post("/", ab.aboutUsService.CreateAboutUsHandler)
	router.Get("/:id", ab.aboutUsService.GetAboutUsByIDHandler)
	router.Put("/:id", ab.aboutUsService.UpdateAboutUsHandler)
	router.Delete("/:id", ab.aboutUsService.DeleteAboutUsHandler)
}
