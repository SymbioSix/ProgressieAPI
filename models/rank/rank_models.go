package models

import (
	"time"

	"github.com/google/uuid"
)

type RankBadge struct {
	RankID          uuid.UUID `gorm:"column:rank_id;primaryKey" json:"rank_id"`
	RankTitle       string    `gorm:"column:rank_title" json:"rank_title"`
	RankIcon        string    `gorm:"column:rank_icon" json:"rank_icon"`
	RankDescription string    `gorm:"column:rank_description" json:"rank_description"`
}

func (rank *RankBadge) TableName() string {
	return "rank_badge"
}

type UserRankBadges struct {
	UserID       uuid.UUID `gorm:"column:user_id;primaryKey" json:"user_id"`
	RankID       uuid.UUID `gorm:"column:rank_id;primaryKey" json:"rank_id"`
	RankData     RankBadge `gorm:"foreignKey:RankID;references:RankID" json:"rank_data"`
	RankCategory string    `gorm:"column:rank_category" json:"rank_category"`
	ObtainedAt   time.Time `gorm:"column:obtained_at" json:"obtained_at"`
}

func (badge *UserRankBadges) TableName() string {
	return "usr_rankbadge"
}
