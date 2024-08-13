package routers

import (
	quiz "github.com/SymbioSix/ProgressieAPI/services/quiz"
	"github.com/gofiber/fiber/v3"
)

type GetQuizAMCRouter struct {
	getQuizAMCController quiz.QuizAnswerMultipleChoiceService
}

func NewGetQuizAMCRouter(getQuizAMCController quiz.QuizAnswerMultipleChoiceService) GetQuizAMCRouter {
	return GetQuizAMCRouter{getQuizAMCController}
}

func (gqAMCs *GetQuizAMCRouter) GetQuizAMCRouter(qAMC fiber.Router) {
	app := qAMC.Group("QuizAnswerMultipleChoice")

	app.Post("/quiz-answer-multiple-choices", gqAMCs.getQuizAMCController.CreateQuizAnswerMultipleChoice)

	// Route to get a QuizAnswerMultipleChoice record by ID
	app.Get("/quiz-answer-multiple-choices/:id", gqAMCs.getQuizAMCController.GetQuizAnswerMultipleChoiceByID)

	// Route to update an existing QuizAnswerMultipleChoice record
	app.Put("/quiz-answer-multiple-choices/:id", gqAMCs.getQuizAMCController.UpdateQuizAnswerMultipleChoice)

	// Route to delete a QuizAnswerMultipleChoice record by ID
	app.Delete("/quiz-answer-multiple-choices/:id", gqAMCs.getQuizAMCController.DeleteQuizAnswerMultipleChoice)
}
