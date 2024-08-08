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

// GetAllHero godoc
//	@Summary		Get all hero components
//	@Description	Get all hero components
//	@Tags			Hero
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		landing.Land_Hero_Response
//	@Failure		500	{object}	object
//	@Router			/hero [get]
func (service LandHeroService) GetAllHero(c fiber.Ctx) error {
	var heros []landing.Land_Hero_Response
	if err := service.DB.Table("land_hero").Find(&heros); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(heros), "result": heros})
}

// CreateHeroRequest godoc
//	@Summary		Create a new hero component
//	@Description	Create a new hero component
//	@Tags			Hero
//	@Accept			json
//	@Produce		json
//	@Param			request	body		landing.Land_Hero_Request	true	"Hero component data"
//	@Success		201		{object}	landing.Land_Hero_Response
//	@Failure		400		{object}	object
//	@Failure		500		{object}	object
//	@Router			/hero [post]

func (service LandHeroService) CreateHeroRequest(c fiber.Ctx) error {
	var request landing.Land_Hero_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_hero").Create(&request).Error; err != nil {
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

// GetHeroRequestByID godoc
//	@Summary		Get a hero component by ID
//	@Description	Get a hero component by ID
//	@Tags			Hero
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Hero component ID"
//	@Success		200	{object}	landing.Land_Hero_Response
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/hero/{id} [get]

func (service LandHeroService) GetHeroRequestByID(c fiber.Ctx) error {
	heroComponentID := c.Params("id")
	var request landing.Land_Hero_Request
	if err := service.DB.Table("land_hero").Where("hero_component_id = ?", heroComponentID).First(&request).Error; err != nil {
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

// UpdateHeroRequest godoc
//	@Summary		Update a hero component
//	@Description	Update a hero component
//	@Tags			Hero
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Hero component ID"
//	@Param			request	body		landing.Land_Hero_Request	true	"Updated hero component data"
//	@Success		200		{object}	landing.Land_Hero_Response
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/hero/{id} [put]

func (service LandHeroService) UpdateHeroRequest(c fiber.Ctx) error {
	heroComponentID := c.Params("id")

	var updatedRequest landing.Land_Hero_Request
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	var request landing.Land_Hero_Request
	if err := service.DB.Table("land_hero").Where("hero_component_id = ?", heroComponentID).First(&request).Error; err != nil {
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

	if err := service.DB.Table("land_hero").Save(&request).Error; err != nil {
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
