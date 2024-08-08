package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LandFaqService struct {
	DB *gorm.DB
}

func NewLandFaqService(db *gorm.DB) LandFaqService {
	return LandFaqService{DB: db}
}

// GetAllFaq godoc
//	@Summary		Get all FAQ components
//	@Description	Get all FAQ components
//	@Tags			FAQ
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		landing.Land_Faq_Response
//	@Failure		500	{object}	object
//	@Router			/faq [get]
func (service LandFaqService) GetAllFaq(c fiber.Ctx) error {
	var faq []landing.Land_Faq_Response
	if err := service.DB.Table("land_faq").Find(&faq); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(faq), "result": faq})
}
// CreateFaqRequest godoc
//	@Summary		Create a new FAQ component
//	@Description	Create a new FAQ component
//	@Tags			FAQ
//	@Accept			json
//	@Produce		json
//	@Param			request	body		landing.Land_Faq_Request	true	"FAQ component data"
//	@Success		201		{object}	landing.Land_Faq_Response
//	@Failure		400		{object}	object
//	@Failure		500		{object}	object
//	@Router			/faq [post]
func (service LandFaqService) CreateFaqRequest(c fiber.Ctx) error {
	var request landing.Land_Faq_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_faq").Create(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Faq_Response{
		FaqID:          request.FaqID,
		FaqCategory:    request.FaqCategory,
		FaqTitle:       request.FaqTitle,
		FaqDescription: request.FaqDescription,
		Tooltip:        request.Tooltip,
		CreatedBy:      request.CreatedBy,
		CreatedAt:      request.CreatedAt,
		UpdatedBy:      request.UpdatedBy,
		UpdatedAt:      request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// GetFaqRequestByID godoc
//	@Summary		Get a FAQ component by ID
//	@Description	Get a FAQ component by ID
//	@Tags			FAQ
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"FAQ component ID"
//	@Success		200	{object}	landing.Land_Faq_Response
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/faq/{id} [get]
func (service LandFaqService) GetFaqRequestByID(c fiber.Ctx) error {
	faqID := c.Params("id")
	var request landing.Land_Faq_Request
	if err := service.DB.Table("land_faq").Where("faq_id = ?", faqID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "FAQ not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Faq_Response{
		FaqID:          request.FaqID,
		FaqCategory:    request.FaqCategory,
		FaqTitle:       request.FaqTitle,
		FaqDescription: request.FaqDescription,
		Tooltip:        request.Tooltip,
		CreatedBy:      request.CreatedBy,
		CreatedAt:      request.CreatedAt,
		UpdatedBy:      request.UpdatedBy,
		UpdatedAt:      request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
// UpdateFaqRequest godoc
//	@Summary		Update a FAQ component
//	@Description	Update a FAQ component
//	@Tags			FAQ
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"FAQ component ID"
//	@Param			request	body		landing.Land_Faq_Request	true	"Updated FAQ component data"
//	@Success		200		{object}	landing.Land_Faq_Response
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/faq/{id} [put]
func (service LandFaqService) UpdateFaqRequest(c fiber.Ctx) error {
	faqID := c.Params("id")

	var updatedRequest landing.Land_Faq_Request
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	var request landing.Land_Faq_Request
	if err := service.DB.Table("land_faq").Where("faq_id = ?", faqID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "FAQ not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	request.FaqCategory = updatedRequest.FaqCategory
	request.FaqTitle = updatedRequest.FaqTitle
	request.FaqDescription = updatedRequest.FaqDescription
	request.Tooltip = updatedRequest.Tooltip
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_faq").Save(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Faq_Response{
		FaqID:          request.FaqID,
		FaqCategory:    request.FaqCategory,
		FaqTitle:       request.FaqTitle,
		FaqDescription: request.FaqDescription,
		Tooltip:        request.Tooltip,
		CreatedBy:      request.CreatedBy,
		CreatedAt:      request.CreatedAt,
		UpdatedBy:      request.UpdatedBy,
		UpdatedAt:      request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
