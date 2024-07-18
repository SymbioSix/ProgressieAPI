package models

import "time"

type land_navbar_request struct {
    NavComponentID     int `json:"nav_component_id"`
    NavComponentName   string `json:"nav_component_name"`
    NavComponentGroup  int `json:"nav_component_group"`
    NavComponentIcon   string `json:"nav_component_icon"`
    Tooltip            string `json:"tooltip"`
    Endpoint           string `json:"endpoint"`
    CreatedBy          string `json:"created_by"`
    CreatedAt          time.Time `json:"created_at"`
    UpdatedBy          string `json:"updated_by"`
    UpdatedAt          time.Time `json:"updated_at"`
}




type land_navbar_response struct {
	NavComponentID     int `json:"nav_component_id"`
    NavComponentName   string `json:"nav_component_name"`
    NavComponentGroup  int `json:"nav_component_group"`
    NavComponentIcon   string `json:"nav_component_icon"`
    Tooltip            string `json:"tooltip"`
    Endpoint           string `json:"endpoint"`
    CreatedBy          string `json:"created_by"`
    CreatedAt          time.Time `json:"created_at"`
    UpdatedBy          string `json:"updated_by"`
    UpdatedAt          time.Time `json:"updated_at"`
}