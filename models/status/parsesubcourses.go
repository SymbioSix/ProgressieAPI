package models

import models "github.com/SymbioSix/ProgressieAPI/models/courses"

type ParseSubcoursesFromCourseID struct {
	Subcourses []models.SubCourseModel `json:"subcourses"`
}
