package models

import (
	"time"
)

type Role struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	// IsActive    bool       `json:"is_active" db:"is_active"`
	CreatedAt string     `json:"created_at" db:"created_at"`
	UpdatedAt string     `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // Optional, for soft delete
}
