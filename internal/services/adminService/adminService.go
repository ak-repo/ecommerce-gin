package adminservice

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/utils"
	"github.com/ak-repo/ecommerce-gin/internal/dto"
	adminrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/adminRepository"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
)

type AdminService interface {
	AdminLoginService(email, password string) (*dto.LoginResponse, error)
	AdminProfileService(email string) (*dto.ProfileDTO, error)
}

type adminService struct {
	adminRepo adminrepository.AdminRepo
	cfg       *config.Config
}

// New admin auth service
func NewAdminService(adminRepo adminrepository.AdminRepo, cfg *config.Config) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		cfg:       cfg,
	}
}

// Admin login JWT token, Password checking
func (s *adminService) AdminLoginService(email, password string) (*dto.LoginResponse, error) {

	admin, err := s.adminRepo.GetAdminInfo(email)
	if err != nil {
		return nil, err
	}

	if ok := utils.CompareHashAndPassword(password, admin.PasswordHash); !ok {
		return nil, errors.New("entered password is not matching")
	}

	accessToken, err := jwtpkg.AccessTokenGenerator(admin, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	refreshToken, err := jwtpkg.RefreshTokenGenerator(admin, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	res := dto.LoginResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		User:         admin,
		RefreshExp:   s.cfg.JWT.RefreshExpiration,
		AccessExp:    s.cfg.JWT.AccessExpiration,
	}

	return &res, nil

}

// admin profile service
func (s *adminService) AdminProfileService(email string) (*dto.ProfileDTO, error) {

	admin, err := s.adminRepo.GetAdminInfo(email)
	if err != nil {
		return nil, err
	}

	address, err := s.adminRepo.GetAdminAddress(admin.ID)
	if err != nil {
		return nil, err
	}

	profile := dto.ProfileDTO{
		ID:    admin.ID,
		Name:  admin.Username,
		Email: admin.Email,
		Role:  admin.Role,
		Address: dto.AddressDTO{
			ID:          address.ID,
			AddressLine: address.AddressLine,
			City:        address.City,
			State:       address.State,
			PostalCode:  address.PostalCode,
			Country:     address.Country,
		},
	}

	return &profile, nil
}
