package service

import (
	"database/sql"
	"fmt"

	"github.com/DiansSopandi/goride_be/dto"
	"github.com/DiansSopandi/goride_be/models"
	"github.com/DiansSopandi/goride_be/pkg/utils"
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
	RoleRepo *repository.RoleRepository
}

func NewUserService(userRepo *repository.UserRepository, roleRepo *repository.RoleRepository) *UserService {
	return &UserService{
		UserRepo: userRepo, // repository.NewUserRepository(),
		RoleRepo: roleRepo, // Initialize RoleRepository if needed
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

// func (s *UserService) CreateUser(tx *sql.Tx, user *models.User) (models.User, error) {
func (s *UserService) CreateUser(tx *sql.Tx, createUserDto *dto.UserCreateRequest) (models.User, error) {
	password, _ := utils.HashPassword(createUserDto.Password)
	user := models.User{
		Username: createUserDto.Username,
		Email:    createUserDto.Email,
		Password: password,
	}

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

	res, err := s.UserRepo.CreateUser(tx, &user)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (s *UserService) LoginUser(tx *sql.Tx, loginDto dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	var roleNames []string

	user, errUser := s.UserRepo.GetUserByEmail(loginDto.Email)
	if errUser != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("failed to get user by email: %w", errUser)
	}

	role, errRole := s.RoleRepo.GetRoleByUserID(int(user.ID))
	if errRole != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("failed to get role by user ID: %w", errRole)
	}

	for _, r := range role {
		roleNames = append(roleNames, r.Name)
	}

	if user == nil || !utils.CheckPasswordHash(loginDto.Password, user.Password) {
		return dto.UserLoginResponse{}, fmt.Errorf("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	// userMap := map[string]interface{}{}
	// userBytes, _ := json.Marshal(user)
	// json.Unmarshal(userBytes, &userMap)
	// userMap["token"] = token

	userResponse := dto.UserResponse{
		ID:       uint(user.ID),
		Username: &user.Username,
		Email:    user.Email,
		Roles:    roleNames,
	}

	return dto.UserLoginResponse{
		User:  userResponse,
		Token: token,
	}, nil
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
