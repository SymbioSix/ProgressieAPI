package routers

import (
	quiz "github.com/SymbioSix/ProgressieAPI/services/quiz"
	"github.com/gofiber/fiber/v3"
)

type GetQuizResultRouter struct {
	getQuizResultController quiz.QuizResultService
}

func NewGetQuizResultRouter(getQuizResultController quiz.QuizResultService) GetQuizResultRouter {
	return GetQuizResultRouter{getQuizResultController}
}

func (gqrs *GetQuizResultRouter) GetQuizResultRouter(qrr fiber.Router) {
	app := qrr.Group("QuizResult")
	// Route to create a new QuizResult record
	app.Post("/quiz-results", gqrs.getQuizResultController.CreateQuizResultHandler)

	// Route to get a QuizResult record by QuizID and UserID
	app.Get("/quiz-results/:quiz_id/:user_id", gqrs.getQuizResultController.GetQuizResultHandler)

	// Route to update an existing QuizResult record
	app.Put("/quiz-results/:quiz_id/:user_id", gqrs.getQuizResultController.UpdateQuizResultHandler)

	// Route to delete a QuizResult record by QuizID and UserID
	app.Delete("/quiz-results/:quiz_id/:user_id", gqrs.getQuizResultController.DeleteQuizResultHandler)
}
