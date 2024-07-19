package models

import "time"

type land_footer_request struct {
    FooterComponentID    int       `json:"footer_component_id"`   // Primary key
    FooterComponentName  string    `json:"footer_component_name"`
    FooterComponentGroup int       `json:"footer_component_group"` // Foreign key
    FooterComponentIcon  string    `json:"footer_component_icon"`
    Tooltip              string    `json:"tooltip"`
    Endpoint             string    `json:"endpoint"`
    CreatedBy            string    `json:"created_by"`
    CreatedAt            time.Time `json:"created_at"`
    UpdatedBy            string    `json:"updated_by"`
    UpdatedAt            time.Time `json:"updated_at"`
}

type land_footer_response struct {
    FooterComponentID    int       `json:"footer_component_id"`
    FooterComponentName  string    `json:"footer_component_name"`
    FooterComponentGroup int       `json:"footer_component_group"`
    FooterComponentIcon  string    `json:"footer_component_icon"`
    Tooltip              string    `json:"tooltip"`
    Endpoint             string    `json:"endpoint"`
    CreatedBy            string    `json:"created_by"`
    CreatedAt            time.Time `json:"created_at"`
    UpdatedBy            string    `json:"updated_by"`
    UpdatedAt            time.Time `json:"updated_at"`
}