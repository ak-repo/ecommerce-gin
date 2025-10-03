package routes

import (
	"net/http"

	"github.com/ak-repo/ecommerce-gin/config"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// Routes handler
func RegisterRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "pages/404/404.html", gin.H{})
	})

	PublicRoute(r, db, cfg)
	RegisterUserRoute(r, db, cfg)
	RegisterAdminRoute(r, db, cfg)

}
