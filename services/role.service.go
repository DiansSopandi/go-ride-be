package service

import (
	"database/sql"

	"github.com/DiansSopandi/goride_be/models"
	"github.com/DiansSopandi/goride_be/repository"
)

type RoleService struct {
	Repo *repository.RoleRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		Repo: roleRepo,
	}
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.Repo.GetAllRoles()
}

func (s *RoleService) CreateRoles(tx *sql.Tx, role *models.Role) (models.Role, error) {
	return s.Repo.CreateRoles(tx, role)
}
