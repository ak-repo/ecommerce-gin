package authservice

import (
	"errors"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internals/auth"
	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
)

type AuthService struct {
	authRepo authinterface.AuthRepoInterface
	cfg      *config.Config
}

func NewAuthService(authRepo authinterface.AuthRepoInterface, cfg *config.Config) authinterface.AuthServiceInterface {
	return &AuthService{authRepo: authRepo, cfg: cfg}
}

// -------------------------------------------- Registration service -------------------------------------------------------
func (s *AuthService) RegisterService(input *auth.RegisterRequest, role string) error {

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	if user, err := s.authRepo.GetUserByEmail(input.Email); user != nil && err != nil {
		return errors.New("email already taken")
	}

	return s.authRepo.CreateUser(input.Username, input.Email, hash, role)
}

// -----------------------------------------------------Login service ----------------------------------------------------
func (s *AuthService) LoginService(input *auth.LoginRequest, role string) (*auth.LoginResponse, error) {

	user, err := s.authRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
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
