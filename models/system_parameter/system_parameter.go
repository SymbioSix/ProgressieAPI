package models

import "time"

type SystemParameter struct {
	ParameterID    int       `gorm:"column:parameter_id;primaryKey"`
	ParameterName  string    `gorm:"column:parameter_name"`
	ParameterValue string    `gorm:"column:parameter_value"`
	Status         string    `gorm:"column:status"`
	CreatedBy      string    `gorm:"column:created_by"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedBy      string    `gorm:"column:updated_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (system *SystemParameter) TableName() string {
	return "system_parameter"
}
