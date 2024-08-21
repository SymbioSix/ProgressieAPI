package models

import (
	"time"

	"gorm.io/gorm"
)

// QuizStatus is a custom type representing the status of a quiz
type QuizStatus string

const (
	QuizActive   QuizStatus = "active"
	QuizInactive QuizStatus = "inactive"
)

// Quiz represents the quiz table
type Quiz struct {
	QuizID      string     `gorm:"column:quiz_id;primaryKey" json:"quiz_id"`
	SubcourseID string     `gorm:"column:subcourse_id" json:"subcourse_id"`
	Description string     `gorm:"column:description" json:"description"`
	CreatedBy   string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy   string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	Status      QuizStatus `gorm:"column:status" json:"status"`
}

// BeforeCreate is a GORM hook that runs before a new record is inserted into the database
func (quiz *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	quiz.CreatedAt = time.Now()
	quiz.UpdatedAt = time.Now()
	return
}

// BeforeUpdate is a GORM hook that runs before an existing record is updated in the database
func (quiz *Quiz) BeforeUpdate(tx *gorm.DB) (err error) {
	quiz.UpdatedAt = time.Now()
	return
}

func (quiz *Quiz) TableName() string {
	return "crs_quiz"
}
