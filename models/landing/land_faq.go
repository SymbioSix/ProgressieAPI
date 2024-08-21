package models

import "time"

// Land_Faq represents a request for FAQ operations.
type Land_Faq struct {
	FaqID          int       `gorm:"column:faq_id;primaryKey" json:"faq_id"`
	FaqCategory    int       `gorm:"column:faq_category" json:"faq_category"`
	FaqTitle       string    `gorm:"column:faq_title" json:"faq_title"`
	FaqDescription string    `gorm:"column:faq_description" json:"faq_description"`
	Tooltip        string    `gorm:"column:tooltip" json:"tooltip"`
	CreatedBy      string    `gorm:"column:created_by" json:"created_by"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (l *Land_Faq) TableName() string {
	return "land_faq"
}
