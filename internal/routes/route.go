package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// Routes handler
func RegisterRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	RegisterUserRoute(r, db, cfg)
	PublicRoute(r, db, cfg)

}
