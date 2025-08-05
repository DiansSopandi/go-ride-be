package repository

import (
	"database/sql"

	"github.com/DiansSopandi/goride_be/db"
	"github.com/DiansSopandi/goride_be/models"
)

type RoleRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

// func NewRoleRepository(withTransaction bool) (*RoleRepository, error) {
func NewRoleRepository(tx *sql.Tx) (*RoleRepository, error) {
	// dbConn := db.InitDatabase()
	// var trx *sql.Tx

	// if withTransaction {
	// 	var err error
	// 	trx, err = dbConn.Begin()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	// 	}
	// }

	return &RoleRepository{
		DB: db.InitDatabase(),
		TX: tx,
	}, nil
}

func (r *RoleRepository) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role

	query := `SELECT id, name, description, created_at, updated_at 
			  FROM roles 
			  WHERE deleted_at IS NULL
			  ORDER BY name ASC`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RoleRepository) CreateRoles(tx *sql.Tx, role *models.Role) (models.Role, error) {
	query := `INSERT INTO roles (name, description) VALUES ($1, $2) 
	RETURNING id, name, description, created_at, updated_at`

	err := tx.QueryRow(query, role.Name, role.Description).Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	return *role, err
}
