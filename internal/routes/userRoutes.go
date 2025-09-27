package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/middleware"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// control all user routes
func RegisterUserRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	userRoute := r.Group("/user")
	userRoute.Use(middleware.AuthMiddleware(cfg))

	{

	}

}
