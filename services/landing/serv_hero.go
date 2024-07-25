package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LandHeroService struct {
	DB *gorm.DB
}

func NewLandHeroService(db *gorm.DB) LandHeroService {
	return LandHeroService{DB: db}
}

func (service LandHeroService) GetAllHero(c fiber.Ctx) error {
	var heros []landing.Land_Hero_Response
	if err := service.DB.Table("land_hero").Find(&heros); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(heros), "result": heros})
}

func (service LandHeroService) CreateHeroRequest(c fiber.Ctx) error {
	var request landing.Land_Hero_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Create(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Hero_Response{
		HeroComponentID:       request.HeroComponentID,
		HeroComponentTitle:    request.HeroComponentTitle,
		HeroComponentSubtitle: request.HeroComponentSubtitle,
		HeroComponentImage:    request.HeroComponentImage,
		HeroComponentCoverImg: request.HeroComponentCoverImg,
		Tooltip:               request.Tooltip,
		CreatedBy:             request.CreatedBy,
		CreatedAt:             request.CreatedAt,
		UpdatedBy:             request.UpdatedBy,
		UpdatedAt:             request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (service LandHeroService) GetHeroRequestByID(c fiber.Ctx) error {
	heroComponentID := c.Params("id")
	var request landing.Land_Hero_Request
	if err := service.DB.Where("hero_component_id = ?", heroComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Hero component not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Hero_Response{
		HeroComponentID:       request.HeroComponentID,
		HeroComponentTitle:    request.HeroComponentTitle,
		HeroComponentSubtitle: request.HeroComponentSubtitle,
		HeroComponentImage:    request.HeroComponentImage,
		HeroComponentCoverImg: request.HeroComponentCoverImg,
		Tooltip:               request.Tooltip,
		CreatedBy:             request.CreatedBy,
		CreatedAt:             request.CreatedAt,
		UpdatedBy:             request.UpdatedBy,
		UpdatedAt:             request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (service LandHeroService) UpdateHeroRequest(c fiber.Ctx) error {
	heroComponentID := c.Params("id")

	var updatedRequest landing.Land_Hero_Request
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	var request landing.Land_Hero_Request
	if err := service.DB.Where("hero_component_id = ?", heroComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Hero component not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	request.HeroComponentTitle = updatedRequest.HeroComponentTitle
	request.HeroComponentSubtitle = updatedRequest.HeroComponentSubtitle
	request.HeroComponentImage = updatedRequest.HeroComponentImage
	request.HeroComponentCoverImg = updatedRequest.HeroComponentCoverImg
	request.Tooltip = updatedRequest.Tooltip
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Save(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Hero_Response{
		HeroComponentID:       request.HeroComponentID,
		HeroComponentTitle:    request.HeroComponentTitle,
		HeroComponentSubtitle: request.HeroComponentSubtitle,
		HeroComponentImage:    request.HeroComponentImage,
		HeroComponentCoverImg: request.HeroComponentCoverImg,
		Tooltip:               request.Tooltip,
		CreatedBy:             request.CreatedBy,
		CreatedAt:             request.CreatedAt,
		UpdatedBy:             request.UpdatedBy,
		UpdatedAt:             request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
