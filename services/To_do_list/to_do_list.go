package To_do_list

import (
	"time"

	todoroki "github.com/SymbioSix/ProgressieAPI/models/Todo"
	profile "github.com/SymbioSix/ProgressieAPI/models/profile"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
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
	Todotg     []todoroki.TdCustomTargetResponse      `json:"custom_target"`
	Todosubcrs []todoroki.TdSubcourseReminderResponse `json:"subcourse_reminder"`
}

func NewTodoController(DB *gorm.DB, API *utils.Client) ToDoListController {
	return ToDoListController{DB, API}
}

// get all td_customtarget and get all td_subcoursereminder

// Getalltodo godoc
//
//	@Summary		Get all custom targets and subcourse reminders
//	@Description	Get all td_customtarget and td_subcoursereminder records
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Todoall
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/todos [get]
func (todall *ToDoListController) Getalltodo(c fiber.Ctx) error {
	var tdAll Todoall
	var tdCustom []todoroki.TdCustomTarget
	var tdSubcourse []todoroki.TdSubcourseReminder
	if res := todall.DB.Table("td_customtarget").Find(&tdCustom); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	if res := todall.DB.Table("td_subcoursereminder").Find(&tdSubcourse); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	for k, v := range tdCustom {
		tdAll.Todotg[k].AchievementID = v.AchievementID
		tdAll.Todotg[k].TargetID = v.TargetID
		tdAll.Todotg[k].TargetTitle = v.TargetTitle
		tdAll.Todotg[k].TargetSubtitle = v.TargetSubtitle
		tdAll.Todotg[k].DailyClockReminder = v.DailyClockReminder
		tdAll.Todotg[k].DueAt = v.DueAt
	}
	for k, v := range tdSubcourse {
		tdAll.Todosubcrs[k].ReminderID = v.ReminderID
		tdAll.Todosubcrs[k].SubcourseprogressID = v.SubcourseprogressID
		tdAll.Todosubcrs[k].ReminderTitle = v.ReminderTitle
		tdAll.Todosubcrs[k].ReminderTime = v.ReminderTime
		tdAll.Todosubcrs[k].StartDate = v.StartDate
		tdAll.Todosubcrs[k].IsFinished = v.IsFinished
	}
	return c.Status(fiber.StatusOK).JSON(tdAll)
}

// all function subcourse reminder
// TdSubcourseReminder godoc
//
//	@Summary		Get all subcourse reminders
//	@Description	Get all td_subcoursereminder records
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]todoroki.TdSubcourseReminder
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/subcourse_reminders [get]
func (td *ToDoListController) TdSubcourseReminder(c fiber.Ctx) error {
	var tdo []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(tdo)
}

// userid todoall
// TdAllbyuserID godoc
//
//	@Summary		Get all todos by user ID
//	@Description	Get all td_customtarget and td_subcoursereminder records for a specific user
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Todoall
//	@Failure		401	{object}	status.StatusModel
//	@Router			/todo/todos/user [get]
func (td *ToDoListController) TdAllbyuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdAll Todoall
	var tdCustom []todoroki.TdCustomTarget
	var tdSubcourse []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_customtarget").Find(&tdCustom, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	if res := td.DB.Table("td_subcoursereminder").Find(&tdSubcourse); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	for k, v := range tdCustom {
		tdAll.Todotg[k].AchievementID = v.AchievementID
		tdAll.Todotg[k].TargetID = v.TargetID
		tdAll.Todotg[k].TargetTitle = v.TargetTitle
		tdAll.Todotg[k].TargetSubtitle = v.TargetSubtitle
		tdAll.Todotg[k].DailyClockReminder = v.DailyClockReminder
		tdAll.Todotg[k].DueAt = v.DueAt
	}
	for k, v := range tdSubcourse {
		tdAll.Todosubcrs[k].ReminderID = v.ReminderID
		tdAll.Todosubcrs[k].SubcourseprogressID = v.SubcourseprogressID
		tdAll.Todosubcrs[k].ReminderTitle = v.ReminderTitle
		tdAll.Todosubcrs[k].ReminderTime = v.ReminderTime
		tdAll.Todosubcrs[k].StartDate = v.StartDate
		tdAll.Todosubcrs[k].IsFinished = v.IsFinished
	}
	return c.Status(fiber.StatusOK).JSON(tdAll)
}

