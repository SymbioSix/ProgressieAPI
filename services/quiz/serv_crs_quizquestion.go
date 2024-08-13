package services

import (
	"errors"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/quiz"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// QuizQuestionService defines methods to handle QuizQuestion operations
type QuizQuestionService struct {
	DB *gorm.DB
}

// NewQuizQuestionService creates a new instance of QuizQuestionService
func NewQuizQuestionService(db *gorm.DB) QuizQuestionService {
	return QuizQuestionService{DB: db}
}

// CreateQuizQuestion creates a new QuizQuestion record
//	@Summary		Create a new QuizQuestion
//	@Description	Create a new QuizQuestion record
//	@Tags			QuizQuestion Service
//	@Accept			json
//	@Produce		json
//	@Param			quizQuestion	body		models.QuizQuestion	true	"QuizQuestion data"
//	@Success		201				{object}	models.QuizQuestion
//	@Failure		400				{object}	status.StatusModel
//	@Failure		500				{object}	status.StatusModel
//	@Router			/quiz-questions [post]
func (s QuizQuestionService) CreateQuizQuestion(c fiber.Ctx) error {
	var question models.QuizQuestion
	if err := c.Bind().JSON(&question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}
	question.CreatedAt = time.Now()
	question.UpdatedAt = time.Now()
	if err := s.DB.Create(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Failed to create quiz question",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(question)
}

// GetQuizQuestionByID retrieves a QuizQuestion record by ID
//	@Summary		Get a QuizQuestion by ID
//	@Description	Retrieve a QuizQuestion record by its ID
//	@Tags			QuizQuestion Service
//	@Produce		json
//	@Param			id	path		string	true	"QuizQuestion ID"
//	@Success		200	{object}	models.QuizQuestion
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/quiz-questions/{id} [get]
func (s QuizQuestionService) GetQuizQuestionByID(c fiber.Ctx) error {
	id := c.Params("id")
	var question models.QuizQuestion
	if err := s.DB.Where("quizquestion_id = ?", id).First(&question).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(status.StatusModel{
				Status:  "error",
				Message: "Quiz question not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Error retrieving quiz question",
		})
	}
	return c.JSON(question)
}

// UpdateQuizQuestion updates an existing QuizQuestion record
//	@Summary		Update an existing QuizQuestion
//	@Description	Update a QuizQuestion record by its ID
//	@Tags			QuizQuestion Service
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string				true	"QuizQuestion ID"
//	@Param			quizQuestion	body		models.QuizQuestion	true	"Updated QuizQuestion data"
//	@Success		200				{object}	models.QuizQuestion
//	@Failure		400				{object}	status.StatusModel
//	@Failure		500				{object}	status.StatusModel
//	@Router			/quiz-questions/{id} [put]
func (s QuizQuestionService) UpdateQuizQuestion(c fiber.Ctx) error {
	id := c.Params("id")
	var question models.QuizQuestion
	if err := c.Bind().JSON(&question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}
	question.QuizQuestionID = id
	question.UpdatedAt = time.Now()
	if err := s.DB.Save(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Error updating quiz question",
		})
	}
	return c.JSON(question)
}

// DeleteQuizQuestion deletes a QuizQuestion record by ID
//	@Summary		Delete a QuizQuestion by ID
//	@Description	Delete a QuizQuestion record by its ID
//	@Tags			QuizQuestion Service
//	@Param			id	path	string	true	"QuizQuestion ID"
//	@Success		204	"No Content"
//	@Failure		500	{object}	status.StatusModel
//	@Router			/quiz-questions/{id} [delete]
func (s QuizQuestionService) DeleteQuizQuestion(c fiber.Ctx) error {
	id := c.Params("id")
	if err := s.DB.Where("quizquestion_id = ?", id).Delete(&models.QuizQuestion{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Error deleting quiz question",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
