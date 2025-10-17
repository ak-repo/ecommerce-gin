package authservice

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internals/auth"
	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
)

type authService struct {
	authRepo authinterface.Repository
	cfg      *config.Config
}

func NewAuthService(authRepo authinterface.Repository, cfg *config.Config) authinterface.Service {
	return &authService{authRepo: authRepo, cfg: cfg}
}

// -------------------------------------------- Registration service -------------------------------------------------------
func (s *authService) Registeration(input *auth.RegisterRequest, role string) error {

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	if user, err := s.authRepo.GetUserByEmail(input.Email); user != nil && err != nil {
		return errors.New("email already taken")
	}

	user := models.User{
		Email:         input.Email,
		Username:      input.Username,
		PasswordHash:  hash,
		Role:          role,
		Status:        "active",
		EmailVerified: false,
	}

	return s.authRepo.Registeration(&user)
}

// -----------------------------------------------------Login service ----------------------------------------------------
func (s *authService) Login(input *auth.LoginRequest, role string) (*auth.LoginResponse, error) {

	user, err := s.authRepo.GetUserByEmail(input.Email)
	if user == nil || err != nil {
		return nil, errors.New("user not found")
	}

	if ok := utils.CompareHashAndPassword(input.Password, user.PasswordHash); !ok {
		return nil, errors.New("entered password is not matching")
	}
	if user.Status != "active" {
		return nil, errors.New("user has been blocked")
	}

	accessToken, err := jwtpkg.AccessTokenGenerator(user, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	refreshToken, err := jwtpkg.RefreshTokenGenerator(user, s.cfg)
	if err != nil {
		return nil, errors.New("token generation failed")
	}

	userDTO := auth.UserDTO{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Role:          user.Role,
		Status:        user.Status,
		CreatedAt:     user.CreatedAt,
	}

	res := auth.LoginResponse{
		User:         &userDTO,
		RefreshToken: refreshToken,
		RefreshExp:   s.cfg.JWT.RefreshExpiration,
		AccessToken:  accessToken,
		AccessExp:    s.cfg.JWT.AccessExpiration,
	}
	return &res, nil

}
