package courses

import (
	"errors"
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/courses"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type CourseController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewCourseController(DB *gorm.DB, API *utils.Client) CourseController {
	return CourseController{DB, API}
}

func (crs *CourseController) GetAllCoursesOnly(c fiber.Ctx) error {
	var courses []models.CourseModel
	if res := crs.DB.Table("crs_course").Find(&courses); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "length": len(courses), "data": courses})
}

func (crs *CourseController) GetAllCoursesAndSubCourses(c fiber.Ctx) error {
	var courses []models.CourseModel
	if res := crs.DB.Table("crs_course").Table("crs_subcourse").Table("crs_subcoursevideo").Table("crs_subcoursereading").Table("crs_subcoursereadingimage").Preload("ReadingImages").Preload("VideoContent").Preload("ReadingContents").Preload("SubCourses").Find(&courses); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": courses})
}

func (crs *CourseController) GetSubCoursesByCourseID(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	var subCourses *models.CourseModel
	if res := crs.DB.Table("crs_course").Table("crs_subcourse").Table("crs_subcoursevideo").Table("crs_subcoursereading").Table("crs_subcoursereadingimage").Preload("ReadingImages").Preload("VideoContent").Preload("ReadingContents").Preload("SubCourses").First(&subCourses, "course_id = ?", courseId); res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": subCourses})
}

func (crs *CourseController) CheckEnrollStatus(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var enrollment *models.EnrollmentModel
	if res := crs.DB.Table("crs_enrollment").Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).First(&enrollment); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "unenrolled", "message": "User Isn't Enrolled!"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusContinue).JSON(fiber.Map{"status": "enrolled", "message": "User Is Enrolled!"})
}

func (crs *CourseController) EnrollUserToACourse(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var enrollment *models.EnrollmentModel
	if res := crs.DB.Table("crs_enrollment").Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).First(&enrollment); res.Error != nil || enrollment != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Already Enrolled!"})
	}

	enrollment = &models.EnrollmentModel{
		UserID:    user.ID,
		CourseID:  courseId,
		Progress:  0.0,
		Status:    "Registered",
		CreatedBy: "API SYSTEM",
		CreatedAt: time.Now(),
	}
	crs.DB.Table("crs_enrollment").Create(&enrollment)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "You are enrolled!"})
}

func (crs *CourseController) GetEnrolledCourseData(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var enrollment *models.EnrollmentModel
	if res := crs.DB.Table("crs_enrollment").Table("crs_course").Table("crs_subcourse").Table("crs_subcoursevideo").Table("crs_subcoursereading").Table("crs_subcoursereadingimage").Preload("ReadingImages").Preload("VideoContent").Preload("ReadingContents").Preload("SubCourses").Preload("TheCourse").Where(&models.EnrollmentModel{UserID: user.ID, CourseID: courseId}).First(&enrollment); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "unenrolled", "message": "User Isn't Enrolled!"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": res.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": enrollment})
}

func (crs *CourseController) UpdateEnrollmentProgress(c fiber.Ctx) error {
	courseId := c.Params("courseid")
	user, err := crs.API.Auth.GetUser()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Unauthorized Access"})
	}
	var updateRequest *models.UpdateEnrollmentProgress
	if err := c.Bind().JSON(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var updateProgress *models.EnrollmentModel

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

// RESTRICTED BY ADMIN MIDDLEWARE

func (crs *CourseController) CreateCourse(c fiber.Ctx) {}

func (crs *CourseController) UpdateCourse(c fiber.Ctx) {}

func (crs *CourseController) DeleteCourse(c fiber.Ctx) {}
