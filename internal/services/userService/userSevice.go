package userservice

import (
	"errors"
	"strconv"
	"time"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/utils"
	"github.com/ak-repo/ecommerce-gin/internal/dto"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"gorm.io/gorm"
)

type Response struct {
	RefreshToken string
	RefreshExp   time.Duration
	AccessToken  string
	AccessExp    time.Duration
	User         *models.User
}
type UserService interface {
	RegisterService(input *dto.RegisterRequest) error
	LoginService(input *dto.LoginRequest) (*dto.LoginResponse, error)
	UserProfileService(email string) (*dto.ProfileDTO, error)
	UserAddressUpdateService(address *dto.AddressDTO, addressID, email string) error
	UserPasswordChangeService(email, newPassword, oldpassword string) error
}

type userService struct {
	userRepo userrepository.UserRepo
	cfg      *config.Config
}

func NewUserService(userRepo userrepository.UserRepo, cfg *config.Config) UserService {
	return &userService{userRepo: userRepo, cfg: cfg}
}

func (s *userService) RegisterService(input *dto.RegisterRequest) error {

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	if user, err := s.userRepo.GetUserByEmail(input.Email); user != nil && err != nil {
		return errors.New("email already taken")
	}

	if err := s.userRepo.CreateUser(input.Username, input.Email, hash); err != nil {
		return err
	}

	return nil

}

func (s *userService) LoginService(input *dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gorm.ErrRecordNotFound
	}

	if ok := utils.CompareHashAndPassword(input.Password, user.PasswordHash); !ok {
		return nil, errors.New("entered password is not matching")
	}

	accessToken, err := jwtpkg.AccessTokenGenerator(user, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	refreshToken, err := jwtpkg.RefreshTokenGenerator(user, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	res := dto.LoginResponse{
		User:         user,
		RefreshToken: refreshToken,
		RefreshExp:   s.cfg.JWT.RefreshExpiration,
		AccessToken:  accessToken,
		AccessExp:    s.cfg.JWT.AccessExpiration,
	}
	return &res, nil

}

func (s *userService) UserProfileService(email string) (*dto.ProfileDTO, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	address, err := s.userRepo.GetUserAddress(user.ID)
	if err != nil {
		return nil, err
	}

	profile := dto.ProfileDTO{
		ID:    user.ID,
		Name:  user.Username,
		Email: user.Email,
		Role:  user.Role,
	}
	if address != nil {
		profile.Address = dto.AddressDTO{
			Phone:       address.Phone,
			ID:          address.ID,
			AddressLine: address.AddressLine,
			City:        address.City,
			State:       address.State,
			PostalCode:  address.PostalCode,
			Country:     address.Country,
		}
	}

	return &profile, nil

}

func (s *userService) UserAddressUpdateService(address *dto.AddressDTO, addressID, email string) error {

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	addressUID, err := strconv.ParseUint(addressID, 10, 64)
	if err != nil {
		return err
	}
	if addressUID == 0 {
		return s.userRepo.AddAddress(address, user.ID)
	} else {
		address.ID = uint(addressUID)
		return s.userRepo.UpdateAddress(address)
	}

}

// User password chnage service
func (s *userService) UserPasswordChangeService(email, newPassword, oldpassword string) error {

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	if user == nil {
		return gorm.ErrRecordNotFound
	}
	if ok := utils.CompareHashAndPassword(oldpassword, user.PasswordHash); !ok {
		return errors.New("incorrect password")
	}

	hashPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return s.userRepo.UpdatePassword(user.ID, string(hashPassword))
}
