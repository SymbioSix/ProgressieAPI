package To_do_list

import (
	"time"

	todoroki "github.com/SymbioSix/ProgressieAPI/models/Todo"
	profile "github.com/SymbioSix/ProgressieAPI/models/profile"
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
	if res := todall.DB.Table("td_customtarget").Table("td_subcoursereminder").Preload("user_id").Preload("subcourse_id").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

//all function subcourse reminder

func (td *ToDoListController) TdSubcourseReminder(c fiber.Ctx) error {
	var tdo []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").Preload("subcourse_id").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(tdo), "data": tdo})
}

// userid todoall
func (td *ToDoListController) TdAllbyuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdo []Todoall
	if res := td.DB.Table("td_customtarget").Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

func (td *ToDoListController) TdSubcourseReminderbyuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdo *todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

func (td *ToDoListController) AutoFinishSubcourseReminders(c fiber.Ctx) error {
	var reminders []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").Find(&reminders); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	for _, reminder := range reminders {
		var progress todoroki.SubcourseProgress
		if res := td.DB.Table("crs_subcourseprogress").Where("subcourseprogress_id = ?", reminder.SubcourseprogressID).First(&progress); res.Error != nil {
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
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized"})
	}

	var subcourses []todoroki.TdSubcourseReminder
	var tdo *todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		if res := td.DB.Table("subcourses").Where("id NOT IN (SELECT subcourse_id FROM td_subcourse_reminders)").Find(&subcourses); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": subcourses})
}

// GetSelectedSubcourses retrieves subcourses that have been selected
func (td *ToDoListController) GetSelectedSubcourses(c fiber.Ctx) error {
<<<<<<< HEAD
    user, err := td.API.Auth.GetUser()
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized"})
    }

    userID := c.Params("userID")
	var tdo *todoroki.TdSubcourseReminder
    var reminders []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
    	if res := td.DB.Table("td_subcoursereminder").Where("user_id = ?", userID).Find(&reminders); res.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
    	}
=======
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized"})
>>>>>>> 5863920bb8ef2128b3539a48b457390de923f8f9
	}

	userID := c.Params("userID")
	var tdo *todoroki.TdSubcourseReminder
	var reminders []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		if res := td.DB.Table("td_subcourse_reminders").Where("user_id = ?", userID).Find(&reminders); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": reminders})
}

// DeleteSelectedSubcourse deletes a selected subcourse reminder
func (td *ToDoListController) DeleteSelectedSubcourse(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized"})
	}

	reminderID := c.Params("reminderID")
	var tdo *todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
<<<<<<< HEAD
    	if res := td.DB.Table("td_subcoursereminder").Where("reminder_id = ?", reminderID).Delete(&todoroki.TdSubcourseReminder{}); res.Error != nil {
       		 return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
    	}
=======
		if res := td.DB.Table("td_subcoursereminder").Where("reminder_id = ?", reminderID).Delete(&todoroki.TdSubcourseReminder{}); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}
>>>>>>> 5863920bb8ef2128b3539a48b457390de923f8f9
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Subcourse reminder deleted successfully"})
}

// SaveReminder saves the reminder with specified date and time
func (td *ToDoListController) SaveReminder(c fiber.Ctx) error {

	var req todoroki.RequestTdSubcourseReminder
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	reminder := todoroki.TdSubcourseReminder{
		ReminderID:          req.ReminderID,
		SubcourseprogressID: req.SubcourseprogressID,
		ReminderTitle:       req.ReminderTitle,
		Icon:                req.Icon,
		ReminderTime:        req.ReminderTime,
		StartDate:           req.StartDate,
		Type:                "Lock&autoonfinishedProgress",
		IsFinished:          req.IsFinished,
		CreatedBy:           user.ID.String(),
		CreatedAt:           time.Now(),
	}

	if res := td.DB.Save(&reminder); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Reminder saved successfully"})
}

//FUNCTION TO DO LIST CUSTOM TARGET

func (td *ToDoListController) TdCustomTarget(c fiber.Ctx) error {
	var tdo []todoroki.TdCustomTarget
	if res := td.DB.Table("td_customtarget").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(tdo), "data": tdo})
}

// GetTdCustomTargetByID retrieves a specific td_customtarget record by target ID
func (td *ToDoListController) GetTdCustomTargetByuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdo profile.UserAchievement
	if res := td.DB.Table("usr_achievement").Preload("CustomTargets").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": tdo})
}

// SaveTdCustomTarget saves a custom target with specified time, title, and description for 30 days
func (td *ToDoListController) SaveTdCustomTarget(c fiber.Ctx) error {

	var req todoroki.RequestCustomTarget
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Calculate the due date 30 days from now
	dueDate := time.Now().AddDate(0, 0, 30)
	target := todoroki.TdCustomTarget{
		AchievementID:      req.AchievementID,
		TargetTitle:        req.TargetTitle,
		TargetSubtitle:     req.TargetSubtitle,
		TargetIcon:         req.TargetIcon,
		DailyClockReminder: req.DailyReminder,
		Type:               "Custom", // Set the appropriate type here
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
	if res := td.DB.Table("td_customtarget").First(&target, "target_id = ?", targetID); res.Error != nil {
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

// CheckCustomTargetProgressForAchievement checks if a target has been checked for 30 consecutive days and awards an achievement
func (td *ToDoListController) CheckCustomTargetProgressForAchievement(c fiber.Ctx) error {
	targetID := c.Params("targetID")
	var ach todoroki.TdCustomTarget

	var checklists []todoroki.Checklist
	if res := td.DB.Table("td_checklists").Where("target_id = ?", targetID).Order("date_checked asc").Find(&checklists); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if len(checklists) < 30 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Not enough checklists to qualify for achievement"})
	}

	// Check if the checklists cover 30 consecutive days
	for i := 0; i <= len(checklists)-30; i++ {
		isConsecutive := true
		start := checklists[i].DateChecked

		// Check if the next 29 entries are consecutive days
		for j := i + 1; j < i+30; j++ {
			expectedDate := start.AddDate(0, 0, j-i)
			if !checklists[j].DateChecked.Equal(expectedDate) {
				isConsecutive = false
				break
			}
		}

		if res := td.DB.Table("td_customtarget").Where("target_id = ?", checklists[i].TargetID).First(&ach); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}

		if isConsecutive {
			// Create a new achievement record
			achievement := profile.UserAchievement{
				AchievementTitle:       ach.TargetTitle, // Replace with actual TargetTitle
				UserID:                 user.ID,
				AchievementIcon:        ach.TargetIcon,     // Replace with actual icon path
				AchievementDescription: ach.TargetSubtitle, // Replace with actual description
				AchievementCategory:    "Custom Target Achievement 30 Days",
				IsAchieved:             true,
			}

			// Save the new achievement record
			if res := td.DB.Table("usr_achievement").Save(&achievement); res.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Achievement awarded", "data": achievement})
		}
	}

	if res := td.DB.Table("td_customtarget").Where("target_id = ?", targetID).First(&ach); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	// Create a new achievement record for failed attempt
	achievement := profile.UserAchievement{
		AchievementTitle:       ach.TargetTitle,    // Replace with actual TargetTitle
		UserID:                 user.ID,            // Use a default user ID if needed
		AchievementIcon:        ach.TargetIcon,     // Replace with actual icon path
		AchievementDescription: ach.TargetSubtitle, // Replace with actual description
		AchievementCategory:    "Custom Target Achievement 30 Days failed",
		IsAchieved:             false,
	}

	// Save the new achievement record for failure
	if res := td.DB.Table("usr_achievement").Save(&achievement); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Achievement not awarded", "data": achievement})
}
