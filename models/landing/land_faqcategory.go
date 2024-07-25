package models

import "time"

// Land_Faqcategory_Request represents a request for FAQ category operations.
type Land_Faqcategory_Request struct {
    FaqCategoryID   int       `gorm:"column:faq_category_id;primaryKey" json:"faq_category_id"`
    FaqCategoryName string    `gorm:"column:faq_category_name" json:"faq_category_name"`
    CreatedBy       string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy       string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// Land_Faqcategory_Response represents a response for FAQ category operations.
type Land_Faqcategory_Response struct {
    FaqCategoryID   int       `gorm:"column:faq_category_id" json:"faq_category_id"`
    FaqCategoryName string    `gorm:"column:faq_category_name" json:"faq_category_name"`
    CreatedBy       string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy       string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}
