package models

// HealthCheckResponse basic /api response
type HealthCheckResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Time    int64  `json:"time"`
}
