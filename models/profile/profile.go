package models

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/Todo"
	"github.com/google/uuid"
)

type UserTitleSkill struct {
	UserID     uuid.UUID `gorm:"column:user_id" json:"user_id"`
	TitleID    uuid.UUID `gorm:"column:title_id" json:"title_id"`
	TitleSkill string    `gorm:"column:title_skill" json:"title_skill"`
	Subtitle   string    `gorm:"column:subtitle" json:"subtitle,omitempty"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (usr *UserTitleSkill) TableName() string {
	return "usr_title"
}

type UserRank struct {
	UserID          uuid.UUID `gorm:"column:user_id;primaryKey" json:"user_id"`
	RankID          uuid.UUID `gorm:"column:rank_id;primaryKey" json:"rank_id"`
	RankTitle       string    `gorm:"column:rank_title" json:"rank_title"`
	RankIcon        string    `gorm:"column:rank_icon" json:"rank_icon"`
	RankDescription string    `gorm:"column:rank_description" json:"rank_description,omitempty"`
	IsObtained      bool      `gorm:"column:is_obtained" json:"is_obtained"`
	ObtainedAt      time.Time `gorm:"column:obtained_at" json:"obtained_at,omitempty"`
}

func (usr *UserRank) TableName() string {
	return "usr_rank"
}

type UserAchievement struct {
	UserID                 uuid.UUID               `gorm:"column:user_id;primaryKey" json:"user_id"`
	AchievementID          uuid.UUID               `gorm:"column:achievement_id;primaryKey" json:"achievement_id"`
	AchievementTitle       string                  `gorm:"column:achievement_title" json:"achievement_title"`
	AchievementIcon        string                  `gorm:"column:achievement_icon" json:"achievement_icon"`
	AchievementDescription string                  `gorm:"column:achievement_description" json:"achievement_description"`
	AchievementCategory    string                  `gorm:"column:achievement_category" json:"achievement_category"`
	IsAchieved             bool                    `gorm:"column:is_achieved" json:"is_achieved"`
	CustomTargets          []models.TdCustomTarget `gorm:"foreignKey:AchievementID;references:AchievementID" json:"custom_targets,omitempty"`
	// AchievementTotalAchieved int       `gorm:"column:achievement_total_achieved"` // New field
}

func (usr *UserAchievement) TableName() string {
	return "usr_achievement"
}

type UpdateUserTitleSkillRequest struct {
	TitleSkill string    `json:"title_skill,omitempty"`
	Subtitle   string    `json:"subtitle,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
