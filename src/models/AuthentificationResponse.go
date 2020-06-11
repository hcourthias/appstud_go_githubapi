package models

// AuthentificationResponse basic /api response
type AuthentificationResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
