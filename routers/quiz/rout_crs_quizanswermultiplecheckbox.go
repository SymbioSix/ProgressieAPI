package routers

import (
	services "github.com/SymbioSix/ProgressieAPI/services/quiz"
	"github.com/gofiber/fiber/v3"
)

func SetupQuizMultCheckBoxRoutes(app *fiber.App, service *services.QuizMultCheckBoxService) {
	// Route untuk membuat record QuizMultCheckBox
	app.Post("/quiz-mult-checkboxes", service.CreateQuizMultCheckBox)

	// Route untuk mendapatkan record QuizMultCheckBox berdasarkan ID
	app.Get("/quiz-mult-checkboxes/:id", service.GetQuizMultCheckBoxByID)

	// Route untuk memperbarui record QuizMultCheckBox berdasarkan ID
	app.Put("/quiz-mult-checkboxes/:id", service.UpdateQuizMultCheckBox)

	// Route untuk menghapus record QuizMultCheckBox berdasarkan ID
	app.Delete("/quiz-mult-checkboxes/:id", service.DeleteQuizMultCheckBox)
}
