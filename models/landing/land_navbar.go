package models

import "time"

type land_navbar_request struct {
    NavComponentID    int       `gorm:"column:nav_component_id" json:"nav_component_id"`
    NavComponentName  string    `gorm:"column:nav_component_name" json:"nav_component_name"`
    NavComponentGroup int       `gorm:"column:nav_component_group" json:"nav_component_group"`
    NavComponentIcon  string    `gorm:"column:nav_component_icon" json:"nav_component_icon"`
    Tooltip           string    `gorm:"column:tooltip" json:"tooltip"`
    Endpoint          string    `gorm:"column:endpoint" json:"endpoint"`
    CreatedBy         string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy         string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type land_navbar_response struct {
    NavComponentID    int       `gorm:"column:nav_component_id" json:"nav_component_id"`
    NavComponentName  string    `gorm:"column:nav_component_name" json:"nav_component_name"`
    NavComponentGroup int       `gorm:"column:nav_component_group" json:"nav_component_group"`
    NavComponentIcon  string    `gorm:"column:nav_component_icon" json:"nav_component_icon"`
    Tooltip           string    `gorm:"column:tooltip" json:"tooltip"`
    Endpoint          string    `gorm:"column:endpoint" json:"endpoint"`
    CreatedBy         string    `gorm:"column:created_by" json:"created_by"`
    CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedBy         string    `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}
