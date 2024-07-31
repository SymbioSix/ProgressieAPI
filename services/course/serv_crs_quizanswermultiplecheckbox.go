package services

import (
	"errors"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/course"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// QuizMultCheckBoxService defines methods to handle QuizMultCheckBox operations
type QuizMultCheckBoxService struct {
	DB *gorm.DB
}

// NewQuizMultCheckBoxService creates a new instance of QuizMultCheckBoxService
func NewQuizMultCheckBoxService(db *gorm.DB) *QuizMultCheckBoxService {
	return &QuizMultCheckBoxService{DB: db}
}

// CreateQuizMultCheckBox creates a new QuizMultCheckBox record
func (s QuizMultCheckBoxService) CreateQuizMultCheckBox(c fiber.Ctx) error {
	var checkbox models.QuizMultCheckBox
	if err := c.Bind().JSON(&checkbox); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	checkbox.CreatedAt = time.Now()
	checkbox.UpdatedAt = time.Now()
	if err := s.DB.Create(&checkbox).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(checkbox)
}

// GetQuizMultCheckBoxByID retrieves a QuizMultCheckBox record by ID
func (s QuizMultCheckBoxService) GetQuizMultCheckBoxByID(c fiber.Ctx) error {
	id := c.Params("id")
	var checkbox models.QuizMultCheckBox
	if err := s.DB.Where("quizquestion_id = ?", id).First(&checkbox).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(checkbox)
}

// UpdateQuizMultCheckBox updates an existing QuizMultCheckBox record
func (s QuizMultCheckBoxService) UpdateQuizMultCheckBox(c fiber.Ctx) error {
	id := c.Params("id")
	var checkbox models.QuizMultCheckBox
	if err := c.Bind().JSON(&checkbox); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	checkbox.QuizQuestionID = id
	checkbox.UpdatedAt = time.Now()
	if err := s.DB.Save(&checkbox).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(checkbox)
}

// DeleteQuizMultCheckBox deletes a QuizMultCheckBox record by ID
func (s QuizMultCheckBoxService) DeleteQuizMultCheckBox(c fiber.Ctx) error {
	id := c.Params("id")
	if err := s.DB.Where("quizquestion_id = ?", id).Delete(&models.QuizMultCheckBox{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
