package services

import (
	models "github.com/SymbioSix/ProgressieAPI/models/profile" // Sesuaikan dengan path project Anda
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/SymbioSix/ProgressieAPI/utils" // Sesuaikan dengan path project Anda
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AchiallController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewAchiALLController(DB *gorm.DB, API *utils.Client) AchiallController {
	return AchiallController{DB, API}
}

// Get All Achievement godoc
//
//	@Summary		Get All Acheivements
//	@Description	Get All Achievements without credentials
//	@Tags			Achievements Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.UserAchievement
//	@Failure		500	{object}	status.StatusModel
//	@Router			/achievements [get]
func (controller *AchiallController) GetAllAchievement(c fiber.Ctx) error {
	var achievements []models.UserAchievement

	if result := controller.DB.Find(&achievements); result.Error != nil {
		msg := status.StatusModel{Status: "fail", Message: "Could not fetch achievements"}
		return c.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return c.Status(fiber.StatusOK).JSON(achievements)
}

// Get All Achievement godoc
//
//	@Summary		Get All Acheivements for logged in user
//	@Description	Get All Achievements for logged in user with credentials
//	@Tags			Achievements Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.UserAchievement
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/achievements/user [get]
func (controller *AchiallController) GetAllAchievementByUserID(c fiber.Ctx) error {
	user, err := controller.API.Auth.GetUser() // Mendapatkan pengguna yang sedang login
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "Unauthorized: request",
		})
	}

	var achievements []models.UserAchievement
	// Menggunakan user.ID untuk mengambil pencapaian yang terkait
	if result := controller.DB.Where("user_id = ?", user.ID).Find(&achievements); result.Error != nil {
		msg := status.StatusModel{Status: "fail", Message: "Could not fetch achievements for the specified user"}
		return c.Status(fiber.StatusInternalServerError).JSON(msg)
	}

	return c.Status(fiber.StatusOK).JSON(achievements)
}
