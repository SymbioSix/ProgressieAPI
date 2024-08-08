package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LandNavbarService struct {
	DB *gorm.DB
}

func NewLandNavbarService(db *gorm.DB) LandNavbarService {
	return LandNavbarService{DB: db}
}

// GetAllNavbar godoc
//
//	@Summary		Get all navbar components
//	@Description	Get all navbar components
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]landing.Land_Navbar_Response
//	@Failure		500	{object}	status.StatusModel
//	@Router			/navbar [get]
func (service LandNavbarService) GetAllNavbar(c fiber.Ctx) error {
	var navbars []landing.Land_Navbar_Response
	if err := service.DB.Table("land_navbar").Find(&navbars); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	return c.Status(fiber.StatusOK).JSON(navbars)
}

// CreateNavbarRequest godoc
//
//	@Summary		Create a new navbar component
//	@Description	Create a new navbar component
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		landing.Land_Navbar_Request	true	"Navbar component data"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/navbar [post]
func (service LandNavbarService) CreateNavbarRequest(c fiber.Ctx) error {
	var request landing.Land_Navbar_Request
	if err := c.Bind().JSON(&request); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_navbar").Create(&request).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Created successfully"})
}

// GetNavbarRequestByID godoc
//
//	@Summary		Get a navbar component by ID
//	@Description	Get a navbar component by ID
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Navbar component ID"
//	@Success		200	{object}	landing.Land_Navbar_Response
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/navbar/{id} [get]
func (service LandNavbarService) GetNavbarRequestByID(c fiber.Ctx) error {
	navComponentID := c.Params("id")
	var request landing.Land_Navbar_Request
	if err := service.DB.Table("land_navbar").Where("nav_component_id = ?", navComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Navbar component not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	response := request

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateNavbarRequest godoc
//
//	@Summary		Update a navbar component
//	@Description	Update a navbar component
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Navbar component ID"
//	@Param			request	body		landing.Land_Navbar_Request	true	"Updated navbar component data"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/navbar/{id} [put]
func (service LandNavbarService) UpdateNavbarRequest(c fiber.Ctx) error {
	navComponentID := c.Params("id")

	var updatedRequest landing.Land_Navbar_Request
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	var request landing.Land_Navbar_Request
	if err := service.DB.Table("land_navbar").Where("nav_component_id = ?", navComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Navbar component not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	request.NavComponentName = updatedRequest.NavComponentName
	request.NavComponentGroup = updatedRequest.NavComponentGroup
	request.NavComponentIcon = updatedRequest.NavComponentIcon
	request.Tooltip = updatedRequest.Tooltip
	request.Endpoint = updatedRequest.Endpoint
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_navbar").Save(&request).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Updated successfully"})
}
