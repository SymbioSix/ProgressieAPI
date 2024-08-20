package models

import (
	"time"
	"github.com/google/uuid"

)

// AchiAll represents the achi_all table
type AchiAll struct {
	AchievementID    uuid.UUID `gorm:"primaryKey;type:uuid;not null"` // PK, FK
	AchievementCrsID uuid.UUID `gorm:"primaryKey;type:uuid;not null"` // PK, FK
	CreatedAt        time.Time `gorm:"autoCreateTime"`                // Optional: if you want to track creation time
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`                // Optional: if you want to track update time
}

