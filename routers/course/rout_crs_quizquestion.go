package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/course"
	"github.com/gofiber/fiber/v3"
)

// SetupQuizQuestionRoutes sets up routes for QuizQuestion operations
func SetupQuizQuestionRoutes(app *fiber.App, service *services.QuizQuestionService) {
	// Route to create a new QuizQuestion record
	app.Post("/quiz-questions", service.CreateQuizQuestion)

	// Route to get a QuizQuestion record by ID
	app.Get("/quiz-questions/:id", service.GetQuizQuestionByID)

	// Route to update an existing QuizQuestion record
	app.Put("/quiz-questions/:id", service.UpdateQuizQuestion)

	// Route to delete a QuizQuestion record by ID
	app.Delete("/quiz-questions/:id", service.DeleteQuizQuestion)
}
