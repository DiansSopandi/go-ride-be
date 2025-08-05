package service

import (
	"database/sql"
	"fmt"

	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/models"
	"github.com/DiansSopandi/goride_be/repository"
)

// type GalleryService struct {
// 	repo *repositories.GalleryRepository
// }

// func NewGalleryService() *GalleryService {
// 	return &GalleryService{
// 		// Initialize any dependencies or configurations here
// 		repo: repositories.NewGalleryRepository(),
// 	}
// }

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo, // repository.NewUserRepository(),
	}
}

func (s *UserService) BeginTransaction() (*sql.Tx, error) {
	return s.UserRepo.BeginTransaction()
}

func (s *UserService) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(tx *sql.Tx, user *models.User) (models.User, error) {
	exists, err := s.UserRepo.CheckUsernameExistsWithTx(tx, user.Username)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to check username: %w", err)
	}
	if exists {
		return models.User{}, fmt.Errorf("username already exists")
	}

	exists, err = s.UserRepo.CheckEmailExistsWithTx(tx, user.Email)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to check email: %w", err)
	}
	if exists {
		return models.User{}, fmt.Errorf("email already exists")
	}

	res, err := s.UserRepo.CreateUser(tx, user)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (s *UserService) UpdateUser(tx *sql.Tx, user *models.User) error {
	err := s.UserRepo.UpdateUser(tx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := s.UserRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) ValidateRolesExist(tx *sql.Tx, roleNames []string) ([]int64, error) {
	return s.UserRepo.ValidateRolesExist(tx, roleNames)
}

func (s *UserService) AssignRolesToUserWithTx(tx *sql.Tx, userID uint, roleIDs []int64) error {
	return s.UserRepo.AssignRolesToUserWithTx(tx, userID, roleIDs)
}
