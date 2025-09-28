package adminservice

import (
	"errors"
	"time"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/utils"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	adminrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/adminRepository"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
)

type AdminAuthService interface {
	AdminLoginService(input *models.InputUser) (*AdminResponse, error)
}

type adminAuthService struct {
	adminAuthRepo adminrepository.AdminAuthRepo
	cfg           *config.Config
}

type AdminResponse struct {
	RefreshToken string
	RefreshExp   time.Duration
	AccessToken  string
	AccessExp    time.Duration
	User         *models.User
}

// New admin auth service
func NewAdminAuthService(adminAuthRepo adminrepository.AdminAuthRepo, cfg *config.Config) AdminAuthService {
	return &adminAuthService{
		adminAuthRepo: adminAuthRepo,
		cfg:           cfg,
	}
}

// Admin login JWT token, Password checking
func (s *adminAuthService) AdminLoginService(input *models.InputUser) (*AdminResponse, error) {

	admin := models.User{
		Email: input.Email,
	}

	if err := s.adminAuthRepo.GetAdminInfo(&admin); err != nil {
		return nil, err
	}

	if ok := utils.CompareHashAndPassword(input.Password, admin.PasswordHash); !ok {
		return nil, errors.New("entered password is not matching")
	}

	accessToken, err := jwtpkg.AccessTokenGenerator(&admin, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	refreshToken, err := jwtpkg.RefreshTokenGenerator(&admin, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	res := AdminResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		User:         &admin,
		RefreshExp:   s.cfg.JWT.RefreshExpiration,
		AccessExp:    s.cfg.JWT.AccessExpiration,
	}

	return &res, nil

}
