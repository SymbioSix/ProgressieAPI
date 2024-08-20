package services

import (
    "net/http"
    "github.com/gofiber/fiber/v3"
    "gorm.io/gorm"
    models "github.com/SymbioSix/ProgressieAPI/models/achievement" // Sesuaikan dengan path project Anda
    "github.com/SymbioSix/ProgressieAPI/utils" // Sesuaikan dengan path project Anda
)

type AchiallController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewAchiALLController(DB *gorm.DB, API *utils.Client) AchiallController {
	return AchiallController{DB, API}
}

// getallachievement handles fetching all achievements
func (controller *AchiallController) GetAllAchievement(c fiber.Ctx) error {
    var achievements []models.AchiAll

    if result := controller.DB.Find(&achievements); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not fetch achievements",
        })
    }

    return c.Status(http.StatusOK).JSON(achievements)
}

// getallachievementByUserID handles fetching all achievements by a specific user ID
func (controller *AchiallController) GetAllAchievementByUserID(c fiber.Ctx) error {
    user, err := controller.API.Auth.GetUser() // Mendapatkan pengguna yang sedang login
    if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status": "fail", 
            "message": "Unauthorized: " + err.Error(),
        })
	}
    
    var achievements []models.AchiAll
    // Menggunakan user.ID untuk mengambil pencapaian yang terkait
    if result := controller.DB.Where("user_id = ?", user.ID).Find(&achievements); result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": "Could not fetch achievements for the specified user",
        })
    }
    
    return c.Status(http.StatusOK).JSON(achievements)
}