// TdSubcourseReminderbyuserID godoc
//
//	@Summary		Get subcourse reminders by user ID
//	@Description	Get td_subcoursereminder records for a specific user
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]todoroki.TdSubcourseReminder
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/subcourse_reminders/user [get]
func (td *ToDoListController) TdSubcourseReminderbyuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdo []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").Find(&tdo, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(tdo)
}

// AutoFinishSubcourseReminders godoc
//
//	@Summary		Auto finish subcourse reminders
//	@Description	Auto finish td_subcoursereminder records if their subcourse is finished
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/subcourse_reminders/auto_finish [post]
func (td *ToDoListController) AutoFinishSubcourseReminders(c fiber.Ctx) error {
	//TODO: STILL IN DEVELOPMENT STATE
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

// GetSubcoursesNotSelected godoc
//
//	@Summary		Get subcourses not selected
//	@Description	Get subcourses that have not been selected by the user
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]todoroki.TdSubcourseReminder
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/subcourses/not_selected [get]
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
	return c.Status(fiber.StatusOK).JSON(subcourses)
}

// GetSelectedSubcourses retrieves subcourses that have been selected
// GetSelectedSubcourses godoc
//
//	@Summary		Get selected subcourses
//	@Description	Get subcourses that have been selected by the user
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]todoroki.TdSubcourseReminder
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/subcourses/selected [get]
func (td *ToDoListController) GetSelectedSubcourses(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized"})
	}

	userID := c.Params("userID")
	var tdo *todoroki.TdSubcourseReminder
	var reminders []todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		if res := td.DB.Table("td_subcourse_reminders").Where("user_id = ?", userID).Find(&reminders); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}
	}
	return c.Status(fiber.StatusOK).JSON(reminders)
}

// DeleteSelectedSubcourse godoc
//
//	@Summary		Delete a selected subcourse reminder
//	@Description	Delete a subcourse reminder for the authenticated user by reminder ID
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Param			reminderID	path		string	true	"Reminder ID"
//	@Success		200			{object}	status.StatusModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/todo/subcourse_reminders/{reminderID} [delete]
func (td *ToDoListController) DeleteSelectedSubcourse(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized"})
	}

	reminderID := c.Params("reminderID")
	var tdo *todoroki.TdSubcourseReminder
	if res := td.DB.Table("td_subcoursereminder").First(&tdo, "user_id = ?", user.ID); res.Error != nil {
		if res := td.DB.Table("td_subcoursereminder").Where("reminder_id = ?", reminderID).Delete(&todoroki.TdSubcourseReminder{}); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Subcourse reminder deleted successfully"})
}

// SaveReminder saves the reminder with specified date and time
// SaveReminder godoc
//
//	@Summary		Save a subcourse reminder
//	@Description	Save a subcourse reminder with specified date and time
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		todoroki.RequestTdSubcourseReminder	true	"Request body"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		401		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/todo/subcourse_reminders [post]
func (td *ToDoListController) SaveReminder(c fiber.Ctx) error {
	//TODO: HOW COULD THE FRONTEND GET THE SUBCOURSEPROGRESSID?
	var req todoroki.RequestTdSubcourseReminder
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	reminder := todoroki.TdSubcourseReminder{
		SubcourseprogressID: req.SubcourseprogressID,
		ReminderTitle:       req.ReminderTitle,
		ReminderTime:        req.ReminderTime,
		StartDate:           req.StartDate,
		Type:                "SubCourse",
		IsFinished:          req.IsFinished,
		CreatedBy:           user.ID.String(),
		CreatedAt:           time.Now(),
	}

	if res := td.DB.Save(&reminder); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Reminder saved successfully"})
}

// FUNCTION TO DO LIST CUSTOM TARGET
// TdCustomTarget godoc
//
//	@Summary		Get all custom targets
//	@Description	Get all td_customtarget records
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]todoroki.TdCustomTarget
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/custom_targets [get]
func (td *ToDoListController) TdCustomTarget(c fiber.Ctx) error {
	var tdo []todoroki.TdCustomTarget
	if res := td.DB.Table("td_customtarget").Find(&tdo); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(tdo)
}

