package userservice

import (
	"errors"
	"log"
	"time"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/utils"
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
	Register(input *models.InputUser) (*models.User, error)
	Login(input *models.InputUser) (*Response, error)
	UserProfileService(email string) (*models.User, error)
	UserProfileUpdateService(email string, addresID string, address *models.Address) error
}

type userService struct {
	userRepo userrepository.UserRepo
	cfg      *config.Config
}

func NewUserService(userRepo userrepository.UserRepo, cfg *config.Config) UserService {
	return &userService{userRepo: userRepo, cfg: cfg}
}

func (s *userService) Register(input *models.InputUser) (*models.User, error) {

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        input.Email,
		PasswordHash: hash,
		Role:         "customer",
		IsActive:     true,
	}

	if err := s.userRepo.GetUserByEmail(&models.User{}, input.Email); err != nil {
		return nil, errors.New("email already taken")

	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (s *userService) Login(input *models.InputUser) (*Response, error) {

	user := models.User{}

	if err := s.userRepo.GetUserByEmail(&user,input.Email); err != nil {
		return nil, err
	}

	if ok := utils.CompareHashAndPassword(input.Password, user.PasswordHash); !ok {
		return nil, errors.New("entered password is not matching")
	}

	log.Println(&user.Username, "user")

	accessToken, err := jwtpkg.AccessTokenGenerator(&user, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	refreshToken, err := jwtpkg.RefreshTokenGenerator(&user, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	res := &Response{
		User:         &user,
		RefreshToken: refreshToken,
		RefreshExp:   s.cfg.JWT.RefreshExpiration,
		AccessToken:  accessToken,
		AccessExp:    s.cfg.JWT.AccessExpiration,
	}
	return res, nil

}

type UserProfle struct {
	User    *models.User
	Address *models.Address
}

func (s *userService) UserProfileService(email string) (*models.User, error) {
	user := &models.User{}

	// Properly fetch user by email
	if err := s.userRepo.GetUserByEmail(user, email); err != nil {
		return nil, err
	}

	// Load profile details (addresses, etc.)
	if err := s.userRepo.GetUserProfile(user); err != nil {
		// not fatal if no addresses found, just return user with empty slice
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Always return a valid slice (avoid nil in template)
	if user.Addresses == nil {
		user.Addresses = []models.Address{}
	}

	return user, nil
}

// User profile add or update
func (s *userService) UserProfileUpdateService(email string, addresID string, address *models.Address) error {

	user := models.User{}
	if err := s.userRepo.GetUserByEmail(&user, email); err != nil {
		return err
	}

	address.UserID = user.ID
	if addresID == "0" {
		if err := s.userRepo.AddUserProfile(address); err != nil {
			return err
		}
	} else {
		if err := s.userRepo.UpdateUserProfile(addresID, address); err != nil {
			return err
		}

	}

	return nil

}
