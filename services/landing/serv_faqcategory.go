package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LandFaqCategoryService struct {
	DB *gorm.DB
}

func NewLandFaqCategoryService(db *gorm.DB) LandFaqCategoryService {
	return LandFaqCategoryService{DB: db}
}

// GetAllFaqCategory godoc
//
//	@Summary		Get all FAQ categories
//	@Description	Get all FAQ categories
//	@Tags			FAQ Category Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		[]landing.Land_Faqcategory
//	@Failure		500	{object}	object
//	@Router			/faqcategory [get]
func (service *LandFaqCategoryService) GetAllFaqCategory(c fiber.Ctx) error {
	var faqCategory []landing.Land_Faqcategory
	if err := service.DB.Preload(clause.Associations).Find(&faqCategory).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	return c.Status(fiber.StatusOK).JSON(faqCategory)
}

// CreateFaqCategoryRequest godoc
//
//	@Summary		Create a new FAQ category
//	@Description	Create a new FAQ category
//	@Tags			FAQ Category Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		landing.Land_Faqcategory	true	"FAQ category data"
//	@Success		201		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/faqcategory [post]
func (service *LandFaqCategoryService) CreateFaqCategoryRequest(c fiber.Ctx) error {
	var request landing.Land_Faqcategory
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

// GetFaqCategoryRequestByID godoc
//
//	@Summary		Get a FAQ category by ID
//	@Description	Get a FAQ category by ID
//	@Tags			FAQ Category Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"FAQ category ID"
//	@Success		200	{object}	landing.Land_Faqcategory
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/faqcategory/{id} [get]
func (service *LandFaqCategoryService) GetFaqCategoryRequestByID(c fiber.Ctx) error {
	faqCategoryID := c.Params("id")
	var request landing.Land_Faqcategory
	if err := service.DB.Where("faq_category_id = ?", faqCategoryID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "FAQ Category not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	response := request

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateFaqCategoryRequest godoc
//
//	@Summary		Update a FAQ category
//	@Description	Update a FAQ category
//	@Tags			FAQ Category Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"FAQ category ID"
//	@Param			request	body		landing.Land_Faqcategory	true	"Updated FAQ category data"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/faqcategory/{id} [put]
func (service *LandFaqCategoryService) UpdateFaqCategoryRequest(c fiber.Ctx) error {
	faqCategoryID := c.Params("id")

	var updatedRequest landing.Land_Faqcategory
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	var request landing.Land_Faqcategory
	if err := service.DB.Where("faq_category_id = ?", faqCategoryID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "FAQ Category not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	request.FaqCategoryName = updatedRequest.FaqCategoryName
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Save(&request).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Updated successfully"})
}
