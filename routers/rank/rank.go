package routers

import (
	"github.com/SymbioSix/ProgressieAPI/services/rank"
	"github.com/gofiber/fiber/v3"
)

type RankRouter struct {
	rankController rank.RankController
}

func NewRouteRankController(rankController rank.RankController) RankRouter {
	return RankRouter{rankController}
}

func (rank *RankRouter) RankRoutes(rg fiber.Router) {
	router := rg.Group("rank")

	router.Get("/", rank.rankController.GetUserRankBadges)
	router.Post("/set", rank.rankController.SetUserRankBadge)
}
