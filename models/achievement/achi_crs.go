package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// AchievementCrs represents the achievement_crs_ table
type AchievementCrs struct {
	AchievementCrsID           uuid.UUID `gorm:"column:achievement_crs_id;primaryKey;type:uuid;not null" json:"achievement_crs_id"`
	UserID                     uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id"`
	Progress                   float32   `gorm:"column:progress;type:float4;not null" json:"progress"`
	SubcourseProgressID        string    `gorm:"column:subcourseprogress_id;type:varchar(255);not null" json:"subcourseprogress_id"`
	SubcourseProgressReadingID string    `gorm:"column:subcourseprogressreading_id;type:varchar(255);not null" json:"subcourseprogressreading_id"`
	AchievementCrsTitle        string    `gorm:"column:achievement_crs_title;type:text" json:"achievement_crs_title"`
	AchievementCrsIcon         string    `gorm:"column:achievement_crs_icon;type:text" json:"achievement_crs_icon"`
	AchievementCrsDescription  string    `gorm:"column:achievement_crs_description;type:text" json:"achievement_crs_description"`
	AchievementCrsCategory     string    `gorm:"column:achievement_crs_category;type:text" json:"achievement_crs_category"`
	IsAchieved                 bool      `gorm:"column:is_achieved;type:boolean" json:"is_achieved"`
	AchievedAt                 time.Time `gorm:"column:achieved_at;type:timestamptz" json:"achieved_at"`
	CreatedAt                  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt                  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUIDs if they are not provided
func (a *AchievementCrs) BeforeCreate(tx *gorm.DB) (err error) {
	if a.AchievementCrsID == uuid.Nil {
		a.AchievementCrsID = uuid.New()
	}
	if a.UserID == uuid.Nil {
		a.UserID = uuid.New()
	}
	return
}
