package rank

import (
	"strconv"
	"time"

	course "github.com/SymbioSix/ProgressieAPI/models/courses"
	rank "github.com/SymbioSix/ProgressieAPI/models/rank"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	system "github.com/SymbioSix/ProgressieAPI/models/system_parameter"
	"github.com/SymbioSix/ProgressieAPI/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type RankController struct {
	DB  *gorm.DB
	API *utils.Client
}

func NewRankController(DB *gorm.DB, API *utils.Client) RankController {
	return RankController{DB, API}
}

// GetUserRankBadges godoc
//
//	@Summary		Get User Rank Badges
//	@Description	Get User Rank Badges
//	@Tags			RankBadge Profile Service
//	@Produce		json
//	@Success		200	{object}	[]rank.UserRankBadges
//	@Failure		401	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/rank/ [get]
func (r *RankController) GetUserRankBadges(c fiber.Ctx) error {
	user, err := r.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Unauthorized"}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}
	var ranks []rank.UserRankBadges
	if res := r.DB.Where(rank.UserRankBadges{UserID: user.ID}).Preload("RankData").Find(&ranks); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	return c.Status(fiber.StatusOK).JSON(ranks)
}

// SetUserRankBadge godoc
//
//	@Summary		Set User Rank Badge
//	@Description	Set User Rank badge
//	@Tags			RankBadge Profile Service
//	@Produce		json
//	@Param			type		query		string	true	"filter by badge type option: Beginner or Growth or Mastery or Enlightened"
//	@Param			category	query		string	true	"filter by category option: Financial or Personal or Social"
//	@Success		200			{object}	status.StatusModel
//	@Failure		400			{object}	status.StatusModel
//	@Failure		401			{object}	status.StatusModel
//	@Failure		500			{object}	status.StatusModel
//	@Router			/rank/set [post]
func (r *RankController) SetUserRankBadge(c fiber.Ctx) error {
	types := c.Query("type")
	if types == "" {
		stat := status.StatusModel{Status: "fail", Message: "type query is empty"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}
	category := c.Query("category")
	if category == "" {
		stat := status.StatusModel{Status: "fail", Message: "category query is empty"}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}
	var parameter system.SystemParameter
	if res := r.DB.Where("parameter_name LIKE ?", "%"+types+"%").First(&parameter); res.Error != nil {
		stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	user, err := r.API.Auth.GetUser()
	if err != nil {
		stat := status.StatusModel{Status: "fail", Message: "Unauthorized"}
		return c.Status(fiber.StatusUnauthorized).JSON(stat)
	}
	var courseIDs []string
	if err := r.DB.Model(&course.CourseModel{}).
		Where("course_category LIKE ?", "%"+category+"%").
		Pluck("course_id", &courseIDs).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	var results struct {
		TotalPointAchieved int
	}
	if err := r.DB.Model(&course.EnrollmentModel{}).
		Select("SUM(point_achieved) as total_point_achieved").
		Where("user_id = ?", user.ID).
		Where("course_id IN ?", courseIDs).
		Find(&results).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	param, err := strconv.Atoi(parameter.ParameterValue)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	if results.TotalPointAchieved >= param {
		var rankk rank.RankBadge
		if res := r.DB.Where("rank_title LIKE ?", "%"+types+"%").First(&rankk); res.Error != nil {
			stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
			return c.Status(fiber.StatusInternalServerError).JSON(stat)
		}
		userRank := rank.UserRankBadges{
			UserID:       user.ID,
			RankID:       rankk.RankID,
			RankCategory: category,
			ObtainedAt:   time.Now(),
		}
		if res := r.DB.Save(&userRank); res.Error != nil {
			stat := status.StatusModel{Status: "fail", Message: res.Error.Error()}
			return c.Status(fiber.StatusInternalServerError).JSON(stat)
		}
		stat := status.StatusModel{Status: "success", Message: "User Rank Badge Has been Successfully Set"}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	} else {
		stat := status.StatusModel{Status: "fail", Message: "Your total point wasn't enough to fulfilled rank requirements"}
		return c.Status(fiber.StatusNotAcceptable).JSON(stat)
	}
}