// GetTdCustomTargetByID retrieves a specific td_customtarget record by target ID
// GetTdCustomTargetByuserID godoc
//
//	@Summary		Get custom targets by user ID
//	@Description	Get td_customtarget records for the authenticated user
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	status.ParseCustomTargetFromAchievement
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/todo/custom_targets/user [get]
func (td *ToDoListController) GetTdCustomTargetByuserID(c fiber.Ctx) error {
	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tdo profile.UserAchievement
	if res := td.DB.Table("usr_achievement").Preload("CustomTargets").Find(&tdo, "user_id = ?", user.ID); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	var tdos status.ParseCustomTargetFromAchievement
	tdos.CustomTargets = tdo.CustomTargets
	return c.Status(fiber.StatusOK).JSON(tdos)
}

// SaveTdCustomTarget saves a custom target with specified time, title, and description for 30 days. Also auto-create related achievement for that custom target
// SaveTdCustomTarget godoc
//
//	@Summary		Save a custom target
//	@Description	Save a custom target with specified time, title, and description for 30 days
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		todoroki.RequestCustomTarget	true	"Request body"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/todo/custom_targets [post]
func (td *ToDoListController) SaveTdCustomTarget(c fiber.Ctx) error {
	var req todoroki.RequestCustomTarget
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := td.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	ach_id := uuid.New()
	achievement := profile.UserAchievement{
		UserID:                 user.ID,
		AchievementID:          ach_id,
		AchievementTitle:       req.TargetTitle,
		AchievementDescription: req.TargetSubtitle,
		AchievementCategory:    "Custom Target Achievement 30 Days OnGoing",
		IsAchieved:             false,
	}
	if res := td.DB.Save(&achievement); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	// Calculate the due date 30 days from now
	dueDate := time.Now().AddDate(0, 0, 30)
	target := todoroki.TdCustomTarget{
		AchievementID:      ach_id,
		TargetTitle:        req.TargetTitle,
		TargetSubtitle:     req.TargetSubtitle,
		DailyClockReminder: req.DailyReminder,
		Type:               "Custom", // Set the appropriate type here
		CreatedAt:          time.Now(),
		DueAt:              dueDate,
	}

	if res := td.DB.Save(&target); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Custom target saved successfully"})
}

// UpdateChecklist updates the checklist status for a custom target
// UpdateChecklist godoc
//
//	@Summary		Update checklist
//	@Description	Update the checklist status for a custom target
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Param			targetID	path		string								true	"Target ID"
//	@Param			request		body		todoroki.UpdateChecklistDateRequest	true	"Request body"
//	@Success		200			{object}	status.StatusModel
//	@Failure		400			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/todo/custom_targets/{targetID}/checklist [put]
func (td *ToDoListController) UpdateChecklist(c fiber.Ctx) error {
	targetID := c.Params("targetID")
	var req todoroki.UpdateChecklistDateRequest

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
// CheckCustomTargetProgressForAchievement godoc
//
//	@Summary		Check custom target progress for achievement
//	@Description	Check if a target has been checked for 30 consecutive days and award an achievement
//	@Tags			TodoList Service
//	@Accept			json
//	@Produce		json
//	@Param			targetID	path		string	true	"Target ID"
//	@Success		200			{object}	status.StatusModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/todo/custom_targets/{targetID}/check_progress [put]
func (td *ToDoListController) CheckCustomTargetProgressForAchievement(c fiber.Ctx) error {
	targetID := c.Params("targetID")
	var ach todoroki.TdCustomTarget

	var checklists []todoroki.Checklist
	if res := td.DB.Table("td_checklists").Where("target_id = ?", targetID).Order("date_checked asc").Find(&checklists); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
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
			// Update IsAchieved value
			achievement := profile.UserAchievement{
				AchievementCategory: "Custom Target Achievement 30 Days Success",
				IsAchieved:          true,
			}

			// Save the new achievement record
			if res := td.DB.Table("usr_achievement").Where("achievement_id = ?", ach.AchievementID).Save(&achievement); res.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Achievement awarded"})
		}
	}

	if res := td.DB.Table("td_customtarget").Where("target_id = ?", targetID).First(&ach); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	// Create a new achievement record for failed attempt
	achievement := profile.UserAchievement{
		AchievementCategory: "Custom Target Achievement 30 Days Failed",
		IsAchieved:          false,
	}

	// Save the new achievement record for failure
	if res := td.DB.Table("usr_achievement").Where("achievement_id = ?", ach.AchievementID).Save(&achievement); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Achievement not awarded"})
}
