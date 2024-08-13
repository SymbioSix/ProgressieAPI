package services

import (
	"errors"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/quiz"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// QuizAnswerMultipleChoiceService defines methods to handle QuizAnswerMultipleChoice operations
type QuizAnswerMultipleChoiceService struct {
	DB *gorm.DB
}

// NewQuizAnswerMultipleChoiceService creates a new instance of QuizAnswerMultipleChoiceService
func NewQuizAnswerMultipleChoiceService(db *gorm.DB) QuizAnswerMultipleChoiceService {
	return QuizAnswerMultipleChoiceService{DB: db}
}

// CreateQuizAnswerMultipleChoice creates a new QuizAnswerMultipleChoice record
//	@Summary		Create a new QuizAnswerMultipleChoice
//	@Description	Create a new QuizAnswerMultipleChoice record in the database
//	@Tags			QuizAnswerMultipleChoice Service
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.QuizAnswerMultipleChoice	true	"QuizAnswerMultipleChoice Object"
//	@Success		201		{object}	models.QuizAnswerMultipleChoice
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/quiz-answer-multiple-choices [post]
func (s QuizAnswerMultipleChoiceService) CreateQuizAnswerMultipleChoice(c fiber.Ctx) error {
	var answer models.QuizAnswerMultipleChoice
	if err := c.Bind().JSON(&answer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: "Invalid input: " + err.Error(),
		})
	}
	answer.CreatedAt = time.Now()
	answer.UpdatedAt = time.Now()
	if err := s.DB.Create(&answer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Failed to create record: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(answer)
}

// GetQuizAnswerMultipleChoiceByID retrieves a QuizAnswerMultipleChoice record by ID
//	@Summary		Get a QuizAnswerMultipleChoice by ID
//	@Description	Retrieve a QuizAnswerMultipleChoice record from the database by its ID
//	@Tags			QuizAnswerMultipleChoice Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"QuizQuestion ID"
//	@Success		200	{object}	models.QuizAnswerMultipleChoice
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/quiz-answer-multiple-choices/{id} [get]
func (s QuizAnswerMultipleChoiceService) GetQuizAnswerMultipleChoiceByID(c fiber.Ctx) error {
	id := c.Params("id")
	var answer models.QuizAnswerMultipleChoice
	if err := s.DB.Where("quizquestion_id = ?", id).First(&answer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(status.StatusModel{
				Status:  "error",
				Message: "Record not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Failed to retrieve record: " + err.Error(),
		})
	}
	return c.JSON(answer)
}

// UpdateQuizAnswerMultipleChoice updates an existing QuizAnswerMultipleChoice record
//	@Summary		Update a QuizAnswerMultipleChoice by ID
//	@Description	Update an existing QuizAnswerMultipleChoice record in the database by its ID
//	@Tags			QuizAnswerMultipleChoice Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string							true	"QuizQuestion ID"
//	@Param			body	body		models.QuizAnswerMultipleChoice	true	"QuizAnswerMultipleChoice Object"
//	@Success		200		{object}	models.QuizAnswerMultipleChoice
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/quiz-answer-multiple-choices/{id} [put]
func (s QuizAnswerMultipleChoiceService) UpdateQuizAnswerMultipleChoice(c fiber.Ctx) error {
	id := c.Params("id")
	var answer models.QuizAnswerMultipleChoice
	if err := c.Bind().JSON(&answer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: "Invalid input: " + err.Error(),
		})
	}
	answer.QuizQuestionID = id
	answer.UpdatedAt = time.Now()
	if err := s.DB.Save(&answer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Failed to update record: " + err.Error(),
		})
	}
	return c.JSON(answer)
}

// DeleteQuizAnswerMultipleChoice deletes a QuizAnswerMultipleChoice record by ID
//	@Summary		Delete a QuizAnswerMultipleChoice by ID
//	@Description	Delete a QuizAnswerMultipleChoice record from the database by its ID
//	@Tags			QuizAnswerMultipleChoice Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"QuizQuestion ID"
//	@Success		204	"No Content"
//	@Failure		500	{object}	status.StatusModel
//	@Router			/quiz-answer-multiple-choices/{id} [delete]
func (s QuizAnswerMultipleChoiceService) DeleteQuizAnswerMultipleChoice(c fiber.Ctx) error {
	id := c.Params("id")
	if err := s.DB.Where("quizquestion_id = ?", id).Delete(&models.QuizAnswerMultipleChoice{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: "Failed to delete record: " + err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
