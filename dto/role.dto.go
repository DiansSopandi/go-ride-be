package dto

type RoleResponse struct {
	ID          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"createdAt" db:"created_at"`
	UpdatedAt   string `json:"updatedAt" db:"updated_at"`
}

type RoleCreateRequest struct {
	Name        string  `json:"name" example:"user"`
	Description *string `json:"description" example:"user"`
}
