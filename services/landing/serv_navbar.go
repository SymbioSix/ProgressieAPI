package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LandNavbarService struct {
	DB *gorm.DB
}

func NewLandNavbarService(db *gorm.DB) LandNavbarService {
	return LandNavbarService{DB: db}
}

func (service LandNavbarService) GetAllNavbar(c fiber.Ctx) error {
	var navbars []landing.Land_Navbar_Response
	if err := service.DB.Table("land_navbar").Find(&navbars); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(navbars), "result": navbars})
}

func (service LandNavbarService) CreateNavbarRequest(c fiber.Ctx) error {
	var request landing.Land_Navbar_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_navbar").Create(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Navbar_Response{
		NavComponentID:    request.NavComponentID,
		NavComponentName:  request.NavComponentName,
		NavComponentGroup: request.NavComponentGroup,
		NavComponentIcon:  request.NavComponentIcon,
		Tooltip:           request.Tooltip,
		Endpoint:          request.Endpoint,
		CreatedBy:         request.CreatedBy,
		CreatedAt:         request.CreatedAt,
		UpdatedBy:         request.UpdatedBy,
		UpdatedAt:         request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (service LandNavbarService) GetNavbarRequestByID(c fiber.Ctx) error {
	navComponentID := c.Params("id")
	var request landing.Land_Navbar_Request
	if err := service.DB.Table("land_navbar").Where("nav_component_id = ?", navComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Navbar component not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Navbar_Response{
		NavComponentID:    request.NavComponentID,
		NavComponentName:  request.NavComponentName,
		NavComponentGroup: request.NavComponentGroup,
		NavComponentIcon:  request.NavComponentIcon,
		Tooltip:           request.Tooltip,
		Endpoint:          request.Endpoint,
		CreatedBy:         request.CreatedBy,
		CreatedAt:         request.CreatedAt,
		UpdatedBy:         request.UpdatedBy,
		UpdatedAt:         request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (service LandNavbarService) UpdateNavbarRequest(c fiber.Ctx) error {
	navComponentID := c.Params("id")

	var updatedRequest landing.Land_Navbar_Request
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}

	var request landing.Land_Navbar_Request
	if err := service.DB.Table("land_navbar").Where("nav_component_id = ?", navComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Navbar component not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	request.NavComponentName = updatedRequest.NavComponentName
	request.NavComponentGroup = updatedRequest.NavComponentGroup
	request.NavComponentIcon = updatedRequest.NavComponentIcon
	request.Tooltip = updatedRequest.Tooltip
	request.Endpoint = updatedRequest.Endpoint
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := service.DB.Table("land_navbar").Save(&request).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	response := landing.Land_Navbar_Response{
		NavComponentID:    request.NavComponentID,
		NavComponentName:  request.NavComponentName,
		NavComponentGroup: request.NavComponentGroup,
		NavComponentIcon:  request.NavComponentIcon,
		Tooltip:           request.Tooltip,
		Endpoint:          request.Endpoint,
		CreatedBy:         request.CreatedBy,
		CreatedAt:         request.CreatedAt,
		UpdatedBy:         request.UpdatedBy,
		UpdatedAt:         request.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
