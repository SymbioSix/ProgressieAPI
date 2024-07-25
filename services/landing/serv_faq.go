package landing

import (
    "errors"
    "time"

    "gorm.io/gorm"
    "github.com/gofiber/fiber/v3"
     landing "github.com/SymbioSix/ProgressieAPI/models/landing"
)

type LandFaqService struct {
    DB *gorm.DB
}

func NewLandFaqService(db *gorm.DB) LandFaqService {
    return LandFaqService{DB: db}
}

func (service LandFaqService) CreateFaqRequest(c fiber.Ctx) error {
    var request landing.Land_Faq_Request
    if err := c.Bind().JSON(&request); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
    }

    request.CreatedAt = time.Now()
    request.UpdatedAt = time.Now()

    if err := service.DB.Create(&request).Error; err != nil {
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

func (service LandFaqService) GetFaqRequestByID(c fiber.Ctx) error {
    faqID := c.Params("id")
    var request landing.Land_Faq_Request
    if err := service.DB.Where("faq_id = ?", faqID).First(&request).Error; err != nil {
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

func (service LandFaqService) UpdateFaqRequest(c fiber.Ctx) error {
    faqID := c.Params("id")

    var updatedRequest landing.Land_Faq_Request
    if err := c.Bind().JSON(&updatedRequest); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
    }

    var request landing.Land_Faq_Request
    if err := service.DB.Where("faq_id = ?", faqID).First(&request).Error; err != nil {
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

    if err := service.DB.Save(&request).Error; err != nil {
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
