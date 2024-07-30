package models

import "time"

// Land_Faq_Request represents a request for FAQ operations.
type Land_Faq_Request struct {
    CREATE TABLE `usr_user` (
        `user_id` PK, uuid,
        `first_name` text,
        `last_name` text,
        `Rank_Id` FK , varchar,
        `email` text,
        `password` text,
        `phone_number` text,
        `photo_profile_link` text,
        `email_verified` bool,
        `phone_verified` bool,
        `last_email_updated` timestamp,
        `total_courses_finished` int,
        `total_subcourses_finished` int,
        `total_point_achieved` int,
        `status` enum(general_status),
        `created_by` text,
        `created_at` timestamp,
        `updated_by` text,
        `updated_at` timestamp
      );
            
}

// Land_Faq_Response represents a response for FAQ operations.
type Land_Faq_Response struct {
    FaqID          int       `gorm:"column:faq_id" json:"faq_id"`
    FaqCategory    int       `gorm:"column:faq_category" json:"faq_category"`
    FaqTitle       string    `gorm:"column:faq_title" json:"faq_title"`
    FaqDescription string    `gorm:"column:faq_description" json:"faq_description"`
    Tooltip        string    `gorm:"column:tooltip" json:"tooltip"`
    CreatedBy      string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy      string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}
