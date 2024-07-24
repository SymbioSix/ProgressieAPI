package landing

import (
	"errors"
	"strconv"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/landing"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AboutUsService struct {
	db *gorm.DB
}

func NewAboutUsService(db *gorm.DB) *AboutUsService {
	return &AboutUsService{db: db}
}

func (s *AboutUsService) CreateAboutUs(request *models.Land_Aboutus_Request) (*models.Land_Aboutus_Response, error) {
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := s.db.Create(request).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Aboutus_Response{
		AboutUsComponentID:      request.AboutUsComponentID,
		AboutUsComponentName:    request.AboutUsComponentName,
		AboutUsComponentJobdesc: request.AboutUsComponentJobdesc,
		AboutUsComponentPhoto:   request.AboutUsComponentPhoto,
		Description:             request.Description,
		Tooltip:                 request.Tooltip,
		Status:                  request.Status,
		CreatedBy:               request.CreatedBy,
		CreatedAt:               request.CreatedAt,
		UpdatedBy:               request.UpdatedBy,
		UpdatedAt:               request.UpdatedAt,
	}

	return response, nil
}

func (s *AboutUsService) GetAboutUsByID(id int) (*models.Land_Aboutus_Response, error) {
	var request models.Land_Aboutus_Request

	if err := s.db.First(&request, id).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Aboutus_Response{
		AboutUsComponentID:      request.AboutUsComponentID,
		AboutUsComponentName:    request.AboutUsComponentName,
		AboutUsComponentJobdesc: request.AboutUsComponentJobdesc,
		AboutUsComponentPhoto:   request.AboutUsComponentPhoto,
		Description:             request.Description,
		Tooltip:                 request.Tooltip,
		Status:                  request.Status,
		CreatedBy:               request.CreatedBy,
		CreatedAt:               request.CreatedAt,
		UpdatedBy:               request.UpdatedBy,
		UpdatedAt:               request.UpdatedAt,
	}

	return response, nil
}

func (s AboutUsService) UpdateAboutUs(id int, updatedRequest *models.Land_Aboutus_Request) (*models.Land_Aboutus_Response, error) {
	var request models.Land_Aboutus_Request

	if err := s.db.First(&request, id).Error; err != nil {
		return nil, err
	}

	request.AboutUsComponentName = updatedRequest.AboutUsComponentName
	request.AboutUsComponentJobdesc = updatedRequest.AboutUsComponentJobdesc
	request.AboutUsComponentPhoto = updatedRequest.AboutUsComponentPhoto
	request.Description = updatedRequest.Description
	request.Tooltip = updatedRequest.Tooltip
	request.Status = updatedRequest.Status
	request.UpdatedBy = updatedRequest.UpdatedBy
	request.UpdatedAt = time.Now()

	if err := s.db.Save(&request).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Aboutus_Response{
		AboutUsComponentID:      request.AboutUsComponentID,
		AboutUsComponentName:    request.AboutUsComponentName,
		AboutUsComponentJobdesc: request.AboutUsComponentJobdesc,
		AboutUsComponentPhoto:   request.AboutUsComponentPhoto,
		Description:             request.Description,
		Tooltip:                 request.Tooltip,
		Status:                  request.Status,
		CreatedBy:               request.CreatedBy,
		CreatedAt:               request.CreatedAt,
		UpdatedBy:               request.UpdatedBy,
		UpdatedAt:               request.UpdatedAt,
	}

	return response, nil
}

func (s *AboutUsService) DeleteAboutUs(id int) error {
	var request models.Land_Aboutus_Request

	if err := s.db.First(&request, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&request).Error; err != nil {
		return err
	}

	return nil
}

func (s AboutUsService) CreateAboutUsHandler(c fiber.Ctx) error {
	var request models.Land_Aboutus_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := s.CreateAboutUs(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (s AboutUsService) GetAboutUsByIDHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	response, err := s.GetAboutUsByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Record not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(response)
}

func (s AboutUsService) UpdateAboutUsHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	var request models.Land_Aboutus_Request
	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := s.UpdateAboutUs(id, &request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(response)
}

func (s AboutUsService) DeleteAboutUsHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	if err := s.DeleteAboutUs(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusNotFound, "Record not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
