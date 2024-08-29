package models

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/Todo"
	quiz "github.com/SymbioSix/ProgressieAPI/models/quiz"
	"github.com/google/uuid"
)

type CourseModel struct {
	CourseID       string           `gorm:"column:course_id;primaryKey" json:"course_id"`
	CourseName     string           `gorm:"column:course_name" json:"course_name"`
	CourseCategory string           `gorm:"column:course_category" json:"course_category"`
	CourseImage    string           `gorm:"column:course_image" json:"course_image"`
	Description    string           `gorm:"column:description" json:"description"`
	Price          float64          `gorm:"column:price" json:"price"`
	Status         string           `gorm:"column:status" json:"status"`
	CreatedBy      string           `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      string           `gorm:"column:updated_by" json:"updated_by,omitempty"`
	UpdatedAt      time.Time        `gorm:"column:updated_at" json:"updated_at,omitempty"`
	SubCourses     []SubCourseModel `gorm:"foreignKey:CourseID;references:CourseID" json:"sub_courses,omitempty"`
}

func (crs *CourseModel) TableName() string {
	return "crs_course"
}

type SubCourseModel struct {
	SubcourseID     string                         `gorm:"column:subcourse_id;primaryKey" json:"subcourse_id"`
	SubcourseName   string                         `gorm:"column:subcourse_name" json:"subcourse_name"`
	CourseID        string                         `gorm:"column:course_id" json:"course_id"`
	Sequence        int                            `gorm:"column:sequence" json:"sequence"`
	Description     string                         `gorm:"column:description" json:"description"`
	Status          string                         `gorm:"column:status" json:"status"`
	CreatedBy       string                         `gorm:"column:created_by" json:"created_by"`
	CreatedAt       time.Time                      `gorm:"column:created_at" json:"created_at"`
	UpdatedBy       string                         `gorm:"column:updated_by" json:"updated_by,omitempty"`
	UpdatedAt       time.Time                      `gorm:"column:updated_at" json:"updated_at,omitempty"`
	Progress        SubcourseProgress              `gorm:"foreignKey:SubcourseID;references:SubcourseID" json:"progress,omitempty"`
	VideoContent    SubCourseVideoContentModel     `gorm:"foreignKey:SubcourseID;references:SubcourseID" json:"video_content,omitempty"`
	ReadingContents []SubCourseReadingContentModel `gorm:"foreignKey:SubcourseID;references:SubcourseID" json:"reading_contents,omitempty"`
	Quizzes         []quiz.Quiz                    `gorm:"foreignKey:SubcourseID;references:SubcourseID" json:"quizzes,omitempty"`
}

func (crs *SubCourseModel) TableName() string {
	return "crs_subcourse"
}

type SubCourseVideoContentModel struct {
	SubcourseID    string `gorm:"column:subcourse_id;primaryKey" json:"subcourse_id"`
	VideoTitle     string `gorm:"column:video_title" json:"video_title"`
	VideoLink      string `gorm:"column:video_link" json:"video_link"`
	Description    string `gorm:"column:description" json:"description"`
	Status         string `gorm:"column:status" json:"status"`
	EarnedPoint    int    `gorm:"column:earned_point" json:"earned_point"`
	Source         string `gorm:"column:source" json:"source"`
	PlatformSource string `gorm:"column:platform_source" json:"platform_source"`
}

func (crs *SubCourseVideoContentModel) TableName() string {
	return "crs_subcoursevideo"
}

type SubCourseReadingContentModel struct {
	SubcoursereadingID string                              `gorm:"column:subcoursereading_id;primaryKey" json:"subcoursereading_id"`
	SubcourseID        string                              `gorm:"column:subcourse_id" json:"subcourse_id"`
	Title              string                              `gorm:"column:title" json:"title"`
	Subtitle           string                              `gorm:"column:subtitle" json:"subtitle"`
	ReadingDuration    int                                 `gorm:"column:reading_duration" json:"reading_duration"`
	EarnedPoint        int                                 `gorm:"column:earned_point" json:"earned_point"`
	Status             string                              `gorm:"column:status" json:"status"`
	CreatedBy          string                              `gorm:"column:created_by" json:"created_by"`
	CreatedAt          time.Time                           `gorm:"column:created_at" json:"created_at"`
	UpdatedBy          string                              `gorm:"column:updated_by" json:"updated_by,omitempty"`
	UpdatedAt          time.Time                           `gorm:"column:updated_at" json:"updated_at,omitempty"`
	ReadingImages      []SubCourseReadingImageContentModel `gorm:"foreignKey:SubcoursereadingID;references:SubcoursereadingID" json:"reading_images"`
}

func (crs *SubCourseReadingContentModel) TableName() string {
	return "crs_subcoursereading"
}

type SubCourseReadingImageContentModel struct {
	SubcoursereadingID string `gorm:"column:subcoursereading_id;primaryKey" json:"subcoursereading_id"`
	ImageLink          string `gorm:"column:image_link" json:"image_link"`
	Description        string `gorm:"column:description" json:"description"`
}

func (crs *SubCourseReadingImageContentModel) TableName() string {
	return "crs_subcoursereadingimage"
}

type SubcourseProgress struct {
	SubcourseprogressID uuid.UUID                    `gorm:"column:subcourseprogress_id;primaryKey" json:"subcourseprogress_id"`
	UserID              uuid.UUID                    `gorm:"column:user_id" json:"user_id"`
	SubcourseID         string                       `gorm:"column:subcourse_id" json:"subcourse_id"`
	IsVideoViewed       bool                         `gorm:"column:is_video_viewed" json:"is_video_viewed"`
	IsSubcourseFinished bool                         `gorm:"column:is_subcourse_finished" json:"is_subcourse_finished"`
	CreatedBy           string                       `gorm:"column:created_by" json:"created_by"`
	CreatedAt           time.Time                    `gorm:"column:created_at" json:"created_at"`
	UpdatedBy           string                       `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt           time.Time                    `gorm:"column:updated_at" json:"updated_at"`
	ReadingProgress     []SubcourseProgressReading   `gorm:"foreignKey:SubcourseprogressID;references:SubcourseprogressID" json:"reading_progress,omitempty"`
	Reminders           []models.TdSubcourseReminder `gorm:"foreignKey:SubcourseprogressID;references:SubcourseprogressID" json:"reminders,omitempty"`
}

func (td *SubcourseProgress) TableName() string {
	return "crs_subcourseprogress"
}

type SubcourseProgressReading struct {
	SubcourseprogressreadingID uuid.UUID `gorm:"column:subcourseprogressreading_id;primaryKey" json:"subcourseprogressreading_id"`
	SubcourseprogressID        uuid.UUID `gorm:"column:subcourseprogress_id" json:"subcourseprogress_id"`
	SubcoursereadingID         string    `gorm:"column:subcoursereading_id" json:"subcoursereading_id"`
	IsReaded                   bool      `gorm:"column:is_readed" json:"is_readed"`
}

func (td *SubcourseProgressReading) TableName() string {
	return "crs_subcourseprogressreading"
}
