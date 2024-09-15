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

type AboutUsService struct {
	db *gorm.DB
}

func NewAboutUsService(db *gorm.DB) AboutUsService {
	return AboutUsService{db: db}
}

func (s *AboutUsService) GetAllAboutUs() ([]models.Land_Aboutus, error) {
	var aboutUs []models.Land_Aboutus
	if err := s.db.Find(&aboutUs); err.Error != nil {
		return nil, err.Error
	}
	return aboutUs, nil
}

func (s *AboutUsService) CreateAboutUs(request *models.LandAboutUsRequest) (*models.Land_Aboutus, error) {
	var status string
	if request.Status == "" {
		status = "Active"
	} else {
		status = request.Status
	}
	response := &models.Land_Aboutus{
		AboutUsComponentName:    request.Name,
		AboutUsComponentJobdesc: request.Jobdesc,
		AboutUsComponentPhoto:   request.PhotoLink,
		Description:             request.Description,
		Tooltip:                 request.Tooltip,
		Status:                  status,
	}

	if err := s.db.Create(response).Error; err != nil {
		return nil, err
	}

	return response, nil
}

func (s *AboutUsService) GetAboutUsByID(id int) (*models.Land_Aboutus, error) {
	var request models.Land_Aboutus

	if err := s.db.First(&request, "aboutus_component_id = ?", id).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Aboutus{
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

func (s AboutUsService) UpdateAboutUs(id int, updatedRequest *models.LandAboutUsRequest) (*models.Land_Aboutus, error) {
	var request models.Land_Aboutus

	if err := s.db.First(&request, "aboutus_component_id = ?", id).Error; err != nil {
		return nil, err
	}

	request.AboutUsComponentName = updatedRequest.Name
	request.AboutUsComponentJobdesc = updatedRequest.Jobdesc
	request.AboutUsComponentPhoto = updatedRequest.PhotoLink
	request.Description = updatedRequest.Description
	request.Tooltip = updatedRequest.Tooltip
	request.Status = updatedRequest.Status
	request.UpdatedBy = "SYSTEM"
	request.UpdatedAt = time.Now()

	if err := s.db.Save(&request).Error; err != nil {
		return nil, err
	}

	response := &models.Land_Aboutus{
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
	var request models.Land_Aboutus

	if err := s.db.First(&request, "aboutus_component_id = ?", id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&request).Error; err != nil {
		return err
	}

	return nil
}

// GetAllAboutUsHandler godoc
//
//	@Summary		Get all About Us components
//	@Description	Get all About Us components
//	@Tags			AboutUs Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Land_Aboutus
//	@Failure		500	{object}	status.StatusModel
//	@Router			/aboutus [get]
func (s AboutUsService) GetAllAboutUsHandler(c fiber.Ctx) error {
	response, err := s.GetAllAboutUs()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateAboutUsHandler godoc
//
//	@Summary		Create a new About Us component
//	@Description	Create a new About Us component
//	@Tags			AboutUs Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.LandAboutUsRequest	true	"About Us component data"
//	@Success		201		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/aboutus [post]
func (s AboutUsService) CreateAboutUsHandler(c fiber.Ctx) error {
	var request models.LandAboutUsRequest
	if err := c.Bind().JSON(&request); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	_, err := s.CreateAboutUs(&request)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Successfully Created"})
}

// GetAboutUsByIDHandler godoc
//
//	@Summary		Get an About Us component by ID
//	@Description	Get an About Us component by ID
//	@Tags			AboutUs Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"About Us component ID"
//	@Success		200	{object}	models.Land_Aboutus
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/aboutus/{id} [get]
func (s AboutUsService) GetAboutUsByIDHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Invalid ID"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	response, err := s.GetAboutUsByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Record not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateAboutUsHandler godoc
//
//	@Summary		Update an About Us component
//	@Description	Update an About Us component
//	@Tags			AboutUs Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"About Us component ID"
//	@Param			request	body		models.LandAboutUsRequest	true	"Updated About Us component data"
//	@Success		200		{object}	models.Land_Aboutus
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/aboutus/{id} [put]
func (s AboutUsService) UpdateAboutUsHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Invalid ID"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	var request models.LandAboutUsRequest
	if err := c.Bind().JSON(&request); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	res, err := s.UpdateAboutUs(id, &request)
	if err != nil || res == nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Successfully Updated"})
}

// DeleteAboutUsHandler godoc
//
//	@Summary		Delete an About Us component
//	@Description	Delete an About Us component
//	@Tags			AboutUs Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"About Us component ID"
//	@Success		204	{object}	status.StatusModel
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/aboutus/{id} [delete]
func (s AboutUsService) DeleteAboutUsHandler(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Invalid ID"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	if err := s.DeleteAboutUs(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Record not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success", "message": "Deleted successfully"})
}
