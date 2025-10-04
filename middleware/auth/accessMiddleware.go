package middleware

import (
	"strings"

	"github.com/ak-repo/ecommerce-gin/config"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/gin-gonic/gin"
)

func AccessMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string

		// 1 First check Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			accessToken = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// 2 Fallback: check cookie if no Bearer token
		if accessToken == "" {
			tokenFromCookie, err := ctx.Cookie("accessToken")
			if err == nil {
				accessToken = tokenFromCookie
			}
		}

		// 3 If no token found, just continue (unauthenticated)
		if accessToken == "" {
			ctx.Next()
			return
		}

		// 4 Validate token
		claims, err := jwtpkg.TokenValidator(accessToken, cfg)
		if err != nil {
			ctx.Next()
			return
		}

		// 5 Save claims to context
		ctx.Set("email", claims.Email)
		ctx.Set("userID", claims.UserID)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
