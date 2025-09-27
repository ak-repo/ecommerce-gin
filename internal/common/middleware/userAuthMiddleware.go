package middleware

import (
	"net/http"

	"github.com/ak-repo/ecommerce-gin/config"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/gin-gonic/gin"
)

func UserAuthMiddleware(cfg *config.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader

		claims, err := jwtpkg.TokenValidator(tokenString, cfg)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}

		ctx.Set("email", claims.Email)
		ctx.Next()

	}
}
