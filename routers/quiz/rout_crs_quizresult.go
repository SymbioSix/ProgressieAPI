package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/quiz"
	"github.com/gofiber/fiber/v3"
)

// SetupQuizResultRoutes sets up routes for QuizResult operations
func SetupQuizResultRoutes(app *fiber.App, service *services.QuizResultService) {
	// Route to create a new QuizResult record
	app.Post("/quiz-results", service.CreateQuizResultHandler)

	// Route to get a QuizResult record by QuizID and UserID
	app.Get("/quiz-results/:quiz_id/:user_id", service.GetQuizResultHandler)

	// Route to update an existing QuizResult record
	app.Put("/quiz-results/:quiz_id/:user_id", service.UpdateQuizResultHandler)

	// Route to delete a QuizResult record by QuizID and UserID
	app.Delete("/quiz-results/:quiz_id/:user_id", service.DeleteQuizResultHandler)
}
