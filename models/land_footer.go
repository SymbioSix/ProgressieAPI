package models

import "time"

type land_footer_request struct {
    FooterComponentID    int       `gorm:"column:footer_component_id" json:"footer_component_id"`   // Primary key
    FooterComponentName  string    `gorm:"column:footer_component_name" json:"footer_component_name"`
    FooterComponentGroup int       `gorm:"column:footer_component_group" json:"footer_component_group"` // Foreign key
    FooterComponentIcon  string    `gorm:"column:footer_component_icon" json:"footer_component_icon"`
    Tooltip              string    `gorm:"column:tooltip" json:"tooltip"`
    Endpoint             string    `gorm:"column:endpoint" json:"endpoint"`
    CreatedBy            string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy            string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type land_footer_response struct {
    FooterComponentID    int       `gorm:"column:footer_component_id" json:"footer_component_id"`
    FooterComponentName  string    `gorm:"column:footer_component_name" json:"footer_component_name"`
    FooterComponentGroup int       `gorm:"column:footer_component_group" json:"footer_component_group"`
    FooterComponentIcon  string    `gorm:"column:footer_component_icon" json:"footer_component_icon"`
    Tooltip              string    `gorm:"column:tooltip" json:"tooltip"`
    Endpoint             string    `gorm:"column:endpoint" json:"endpoint"`
    CreatedBy            string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy            string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
}
