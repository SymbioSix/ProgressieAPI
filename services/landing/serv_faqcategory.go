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

func (service LandFaqCategoryService) GetAllFaqCategory(c fiber.Ctx) error {
	var faqCategory []landing.Land_Faqcategory_Response
	if err := service.DB.Table("land_faqcategory").Find(&faqCategory); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(faqCategory), "result": faqCategory})
}

func (service LandFaqCategoryService) CreateFaqCategoryRequest(c fiber.Ctx) error {
	var request landing.Land_Faqcategory_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Create(&request).Error; err != nil {
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

func (service LandFaqCategoryService) GetFaqCategoryRequestByID(c fiber.Ctx) error {
	faqCategoryID := c.Params("id")
	var request landing.Land_Faqcategory_Request
	if err := service.DB.Where("faq_category_id = ?", faqCategoryID).First(&request).Error; err != nil {
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

func (service LandFaqCategoryService) UpdateFaqCategoryRequest(c fiber.Ctx) error {
	faqCategoryID := c.Params("id")

	var updatedRequest landing.Land_Faqcategory_Request
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	var request landing.Land_Faqcategory_Request
	if err := service.DB.Where("faq_category_id = ?", faqCategoryID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "FAQ category not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	request.FaqCategoryName = updatedRequest.FaqCategoryName
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Save(&request).Error; err != nil {
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
