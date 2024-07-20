package models

import "time"

type Land_Aboutus_Request struct {
    AboutUsComponentID      int       `gorm:"column:aboutus_component_id" json:"aboutus_component_id"` // Primary key
    AboutUsComponentName    string    `gorm:"column:aboutus_component_name" json:"aboutus_component_name"`
    AboutUsComponentJobdesc string    `gorm:"column:aboutus_component_jobdesc" json:"aboutus_component_jobdesc"`
    AboutUsComponentPhoto   string    `gorm:"column:aboutus_component_photo" json:"aboutus_component_photo"`
    Description             string    `gorm:"column:description" json:"description"`
    Tooltip                 string    `gorm:"column:tooltip" json:"tooltip"`
    Status                  string    `gorm:"column:status" json:"status"`
    CreatedBy               string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt               time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy               string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt               time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Land_Aboutus_Response struct {
    AboutUsComponentID      int       `gorm:"column:aboutus_component_id" json:"aboutus_component_id"`
    AboutUsComponentName    string    `gorm:"column:aboutus_component_name" json:"aboutus_component_name"`
    AboutUsComponentJobdesc string    `gorm:"column:aboutus_component_jobdesc" json:"aboutus_component_jobdesc"`
    AboutUsComponentPhoto   string    `gorm:"column:aboutus_component_photo" json:"aboutus_component_photo"`
    Description             string    `gorm:"column:description" json:"description"`
    Tooltip                 string    `gorm:"column:tooltip" json:"tooltip"`
    Status                  string    `gorm:"column:status" json:"status"`
    CreatedBy               string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt               time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy               string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt               time.Time `gorm:"column:updated_at" json:"updated_at"`
}
