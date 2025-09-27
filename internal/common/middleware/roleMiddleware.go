package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		currentRole, exists := ctx.Get("role")

		if !exists {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if role != currentRole {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()

	}
}
