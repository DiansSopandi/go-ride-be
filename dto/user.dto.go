package dto

import "time"

type UserCreateRequest struct {
	Username string   `json:"username" example:"John Doe"`
	Email    string   `json:"email" example:"Q2Sb9@example.com"`
	Password string   `json:"password" example:"Cilok99!@"`
	Roles    []string `json:"roles" example:"admin,driver,user,superadmin"` // multiple roles
	// AvatarUrl  string `json:"avatarUrl" db:"avatar_url"`
	// AvatarName string `json:"avatarName" db:"avatar_name"`
	// FirstName  string `json:"firstName" db:"first_name"`
	// LastName   string `json:"lastName" db:"last_name"`
	// Phone      string `json:"phone" db:"phone"`
	// Address    string `json:"address" db:"address"`
}

type UserRegisterRequest struct {
	Username        string   `json:"username" validate:"required,min=3,max=50" example:"John Doe"`
	Email           string   `json:"email" validate:"required,email" example:"Q2Sb9@example.com"`
	Password        string   `json:"password" validate:"required,min=8" example:"Cilok99!@"`
	PasswordConfirm string   `json:"password_confirm" validate:"required,eqfield=Password" example:"Cilok99!@"`
	Roles           []string `json:"roles" example:"user" validate:"required"`
	// AvatarUrl  string `json:"avatarUrl" db:"avatar_url"`
	// AvatarName string `json:"avatarName" db:"avatar_name"`
	// FirstName  string `json:"firstName" db:"first_name"`
	// LastName   string `json:"lastName" db:"last_name"`
	// Phone      string `json:"phone" db:"phone"`
	// Address    string `json:"address" db:"address"`
}

type UserUpdateRequest struct {
	Username string   `json:"username,omitempty" validate:"omitempty,min=3,max=50"`
	Email    string   `json:"email,omitempty" validate:"omitempty,email"`
	Password string   `json:"password,omitempty" validate:"omitempty,min=8"`
	Roles    []string `json:"roles,omitempty"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UserResponse struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Roles     []string   `json:"roles"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
