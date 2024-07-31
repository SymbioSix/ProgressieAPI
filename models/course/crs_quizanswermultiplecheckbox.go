package models

import (
	"time"

	"gorm.io/gorm"
)

// MultCheckBoxStatus is a custom type representing the status of a multi-checkbox entity
type MultCheckBoxStatus string

const (
	MultCheckBoxActive   MultCheckBoxStatus = "active"
	MultCheckBoxInactive MultCheckBoxStatus = "inactive"
)

// QuizMultCheckBox represents the quiz_mult_checkbox table
type QuizMultCheckBox struct {
	QuizQuestionID string             `gorm:"column:quizquestion_id;primaryKey" json:"quizquestion_id"`
	IsAnswer       bool               `gorm:"column:is_answer" json:"is_answer"`
	AnswerText     string             `gorm:"column:answer_text" json:"answer_text"`
	Status         MultCheckBoxStatus `gorm:"column:status" json:"status"`
	CreatedBy      string             `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      string             `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time          `gorm:"column:updated_at" json:"updated_at"`
}

// BeforeCreate is a GORM hook that runs before a new record is inserted into the database
func (question *QuizMultCheckBox) BeforeCreate(tx *gorm.DB) (err error) {
	question.CreatedAt = time.Now()
	question.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that runs before an existing record is updated in the database
func (question *QuizMultCheckBox) BeforeUpdate(tx *gorm.DB) (err error) {
	question.UpdatedAt = time.Now()
	return
}
