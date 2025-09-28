package middleware

import (
	"log"
	"net/http"

	"github.com/ak-repo/ecommerce-gin/config"
	jwtpkg "github.com/ak-repo/ecommerce-gin/pkg/jwt_pkg"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		tokenString, err := ctx.Cookie("accessToken")
		log.Println("token:", tokenString)
		if err != nil {
			ctx.HTML(http.StatusUnauthorized, "pages/404/404.html", gin.H{})
			ctx.Abort()
			return
		}

		claims, err := jwtpkg.TokenValidator(tokenString, cfg)
		if err != nil {
			ctx.HTML(http.StatusUnauthorized, "pages/404/404.html", gin.H{})
			ctx.Abort()
		}

		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Next()

	}
}
