package jwtpkg

import (
	"time"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	UserID   uint   `json:"userID"`
	jwt.RegisteredClaims
}

// Accesstoken short time
func AccessTokenGenerator(user *models.User, cfg *config.Config) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email:    user.Email,
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.AccessExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString([]byte(cfg.JWT.SecretKey))
}

// Refreshtoken long
func RefreshTokenGenerator(user *models.User, cfg *config.Config) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email:    user.Email,
		Role:     user.Role,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.RefreshExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString([]byte(cfg.JWT.SecretKey))
}
