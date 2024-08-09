package routers

import (
	"github.com/SymbioSix/ProgressieAPI/services/leaderboard"
	"github.com/gofiber/fiber/v3"
)

type LeaderboardRouter struct {
	leaderboardRouter leaderboard.LeaderboardController
}

func NewRouteLeaderboardController(leaderboard leaderboard.LeaderboardController) LeaderboardRouter {
	return LeaderboardRouter{leaderboard}
}

func (leader *LeaderboardRouter) LeaderboardRoutes(rg fiber.Router) {
	router := rg.Group("leaderboard")

	router.Get("/ranks", leader.leaderboardRouter.Get100TopRanks)
	router.Get("/rank", leader.leaderboardRouter.Get100TopFilteredRanks)
}
