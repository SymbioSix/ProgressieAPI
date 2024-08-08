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

func NewAboutUsService(db *gorm.DB) AboutUsService {
	return AboutUsService{db: db}
}

// GetAllAboutUs godoc
//	@Summary		Get all About Us components
//	@Description	Get all About Us components
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Land_Aboutus_Response
//	@Failure		500	{object}	object
//	@Router			/aboutus [get]
func (s *AboutUsService) GetAllAboutUs() ([]models.Land_Aboutus_Response, error) {
	var aboutUs []models.Land_Aboutus_Response
	if err := s.db.Table("land_aboutus").Find(&aboutUs); err.Error != nil {
		return nil, err.Error
	}
	return aboutUs, nil
}

// CreateAboutUs godoc
//	@Summary		Create a new About Us component
//	@Description	Create a new About Us component
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Land_Aboutus_Request	true	"About Us component data"
//	@Success		201		{object}	models.Land_Aboutus_Response
//	@Failure		400		{object}	object
//	@Failure		500		{object}	object
//	@Router			/aboutus [post]
func (s *AboutUsService) CreateAboutUs(request *models.Land_Aboutus_Request) (*models.Land_Aboutus_Response, error) {
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	if err := s.db.Table("land_aboutus").Create(request).Error; err != nil {
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

// GetAboutUsByID godoc
//	@Summary		Get an About Us component by ID
//	@Description	Get an About Us component by ID
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"About Us component ID"
//	@Success		200	{object}	models.Land_Aboutus_Response
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/aboutus/{id} [get]

func (s *AboutUsService) GetAboutUsByID(id int) (*models.Land_Aboutus_Response, error) {
	var request models.Land_Aboutus_Request

	if err := s.db.Table("land_aboutus").First(&request, id).Error; err != nil {
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

// UpdateAboutUs godoc
//	@Summary		Update an About Us component
//	@Description	Update an About Us component
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"About Us component ID"
//	@Param			request	body		models.Land_Aboutus_Request	true	"Updated About Us component data"
//	@Success		200		{object}	models.Land_Aboutus_Response
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/aboutus/{id} [put]

func (s AboutUsService) UpdateAboutUs(id int, updatedRequest *models.Land_Aboutus_Request) (*models.Land_Aboutus_Response, error) {
	var request models.Land_Aboutus_Request

	if err := s.db.Table("land_aboutus").First(&request, id).Error; err != nil {
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

	if err := s.db.Table("land_aboutus").Save(&request).Error; err != nil {
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

// DeleteAboutUs godoc
//	@Summary		Delete an About Us component
//	@Description	Delete an About Us component
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"About Us component ID"
//	@Success		204
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/aboutus/{id} [delete]

func (s *AboutUsService) DeleteAboutUs(id int) error {
	var request models.Land_Aboutus_Request

	if err := s.db.Table("land_aboutus").First(&request, id).Error; err != nil {
		return err
	}

	if err := s.db.Table("land_aboutus").Delete(&request).Error; err != nil {
		return err
	}

	return nil
}

// GetAllAboutUsHandler godoc
//	@Summary		Get all About Us components
//	@Description	Get all About Us components
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Land_Aboutus_Response
//	@Failure		500	{object}	object
//	@Router			/aboutus [get]

func (s AboutUsService) GetAllAboutUsHandler(c fiber.Ctx) error {
	response, err := s.GetAllAboutUs()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateAboutUsHandler godoc
//	@Summary		Create a new About Us component
//	@Description	Create a new About Us component
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.Land_Aboutus_Request	true	"About Us component data"
//	@Success		201		{object}	models.Land_Aboutus_Response
//	@Failure		400		{object}	object
//	@Failure		500		{object}	object
//	@Router			/aboutus [post]

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

// GetAboutUsByIDHandler godoc
//	@Summary		Get an About Us component by ID
//	@Description	Get an About Us component by ID
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"About Us component ID"
//	@Success		200	{object}	models.Land_Aboutus_Response
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/aboutus/{id} [get]

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

// UpdateAboutUsHandler godoc
//	@Summary		Update an About Us component
//	@Description	Update an About Us component
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"About Us component ID"
//	@Param			request	body		models.Land_Aboutus_Request	true	"Updated About Us component data"
//	@Success		200		{object}	models.Land_Aboutus_Response
//	@Failure		400		{object}	object
//	@Failure		404		{object}	object
//	@Failure		500		{object}	object
//	@Router			/aboutus/{id} [put]

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

// DeleteAboutUsHandler godoc
//	@Summary		Delete an About Us component
//	@Description	Delete an About Us component
//	@Tags			AboutUs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"About Us component ID"
//	@Success		204
//	@Failure		400	{object}	object
//	@Failure		404	{object}	object
//	@Failure		500	{object}	object
//	@Router			/aboutus/{id} [delete]

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
