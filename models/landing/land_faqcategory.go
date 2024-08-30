package models

import "time"

// Land_Faqcategory_Request represents a request for FAQ category operations.
type Land_Faqcategory struct {
	FaqCategoryID   int        `gorm:"column:faq_category_id;primaryKey" json:"faq_category_id"`
	FaqCategoryName string     `gorm:"column:faq_categoryname" json:"faq_categoryname"`
	CreatedBy       string     `gorm:"column:created_by" json:"created_by"`
	CreatedAt       time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedBy       string     `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt       time.Time  `gorm:"column:updated_at" json:"updated_at"`
	Faqs            []Land_Faq `gorm:"foreignKey:FaqCategory;referencesFaqCategoryID" json:"faqs,omitempty"`
}

func (l *Land_Faqcategory) TableName() string {
	return "land_faqcategory"
}

type LandFaqCategoryRequest struct {
	FaqCategoryName string `json:"faq_categoryname"`
}
