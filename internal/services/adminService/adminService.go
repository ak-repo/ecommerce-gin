package adminservice

import (
	"errors"
	"strconv"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/utils"
	"github.com/ak-repo/ecommerce-gin/internal/dto"
	adminrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/adminRepository"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"gorm.io/gorm"
)

type AdminService interface {
	AdminLoginService(email, password string) (*dto.LoginResponse, error)
	AdminProfileService(email string) (*dto.ProfileDTO, error)
	AdminAddressUpdateService(email, addressID string, address *dto.AddressDTO) error
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
	if admin == nil {
		return nil, gorm.ErrRecordNotFound
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
	}
	if address != nil {
		profile.Address =
			dto.AddressDTO{
				ID:          address.ID,
				AddressLine: address.AddressLine,
				City:        address.City,
				State:       address.State,
				PostalCode:  address.PostalCode,
				Country:     address.Country,
				Phone:       address.Phone,
			}
	}

	return &profile, nil
}

// admin profile update or add
func (s *adminService) AdminAddressUpdateService(email, addressID string, address *dto.AddressDTO) error {

	admin, err := s.adminRepo.GetAdminInfo(email)
	if err != nil {
		return err
	}
	if admin == nil {
		return gorm.ErrRecordNotFound
	}

	addressUID, _ := strconv.ParseUint(addressID, 10, 64)

	if addressUID == 0 {

		return s.adminRepo.AddAdminAdress(admin.ID, address)

	} else {
		uID := uint(addressUID)
		return s.adminRepo.UpdateAdminAddress(uID, address)
	}
}
