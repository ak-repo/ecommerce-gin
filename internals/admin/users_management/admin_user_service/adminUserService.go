package adminuserservice

import (
	"errors"

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
func (s *AdminUserService) AdminAllUsersService() ([]usersmanagement.AdminUserListDTO, error) {
	users, err := s.AdminUserRepo.AdminGetAllUsers()
	if err != nil {
		return nil, err
	}
	var usersDTO []usersmanagement.AdminUserListDTO
	for _, u := range users {
		user := usersmanagement.AdminUserListDTO{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Role:     u.Role,
			Status:   u.Status,
		}
		usersDTO = append(usersDTO, user)
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
