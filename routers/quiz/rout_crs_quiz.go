package routers

import (
	quiz "github.com/SymbioSix/ProgressieAPI/services/quiz"
	"github.com/gofiber/fiber/v3"
)

type GetQuizRouter struct {
	getQuizController quiz.QuizService
}

func NewGetQuizRouter(getQuizController quiz.QuizService) GetQuizRouter {
	return GetQuizRouter{getQuizController}
}

func (gqs *GetQuizRouter) GetQuizRouter(qr fiber.Router) {
	app := qr.Group("quiz")

	// Route untuk membuat quiz
	app.Post("/quizzes", gqs.getQuizController.CreateQuizHandler)

	// Route untuk mendapatkan quiz berdasarkan ID
	app.Get("/quizzes/:id", gqs.getQuizController.GetQuizHandler)

	// Route untuk memperbarui quiz berdasarkan ID
	app.Put("/quizzes/:id", gqs.getQuizController.UpdateQuizHandler)

	// Route untuk menghapus quiz berdasarkan ID
	app.Delete("/quizzes/:id", gqs.getQuizController.DeleteQuizHandler)
}
