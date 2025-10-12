package routes

import (
	"errors"
	"net/http"
	"time"

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
		role, _ := ctx.Get("role")
		if role == "admin" {
			ctx.HTML(http.StatusNotFound, "pages/response/404.html", gin.H{
				"CurrentYear": time.Now(),
			})
			return
		}
		if role != nil {
			utils.RenderError(ctx, http.StatusNotFound, role.(string), "no page found", errors.New("invalid url"))
		}
		utils.RenderError(ctx, http.StatusNotFound, "no role", "no page found", errors.New("invalid url"))
	})
	r.Use(middleware.AccessMiddleware(cfg))

	RegisterAdminRoutes(r, db, cfg)
	RegisterCustomerRoutes(r, db, cfg)

}
