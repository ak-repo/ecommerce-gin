package usersservice

import (
	"errors"
	"fmt"
	"strings"

	userdto "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/user_dto"
	usersinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/user_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"gorm.io/gorm"
)

type service struct {
	UsersRepo usersinterface.Repository
}

func NewUsersService(repo usersinterface.Repository) usersinterface.Service {
	return &service{UsersRepo: repo}
}

func (s *service) GetAllUsers(query string) ([]userdto.AdminUserListDTO, error) {
	data, err := s.UsersRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var users []userdto.AdminUserListDTO
	for _, u := range data {
		if strings.Contains(strings.ToLower(u.Username), strings.ToLower(query)) {
			user := userdto.AdminUserListDTO{
				ID:            u.ID,
				Username:      u.Username,
				Email:         u.Email,
				Role:          u.Role,
				Status:        u.Status,
				EmailVerified: u.EmailVerified,
			}
			users = append(users, user)
		}

	}
	return users, nil
}

func (s *service) GetUserByID(userID uint) (*userdto.AdminUserDTO, error) {

	data, err := s.UsersRepo.GetUserByID(userID)
	if data == nil || err != nil {
		return nil, errors.New("user not found")
	}
	user := userdto.AdminUserDTO{
		ID:            data.ID,
		Username:      data.Username,
		Email:         data.Email,
		Role:          data.Role,
		Status:        data.Status,
		EmailVerified: data.EmailVerified,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	return &user, nil

}

func (s *service) ChangeUserRole(req *userdto.AdminUserRoleChange) error {

	validRoles := map[string]bool{"customer": true, "admin": true}
	if !validRoles[req.Role] {
		return fmt.Errorf("invalid role: %s", req.Role)
	}
	user := models.User{
		Model: gorm.Model{ID: req.ID},
		Role:  req.Role,
	}
	return s.UsersRepo.UpdateUser(&user)
}

func (s *service) BlockUser(userID uint) error {

	user, err := s.UsersRepo.GetUserByID(userID)
	if user == nil || err != nil {
		return errors.New("user not found")
	}
	if user.Status == "active" {
		user.Status = "inactive"
	} else {
		user.Status = "active"
	}
	return s.UsersRepo.UpdateUser(user)
}

func (s *service) CreateUser(req *userdto.CreateUserRequest) (uint, error) {

	hash, err := utils.HashPassword(req.ConfirmPassword)
	if err != nil {
		return 0, err
	}

	user := models.User{
		Email:         req.Email,
		PasswordHash:  hash,
		Username:      req.Username,
		Role:          req.Role,
		Status:        req.Status,
		EmailVerified: req.EmailVerified,
	}

	if err := s.UsersRepo.CreateUser(&user); err != nil {
		return 0, err
	}
	return user.ID, nil
}
