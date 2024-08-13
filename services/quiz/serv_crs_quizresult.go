package services

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/quiz"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

// QuizResultService struct
type QuizResultService struct {
	db *gorm.DB
}

// NewQuizResultService creates a new QuizResultService
func NewQuizResultService(db *gorm.DB) QuizResultService {
	return QuizResultService{db: db}
}

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

func (service *QuizResultService) GetQuizResult(quizID string, userID uuid.UUID) (*models.QuizResult, error) {
	var quizResult models.QuizResult
	if err := service.db.First(&quizResult, "quiz_id = ? AND user_id = ?", quizID, userID).Error; err != nil {
		return nil, err
	}

	return &quizResult, nil
}

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

func (service *QuizResultService) DeleteQuizResult(quizID string, userID uuid.UUID) error {
	if err := service.db.Delete(&models.QuizResult{}, "quiz_id = ? AND user_id = ?", quizID, userID).Error; err != nil {
		return err
	}
	return nil
}

// CreateQuizResultHandler handles the creation of a new quiz result
//
//	@Summary		Create a new QuizResult
//	@Description	Handles the creation of a new quiz result.
//	@Tags			QuizResult Service
//	@Accept			json
//	@Produce		json
//	@Param			quizResult	body		models.QuizResult	true	"QuizResult Data"
//	@Success		201			{object}	models.QuizResult
//	@Failure		400			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/quiz-results [post]
func (service QuizResultService) CreateQuizResultHandler(c fiber.Ctx) error {
	var req models.QuizResult
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	quizResult, err := service.CreateQuizResult(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(quizResult)
}

// GetQuizResultHandler handles the retrieval of a quiz result by QuizID and UserID
//
//	@Summary		Get a QuizResult
//	@Description	Handles the retrieval of a quiz result by QuizID and UserID.
//	@Tags			QuizResult Service
//	@Accept			json
//	@Produce		json
//	@Param			quiz_id	path		string	true	"Quiz ID"
//	@Param			user_id	path		string	true	"User ID"
//	@Success		200		{object}	models.QuizResult
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Router			/quiz-results/{quiz_id}/{user_id} [get]
func (service QuizResultService) GetQuizResultHandler(c fiber.Ctx) error {
	quizID := c.Params("quiz_id")
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: "Invalid user_id",
		})
	}

	quizResult, err := service.GetQuizResult(quizID, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(status.StatusModel{
				Status:  "error",
				Message: "Quiz result not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(quizResult)
}

// UpdateQuizResultHandler handles the update of an existing quiz result
//
//	@Summary		Update a QuizResult
//	@Description	Handles the update of an existing quiz result.
//	@Tags			QuizResult Service
//	@Accept			json
//	@Produce		json
//	@Param			quiz_id		path		string				true	"Quiz ID"
//	@Param			user_id		path		string				true	"User ID"
//	@Param			quizResult	body		models.QuizResult	true	"Updated QuizResult Data"
//	@Success		200			{object}	models.QuizResult
//	@Failure		400			{object}	status.StatusModel
//	@Failure		404			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/quiz-results/{quiz_id}/{user_id} [put]
func (service QuizResultService) UpdateQuizResultHandler(c fiber.Ctx) error {
	quizID := c.Params("quiz_id")
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: "Invalid user_id",
		})
	}

	var req models.QuizResult
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	quizResult, err := service.UpdateQuizResult(quizID, userID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(quizResult)
}

// DeleteQuizResultHandler handles the deletion of a quiz result by QuizID and UserID
//
//	@Summary		Delete a QuizResult
//	@Description	Handles the deletion of a quiz result by QuizID and UserID.
//	@Tags			QuizResult Service
//	@Accept			json
//	@Produce		json
//	@Param			quiz_id	path	string	true	"Quiz ID"
//	@Param			user_id	path	string	true	"User ID"
//	@Success		204
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/quiz-results/{quiz_id}/{user_id} [delete]
func (service QuizResultService) DeleteQuizResultHandler(c fiber.Ctx) error {
	quizID := c.Params("quiz_id")
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: "Invalid user_id",
		})
	}

	if err := service.DeleteQuizResult(quizID, userID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(status.StatusModel{
				Status:  "error",
				Message: "Quiz result not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
