package authmiddleware

import (
	"strings"

	"github.com/ak-repo/ecommerce-gin/config"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/gin-gonic/gin"
)

func AccessMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string

		authHeader := ctx.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			accessToken = strings.TrimPrefix(authHeader, "Bearer ")
		}

		if accessToken == "" {
			tokenFromCookie, err := ctx.Cookie("accessToken")
			if err == nil {
				accessToken = tokenFromCookie
			}
		}

		if accessToken == "" {
			ctx.Next()
			return
		}

		claims, err := jwtpkg.TokenValidator(accessToken, cfg)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("userID", claims.UserID)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
