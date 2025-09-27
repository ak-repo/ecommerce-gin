package jwtpkg

import (
	"errors"
	"fmt"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/golang-jwt/jwt/v4"
)

func TokenValidator(tokenStr string, cfg *config.Config) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(cfg.JWT.SecretKey), nil
	})

	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token expired")
			}
		}
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}

	return claims, nil
}
