package models

import (
	"time"

	course "github.com/SymbioSix/ProgressieAPI/models/courses"
	models "github.com/SymbioSix/ProgressieAPI/models/profile"
	quiz "github.com/SymbioSix/ProgressieAPI/models/quiz"
	"github.com/google/uuid"
)

type UserModel struct {
	UserID                  uuid.UUID                  `gorm:"column:user_id" json:"id"`
	FirstName               string                     `gorm:"column:first_name" json:"first_name"`
	LastName                string                     `gorm:"column:last_name" json:"last_name"`
	Email                   string                     `gorm:"column:email" json:"email"`
	PhoneNumber             string                     `gorm:"column:phone_number" json:"phone_number"`
	PhotoProfile            string                     `gorm:"column:photo_profile_link" json:"photo_profile_link"`
	TitleProfile            string                     `gorm:"column:title_desc_profile" json:"title_desc_profile"`
	Description             string                     `gorm:"column:description_profile" json:"description,omitempty"`
	EmailVerified           bool                       `gorm:"column:email_verified" json:"email_verified,omitempty"`
	PhoneVerified           bool                       `gorm:"column:phone_verified" json:"phone_verified,omitempty"`
	Username                string                     `gorm:"column:username" json:"username"`
	Gender                  string                     `gorm:"column:gender" json:"gender"`
	TotalCoursesFinished    int                        `gorm:"column:total_courses_finished" json:"total_courses_finished"`
	TotalSubcoursesFinished int                        `gorm:"column:total_subcourses_finished" json:"total_subcourses_finished"`
	TotalPointAchieved      int                        `gorm:"column:total_point_achieved" json:"total_point_achieved"`
	Status                  string                     `gorm:"column:status" json:"status"`
	CreatedBy               string                     `gorm:"column:created_by" json:"created_by,omitempty"`
	UpdatedBy               string                     `gorm:"column:updated_by" json:"updated_by,omitempty"`
	UpdatedAt               time.Time                  `gorm:"column:updated_at" json:"updated_at,omitempty"`
	Achievements            []models.UserAchievement   `gorm:"foreignKey:UserID;references:UserID" json:"achievements,omitempty"`
	Ranks                   []models.UserRank          `gorm:"foreignKey:UserID;references:UserID" json:"ranks,omitempty"`
	TitleSkills             []models.UserTitleSkill    `gorm:"foreignKey:UserID;references:UserID" json:"title_skills,omitempty"`
	QuizResults             []quiz.QuizResult          `gorm:"foreignKey:UserID;references:UserID" json:"quiz_results,omitempty"`
	CourseEnrollment        []course.EnrollmentModel   `gorm:"foreignKey:UserID;references:UserID" json:"course_enrollment,omitempty"`
	SubcourseProgress       []course.SubcourseProgress `gorm:"foreignKey:UserID;references:UserID" json:"subcourse_progress,omitempty"`
}

func (user *UserModel) TableName() string {
	return "usr_user"
}

type RoleModel struct {
	RoleID      string    `gorm:"column:role_id" json:"id"`
	RoleName    string    `gorm:"column:role_name" json:"role_name"`
	Description string    `gorm:"column:description" json:"description,omitempty"`
	Status      string    `gorm:"column:status" json:"status,omitempty"`
	CreatedBy   string    `gorm:"column:created_by" json:"created_by,omitempty"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (role *RoleModel) TableName() string {
	return "usr_role"
}

type UserRoleResponse struct {
	UserID   string    `gorm:"column:user_id;primaryKey" json:"user_id,omitempty"`
	UserData UserModel `gorm:"foreignKey:UserID;references:UserID" json:"user_data"`
	RoleID   string    `gorm:"column:role_id;primaryKey" json:"role_id,omitempty"`
	RoleData RoleModel `gorm:"foreignKey:RoleID;references:RoleID" json:"role_data"`
}

type InsertUserRole struct {
	UserID    uuid.UUID `gorm:"column:user_id"`
	RoleID    string    `gorm:"column:role_id"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedBy string    `gorm:"column:updated_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type UpdateUserProfileRequest struct {
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	PhotoProfile string `json:"photo_profile_link,omitempty"`
	TitleProfile string `json:"title_desc_profile,omitempty"`
	Description  string `json:"description,omitempty"`
	Gender       string `json:"gender,omitempty"`
}
