package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/course"
	"github.com/gofiber/fiber/v3"
)

func SetupQuizRoutes(app *fiber.App, quizService *services.QuizService) {
	// Route untuk membuat quiz
	app.Post("/quizzes", quizService.CreateQuizHandler)

	// Route untuk mendapatkan quiz berdasarkan ID
	app.Get("/quizzes/:id", quizService.GetQuizHandler)

	// Route untuk memperbarui quiz berdasarkan ID
	app.Put("/quizzes/:id", quizService.UpdateQuizHandler)

	// Route untuk menghapus quiz berdasarkan ID
	app.Delete("/quizzes/:id", quizService.DeleteQuizHandler)
}
