package models

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/auth"
)

type SidebarModel struct {
	SidebarID    string    `gorm:"column:sidebarmenu_id" json:"sidebarmenu_id"`
	SidebarName  string    `gorm:"column:sidebarmenu_name" json:"sidebar_name"`
	SidebarGroup string    `gorm:"column:sidebarmenu_group" json:"sidebar_group,omitempty"`
	Endpoint     string    `gorm:"column:endpoint" json:"endpoint"`
	Status       string    `gorm:"column:status" json:"status,omitempty"`
	CreatedBy    string    `gorm:"column:created_by"`
	UpdatedBy    string    `gorm:"column:updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (dash *SidebarModel) TableName() string {
	return "usr_sidebarmenu"
}

type RoleSidebarResponse struct {
	RoleID      string           `gorm:"column:role_id;primaryKey" json:"role_id,omitempty"`
	RoleData    models.RoleModel `gorm:"foreignKey:RoleID;references:RoleID" json:"role_data"`
	SidebarID   string           `gorm:"column:sidebarmenu_id;primaryKey" json:"sidebarmenu_id,omitempty"`
	SidebarData SidebarModel     `gorm:"foreignKey:SidebarID;references:SidebarID" json:"sidebar_data"`
	AllowView   bool             `json:"allow_view,omitempty"`
	AllowAdd    bool             `json:"allow_add,omitempty"`
	AllowEdit   bool             `json:"allow_edit,omitempty"`
	AllowDelete bool             `json:"allow_delete,omitempty"`
}

func (dash *RoleSidebarResponse) TableName() string {
	return "usr_rolesidebarmenu"
}
