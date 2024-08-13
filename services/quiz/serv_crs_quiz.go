package services

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/quiz"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// CustomError to wrap status.StatusModel
type CustomError struct {
	StatusModel status.StatusModel
}

func (e *CustomError) Error() string {
	return e.StatusModel.Message
}

// QuizService struct
type QuizService struct {
	db *gorm.DB
}

// NewQuizService creates a new QuizService
func NewQuizService(db *gorm.DB) QuizService {
	return QuizService{db: db}
}

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
		return nil, &CustomError{
			StatusModel: status.StatusModel{
				Status:  "error",
				Message: "Failed to create quiz",
			},
		}
	}

	return quiz, nil
}

func (service *QuizService) GetQuiz(id string) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := service.db.First(&quiz, "quiz_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &CustomError{
				StatusModel: status.StatusModel{
					Status:  "error",
					Message: "Quiz not found",
				},
			}
		}
		return nil, &CustomError{
			StatusModel: status.StatusModel{
				Status:  "error",
				Message: "Error retrieving quiz",
			},
		}
	}

	return &quiz, nil
}

func (service *QuizService) UpdateQuiz(id string, req *models.Quiz) (*models.Quiz, error) {
	var quiz models.Quiz
	if err := service.db.First(&quiz, "quiz_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &CustomError{
				StatusModel: status.StatusModel{
					Status:  "error",
					Message: "Quiz not found",
				},
			}
		}
		return nil, &CustomError{
			StatusModel: status.StatusModel{
				Status:  "error",
				Message: "Error updating quiz",
			},
		}
	}

	quiz.SubcourseID = req.SubcourseID
	quiz.Description = req.Description
	quiz.UpdatedBy = req.UpdatedBy
	quiz.UpdatedAt = time.Now()
	quiz.Status = req.Status

	if err := service.db.Save(&quiz).Error; err != nil {
		return nil, &CustomError{
			StatusModel: status.StatusModel{
				Status:  "error",
				Message: "Error updating quiz",
			},
		}
	}

	return &quiz, nil
}

func (service *QuizService) DeleteQuiz(id string) error {
	if err := service.db.Delete(&models.Quiz{}, "quiz_id = ?", id).Error; err != nil {
		return &CustomError{
			StatusModel: status.StatusModel{
				Status:  "error",
				Message: "Error deleting quiz",
			},
		}
	}
	return nil
}

// Handlers

// CreateQuizHandler handles the creation of a new quiz
//
//	@Summary		Create a new quiz
//	@Description	Handle the request to create a new quiz
//	@Tags			Quiz Service
//	@Accept			json
//	@Produce		json
//	@Param			quiz	body		models.Quiz	true	"Quiz object"
//	@Success		201		{object}	models.Quiz
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/quizzes [post]
func (service QuizService) CreateQuizHandler(c fiber.Ctx) error {
	var req models.Quiz
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	quiz, err := service.CreateQuiz(&req)
	if err != nil {
		if customErr, ok := err.(*CustomError); ok {
			return c.Status(fiber.StatusInternalServerError).JSON(customErr.StatusModel)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(quiz)
}

// GetQuizHandler handles the retrieval of a quiz by ID
//
//	@Summary		Get a quiz by ID
//	@Description	Handle the request to get a quiz by ID
//	@Tags			Quiz Service
//	@Produce		json
//	@Param			id	path		string	true	"Quiz ID"
//	@Success		200	{object}	models.Quiz
//	@Failure		404	{object}	status.StatusModel
//	@Router			/quizzes/{id} [get]
func (service QuizService) GetQuizHandler(c fiber.Ctx) error {
	id := c.Params("id")
	quiz, err := service.GetQuiz(id)
	if err != nil {
		if customErr, ok := err.(*CustomError); ok {
			if customErr.StatusModel.Status == "error" && customErr.StatusModel.Message == "Quiz not found" {
				return c.Status(fiber.StatusNotFound).JSON(customErr.StatusModel)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(customErr.StatusModel)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(quiz)
}

// UpdateQuizHandler handles the update of an existing quiz
//
//	@Summary		Update a quiz by ID
//	@Description	Handle the request to update a quiz by ID
//	@Tags			Quiz Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Quiz ID"
//	@Param			quiz	body		models.Quiz	true	"Updated quiz object"
//	@Success		200		{object}	models.Quiz
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/quizzes/{id} [put]
func (service QuizService) UpdateQuizHandler(c fiber.Ctx) error {
	id := c.Params("id")
	var req models.Quiz
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	quiz, err := service.UpdateQuiz(id, &req)
	if err != nil {
		if customErr, ok := err.(*CustomError); ok {
			if customErr.StatusModel.Status == "error" && customErr.StatusModel.Message == "Quiz not found" {
				return c.Status(fiber.StatusNotFound).JSON(customErr.StatusModel)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(customErr.StatusModel)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.JSON(quiz)
}

// DeleteQuizHandler handles the deletion of a quiz by ID
//
//	@Summary		Delete a quiz by ID
//	@Description	Handle the request to delete a quiz by ID
//	@Tags			Quiz Service
//	@Produce		json
//	@Param			id	path	string	true	"Quiz ID"
//	@Success		204	"No Content"
//	@Failure		500	{object}	status.StatusModel
//	@Router			/quizzes/{id} [delete]
func (service QuizService) DeleteQuizHandler(c fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteQuiz(id); err != nil {
		if customErr, ok := err.(*CustomError); ok {
			return c.Status(fiber.StatusInternalServerError).JSON(customErr.StatusModel)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(status.StatusModel{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
