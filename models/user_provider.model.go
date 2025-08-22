package models

type UserProvider struct {
	ID            uint   `json:"id" db:"id"`
	UserID        uint   `json:"user_id" db:"user_id"`
	Provider      string `json:"provider" db:"provider"`
	ProviderID    string `json:"provider_id" db:"provider_id"`
	ProviderEmail string `json:"provider_email" db:"provider_email"`
}
