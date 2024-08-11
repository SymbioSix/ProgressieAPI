package models

import (
	"time"

	"github.com/google/uuid"
)

type EnrollmentModel struct {
	UserID        uuid.UUID   `gorm:"column:user_id;primaryKey" json:"user_id"`
	CourseID      string      `gorm:"column:course_id;primaryKey" json:"course_id"`
	TheCourse     CourseModel `gorm:"foreignKey:CourseID;references:CourseID" json:"the_course,omitempty"`
	Progress      float64     `gorm:"column:progress" json:"progress"`
	PointAchieved int         `gorm:"column:point_achieved" json:"point_achieved"`
	Status        string      `gorm:"column:status" json:"status"`
	CreatedBy     string      `gorm:"column:created_by" json:"created_by"`
	CreatedAt     time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedBy     string      `gorm:"column:updated_by" json:"updated_by,omitempty"`
	UpdatedAt     time.Time   `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (crs *EnrollmentModel) TableName() string {
	return "crs_enrollment"
}

type UpdateEnrollmentProgress struct {
	Progress float64 `json:"progress" binding:"required"`
}

type UpdateEnrollmentPoint struct {
	Point int `json:"point" binding:"required"`
}
