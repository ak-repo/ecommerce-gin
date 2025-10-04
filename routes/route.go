package routes

import (
	"errors"
	"net/http"

	"github.com/ak-repo/ecommerce-gin/config"
	middleware "github.com/ak-repo/ecommerce-gin/middleware/auth"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Routes handler
func RegisterRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	// 404 handling
	r.NoRoute(func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists || role == "" {
			utils.RenderError(ctx, http.StatusNotFound, "no role specified", "no page found", errors.New("invalid url"))
			return
		}
		utils.RenderError(ctx, http.StatusNotFound, role.(string), "no page found", errors.New("invalid url"))

	})
	r.Use(middleware.AccessMiddleware(cfg))

	RegisterAdminRoute(r, db, cfg)
	RegisterCustomerRoute(r, db, cfg)

}
