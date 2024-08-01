package services

import (
	"errors"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/course"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// QuizAnswerMultipleChoiceService defines methods to handle QuizAnswerMultipleChoice operations
type QuizAnswerMultipleChoiceService struct {
	DB *gorm.DB
}

// NewQuizAnswerMultipleChoiceService creates a new instance of QuizAnswerMultipleChoiceService
func NewQuizAnswerMultipleChoiceService(db *gorm.DB) *QuizAnswerMultipleChoiceService {
	return &QuizAnswerMultipleChoiceService{DB: db}
}

// CreateQuizAnswerMultipleChoice creates a new QuizAnswerMultipleChoice record
func (s QuizAnswerMultipleChoiceService) CreateQuizAnswerMultipleChoice(c fiber.Ctx) error {
	var answer models.QuizAnswerMultipleChoice
	if err := c.Bind().JSON(&answer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	answer.CreatedAt = time.Now()
	answer.UpdatedAt = time.Now()
	if err := s.DB.Create(&answer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(answer)
}

// GetQuizAnswerMultipleChoiceByID retrieves a QuizAnswerMultipleChoice record by ID
func (s QuizAnswerMultipleChoiceService) GetQuizAnswerMultipleChoiceByID(c fiber.Ctx) error {
	id := c.Params("id")
	var answer models.QuizAnswerMultipleChoice
	if err := s.DB.Where("quizquestion_id = ?", id).First(&answer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(answer)
}

// UpdateQuizAnswerMultipleChoice updates an existing QuizAnswerMultipleChoice record
func (s QuizAnswerMultipleChoiceService) UpdateQuizAnswerMultipleChoice(c fiber.Ctx) error {
	id := c.Params("id")
	var answer models.QuizAnswerMultipleChoice
	if err := c.Bind().JSON(&answer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	answer.QuizQuestionID = id
	answer.UpdatedAt = time.Now()
	if err := s.DB.Save(&answer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(answer)
}

// DeleteQuizAnswerMultipleChoice deletes a QuizAnswerMultipleChoice record by ID
func (s QuizAnswerMultipleChoiceService) DeleteQuizAnswerMultipleChoice(c fiber.Ctx) error {
	id := c.Params("id")
	if err := s.DB.Where("quizquestion_id = ?", id).Delete(&models.QuizAnswerMultipleChoice{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
