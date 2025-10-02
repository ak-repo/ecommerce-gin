package userservice

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/internal/dto"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
)

type AdminUserService interface {
	AdminAllUsersService() ([]dto.AdminUserListDTO, error)
	AdminGetUserByIDService(userID uint) (*dto.AdminUserDTO, error)
}

type adminUserService struct {
	adminUserRepo userrepository.AdminUserRepo
}

func NewAdminUserService(adminUserRepo userrepository.AdminUserRepo) AdminUserService {
	return &adminUserService{adminUserRepo: adminUserRepo}
}

// all users service
func (s *adminUserService) AdminAllUsersService() ([]dto.AdminUserListDTO, error) {
	users, err := s.adminUserRepo.AdminGetAllUsers()
	if err != nil {
		return nil, err
	}
	var usersDTO []dto.AdminUserListDTO
	for _, u := range users {
		user := dto.AdminUserListDTO{
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
func (s *adminUserService) AdminGetUserByIDService(userID uint) (*dto.AdminUserDTO, error) {

	u, err := s.adminUserRepo.AdminFindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}
	userDTO := dto.AdminUserDTO{
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
