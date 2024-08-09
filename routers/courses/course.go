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

	router.Get("/", gcrs.getCourseController.GetAllCoursesOnly)
	router.Get("/all", gcrs.getCourseController.GetAllCoursesAndSubCourses)
	router.Get("/:courseid/subcourses", gcrs.getCourseController.GetSubCoursesByCourseID)
	router.Get("/:courseid/enrollment/status", gcrs.getCourseController.CheckEnrollStatus, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
	router.Post("/:courseid/enroll", gcrs.getCourseController.EnrollUserToACourse, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
	router.Get("/:courseid/enrollment/data", gcrs.getCourseController.GetEnrolledCourseData, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
	router.Put("/:courseid/enrollment/progress", gcrs.getCourseController.UpdateEnrollmentProgress, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
	router.Put("/:courseid/enrollment/point", gcrs.getCourseController.UpdateEnrollmentPoint, middleware.RestrictUnauthenticatedUser(), middleware.RestrictUserWithUnusualStatus())
}
