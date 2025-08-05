package models

import (
	"encoding/json"
	"fmt"
)

type UserRole struct {
	ID     uint `json:"id" db:"id"`
	UserID uint `json:"user_id" db:"user_id"`
	RoleID uint `json:"role_id" db:"role_id"`
}

type UserWithRoles struct {
	User
	Roles []Role `json:"roles,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}

func (r *Role) TableName() string {
	return "roles"
}

func (ur *UserRole) TableName() string {
	return "user_roles"
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Password string `json:"-"`
		*Alias
	}{
		Alias: (*Alias)(&u),
	})
}

func (u *User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}
	if len(u.Username) < 3 || len(u.Username) > 50 {
		return fmt.Errorf("username must be between 3 and 50 characters")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (r *Role) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("role name is required")
	}
	if len(r.Name) < 2 || len(r.Name) > 50 {
		return fmt.Errorf("role name must be between 2 and 50 characters")
	}
	return nil
}
