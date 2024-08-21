package models

import (
	"time"

	"gorm.io/gorm"
)

// GeneralStatus is a custom type representing the status of a quiz question
type QuizQuestionStatus string

const (
	QuizQuestionEnabled  QuizQuestionStatus = "enabled"
	QuizQuestionDisabled QuizQuestionStatus = "disabled"
)

// QuizQuestion represents the quiz_question table
type QuizQuestion struct {
	QuizID         string             `gorm:"column:quiz_id" json:"quiz_id"`
	QuizQuestionID string             `gorm:"column:quizquestion_id;primaryKey" json:"quizquestion_id"`
	QuestionText   string             `gorm:"column:question_text" json:"question_text"`
	CreatedBy      string             `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      string             `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time          `gorm:"column:updated_at" json:"updated_at"`
	Status         QuizQuestionStatus `gorm:"column:status" json:"status"`
}

// BeforeCreate is a GORM hook that runs before a new record is inserted into the database
func (q *QuizQuestion) BeforeCreate(tx *gorm.DB) (err error) {
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that runs before an existing record is updated in the database
func (q *QuizQuestion) BeforeUpdate(tx *gorm.DB) (err error) {
	q.UpdatedAt = time.Now()
	return
}

func (q *QuizQuestion) TableName() string {
	return "crs_quizquestion"
}
