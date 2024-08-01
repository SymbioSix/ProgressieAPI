package To_do_list

import (
	"time"

	todoroki "github.com/SymbioSix/ProgressieAPI/models/Todo"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ToDoListController struct {
	DB  *gorm.DB
	API *utils.Client
}
type Todoall struct {
	Todotg     todoroki.TdCustomTarget      `json:"Customremind"`
	Todosubcrs todoroki.TdSubcourseReminder `json:"SubCourseremind"`
}

func NewTodoController(DB *gorm.DB, API *utils.Client) ToDoListController {
	return ToDoListController{DB, API}
}

// get all td_customtarget and get all td_subcoursereminder
func (todall *ToDoListController) Getalltodo(c fiber.Ctx) error {
	var tdo []Todoall
	if res := todall.DB.Table("TdCustomTarget").Table("TdSubcourseReminder").Preload("TdCustomTarget").Preload("TdSubcourseReminder").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

//all function subcourse reminder

func (td *ToDoListController) TdSubcourseReminder(c fiber.Ctx) error {
	var tdo []todoroki.TdSubcourseReminder
	if res := td.DB.Table("TdSubcourseReminder").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(tdo), "data": tdo})
}

// userid todoall
func (td *ToDoListController) TdSubcourseReminderbyuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdo *Todoall
	if res := td.DB.Table("TdSubcourseReminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

func (td *ToDoListController) TdAllbyuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdo *todoroki.TdSubcourseReminder
	if res := td.DB.Table("TdSubcourseReminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

func (td *ToDoListController) AutoFinishSubcourseReminders(c fiber.Ctx) error {
	var reminders []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcourse_reminders").Find(&reminders); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	for _, reminder := range reminders {
		var progress todoroki.SubcourseProgress
		if res := td.DB.Table("subcourse_progresses").Where("subcourse_id = ? AND user_id = ?", reminder.SubcourseId, reminder.UserID).First(&progress); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}

		if progress.IsSubcourseFinished {
			reminder.IsFinished = true
			if res := td.DB.Save(&reminder); res.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Subcourse reminders updated successfully"})
}

//edit mode user can edit reminder time and day and save

// GetSubcoursesNotSelected retrieves subcourses that have not been selected
func (td *ToDoListController) GetSubcoursesNotSelected(c fiber.Ctx) error {
	var subcourses []todoroki.TdSubcourseReminder
	if res := td.DB.Table("subcourses").Where("id NOT IN (SELECT subcourse_id FROM td_subcourse_reminders)").Find(&subcourses); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": subcourses})
}

// GetSelectedSubcourses retrieves subcourses that have been selected
func (td *ToDoListController) GetSelectedSubcourses(c fiber.Ctx) error {
	userID := c.Params("userID")
	var reminders []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcourse_reminders").Where("user_id = ?", userID).Find(&reminders); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": reminders})
}

// DeleteSelectedSubcourse deletes a selected subcourse reminder
func (td *ToDoListController) DeleteSelectedSubcourse(c fiber.Ctx) error {
	reminderID := c.Params("reminderID")
	if res := td.DB.Table("td_subcourse_reminders").Where("reminder_id = ?", reminderID).Delete(&todoroki.TdSubcourseReminder{}); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Subcourse reminder deleted successfully"})
}

// SaveReminder saves the reminder with specified date and time
func (td *ToDoListController) SaveReminder(c fiber.Ctx) error {
	type RequestTdSubcourseReminder struct {
		ReminderID    string    `json:"reminder_id"`
		UserID        uuid.UUID `json:"user_id"`
		SubcourseID   string    `json:"subcourse_id"`
		ReminderTitle string    `json:"reminder_title"`
		Icon          string    `json:"icon"`
		ReminderTime  time.Time `json:"reminder_time"`
		StartDate     time.Time `json:"start_date"`
		IsFinished    bool      `json:"is_finished"`
	}

	var req RequestTdSubcourseReminder
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	reminder := todoroki.TdSubcourseReminder{
		ReminderID:    req.ReminderID,
		SubcourseId:   req.SubcourseID,
		UserID:        req.UserID,
		ReminderTitle: req.ReminderTitle,
		Icon:          req.Icon,
		ReminderTime:  req.ReminderTime,
		StartDate:     req.StartDate,
		IsFinished:    req.IsFinished,
		CreatedBy:     req.UserID.String(),
		CreatedAt:     time.Now(),
	}

	if res := td.DB.Save(&reminder); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Reminder saved successfully"})
}

