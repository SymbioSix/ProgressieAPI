package courses

import (
	"errors"
	"time"

	todo "github.com/SymbioSix/ProgressieAPI/models/Todo"
	models "github.com/SymbioSix/ProgressieAPI/models/courses"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewCourseController(DB *gorm.DB, API *utils.Client) CourseController {
	return CourseController{DB, API}
}

// GetAllCoursesOnly godoc
//
//	@Summary		Get all courses only
//	@Description	Get all courses only
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		[]models.CourseModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/courses/ [get]
func (crs *CourseController) GetAllCoursesOnly(c fiber.Ctx) error {
	var courses []models.CourseModel
	if res := crs.DB.Table("crs_course").Find(&courses); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(courses)
}

// GetAllCoursesAndSubCourses godoc
//
//	@Summary		Get all courses and their sub-courses
//	@Description	Get all courses and their sub-courses
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		[]models.CourseModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/courses/all [get]
func (crs *CourseController) GetAllCoursesAndSubCourses(c fiber.Ctx) error {
	var courses []models.CourseModel
	if res := crs.DB.
		Preload("SubCourses.VideoContent").
		Preload("SubCourses.ReadingContents.ReadingImages").
		Preload("SubCourses.Quizzes.Questions.Answers").
		Find(&courses); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(courses)
}

// GetSubCoursesByCourseID godoc
//
//	@Summary		Get sub-courses by course ID
//	@Description	Get sub-courses by course ID
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Param			courseid	path		string	true	"Course ID"
//	@Success		200			{object}	status.ParseSubcoursesFromCourseID
//	@Failure		500			{object}	status.StatusModel
//	@Router			/courses/{courseid}/subcourses [get]
func (crs *CourseController) GetSubCoursesByCourseID(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	var subCourses models.CourseModel
	if res := crs.DB.
		Preload("SubCourses.VideoContent").
		Preload("SubCourses.ReadingContents.ReadingImages").
		Preload("SubCourses.Quizzes.Questions.Answers").
		First(&subCourses, "course_id = ?", courseId); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	var subcourse status.ParseSubcoursesFromCourseID
	subcourse.Subcourses = subCourses.SubCourses
	return c.Status(fiber.StatusOK).JSON(subcourse)
}

// CheckEnrollStatus godoc
//
//	@Summary		Check user enrollment status for a specific course
//	@Description	Check user enrollment status for a specific course (Will return either, Unauthorized Access, Error, User Isn't Enrolled!, or User Is Enrolled!)
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Param			courseid	path		string	true	"Course ID"
//	@Success		200			{object}	status.StatusModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		403			{object}	status.StatusModel
//	@Failure		400			{object}	status.StatusModel
//	@Router			/courses/{courseid}/enrollment/status [get]
func (crs *CourseController) CheckEnrollStatus(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var enrollment models.EnrollmentModel
	if res := crs.DB.Table("crs_enrollment").Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).First(&enrollment); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "unenrolled", "message": "User Isn't Enrolled!"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "enrolled", "message": "User Is Enrolled!"})
}

// EnrollUserToACourse godoc
//
//	@Summary		Enroll a user to a course. Also auto-generate new subcourse progress data. This action might take a while because of nested for loops to iterate list of data
//	@Description	Enroll a user to a course. Also auto-generate new subcourse progress data. This action might take a while because of nested for loops to iterate list of data
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Param			courseid	path		string	true	"Course ID"
//	@Success		201			{object}	status.StatusModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		400			{object}	status.StatusModel
//	@Router			/courses/{courseid}/enroll [post]
func (crs *CourseController) EnrollUserToACourse(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var enrollment models.EnrollmentModel
	if res := crs.DB.Table("crs_enrollment").Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).First(&enrollment); res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}
		enrollment = models.EnrollmentModel{
			UserID:    user.ID,
			CourseID:  courseId,
			Progress:  0.0,
			Status:    "Registered",
			CreatedBy: "API SYSTEM",
			CreatedAt: time.Now(),
		}
		crs.DB.Table("crs_enrollment").Create(&enrollment)

		var subcourse []models.SubCourseModel
		if res := crs.DB.Table("crs_subcourse").Where(&models.SubCourseModel{CourseID: courseId}).Find(&subcourse); res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}

		for i, v := range subcourse {
			progressId := uuid.New()
			progress := models.SubcourseProgress{
				SubcourseprogressID: progressId,
				UserID:              user.ID,
				SubcourseID:         v.SubcourseID,
				IsVideoViewed:       false,
				IsSubcourseFinished: false,
				CreatedBy:           "SYSTEM",
				CreatedAt:           time.Now(),
			}
			if res := crs.DB.Table("crs_subcourseprogress").Create(&progress); res.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
			}
			for _, k := range subcourse[i].ReadingContents {
				progressReading := models.SubcourseProgressReading{
					SubcourseprogressID: progressId,
					SubcoursereadingID:  k.SubcoursereadingID,
					IsReaded:            false,
				}
				if res := crs.DB.Table("crs_subcourseprogressreading").Create(&progressReading); res.Error != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
				}
			}

			reminder := todo.TdSubcourseReminder{
				SubcourseprogressID: progressId,
				ReminderTitle:       "SubCourse Reminder: " + v.SubcourseName,
				VideoTitle:          v.VideoContent.VideoTitle,
				QuizTitle:           "SubCourse Quiz: " + v.SubcourseName,
				Type:                "SubCourse",
				IsFinished:          false,
			}
			if res := crs.DB.Table("td_subcoursereminder").Create(&reminder); res.Error != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
			}

		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "You are enrolled!"})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Already Enrolled!"})
}

