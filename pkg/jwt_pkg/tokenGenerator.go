package jwtpkg

import (
	"time"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Accesstoken short time
func AccessTokenGenerator(email, username string, cfg *config.Config) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.AccessExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString([]byte(cfg.JWT.SecretKey))
}

// Refreshtoken long
func RefreshTokenGenerator(email, role string, cfg *config.Config) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.RefreshExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString([]byte(cfg.JWT.SecretKey))
}
