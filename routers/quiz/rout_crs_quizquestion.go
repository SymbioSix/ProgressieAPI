package routers

import (
	quiz "github.com/SymbioSix/ProgressieAPI/services/quiz"
	"github.com/gofiber/fiber/v3"
)

type GetQuizQuestionRouter struct {
	getQuizQuestionController quiz.QuizQuestionService
}

func NewGetQuizQuestionRouter(getQuizQuestionController quiz.QuizQuestionService) GetQuizQuestionRouter {
	return GetQuizQuestionRouter{getQuizQuestionController}
}

func (gqqs *GetQuizQuestionRouter) GetQuizQuestionRouter(qq fiber.Router) {
	app := qq.Group("QuizQuestion")
	// Route to create a new QuizQuestion record
	app.Post("/quiz-questions", gqqs.getQuizQuestionController.CreateQuizQuestion)

	// Route to get a QuizQuestion record by ID
	app.Get("/quiz-questions/:id", gqqs.getQuizQuestionController.GetQuizQuestionByID)

	// Route to update an existing QuizQuestion record
	app.Put("/quiz-questions/:id", gqqs.getQuizQuestionController.UpdateQuizQuestion)

	// Route to delete a QuizQuestion record by ID
	app.Delete("/quiz-questions/:id", gqqs.getQuizQuestionController.DeleteQuizQuestion)
}
