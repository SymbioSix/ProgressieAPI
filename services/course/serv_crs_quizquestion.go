package services

import (
	"errors"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/course"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// QuizQuestionService defines methods to handle QuizQuestion operations
type QuizQuestionService struct {
	DB *gorm.DB
}

// NewQuizQuestionService creates a new instance of QuizQuestionService
func NewQuizQuestionService(db *gorm.DB) *QuizQuestionService {
	return &QuizQuestionService{DB: db}
}

// CreateQuizQuestion creates a new QuizQuestion record
func (s QuizQuestionService) CreateQuizQuestion(c fiber.Ctx) error {
	var question models.QuizQuestion
	if err := c.Bind().JSON(&question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	question.CreatedAt = time.Now()
	question.UpdatedAt = time.Now()
	if err := s.DB.Create(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(question)
}

// GetQuizQuestionByID retrieves a QuizQuestion record by ID
func (s QuizQuestionService) GetQuizQuestionByID(c fiber.Ctx) error {
	id := c.Params("id")
	var question models.QuizQuestion
	if err := s.DB.Where("quizquestion_id = ?", id).First(&question).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(question)
}

// UpdateQuizQuestion updates an existing QuizQuestion record
func (s QuizQuestionService) UpdateQuizQuestion(c fiber.Ctx) error {
	id := c.Params("id")
	var question models.QuizQuestion
	if err := c.Bind().JSON(&question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	question.QuizQuestionID = id
	question.UpdatedAt = time.Now()
	if err := s.DB.Save(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(question)
}

// DeleteQuizQuestion deletes a QuizQuestion record by ID
func (s QuizQuestionService) DeleteQuizQuestion(c fiber.Ctx) error {
	id := c.Params("id")
	if err := s.DB.Where("quizquestion_id = ?", id).Delete(&models.QuizQuestion{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
