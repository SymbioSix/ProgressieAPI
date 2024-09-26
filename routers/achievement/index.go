package routers

import (
	achievement "github.com/SymbioSix/ProgressieAPI/services/achievement"
	"github.com/gofiber/fiber/v3"
)

type AchievementRouter struct {
	achievementRouter achievement.AchiallController
}

func NewRouteAchievementController(achievementRouter achievement.AchiallController) AchievementRouter {
	return AchievementRouter{achievementRouter}
}

func (achi *AchievementRouter) AchievementRoutes(rg fiber.Router) {
	router := rg.Group("achievements")

	router.Get("/", achi.achievementRouter.GetAllAchievement)
	router.Get("/user", achi.achievementRouter.GetAllAchievementByUserID)
}
