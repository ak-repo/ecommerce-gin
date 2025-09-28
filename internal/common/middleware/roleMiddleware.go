package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		currentRole, exists := ctx.Get("role")
		log.Println("role:", role)

		if !exists {
			ctx.HTML(http.StatusUnauthorized, "pages/404/404.html", gin.H{})
			ctx.Abort()
			return
		}

		if role != currentRole {
			ctx.HTML(http.StatusUnauthorized, "pages/404/404.html", gin.H{})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}
