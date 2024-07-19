package models

import "time"
  
type land_hero_request struct {
	HeroComponentID      int       `json:"hero_component_id"`      // Primary key
	HeroComponentTitle   string    `json:"hero_component_title"`
	HeroComponentSubtitle string   `json:"hero_component_subtitle"`
	HeroComponentImage   string    `json:"hero_component_image"`
	HeroComponentCoverImg string   `json:"hero_component_cover_img"`
	Tooltip              string    `json:"tooltip"`
	CreatedBy            string    `json:"created_by"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedBy            string    `json:"updated_by"`
	UpdatedAt            time.Time `json:"updated_at"`
}
  
type land_hero_response struct {
	HeroComponentID      int       `json:"hero_component_id"`
	HeroComponentTitle   string    `json:"hero_component_title"`
	HeroComponentSubtitle string   `json:"hero_component_subtitle"`
	HeroComponentImage   string    `json:"hero_component_image"`
	HeroComponentCoverImg string   `json:"hero_component_cover_img"`
	Tooltip              string    `json:"tooltip"`
	CreatedBy            string    `json:"created_by"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedBy            string    `json:"updated_by"`
	UpdatedAt            time.Time `json:"updated_at"`
}
  