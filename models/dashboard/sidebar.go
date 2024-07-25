package models

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/auth"
)

type SidebarModel struct {
	ID           string    `gorm:"column:id"`
	SidebarName  string    `gorm:"column:sidebarmenu_name"`
	SidebarGroup string    `gorm:"column:sidebarmenu_group"`
	Endpoint     string    `gorm:"column:endpoint"`
	Status       string    `gorm:"column:status"`
	CreatedBy    string    `gorm:"column:created_by"`
	UpdatedBy    string    `gorm:"column:updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type RoleSidebarResponse struct {
	RoleData    models.RoleModel `json:"role_data"`
	SidebarData []SidebarModel   `json:"sidebar_data"`
}
