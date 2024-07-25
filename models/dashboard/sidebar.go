package models

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/auth"
)

type SidebarModel struct {
	ID           string    `gorm:"column:id" json:"id"`
	SidebarName  string    `gorm:"column:sidebarmenu_name" json:"sidebar_name"`
	SidebarGroup string    `gorm:"column:sidebarmenu_group" json:"sidebar_group,omitempty"`
	Endpoint     string    `gorm:"column:endpoint" json:"endpoint"`
	Status       string    `gorm:"column:status" json:"status,omitempty"`
	CreatedBy    string    `gorm:"column:created_by"`
	UpdatedBy    string    `gorm:"column:updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type RoleSidebarResponse struct {
	RoleData    models.RoleModel `json:"role_data"`
	SidebarData []SidebarModel   `json:"sidebar_data"`
}
