package landing

import (
	"errors"
	"strconv"
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

// GetAllFooter godoc
//	@Summary		Get all Footer components
//	@Description	Get all Footer components
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Land_Footer_Response
//	@Failure		500	{object}	object
//	@Router			/footer [get]

func (service *FooterService) GetAllFooter() ([]models.Land_Footer_Response, error) {
	var footer []models.Land_Footer_Response
	if err := service.db.Table("land_footer").Find(&footer); err.Error != nil {
		return nil, err.Error
	}

	return footer, nil
}

// CreateFooter creates a new footer component

// CreateFooter godoc
//	@Summary		Create a new Footer component
//	@Description	Create a new Footer component
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Land_Footer_Request	true	"Footer component data"
//	@Success		201		{object}	models.Land_Footer_Response
//	@Failure		400		{object}	object
//	@Failure		500		{object}	object
//	@Router			/footer [post]
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

	if err := service.db.Table("land_footer").Create(footer).Error; err != nil {
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

// GetFooter godoc
//	@Summary		Get a Footer component by ID
//	@Description	Get a Footer component by ID
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Footer component ID"
//	@Success		200	{object}	models.Land_Footer_Response
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/footer/{id} [get]

func (service *FooterService) GetFooter(id int) (*models.Land_Footer_Response, error) {
	var footer models.Land_Footer_Request
	if err := service.db.Table("land_footer").First(&footer, id).Error; err != nil {
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

// UpdateFooter godoc
//	@Summary		Update a Footer component
//	@Description	Update a Footer component
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Footer component ID"
//	@Param			request	body		models.Land_Footer_Request	true	"Updated Footer component data"
//	@Success		200		{object}	models.Land_Footer_Response
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/footer/{id} [put]

func (service *FooterService) UpdateFooter(id int, req *models.Land_Footer_Request) (*models.Land_Footer_Response, error) {
	var footer models.Land_Footer_Request
	if err := service.db.Table("land_footer").First(&footer, id).Error; err != nil {
		return nil, err
	}

	footer.FooterComponentName = req.FooterComponentName
	footer.FooterComponentGroup = req.FooterComponentGroup
	footer.FooterComponentIcon = req.FooterComponentIcon
	footer.Tooltip = req.Tooltip
	footer.Endpoint = req.Endpoint
	footer.UpdatedBy = req.UpdatedBy
	footer.UpdatedAt = time.Now()

	if err := service.db.Table("land_footer").Save(&footer).Error; err != nil {
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

// DeleteFooter godoc
//	@Summary		Delete a Footer component
//	@Description	Delete a Footer component
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Footer component ID"
//	@Success		204
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/footer/{id} [delete]

func (service *FooterService) DeleteFooter(id int) error {
	if err := service.db.Table("land_footer").Delete(&models.Land_Footer_Request{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Handlers
// CreateFooterHandler godoc
//	@Summary		Create a new Footer component
//	@Description	Create a new Footer component
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Land_Footer_Request	true	"Footer component data"
//	@Success		201		{object}	models.Land_Footer_Response
//	@Failure		400		{object}	object
//	@Failure		500		{object}	object
//	@Router			/footer [post]

func (service FooterService) CreateFooterHandler(c fiber.Ctx) error {
	req := new(models.Land_Footer_Request)
	if err := c.Bind().JSON(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := service.CreateFooter(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// GetFooterHandler godoc
//	@Summary		Get a Footer component by ID
//	@Description	Get a Footer component by ID
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Footer component ID"
//	@Success		200	{object}	models.Land_Footer_Response
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/footer/{id} [get]

func (service FooterService) GetFooterHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid footer ID")
	}

	response, err := service.GetFooter(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Footer not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateFooterHandler godoc
//	@Summary		Update a Footer component
//	@Description	Update a Footer component
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Footer component ID"
//	@Param			request	body		models.Land_Footer_Request	true	"Updated Footer component data"
//	@Success		200		{object}	models.Land_Footer_Response
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/footer/{id} [put]

func (service FooterService) UpdateFooterHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid footer ID")
	}

	req := new(models.Land_Footer_Request)
	if err := c.Bind().JSON(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := service.UpdateFooter(id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Footer not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteFooterHandler godoc
//	@Summary		Delete a Footer component
//	@Description	Delete a Footer component
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Footer component ID"
//	@Success		204
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/footer/{id} [delete]

func (service FooterService) DeleteFooterHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid footer ID")
	}

	if err := service.DeleteFooter(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Footer not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// func (service FooterService) GetAllFooterHandler(c fiber.Ctx) error {
// 	var footers []models.Land_Footer_Request
// 	if err := service.db.Find(&footers).Error; err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}

// 	response := make([]*models.Land_Footer_Response, len(footers))
// 	for i, footer := range footers {
// 		response[i] = &models.Land_Footer_Response{
// 			FooterComponentID:    footer.FooterComponentID,
// 			FooterComponentName:  footer.FooterComponentName,
// 			FooterComponentGroup: footer.FooterComponentGroup,
// 			FooterComponentIcon:  footer.FooterComponentIcon,
// 			Tooltip:              footer.Tooltip,
// 			Endpoint:             footer.Endpoint,
// 			CreatedBy:            footer.CreatedBy,
// 			CreatedAt:            footer.CreatedAt,
// 			UpdatedBy:            footer.UpdatedBy,
// 			UpdatedAt:            footer.UpdatedAt,
// 		}
// 	}

// GetAllFooterHandler godoc
//	@Summary		Get all Footer components
//	@Description	Get all Footer components
//	@Tags			Footer
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Land_Footer_Response
//	@Failure		500	{object}	object
//	@Router			/footer [get]
func (service FooterService) GetAllFooterHandler(c fiber.Ctx) error {
	response, err := service.GetAllFooter()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
