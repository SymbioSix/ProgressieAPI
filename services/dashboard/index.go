package dashboard

import (
	"time"

	auth "github.com/SymbioSix/ProgressieAPI/models/auth"
	dashb "github.com/SymbioSix/ProgressieAPI/models/dashboard"
	profile "github.com/SymbioSix/ProgressieAPI/models/profile"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DashboardController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewDashboardController(DB *gorm.DB, API *utils.Client) DashboardController {
	return DashboardController{DB, API}
}

// TODO: Add Another Feature/Service Relatable to Dashboard Below this Line

// SidebarMapper godoc
//
//	@Summary		Get sidebar mapping for the user
//	@Description	Get sidebar mapping for the authenticated user based on their roles
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dashb.RoleSidebarResponse
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/dashboard/sidebar [get]
func (dash *DashboardController) SidebarMapperForAuthenticatedUser(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}
	var userRoleResponse []auth.UserRoleResponse
	if getUserRole := dash.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", getUser.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}

	var roleSidebarResponse []dashb.RoleSidebarResponse
	for _, v := range userRoleResponse {
		if getRoleSidebarFromAuthenticatedUser := dash.DB.Table("usr_rolesidebar").Preload(clause.Associations).Find(&roleSidebarResponse, "role_id = ?", v.RoleID); getRoleSidebarFromAuthenticatedUser.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getRoleSidebarFromAuthenticatedUser.Error.Error()})
		}

	}

	return c.Status(fiber.StatusOK).JSON(roleSidebarResponse)
}

// GetUserProfile godoc
//
//	@Summary		Get user profile
//	@Description	Get the profile of the authenticated user
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	auth.UserRoleResponse
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/dashboard/profile [get]
func (dash *DashboardController) GetUserProfile(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}
	var userRoleResponse auth.UserRoleResponse
	if getUserRole := dash.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", getUser.ID); getUserRole.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(userRoleResponse)
}

// UpdateUserProfile godoc
//
//	@Summary		Update user profile
//	@Description	Update the profile of the authenticated user
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		auth.UpdateUserProfileRequest	true	"Update User Profile Request"
//	@Success		200		{object}	status.StatusModel
//	@Failure		401		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/dashboard/profile [put]
func (dash *DashboardController) UpdateUserProfile(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}

	var request auth.UpdateUserProfileRequest
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var updateUser auth.UserModel
	if res := dash.DB.Where(auth.UserModel{UserID: getUser.ID}).First(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	if request.FirstName != "" {
		updateUser.FirstName = request.FirstName
	}
	if request.LastName != "" {
		updateUser.LastName = request.LastName
	}
	if request.Email != "" {
		updateUser.Email = request.Email
	}
	if request.PhotoProfile != "" {
		updateUser.PhotoProfile = request.PhotoProfile
	}
	if request.TitleProfile != "" {
		updateUser.TitleProfile = request.TitleProfile
	}
	if request.PhoneNumber != "" {
		updateUser.PhoneNumber = request.PhoneNumber
	}
	if request.Description != "" {
		updateUser.Description = request.Description
	}
	if request.Gender != "" {
		updateUser.Gender = request.Gender
	}

	if res := dash.DB.Save(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Profile Updated Successfully!"})
}

// UpdateUserSkill godoc
//
//	@Summary		Update user skill
//	@Description	Update the skill of the authenticated user
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		profile.UpdateUserTitleSkillRequest	true	"Update User Skill Request"
//	@Success		200		{object}	status.StatusModel
//	@Failure		401		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/dashboard/skill [put]
func (dash *DashboardController) CreateOrUpdateUserSkill(c fiber.Ctx) error {
	getUser, err := dash.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}

	var request profile.UpdateUserTitleSkillRequest
	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var updateUser profile.UserTitleSkill
	if res := dash.DB.Where(auth.UserModel{UserID: getUser.ID}).First(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	if request.TitleSkill != "" {
		updateUser.TitleSkill = request.TitleSkill
	}
	if request.Subtitle != "" {
		updateUser.Subtitle = request.Subtitle
	}
	if request.UpdatedAt.String() != "" {
		updateUser.UpdatedAt = request.UpdatedAt
	}

	if res := dash.DB.Save(&updateUser); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Skills Updated Successfully!"})
}

// SoftDeleteUser godoc
//
//	@Summary		Soft Delete User
//	@Description	Soft Delete User (Can only be performed by user with at least Administrator role)
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of the user who wants to be deleted"
//	@Success		200	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/dashboard/{id}/soft [delete]
func (dash *DashboardController) SoftDeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	var now time.Time = time.Now()
	_, _, err := dash.API.Rest.From("auth.users").Update("deleted_at = "+now.String(), "", "exact").Eq("id", id).Execute()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	var user auth.UserModel
	if res := dash.DB.Table("usr_user").Where("user_id = ?", id).First(&user); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	user.Status = "Deactivated"

	if res := dash.DB.Table("usr_user").Save(&user); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User Is Successfully Soft Deleted!"})
}

// GetUserActivityChart godoc
//
//	@Summary		Get User Activity Chart
//	@Description	Get User Activity Chart
//	@Tags			Dashboard Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dashb.ActivityCount
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/dashboard/activity-chart [get]
func (dash *DashboardController) GetUserActivityChart(c fiber.Ctx) error {
	db := dash.DB
	user, err := dash.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Request"})
	}
	// Map to store activity count by day of the week
	activityByDay := map[string]int{
		"Monday":    0,
		"Tuesday":   0,
		"Wednesday": 0,
		"Thursday":  0,
		"Friday":    0,
		"Saturday":  0,
		"Sunday":    0,
	}

	// Loop over the last 7 days
	for i := 0; i < 7; i++ {
		dayTime := time.Now().AddDate(0, 0, -i)
		day := dayTime.Weekday().String()

		// Count completed TodoList items for the day
		var todoCount int64
		if res := db.Table("td_customtargetchecklist").
			Joins("JOIN td_customtarget ON td_customtargetchecklist.target_id = td_customtarget.target_id").
			Joins("JOIN usr_achievement ON td_customtarget.achievement_id = usr_achievement.achievement_id").
			Where("usr_achievement.user_id = ? AND DATE(td_customtargetchecklist.date_checked) = ?", user.ID, dayTime.Format("2006-01-02")).
			Count(&todoCount); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}

		// Count finished SubcourseReminders for the day
		var reminderCount int64
		if res := db.Table("td_subcoursereminder").
			Joins("JOIN crs_subcourseprogress ON td_subcoursereminder.subcourseprogress_id = crs_subcourseprogress.subcourseprogress_id").
			Joins("JOIN crs_subcourse ON crs_subcourseprogress.subcourse_id = crs_subcourse.subcourse_id").
			Joins("JOIN crs_course ON crs_subcourse.course_id = crs_course.course_id").
			Joins("JOIN crs_enrollment ON crs_course.course_id = crs_enrollment.course_id").
			Where("crs_enrollment.user_id = ? AND DATE(td_subcoursereminder.updated_at) = ? AND td_subcoursereminder.is_finished = ?", user.ID, dayTime.Format("2006-01-02"), true).
			Count(&reminderCount); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}

		// Sum the counts for the day of the week
		activityByDay[day] += int(todoCount + reminderCount)
	}

	// Prepare the final response array
	var results []dashb.ActivityCount
	for day, count := range activityByDay {
		results = append(results, dashb.ActivityCount{
			DayOfWeek: day,
			Count:     count,
		})
	}

	return c.Status(fiber.StatusOK).JSON(results)
}
