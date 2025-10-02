package middleware

import (
	"github.com/ak-repo/ecommerce-gin/config"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/gin-gonic/gin"
)

func AccessMiddleware(cfg *config.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("accessToken")
		if err != nil || accessToken == "" {
			ctx.Next()
			return
		}

		claims, err := jwtpkg.TokenValidator(accessToken, cfg)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("username", &claims.Username)
		ctx.Set("email", claims.Email)
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
