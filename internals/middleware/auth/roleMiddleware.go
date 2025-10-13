package authmiddleware

import (
	"errors"
	"net/http"

	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		currentRole, exists := ctx.Get("role")

		if !exists {
			utils.RenderError(ctx, http.StatusUnauthorized, "no role assigined", "no access for this page", errors.New("no role found"))
			ctx.Abort()
			return
		}

		if role != currentRole {
			utils.RenderError(ctx, http.StatusUnauthorized, currentRole.(string), "no access for this page", errors.New("you have no access this page"))
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}
