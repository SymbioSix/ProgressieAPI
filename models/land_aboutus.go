package models

import "time"



type land_aboutus_request struct {
    AboutUsComponentID   int            `json:"aboutus_component_id"` // Primary key
    AboutUsComponentName string         `json:"aboutus_component_name"`
    AboutUsComponentJobdesc string      `json:"aboutus_component_jobdesc"`
    AboutUsComponentPhoto string        `json:"aboutus_component_photo"`
    Description          string         `json:"description"`
    Tooltip              string         `json:"tooltip"`
    Status               string			`json:"status"`
    CreatedBy            string         `json:"created_by"`
    CreatedAt            time.Time      `json:"created_at"`
    UpdatedBy            string         `json:"updated_by"`
    UpdatedAt            time.Time      `json:"updated_at"`
}

type land_aboutus_response struct {
    AboutUsComponentID   int            `json:"aboutus_component_id"`
    AboutUsComponentName string         `json:"aboutus_component_name"`
    AboutUsComponentJobdesc string      `json:"aboutus_component_jobdesc"`
    AboutUsComponentPhoto string        `json:"aboutus_component_photo"`
    Description          string         `json:"description"`
    Tooltip              string         `json:"tooltip"`
    Status               string 		`json:"status"`
    CreatedBy            string         `json:"created_by"`
    CreatedAt            time.Time      `json:"created_at"`
    UpdatedBy            string         `json:"updated_by"`
    UpdatedAt            time.Time      `json:"updated_at"`
}
