package services

import (
    "net/http"
    "github.com/gofiber/fiber/v3"
    "gorm.io/gorm"
    models "github.com/SymbioSix/ProgressieAPI/models/achievement" // Sesuaikan dengan path project Anda
    "github.com/SymbioSix/ProgressieAPI/utils" // Sesuaikan dengan path project Anda
)

type AchicrsController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewAchiCRSController(DB *gorm.DB, API *utils.Client) AchicrsController {
	return AchicrsController{DB, API}
}

// getachicourse handles fetching achievements related to a course
func (controller *AchicrsController) GetAchiCourse(c fiber.Ctx) error {
    var achievements []models.AchievementCrs // Assuming there's a model called Achievement
    courseId := c.Params("courseId")

    if result := controller.DB.Where("course_id = ?", courseId).Find(&achievements); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not fetch achievements for the course",
        })
    }

    return c.Status(http.StatusOK).JSON(achievements)
}

// getachisubcourse handles fetching achievements related to a subcourse
func (controller *AchicrsController) GetAchiSubCourse(c fiber.Ctx) error {
    var achievements []models.AchievementCrs // Assuming there's a model called Achievement
    subCourseId := c.Params("subCourseId")

    if result := controller.DB.Where("sub_course_id = ?", subCourseId).Find(&achievements); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not fetch achievements for the subcourse",
        })
    }

    return c.Status(http.StatusOK).JSON(achievements)
}

// getachireading handles fetching achievements related to a reading
func (controller *AchicrsController) GetAchiReading(c fiber.Ctx) error {
    var achievements []models.AchievementCrs // Assuming there's a model called Achievement
    readingId := c.Params("readingId")

    if result := controller.DB.Where("reading_id = ?", readingId).Find(&achievements); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not fetch achievements for the reading",
        })
    }

    return c.Status(http.StatusOK).JSON(achievements)
}

// getachi handles fetching a specific achievement
func (controller *AchicrsController) GetAchi(c fiber.Ctx) error {
    var achievement models.AchievementCrs // Assuming there's a model called Achievement
    achievementId := c.Params("achievementId")

    if result := controller.DB.First(&achievement, "id = ?", achievementId); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not fetch the achievement",
        })
    }

    return c.Status(http.StatusOK).JSON(achievement)
}

// getachipointbyrank handles fetching achievements based on a rank
func (controller *AchicrsController) GetAchiPointByRank(c fiber.Ctx) error {
    var achievements []models.AchievementCrs // Assuming there's a model called Achievement
    rank := c.Params("rank")

    if result := controller.DB.Where("rank = ?", rank).Find(&achievements); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not fetch achievements by rank",
        })
    }

    return c.Status(http.StatusOK).JSON(achievements)
}
