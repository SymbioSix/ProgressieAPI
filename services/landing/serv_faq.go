package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LandFaqService struct {
	DB *gorm.DB
}

func NewLandFaqService(DB *gorm.DB) LandFaqService {
	return LandFaqService{DB}
}

// GetAllFaq godoc
//
//	@Summary		Get all FAQ components
//	@Description	Get all FAQ components
//	@Tags			FAQ Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}		[]landing.Land_Faq
//	@Failure		500	{object}	status.StatusModel
//	@Router			/faq [get]
func (service *LandFaqService) GetAllFaq(c fiber.Ctx) error {
	var faq []landing.Land_Faq
	if res := service.DB.Table("land_faq").Find(&faq); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	return c.Status(fiber.StatusOK).JSON(faq)
}

// CreateFaqRequest godoc
//
//	@Summary		Create a new FAQ component
//	@Description	Create a new FAQ component
//	@Tags			FAQ Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		landing.Land_Faq	true	"FAQ component data"
//	@Success		201		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/faq [post]
func (service *LandFaqService) CreateFaqRequest(c fiber.Ctx) error {
	var request landing.Land_Faq
	if err := c.Bind().JSON(&request); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Create(&request).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Created successfully"})
}

// GetFaqRequestByID godoc
//
//	@Summary		Get a FAQ component by ID
//	@Description	Get a FAQ component by ID
//	@Tags			FAQ Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"FAQ component ID"
//	@Success		200	{object}	landing.Land_Faq
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/faq/{id} [get]
func (service *LandFaqService) GetFaqRequestByID(c fiber.Ctx) error {
	faqID := c.Params("id")
	var request landing.Land_Faq
	if err := service.DB.Where("faq_id = ?", faqID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "FAQ not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	response := request

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateFaqRequest godoc
//
//	@Summary		Update a FAQ component
//	@Description	Update a FAQ component
//	@Tags			FAQ Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"FAQ component ID"
//	@Param			request	body		landing.Land_Faq	true	"Updated FAQ component data"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/faq/{id} [put]
func (service *LandFaqService) UpdateFaqRequest(c fiber.Ctx) error {
	faqID := c.Params("id")

	var updatedRequest landing.Land_Faq
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	var request landing.Land_Faq
	if err := service.DB.Where("faq_id = ?", faqID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "FAQ not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	request.FaqCategory = updatedRequest.FaqCategory
	request.FaqTitle = updatedRequest.FaqTitle
	request.FaqDescription = updatedRequest.FaqDescription
	request.Tooltip = updatedRequest.Tooltip
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Save(&request).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Updated successfully"})
}
