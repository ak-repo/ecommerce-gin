package middleware

import (
	"net/http"
	"strings"

	"github.com/ak-repo/ecommerce-gin/config"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string

		authHeader := ctx.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		if tokenString == "" {
			cookieToken, err := ctx.Cookie("accessToken")
			if err != nil || cookieToken == "" {
				utils.RenderError(ctx, http.StatusUnauthorized, "no role assigned", "token not available", err)
				ctx.Abort()
				return
			}
			tokenString = cookieToken
		}

		claims, err := jwtpkg.TokenValidator(tokenString, cfg)
		if err != nil {
			utils.RenderError(ctx, http.StatusUnauthorized, "no role assigned", "session has ended or expired", err)
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Set("userID", claims.UserID)

		ctx.Next()
	}
}
