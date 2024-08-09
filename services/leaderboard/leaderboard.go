package leaderboard

import (
	user "github.com/SymbioSix/ProgressieAPI/models/auth"
	point "github.com/SymbioSix/ProgressieAPI/models/courses"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LeaderboardController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewLeaderboardController(DB *gorm.DB, API *utils.Client) LeaderboardController {
	return LeaderboardController{DB, API}
}

// Get100TopRanks godoc
//
//	@Summary		Get 100 Top Users With Highest Achieved Points
//	@Description	Get 100 Top Users With Highest Achieved Points
//	@Tags			Leaderboard Service
//	@Produce		json
//	@Success		200	{object}	[]user.UserModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/leaderboard/ranks [get]
func (lead *LeaderboardController) Get100TopRanks(c fiber.Ctx) error {
	var user []user.UserModel
	if res := lead.DB.Limit(100).Order(clause.OrderByColumn{Column: clause.Column{Table: "usr_user", Name: "total_point_achieved"}, Desc: true}).Find(&user); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// Get100TopFilteredRanks godoc
//
//	@Summary		Get 100 Top Users With Highest Achieved Points Based On Category Filter
//	@Description	Get 100 Top Users With Highest Achieved Points Based On Category Filter
//	@Tags			Leaderboard Service
//	@Produce		json
//	@Param			category query string true "filter by category"
//	@Success		200	{object}	[]user.UserModel
//	@Failute		400 {object} status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/leaderboard/rank [get]
func (lead *LeaderboardController) Get100TopFilteredRanks(c fiber.Ctx) error {
	// Ambil filter kategori dari query
	filter := c.Query("category")
	if filter == "" {
		stat := status.StatusModel{Status: "fail", Message: "Filter category is empty"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	// Ambil semua CourseID yang sesuai dengan kategori
	var courseIDs []string
	if err := lead.DB.Model(&point.CourseModel{}).
		Where("course_category = ?", filter).
		Pluck("course_id", &courseIDs).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	// Jika tidak ada course yang sesuai dengan kategori, langsung return
	if len(courseIDs) == 0 {
		return c.Status(fiber.StatusOK).JSON([]user.UserModel{})
	}

	// Ambil UserID dan total point yang di-achieve oleh masing-masing user berdasarkan CourseID yang telah di-filter
	var results []struct {
		UserID             uuid.UUID
		TotalPointAchieved int
	}
	if err := lead.DB.Model(&point.EnrollmentModel{}).
		Select("user_id, SUM(point_achieved) as total_point_achieved").
		Where("course_id IN ?", courseIDs).
		Group("user_id").
		Order("total_point_achieved DESC").
		Limit(100).
		Find(&results).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	// Ambil informasi user berdasarkan hasil di atas
	var topUsers []user.UserModel
	for _, result := range results {
		var user user.UserModel
		if err := lead.DB.Where("user_id = ?", result.UserID).First(&user).Error; err != nil {
			stat := status.StatusModel{Status: "fail", Message: err.Error()}
			return c.Status(fiber.StatusInternalServerError).JSON(stat)
		}
		user.TotalPointAchieved = result.TotalPointAchieved
		topUsers = append(topUsers, user)
	}

	// Return top 100 users dengan total point mereka
	return c.Status(fiber.StatusOK).JSON(topUsers)
}
