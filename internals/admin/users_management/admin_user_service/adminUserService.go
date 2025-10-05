package adminuserservice

import (
	"errors"
	"fmt"
	"strings"

	usersmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/users_management"
	adminuserinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_interface"
)

type AdminUserService struct {
	AdminUserRepo adminuserinterface.RepoInterface
}

func NewAdminUserService(adminUserRepo adminuserinterface.RepoInterface) adminuserinterface.ServiceInterface {
	return &AdminUserService{AdminUserRepo: adminUserRepo}
}

// all users service
func (s *AdminUserService) AdminAllUsersService(query string) ([]usersmanagement.AdminUserListDTO, error) {
	users, err := s.AdminUserRepo.AdminGetAllUsers()
	if err != nil {
		return nil, err
	}
	var usersDTO []usersmanagement.AdminUserListDTO
	for _, u := range users {
		if strings.Contains(strings.ToLower(u.Username), strings.ToLower(query)) {
			user := usersmanagement.AdminUserListDTO{
				ID:            u.ID,
				Username:      u.Username,
				Email:         u.Email,
				Role:          u.Role,
				Status:        u.Status,
				EmailVerified: u.EmailVerified,
			}
			usersDTO = append(usersDTO, user)
		}

	}

	return usersDTO, nil
}

// Get on user by id
func (s *AdminUserService) AdminGetUserByIDService(userID uint) (*usersmanagement.AdminUserDTO, error) {

	u, err := s.AdminUserRepo.AdminFindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}
	userDTO := usersmanagement.AdminUserDTO{
		ID:            u.ID,
		Username:      u.Username,
		Email:         u.Email,
		Role:          u.Role,
		Status:        u.Status,
		EmailVerified: u.EmailVerified,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}

	return &userDTO, nil

}

// user role change
func (s *AdminUserService) AdminUserRoleChangeService(user *usersmanagement.AdminUserRoleChange) error {

	validRoles := map[string]bool{"user": true, "store": true, "admin": true, "delivery": true}
	if !validRoles[user.Role] {
		return fmt.Errorf("invalid role: %s", user.Role)
	}
	return s.AdminUserRepo.AdminUserRoleChange(user)
}

// user block and unblock
func (s *AdminUserService) AdminUserBlockService(userID uint) error {

	user, err := s.AdminUserRepo.AdminFindUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if user.Status == "active" {
		user.Status = "blocked"
	} else {
		user.Status = "active"
	}
	return s.AdminUserRepo.AdminUserBlock(user)
}

// user seach \
func (s *AdminUserService) AdminUserSearchService(query string) ([]usersmanagement.AdminUserListDTO, error) {

	data, err := s.AdminUserRepo.AdminGetAllUsers()
	if err != nil {
		return nil, err
	}

	var users []usersmanagement.AdminUserListDTO
	for _, u := range data {
		if strings.Contains(strings.ToLower(u.Username), strings.ToLower(query)) {
			user := usersmanagement.AdminUserListDTO{
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
