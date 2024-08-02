package models



type Achievement struct {
    UserID                   string `gorm:"user_id;column:user_id"`
    AchievementTitle         string `gorm:"achievement_title;column:achievement_title"`
    AchievementIcon          string `gorm:"column:achievement_icon"`
    AchievementDescription   string `gorm:"column:achievement_description"`
    AchievementCategory      string `gorm:"column:achievement_category"`
    IsAchieved               bool   `gorm:"column:is_achieved"`
    AchievementTotalAchieved int    `gorm:"column:achievement_total_achieved"` // New field
}
