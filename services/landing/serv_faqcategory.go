package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LandFaqCategoryService struct {
	DB *gorm.DB
}

func NewLandFaqCategoryService(db *gorm.DB) LandFaqCategoryService {
	return LandFaqCategoryService{DB: db}
}

// GetAllFaqCategory godoc
//	@Summary		Get all FAQ categories
//	@Description	Get all FAQ categories
//	@Tags			FAQ Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		landing.Land_Faqcategory_Response
//	@Failure		500	{object}	object
//	@Router			/faqcategory [get]

func (service LandFaqCategoryService) GetAllFaqCategory(c fiber.Ctx) error {
	var faqCategory []landing.Land_Faqcategory_Response
	if err := service.DB.Table("land_faqcategory").Find(&faqCategory); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(faqCategory), "result": faqCategory})
}

// CreateFaqCategoryRequest godoc
//	@Summary		Create a new FAQ category
//	@Description	Create a new FAQ category
//	@Tags			FAQ Category
//	@Accept			json
//	@Produce		json
//	@Param			request	body		landing.Land_Faqcategory_Request	true	"FAQ category data"
//	@Success		201		{object}	landing.Land_Faqcategory_Response
//	@Failure		400		{object}	object
//	@Failure		500		{object}	object
//	@Router			/faqcategory [post]

func (service LandFaqCategoryService) CreateFaqCategoryRequest(c fiber.Ctx) error {
	var request landing.Land_Faqcategory_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_faqcategory").Create(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Faqcategory_Response{
		FaqCategoryID:   request.FaqCategoryID,
		FaqCategoryName: request.FaqCategoryName,
		CreatedBy:       request.CreatedBy,
		CreatedAt:       request.CreatedAt,
		UpdatedBy:       request.UpdatedBy,
		UpdatedAt:       request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// GetFaqCategoryRequestByID godoc
//	@Summary		Get a FAQ category by ID
//	@Description	Get a FAQ category by ID
//	@Tags			FAQ Category
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"FAQ category ID"
//	@Success		200	{object}	landing.Land_Faqcategory_Response
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/faqcategory/{id} [get]

func (service LandFaqCategoryService) GetFaqCategoryRequestByID(c fiber.Ctx) error {
	faqCategoryID := c.Params("id")
	var request landing.Land_Faqcategory_Request
	if err := service.DB.Table("land_faqcategory").Where("faq_category_id = ?", faqCategoryID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "FAQ category not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Faqcategory_Response{
		FaqCategoryID:   request.FaqCategoryID,
		FaqCategoryName: request.FaqCategoryName,
		CreatedBy:       request.CreatedBy,
		CreatedAt:       request.CreatedAt,
		UpdatedBy:       request.UpdatedBy,
		UpdatedAt:       request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateFaqCategoryRequest godoc
//	@Summary		Update a FAQ category
//	@Description	Update a FAQ category
//	@Tags			FAQ Category
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int									true	"FAQ category ID"
//	@Param			request	body		landing.Land_Faqcategory_Request	true	"Updated FAQ category data"
//	@Success		200		{object}	landing.Land_Faqcategory_Response
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/faqcategory/{id} [put]

func (service LandFaqCategoryService) UpdateFaqCategoryRequest(c fiber.Ctx) error {
	faqCategoryID := c.Params("id")

	var updatedRequest landing.Land_Faqcategory_Request
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	var request landing.Land_Faqcategory_Request
	if err := service.DB.Table("land_faqcategory").Where("faq_category_id = ?", faqCategoryID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "FAQ category not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	request.FaqCategoryName = updatedRequest.FaqCategoryName
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_faqcategory").Save(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Faqcategory_Response{
		FaqCategoryID:   request.FaqCategoryID,
		FaqCategoryName: request.FaqCategoryName,
		CreatedBy:       request.CreatedBy,
		CreatedAt:       request.CreatedAt,
		UpdatedBy:       request.UpdatedBy,
		UpdatedAt:       request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
