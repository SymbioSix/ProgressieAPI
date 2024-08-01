package models

import (
	"time"

	"gorm.io/gorm"
)

// GeneralStatus is a custom type representing the status of a quiz question
type MultChoiceStatus string

const (
	MultChoiceEnabled  MultChoiceStatus = "enabled"
	MultChoiceDisabled MultChoiceStatus = "disabled"
)

// QuizAnswerMultipleChoice represents the quiz_answer_multiple_choice table
type QuizAnswerMultipleChoice struct {
	QuizQuestionID string             `gorm:"column:quizquestion_id;primaryKey" json:"quizquestion_id"`
	Answer         string             `gorm:"column:answer" json:"answer"`
	OptionA        string             `gorm:"column:option_a" json:"option_a"`
	OptionB        string             `gorm:"column:option_b" json:"option_b"`
	OptionC        string             `gorm:"column:option_c" json:"option_c"`
	OptionD        string             `gorm:"column:option_d" json:"option_d"`
	OptionE        string             `gorm:"column:option_e" json:"option_e"`
	Status         MultCheckBoxStatus `gorm:"column:status" json:"status"`
	CreatedBy      string             `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      string             `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time          `gorm:"column:updated_at" json:"updated_at"`
}

// BeforeCreate is a GORM hook that runs before a new record is inserted into the database
func (q *QuizAnswerMultipleChoice) BeforeCreate(tx *gorm.DB) (err error) {
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that runs before an existing record is updated in the database
func (q *QuizAnswerMultipleChoice) BeforeUpdate(tx *gorm.DB) (err error) {
	q.UpdatedAt = time.Now()
	return
}
