package landing

import (
	"errors"
	"strconv"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/landing"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
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

func (service *FooterService) GetAllFooter() ([]models.Land_Footer, error) {
	var footer []models.Land_Footer
	if err := service.db.Find(&footer); err.Error != nil {
		return nil, err.Error
	}

	return footer, nil
}

// CreateFooter creates a new footer component

func (service *FooterService) CreateFooter(req *models.Land_Footer) (*models.Land_Footer, error) {
	footer := &models.Land_Footer{
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

	response := &models.Land_Footer{
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

func (service *FooterService) GetFooter(id int) (*models.Land_Footer, error) {
	var footer models.Land_Footer
	if err := service.db.First(&footer, id).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Footer{
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

func (service *FooterService) UpdateFooter(id int, req *models.Land_Footer) (*models.Land_Footer, error) {
	var footer models.Land_Footer
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

	response := &models.Land_Footer{
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
	if err := service.db.Delete(&models.Land_Footer{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Handlers
// CreateFooterHandler godoc
//
//	@Summary		Create a new Footer component
//	@Description	Create a new Footer component
//	@Tags			Footer Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Land_Footer	true	"Footer component data"
//	@Success		201		{object}	models.Land_Footer
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/footer [post]
func (service FooterService) CreateFooterHandler(c fiber.Ctx) error {
	req := new(models.Land_Footer)
	if err := c.Bind().JSON(&req); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	response, err := service.CreateFooter(req)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetFooterHandler godoc
//
//	@Summary		Get a Footer component by ID
//	@Description	Get a Footer component by ID
//	@Tags			Footer Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Footer component ID"
//	@Success		200	{object}	models.Land_Footer
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/footer/{id} [get]
func (service FooterService) GetFooterHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Invalid footer ID"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	response, err := service.GetFooter(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Footer not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateFooterHandler godoc
//
//	@Summary		Update a Footer component
//	@Description	Update a Footer component
//	@Tags			Footer Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Footer component ID"
//	@Param			request	body		models.Land_Footer	true	"Updated Footer component data"
//	@Success		200		{object}	models.Land_Footer
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/footer/{id} [put]
func (service FooterService) UpdateFooterHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Invalid footer ID"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	req := new(models.Land_Footer)
	if err := c.Bind().JSON(&req); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	response, err := service.UpdateFooter(id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Footer not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteFooterHandler godoc
//
//	@Summary		Delete a Footer component
//	@Description	Delete a Footer component
//	@Tags			Footer Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Footer component ID"
//	@Success		204
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/footer/{id} [delete]
func (service FooterService) DeleteFooterHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Invalid footer ID"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	if err := service.DeleteFooter(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Footer not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetAllFooterHandler godoc
//
//	@Summary		Get all Footer components
//	@Description	Get all Footer components
//	@Tags			Footer Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Land_Footer
//	@Failure		500	{object}	status.StatusModel
//	@Router			/footer [get]
func (service FooterService) GetAllFooterHandler(c fiber.Ctx) error {
	response, err := service.GetAllFooter()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
