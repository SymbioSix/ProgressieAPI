package services

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/quiz"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// QuizService struct
type QuizService struct {
	db *gorm.DB
}

// NewQuizService creates a new QuizService
func NewQuizService(db *gorm.DB) *QuizService {
	return &QuizService{db: db}
}

// CreateQuiz creates a new quiz
func (service *QuizService) CreateQuiz(req *models.Quiz) (*models.Quiz, error) {
	quiz := &models.Quiz{
		QuizID:      req.QuizID,
		SubcourseID: req.SubcourseID,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		CreatedAt:   time.Now(),
		UpdatedBy:   req.UpdatedBy,
		UpdatedAt:   time.Now(),
		Status:      req.Status,
	}

	if err := service.db.Create(quiz).Error; err != nil {
		return nil, err
	}

	return quiz, nil
}

// GetQuiz retrieves a quiz by ID
func (service *QuizService) GetQuiz(id string) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := service.db.First(&quiz, "quiz_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &quiz, nil
}

// UpdateQuiz updates an existing quiz
func (service *QuizService) UpdateQuiz(id string, req *models.Quiz) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := service.db.First(&quiz, "quiz_id = ?", id).Error; err != nil {
		return nil, err
	}

	quiz.SubcourseID = req.SubcourseID
	quiz.Description = req.Description
	quiz.UpdatedBy = req.UpdatedBy
	quiz.UpdatedAt = time.Now()
	quiz.Status = req.Status

	if err := service.db.Save(&quiz).Error; err != nil {
		return nil, err
	}

	return &quiz, nil
}

// DeleteQuiz deletes a quiz by ID
func (service *QuizService) DeleteQuiz(id string) error {
	if err := service.db.Delete(&models.Quiz{}, "quiz_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// Handlers

// CreateQuizHandler handles the creation of a new quiz
func (service QuizService) CreateQuizHandler(c fiber.Ctx) error {
	var req models.Quiz
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	quiz, err := service.CreateQuiz(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(quiz)
}

// GetQuizHandler handles the retrieval of a quiz by ID
func (service QuizService) GetQuizHandler(c fiber.Ctx) error {
	id := c.Params("id")
	quiz, err := service.GetQuiz(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(quiz)
}

// UpdateQuizHandler handles the update of an existing quiz
func (service QuizService) UpdateQuizHandler(c fiber.Ctx) error {
	id := c.Params("id")
	var req models.Quiz
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	quiz, err := service.UpdateQuiz(id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(quiz)
}

// DeleteQuizHandler handles the deletion of a quiz by ID
func (service QuizService) DeleteQuizHandler(c fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteQuiz(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
