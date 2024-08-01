package services

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/course"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

// QuizResultService struct
type QuizResultService struct {
	db *gorm.DB
}

// NewQuizResultService creates a new QuizResultService
func NewQuizResultService(db *gorm.DB) *QuizResultService {
	return &QuizResultService{db: db}
}

// CreateQuizResult creates a new quiz result
func (service *QuizResultService) CreateQuizResult(req *models.QuizResult) (*models.QuizResult, error) {
	quizResult := &models.QuizResult{
		QuizID:       req.QuizID,
		UserID:       req.UserID,
		Progress:     req.Progress,
		HighestScore: req.HighestScore,
		LastScore:    req.LastScore,
		Status:       req.Status,
		CompletedAt:  time.Now(),
		UpdatedBy:    req.UpdatedBy,
		UpdatedAt:    time.Now(),
	}

	if err := service.db.Create(quizResult).Error; err != nil {
		return nil, err
	}

	return quizResult, nil
}

// GetQuizResult retrieves a quiz result by QuizID and UserID
func (service *QuizResultService) GetQuizResult(quizID string, userID uuid.UUID) (*models.QuizResult, error) {
	var quizResult models.QuizResult
	if err := service.db.First(&quizResult, "quiz_id = ? AND user_id = ?", quizID, userID).Error; err != nil {
		return nil, err
	}

	return &quizResult, nil
}

// UpdateQuizResult updates an existing quiz result
func (service *QuizResultService) UpdateQuizResult(quizID string, userID uuid.UUID, req *models.QuizResult) (*models.QuizResult, error) {
	var quizResult models.QuizResult
	if err := service.db.First(&quizResult, "quiz_id = ? AND user_id = ?", quizID, userID).Error; err != nil {
		return nil, err
	}

	quizResult.Progress = req.Progress
	quizResult.HighestScore = req.HighestScore
	quizResult.LastScore = req.LastScore
	quizResult.Status = req.Status
	quizResult.UpdatedBy = req.UpdatedBy
	quizResult.UpdatedAt = time.Now()

	if err := service.db.Save(&quizResult).Error; err != nil {
		return nil, err
	}

	return &quizResult, nil
}

// DeleteQuizResult deletes a quiz result by QuizID and UserID
func (service *QuizResultService) DeleteQuizResult(quizID string, userID uuid.UUID) error {
	if err := service.db.Delete(&models.QuizResult{}, "quiz_id = ? AND user_id = ?", quizID, userID).Error; err != nil {
		return err
	}
	return nil
}

// CreateQuizResultHandler handles the creation of a new quiz result
func (service QuizResultService) CreateQuizResultHandler(c fiber.Ctx) error {
	var req models.QuizResult
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	quizResult, err := service.CreateQuizResult(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(quizResult)
}

// GetQuizResultHandler handles the retrieval of a quiz result by QuizID and UserID
func (service QuizResultService) GetQuizResultHandler(c fiber.Ctx) error {
	quizID := c.Params("quiz_id")
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	quizResult, err := service.GetQuizResult(quizID, userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(quizResult)
}

// UpdateQuizResultHandler handles the update of an existing quiz result
func (service QuizResultService) UpdateQuizResultHandler(c fiber.Ctx) error {
	quizID := c.Params("quiz_id")
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	var req models.QuizResult
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	quizResult, err := service.UpdateQuizResult(quizID, userID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(quizResult)
}

// DeleteQuizResultHandler handles the deletion of a quiz result by QuizID and UserID
func (service QuizResultService) DeleteQuizResultHandler(c fiber.Ctx) error {
	quizID := c.Params("quiz_id")
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	if err := service.DeleteQuizResult(quizID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