// GetEnrolledCourseData godoc
//
//	@Summary		Get enrolled course data
//	@Description	Get enrolled course data
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Param			courseid	path		string	true	"Course ID"
//	@Success		200			{object}	models.EnrollmentModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		403			{object}	status.StatusModel
//	@Failure		400			{object}	status.StatusModel
//	@Router			/courses/{courseid}/enrollment/data [get]
func (crs *CourseController) GetEnrolledCourseData(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var enrollment models.EnrollmentModel
	if res := crs.DB.
		Preload("TheCourse.SubCourses.VideoContent").
		Preload("TheCourse.SubCourses.ReadingContents.ReadingImages").
		Preload("TheCourse.SubCourses.Quizzes.Questions.Answers").
		Preload("TheCourse.SubCourses.Progress.ReadingProgress").
		Preload("TheCourse.SubCourses.Progress.Reminders").
		Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).
		First(&enrollment); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "unenrolled", "message": "User Isn't Enrolled!"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(enrollment)
}

// UpdateEnrollmentProgress godoc
//
//	@Summary		Update enrollment progress
//	@Description	Update enrollment progress
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Param			courseid	path		string							true	"Course ID"
//	@Param			request		body		models.UpdateEnrollmentProgress	true	"Updated enrollment progress"
//	@Success		200			{object}	status.StatusModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		400			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/courses/{courseid}/enrollment/progress [patch]
func (crs *CourseController) UpdateEnrollmentProgress(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var updateRequest models.UpdateEnrollmentProgress
	if err := c.Bind().JSON(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var updateProgress models.EnrollmentModel

	if res := crs.DB.Table("crs_enrollment").Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).First(&updateProgress); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	updateProgress.Progress = updateRequest.Progress
	updateProgress.UpdatedBy = "API SYSTEM"
	updateProgress.UpdatedAt = time.Now()

	if res := crs.DB.Table("crs_enrollment").Save(&updateProgress); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Succeed updating enrollment progress"})
}

// UpdateEnrollmentPoint godoc
//
//	@Summary		Update enrollment point
//	@Description	Update enrollment point
//	@Tags			Courses Service
//	@Accept			json
//	@Produce		json
//	@Param			courseid	path		string							true	"Course ID"
//	@Param			request		body		models.UpdateEnrollmentPoint	true	"Updated enrollment point"
//	@Success		200			{object}	status.StatusModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		400			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/courses/{courseid}/enrollment/point [patch]
func (crs *CourseController) UpdateEnrollmentPoint(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var updateRequest models.UpdateEnrollmentPoint
	if err := c.Bind().JSON(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var updateProgress models.EnrollmentModel

	if res := crs.DB.Table("crs_enrollment").Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).First(&updateProgress); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	updateProgress.PointAchieved = updateProgress.PointAchieved + updateRequest.Point
	updateProgress.UpdatedBy = "API SYSTEM"
	updateProgress.UpdatedAt = time.Now()

	if res := crs.DB.Table("crs_enrollment").Save(&updateProgress); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Succeed updating enrollment point"})
}

// RESTRICTED BY ADMIN MIDDLEWARE

func (crs *CourseController) CreateCourse(c fiber.Ctx) {}

func (crs *CourseController) UpdateCourse(c fiber.Ctx) {}

func (crs *CourseController) DeleteCourse(c fiber.Ctx) {}
