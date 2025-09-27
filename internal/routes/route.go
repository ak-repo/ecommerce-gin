package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// Routes handler
func RegisterRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	PublicRoute(r, db, cfg)
	RegisterUserRoute(r, db, cfg)
	RegisterAdminRoute(r,db,cfg)

}
