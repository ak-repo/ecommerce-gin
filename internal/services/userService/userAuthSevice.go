package userservice

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
)

type UserAuthService interface {
}

type userAuthService struct {
	userAuthRepo userrepository.UserAuthRepo
	cfg          *config.Config
}

func NewUserAuthService(userAuthRepo userrepository.UserAuthRepo, cfg *config.Config) UserAuthService {
	return &userAuthService{userAuthRepo: userAuthRepo, cfg: cfg}
}

func (s *userAuthService) Register(input *models.InputUser) (*models.User, error) {

	// hash password
	hash := input.Password

	user := &models.User{
		Email:        input.Email,
		PasswordHash: hash,
		Role:         "customer",
		IsActive:     true,
	}

	if err := s.userAuthRepo.CreateUser(user); err != nil {

		// response
	}
	return user, nil

}
