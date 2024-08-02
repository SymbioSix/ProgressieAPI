// controllers/achievement.go
package dashboard

import (
    "encoding/json"
    "net/http"
    "strconv"
    models "github.com/SymbioSix/ProgressieAPI/models/dashboard"
)

var achievements []models.Achievement

// CreateAchievement handles POST requests to create a new achievement
func CreateAchievement(w http.ResponseWriter, r *http.Request) {
    var achievement models.Achievement
    _ = json.NewDecoder(r.Body).Decode(&achievement)
    achievement.ID = len(achievements) + 1
    achievements = append(achievements, achievement)
    json.NewEncoder(w).Encode(achievement)
}

// GetAchievements handles GET requests to fetch all achievements
func GetAchievements(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(achievements)
}

// GetAchievement handles GET requests to fetch a single achievement by ID
func GetAchievement(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, item := range achievements {
        if item.ID == id {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.NotFound(w, r)
}

// UpdateAchievement handles PUT requests to update an existing achievement
func UpdateAchievement(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, item := range achievements {
        if item.ID == id {
            achievements = append(achievements[:index], achievements[index+1:]...)
            var achievement models.Achievement
            _ = json.NewDecoder(r.Body).Decode(&achievement)
            achievement.ID = id
            achievements = append(achievements, achievement)
            json.NewEncoder(w).Encode(achievement)
            return
        }
    }
    http.NotFound(w, r)
}

// DeleteAchievement handles DELETE requests to remove an achievement by ID
func DeleteAchievement(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, item := range achievements {
        if item.ID == id {
            achievements = append(achievements[:index], achievements[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(achievements)
}

// GetAchievementProgress handles GET requests to fetch the progress of an achievement by ID
func GetAchievementProgress(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, item := range achievements {
        if item.ID == id {
            progress := item.CalculateProgress()
            json.NewEncoder(w).Encode(map[string]float64{"progress": progress})
            return
        }
    }
    http.NotFound(w, r)
}
