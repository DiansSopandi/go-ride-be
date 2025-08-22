package models

import (
	"time"
)

type User struct {
	ID         int     `json:"id" db:"id"`
	Username   string  `json:"username" db:"username"`
	Email      string  `json:"email" db:"email"`
	Password   string  `json:"password,omitempty" db:"password"`
	Provider   string  `json:"provider,omitempty" db:"provider"`       // e.g., "google", "local"
	ProviderID *string `json:"provider_id,omitempty" db:"provider_id"` // ID from the provider (e.g., Google ID)
	Picture    string  `json:"picture,omitempty" db:"picture"`         // URL to the user's profile picture
	// Roles    string `json:"role" db:"role"` // e.g., "user", "driver", "admin"
	// AvatarUrl  sql.NullString `json:"avatar_url" db:"avatar_url"`
	// AvatarName sql.NullString `json:"avatar_name" db:"avatar_name"`
	// FirstName  sql.NullString `json:"first_name" db:"first_name"`
	// LastName   sql.NullString `json:"last_name" db:"last_name"`
	// Phone      sql.NullString `json:"phone" db:"phone"`
	// Address    sql.NullString `json:"address" db:"address"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // Optional, for soft delete
}
