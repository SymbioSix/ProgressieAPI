package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GeneralStatus is a custom type representing the status of a quiz result
type QuizResultStatus string

const (
	QuizResultEnable   QuizResultStatus = "enable"
	QuizResultDisabled QuizResultStatus = "disabled"
)

// QuizResult represents the quiz_result table
type QuizResult struct {
	QuizID       string           `gorm:"column:quiz_id;primaryKey" json:"quiz_id"`
	QuizData     Quiz             `gorm:"foreignKey:QuizID;references:QuizID" json:"quiz_data,omitempty"`
	UserID       uuid.UUID        `gorm:"column:user_id" json:"user_id"`
	Progress     float32          `gorm:"column:progress" json:"progress"`
	HighestScore float32          `gorm:"column:highest_score" json:"highest_score"`
	LastScore    float32          `gorm:"column:last_score" json:"last_score"`
	Status       QuizResultStatus `gorm:"column:status" json:"status"`
	CompletedAt  time.Time        `gorm:"column:completed_at" json:"completed_at"`
	UpdatedBy    string           `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time        `gorm:"column:updated_at" json:"updated_at"`
}

// BeforeCreate is a GORM hook that runs before a new record is inserted into the database
func (qr *QuizResult) BeforeCreate(tx *gorm.DB) (err error) {
	qr.CompletedAt = time.Now()
	qr.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that runs before an existing record is updated in the database
func (qr *QuizResult) BeforeUpdate(tx *gorm.DB) (err error) {
	qr.UpdatedAt = time.Now()
	return
}

func (qr *QuizResult) TableName() string {
	return "crs_quizresult"
}
