package models

import "time"

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
	VideoContent    SubCourseVideoContentModel     `gorm:"foreignKey:SubcourseID;references:SubcourseID" json:"video_content,omitempty"`
	ReadingContents []SubCourseReadingContentModel `gorm:"foreignKey:SubcourseID;references:SubcourseID" json:"reading_contents,omitempty"`
}

func (crs *SubCourseModel) TableName() string {
	return "crs_subcourse"
}

type SubCourseVideoContentModel struct {
	SubcourseID string `gorm:"column:subcourse_id;primaryKey" json:"subcourse_id"`
	VideoLink   string `gorm:"column:video_link" json:"video_link"`
	Description string `gorm:"column:description" json:"description"`
	Status      string `gorm:"column:status" json:"status"`
	EarnedPoint int    `gorm:"column:earned_point" json:"earned_point"`
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