//FUNCTION TO DO LIST CUSTOM TARGET

func (td *ToDoListController) TdCustomTarget(c fiber.Ctx) error {
	var tdo []todoroki.TdCustomTarget
	if res := td.DB.Table("todo_SubCourseReminder").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(tdo), "data": tdo})
}

// GetTdCustomTargetByID retrieves a specific td_customtarget record by target ID
func (td *ToDoListController) GetTdCustomTargetByID(c fiber.Ctx) error {
	targetID := c.Params("targetID")
	var tdo todoroki.TdCustomTarget
	if res := td.DB.Table("td_custom_targets").First(&tdo, "target_id = ?", targetID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

// SaveTdCustomTarget saves a custom target with specified time, title, and description for 30 days
func (td *ToDoListController) SaveTdCustomTarget(c fiber.Ctx) error {
	type Request struct {
		UserID         string    `json:"user_id"`
		TargetTitle    string    `json:"target_title"`
		TargetSubtitle string    `json:"target_subtitle"`
		TargetIcon     string    `json:"target_icon"`
		DailyReminder  time.Time `json:"daily_reminder"`
	}

	var req Request
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Calculate the due date 30 days from now
	dueDate := time.Now().AddDate(0, 0, 30)

	target := todoroki.TdCustomTarget{
		TargetID:           uuid.New().String(),
		UserID:             req.UserID,
		TargetTitle:        req.TargetTitle,
		TargetSubtitle:     req.TargetSubtitle,
		TargetIcon:         req.TargetIcon,
		DailyClockReminder: req.DailyReminder,
		Type:               "Custom", // Set the appropriate type here
		CreatedBy:          req.UserID,
		CreatedAt:          time.Now(),
		DueAt:              dueDate,
	}

	if res := td.DB.Save(&target); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Custom target saved successfully", "data": target})
}

// UpdateChecklist updates the checklist status for a custom target
func (td *ToDoListController) UpdateChecklist(c fiber.Ctx) error {
	targetID := c.Params("targetID")
	var req struct {
		DateChecked time.Time `json:"date_checked"`
	}

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var target todoroki.TdCustomTarget
	if res := td.DB.Table("td_custom_targets").First(&target, "target_id = ?", targetID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	// Create a new checklist entry
	checklist := todoroki.Checklist{
		ChecklistID: uuid.New().String(),
		TargetID:    targetID,
		DateChecked: req.DateChecked,
	}

	if res := td.DB.Save(&checklist); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Checklist updated successfully"})
}

// CheckAchievement checks if a target has been checked for 30 consecutive days and awards an achievement
func (td *ToDoListController) CheckAchievement(c fiber.Ctx) error {
	targetID := c.Params("targetID")

	var checklists []todoroki.Checklist
	if res := td.DB.Table("checklists").Where("target_id = ?", targetID).Order("date_checked asc").Find(&checklists); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	if len(checklists) < 30 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Not enough checklists to qualify for achievement"})
	}

	// Check if the checklists cover 30 consecutive days
	for i := 0; i <= len(checklists)-30; i++ {
		start := checklists[i].DateChecked
		end := checklists[i+29].DateChecked

		if end.Sub(start).Hours() <= 24*29 {
			// Award achievement
			achievement := todoroki.Achievement{
				AchievementID: uuid.New().String(),
				UserID:        checklists[i].UserID,
				TargetID:      targetID,
				Achievement:   "30 Day Checklist Completion",
				AwardedAt:     time.Now(),
			}

			if res := td.DB.Save(&achievement); res.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Achievement awarded", "data": achievement})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "No achievement awarded"})
}
