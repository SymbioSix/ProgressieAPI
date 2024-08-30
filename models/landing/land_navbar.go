package models

import "time"

type Land_Navbar struct {
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

func (l *Land_Navbar) TableName() string {
	return "land_navbar"
}

type LandNavbarRequest struct {
	NavName  string `json:"nav_name"`
	NavGroup int    `json:"nav_group"`
	NavIcon  string `json:"nav_icon"`
	Tooltip  string `json:"tooltip"`
	Endpoint string `json:"endpoint"`
}
