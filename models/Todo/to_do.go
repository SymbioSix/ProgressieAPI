package models

import (
	"time"

	"github.com/google/uuid"
)

type TdCustomTarget struct {
	TargetID           string    `gorm:"column:target_id" json:"target_id"`
	AchievementTitle   string `gorm:"achievement_title;column:achievement_title"`
	UserID             string    `gorm:"column:user_id" json:"user_id"`
	TargetTitle        string    `gorm:"column:target_title" json:"target_title"`
	TargetSubtitle     string    `gorm:"column:target_subtitle" json:"target_subtitle"`
	TargetIcon         string    `gorm:"column:target_icon" json:"target_icon"`
	DailyClockReminder time.Time `gorm:"column:daily_clock_reminder" json:"daily_clock_reminder"`
	Type               string    `gorm:"column:type" json:"type"`
	CreatedBy          string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy          string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
	DueAt              time.Time `gorm:"column:due_at" json:"due_at"`
}

type TdSubcourseReminder struct {
	ReminderID    string    `gorm:"column:reminder_id" json:"reminder_id"`
	SubcourseId   string    `gorm:"column:subcourse_id" json:"subcourse_id"`
	UserID        uuid.UUID `gorm:"column:user_id" json:"user_id"`
	ReminderTitle string    `gorm:"column:reminder_title" json:"reminder_title"`
	Icon          string    `gorm:"column:icon" json:"icon"`
	ReminderTime  time.Time `gorm:"column:reminder_time" json:"reminder_time"`
	StartDate     time.Time `gorm:"column:start_date" json:"start_date"`
	Type          string    `gorm:"column:type" json:"type"`
	IsFinished    bool      `gorm:"column:is_finished" json:"is_finished"`
	CreatedBy     string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy     string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SubcourseProgress struct {
	SubcourseProgressID string    `gorm:"column:subcourseprogress_id" json:"subcourse_progress_id"`
	UserID              uuid.UUID `gorm:"column:user_id" json:"user_id"`
	SubcourseID         string    `gorm:"column:subcourse_id" json:"subcourse_id"`
	IsVideoViewed       bool      `gorm:"column:is_video_viewed" json:"is_video_viewed"`
	IsSubcourseFinished bool      `gorm:"column:is_subcourse_finished" json:"is_subcourse_finished"`
	CreatedBy           string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy           string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// Checklist represents the daily checklist entries for custom targets
type Checklist struct {
    ChecklistID string    `gorm:"column:checklist_id;primaryKey" json:"checklist_id"`
    TargetID    string    `gorm:"column:target_id" json:"target_id"`
    DateChecked time.Time `gorm:"column:date_checked" json:"date_checked"`
    UserID      string    `gorm:"column:user_id" json:"user_id"`
}

type Achievementtar struct {
    UserID                   string `gorm:"user_id;column:user_id"`
    AchievementTitle         string `gorm:"achievement_title;column:achievement_title"`
    AchievementIcon          string `gorm:"column:achievement_icon"`
    AchievementDescription   string `gorm:"column:achievement_description"`
    AchievementCategory      string `gorm:"column:achievement_category"`
    IsAchieved               bool   `gorm:"column:is_achieved"`
    AchievementTotalAchieved int    `gorm:"column:achievement_total_achieved"` // New field
}

