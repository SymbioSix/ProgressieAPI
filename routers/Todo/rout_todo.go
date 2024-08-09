package routers

import (
	"github.com/SymbioSix/ProgressieAPI/services/To_do_list"
	"github.com/gofiber/fiber/v3"
)

type SetupToDoListRoutes struct {
	gettodoController To_do_list.ToDoListController
}

func NewSetupToDoListRoutes(gettodoController To_do_list.ToDoListController) SetupToDoListRoutes {
	return SetupToDoListRoutes{gettodoController}
}

func (controller *SetupToDoListRoutes) GetSetupToDoListRoutes(app fiber.Router) {
	// Grouping all TodoList routes under /api/todo
	todo := app.Group("todo")

	// Route to get all custom targets and subcourse reminders
	todo.Get("/todos", controller.gettodoController.Getalltodo)

	// Route to get all subcourse reminders
	todo.Get("/subcourse_reminders", controller.gettodoController.TdSubcourseReminder)

	// Route to get all todos by user ID
	todo.Get("/todos/user", controller.gettodoController.TdAllbyuserID)

	// Route to get subcourse reminders by user ID
	todo.Get("/subcourse_reminders/user", controller.gettodoController.TdSubcourseReminderbyuserID)

	// Route to auto finish subcourse reminders
	todo.Post("/subcourse_reminders/auto_finish", controller.gettodoController.AutoFinishSubcourseReminders)

	// Route to get subcourses not selected
	todo.Get("/subcourses/not_selected", controller.gettodoController.GetSubcoursesNotSelected)

	// Route to get selected subcourses
	todo.Get("/subcourses/selected", controller.gettodoController.GetSelectedSubcourses)

	// Route to delete selected subcourse reminder
	todo.Delete("/subcourse_reminders/:reminderID", controller.gettodoController.DeleteSelectedSubcourse)

	// Route to save a subcourse reminder
	todo.Post("/subcourse_reminders", controller.gettodoController.SaveReminder)

	// Route to get all custom targets
	todo.Get("/custom_targets", controller.gettodoController.TdCustomTarget)

	// Route to get custom targets by user ID
	todo.Get("/custom_targets/user", controller.gettodoController.GetTdCustomTargetByuserID)

	// Route to save a custom target
	todo.Post("/custom_targets", controller.gettodoController.SaveTdCustomTarget)

	// Route to update checklist
	todo.Put("/custom_targets/:targetID/checklist", controller.gettodoController.UpdateChecklist)

	// Route to check custom target progress for achievement
	todo.Put("/custom_targets/:targetID/check_progress", controller.gettodoController.CheckCustomTargetProgressForAchievement)
}
