package routers

import (
	"github.com/SymbioSix/ProgressieAPI/middleware"
	"github.com/SymbioSix/ProgressieAPI/services/courses"
	"github.com/gofiber/fiber/v3"
)

type GetCourseRouter struct {
	getCourseController courses.CourseController
}

func NewGetCourseRouter(getCourseController courses.CourseController) GetCourseRouter {
	return GetCourseRouter{getCourseController}
}

func (gcrs *GetCourseRouter) GetCourseRoutes(rg fiber.Router) {
	router := rg.Group("courses")
	// router.Use(middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())

	router.Get("/all", gcrs.getCourseController.GetAllCoursesAndSubCourses)
	router.Get("/all-course-only", gcrs.getCourseController.GetAllCoursesOnly)
	router.Get("/get-by-courseid/:courseid", gcrs.getCourseController.GetSubCoursesByCourseID)
	router.Get("/check-enrollment-status/:courseid", gcrs.getCourseController.CheckEnrollStatus, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
	router.Post("/enroll-user-to-a-course/:courseid", gcrs.getCourseController.EnrollUserToACourse, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
	router.Get("/get-user-enrolled-course/:courseid", gcrs.getCourseController.GetEnrolledCourseData, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
	router.Put("/update-enrollment-progress/:courseid", gcrs.getCourseController.UpdateEnrollmentProgress, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
}
