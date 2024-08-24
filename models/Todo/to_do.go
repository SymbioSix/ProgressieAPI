package models

import (
	"time"

	"github.com/google/uuid"
)

type TdCustomTarget struct {
	TargetID           uuid.UUID   `gorm:"column:target_id;primaryKey" json:"target_id"`
	AchievementID      uuid.UUID   `gorm:"column:achievement_title" json:"achievement_id"`
	TargetTitle        string      `gorm:"column:target_title" json:"target_title"`
	TargetSubtitle     string      `gorm:"column:target_subtitle" json:"target_subtitle"`
	DailyClockReminder time.Time   `gorm:"column:daily_clock_reminder" json:"daily_clock_reminder"`
	Type               string      `gorm:"column:type" json:"type"`
	CreatedAt          time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time   `gorm:"column:updated_at" json:"updated_at"`
	DueAt              time.Time   `gorm:"column:due_at" json:"due_at"`
	CheckLists         []Checklist `gorm:"foreignKey:TargetID;references:TargetID" json:"checklists,omitempty"`
}

func (td *TdCustomTarget) TableName() string {
	return "td_customtarget"
}

type TdSubcourseReminder struct {
	ReminderID          string    `gorm:"column:reminder_id;primaryKey" json:"reminder_id"`
	SubcourseprogressID string    `gorm:"column:subcourseprogress_id" json:"subcourseprogress_id"`
	ReminderTitle       string    `gorm:"column:reminder_title" json:"reminder_title"`
	ReminderTime        time.Time `gorm:"column:reminder_time" json:"reminder_time"`
	StartDate           time.Time `gorm:"column:start_date" json:"start_date"`
	IsFinished          bool      `gorm:"column:is_finished" json:"is_finished"`
	Type                string    `gorm:"column:type" json:"type"`
	CreatedBy           string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy           string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (td *TdSubcourseReminder) TableName() string {
	return "td_subcoursereminder"
}

type SubcourseProgress struct {
	SubcourseprogressID string                `gorm:"column:subcourseprogress_id;primaryKey" json:"subcourseprogress_id"`
	UserID              uuid.UUID             `gorm:"column:user_id" json:"user_id"`
	SubcourseID         string                `gorm:"column:subcourse_id" json:"subcourse_id"`
	IsVideoViewed       bool                  `gorm:"column:is_video_viewed" json:"is_video_viewed"`
	IsSubcourseFinished bool                  `gorm:"column:is_subcourse_finished" json:"is_subcourse_finished"`
	CreatedBy           string                `gorm:"column:created_by" json:"created_by"`
	CreatedAt           time.Time             `gorm:"column:created_at" json:"created_at"`
	UpdatedBy           string                `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt           time.Time             `gorm:"column:updated_at" json:"updated_at"`
	Reminders           []TdSubcourseReminder `gorm:"foreignKey:SubcourseprogressID;references:SubcourseprogressID" json:"reminders,omitempty"`
}

func (td *SubcourseProgress) TableName() string {
	return "crs_subcourseprogress"
}

// Checklist represents the daily checklist entries for custom targets
type Checklist struct {
	ChecklistID string    `gorm:"column:checklist_id;primaryKey" json:"checklist_id"`
	TargetID    string    `gorm:"column:target_id" json:"target_id"`
	DateChecked time.Time `gorm:"column:date_checked" json:"date_checked"`
}

func (td *Checklist) TableName() string {
	return "td_customtargetchecklist"
}

type RequestCustomTarget struct {
	TargetTitle    string    `json:"target_title"`
	TargetSubtitle string    `json:"target_subtitle"`
	DailyReminder  time.Time `json:"daily_reminder"`
}

type RequestTdSubcourseReminder struct {
	SubcourseprogressID string    `json:"subcourseprogress_id"`
	ReminderTitle       string    `json:"reminder_title"`
	ReminderTime        time.Time `json:"reminder_time"`
	StartDate           time.Time `json:"start_date"`
	IsFinished          bool      `json:"is_finished"`
}

type TdCustomTargetResponse struct {
	TargetID           uuid.UUID `json:"target_id"`
	AchievementID      uuid.UUID `json:"achievement_id"`
	TargetTitle        string    `json:"target_title"`
	TargetSubtitle     string    `json:"target_subtitle"`
	DailyClockReminder time.Time `json:"daily_clock_reminder"`
	DueAt              time.Time `json:"due_at"`
}

type TdSubcourseReminderResponse struct {
	ReminderID          string    `gorm:"column:reminder_id;primaryKey" json:"reminder_id"`
	SubcourseprogressID string    `gorm:"column:subcourseprogress_id" json:"subcourseprogress_id"`
	ReminderTitle       string    `gorm:"column:reminder_title" json:"reminder_title"`
	ReminderTime        time.Time `gorm:"column:reminder_time" json:"reminder_time"`
	StartDate           time.Time `gorm:"column:start_date" json:"start_date"`
	IsFinished          bool      `gorm:"column:is_finished" json:"is_finished"`
}

type UpdateChecklistDateRequest struct {
	DateChecked time.Time `json:"date_checked"`
}
