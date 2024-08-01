package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/course"
	"github.com/gofiber/fiber/v3"
)

// SetupQuizAnswerMultipleChoiceRoutes sets up routes for QuizAnswerMultipleChoice operations
func SetupQuizAnswerMultipleChoiceRoutes(app *fiber.App, service *services.QuizAnswerMultipleChoiceService) {
	// Route to create a new QuizAnswerMultipleChoice record
	app.Post("/quiz-answer-multiple-choices", service.CreateQuizAnswerMultipleChoice)

	// Route to get a QuizAnswerMultipleChoice record by ID
	app.Get("/quiz-answer-multiple-choices/:id", service.GetQuizAnswerMultipleChoiceByID)

	// Route to update an existing QuizAnswerMultipleChoice record
	app.Put("/quiz-answer-multiple-choices/:id", service.UpdateQuizAnswerMultipleChoice)

	// Route to delete a QuizAnswerMultipleChoice record by ID
	app.Delete("/quiz-answer-multiple-choices/:id", service.DeleteQuizAnswerMultipleChoice)
}
