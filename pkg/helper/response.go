package pkg

import (
	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/models"
)

func ToUserResponse(user models.User, roles []string) dto.UserResponse {
	return dto.UserResponse{
		ID:    uint(user.ID),
		Email: user.Email,
		Roles: roles,
	}
}
