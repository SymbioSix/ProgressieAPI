package landing

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/landing"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// FooterService struct
type FooterService struct {
	db *gorm.DB
}

// NewFooterService creates a new FooterService
func NewFooterService(db *gorm.DB) FooterService {
	return FooterService{db: db}
}

func (service *FooterService) GetAllFooter() ([]models.Land_Footer_Response, error) {
	var footer []models.Land_Footer_Response
	if err := service.db.Table("land_footer").Find(&footer); err.Error != nil {
		return nil, err.Error
	}

	return footer, nil
}

// CreateFooter creates a new footer component
func (service *FooterService) CreateFooter(req *models.Land_Footer_Request) (*models.Land_Footer_Response, error) {
	footer := &models.Land_Footer_Request{
		FooterComponentName:  req.FooterComponentName,
		FooterComponentGroup: req.FooterComponentGroup,
		FooterComponentIcon:  req.FooterComponentIcon,
		Tooltip:              req.Tooltip,
		Endpoint:             req.Endpoint,
		CreatedBy:            req.CreatedBy,
		CreatedAt:            time.Now(),
	}

	if err := service.db.Create(footer).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Footer_Response{
		FooterComponentID:    footer.FooterComponentID,
		FooterComponentName:  footer.FooterComponentName,
		FooterComponentGroup: footer.FooterComponentGroup,
		FooterComponentIcon:  footer.FooterComponentIcon,
		Tooltip:              footer.Tooltip,
		Endpoint:             footer.Endpoint,
		CreatedBy:            footer.CreatedBy,
		CreatedAt:            footer.CreatedAt,
	}

	return response, nil
}

// GetFooter retrieves a footer component by ID
func (service *FooterService) GetFooter(id int) (*models.Land_Footer_Response, error) {
	var footer models.Land_Footer_Request
	if err := service.db.First(&footer, id).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Footer_Response{
		FooterComponentID:    footer.FooterComponentID,
		FooterComponentName:  footer.FooterComponentName,
		FooterComponentGroup: footer.FooterComponentGroup,
		FooterComponentIcon:  footer.FooterComponentIcon,
		Tooltip:              footer.Tooltip,
		Endpoint:             footer.Endpoint,
		CreatedBy:            footer.CreatedBy,
		CreatedAt:            footer.CreatedAt,
		UpdatedBy:            footer.UpdatedBy,
		UpdatedAt:            footer.UpdatedAt,
	}

	return response, nil
}

// UpdateFooter updates an existing footer component
func (service *FooterService) UpdateFooter(id int, req *models.Land_Footer_Request) (*models.Land_Footer_Response, error) {
	var footer models.Land_Footer_Request
	if err := service.db.First(&footer, id).Error; err != nil {
		return nil, err
	}

	footer.FooterComponentName = req.FooterComponentName
	footer.FooterComponentGroup = req.FooterComponentGroup
	footer.FooterComponentIcon = req.FooterComponentIcon
	footer.Tooltip = req.Tooltip
	footer.Endpoint = req.Endpoint
	footer.UpdatedBy = req.UpdatedBy
	footer.UpdatedAt = time.Now()

	if err := service.db.Save(&footer).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Footer_Response{
		FooterComponentID:    footer.FooterComponentID,
		FooterComponentName:  footer.FooterComponentName,
		FooterComponentGroup: footer.FooterComponentGroup,
		FooterComponentIcon:  footer.FooterComponentIcon,
		Tooltip:              footer.Tooltip,
		Endpoint:             footer.Endpoint,
		CreatedBy:            footer.CreatedBy,
		CreatedAt:            footer.CreatedAt,
		UpdatedBy:            footer.UpdatedBy,
		UpdatedAt:            footer.UpdatedAt,
	}

	return response, nil
}

// DeleteFooter deletes a footer component by ID
func (service *FooterService) DeleteFooter(id int) error {
	if err := service.db.Delete(&models.Land_Footer_Request{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (service FooterService) GetAllFooterHandler(c fiber.Ctx) error {
	response, err := service.GetAllFooter()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
