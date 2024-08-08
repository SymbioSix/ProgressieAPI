package models

import models "github.com/SymbioSix/ProgressieAPI/models/Todo"

type ParseCustomTargetFromAchievement struct {
	CustomTargets []models.TdCustomTarget `json:"custom_targets"`
}
