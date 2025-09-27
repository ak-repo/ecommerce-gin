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
)

type Response struct {
	RefreshToken string
	RefreshExp   time.Duration
	AccessToken  string
	AccessExp    time.Duration
	User         *models.User
}
type UserAuthService interface {
	Register(input *models.InputUser) (*models.User, error)
	Login(input *models.InputUser) (*Response, error)
}

type userAuthService struct {
	userAuthRepo userrepository.UserAuthRepo
	cfg          *config.Config
}

func NewUserAuthService(userAuthRepo userrepository.UserAuthRepo, cfg *config.Config) UserAuthService {
	return &userAuthService{userAuthRepo: userAuthRepo, cfg: cfg}
}

func (s *userAuthService) Register(input *models.InputUser) (*models.User, error) {

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

	if ok := s.userAuthRepo.GetUserByEmail(user); ok {
		return nil, errors.New("email already taken")

	}

	if err := s.userAuthRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (s *userAuthService) Login(input *models.InputUser) (*Response, error) {

	user := &models.User{
		Email: input.Email,
	}

	if ok := s.userAuthRepo.GetUserByEmail(user); !ok {
		return nil, errors.New("no user found in db")
	}
	log.Println("inpu:", input.Password)
	log.Println("db:", user.PasswordHash)

	if ok := utils.CompareHashAndPassword(input.Password, user.PasswordHash); !ok {
		return nil, errors.New("entered password is not matching")
	}

	accessToken, err := jwtpkg.AccessTokenGenerator(user.Email, user.Username, user.Role, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	refreshToken, err := jwtpkg.RefreshTokenGenerator(user.Email, user.Username, user.Role, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	res := &Response{
		User:         user,
		RefreshToken: refreshToken,
		RefreshExp:   s.cfg.JWT.RefreshExpiration,
		AccessToken:  accessToken,
		AccessExp:    s.cfg.JWT.AccessExpiration,
	}
	return res, nil

}
