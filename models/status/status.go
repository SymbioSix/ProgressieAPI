package models

type StatusModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type HealthMap struct {
	DatabaseStatus    string `json:"database_status"`
	SupabaseAPIStatus string `json:"supabase_api_status"`
	OverallStatus     string `json:"overall_status"`
}
