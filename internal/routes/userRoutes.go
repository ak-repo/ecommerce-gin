package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/middleware"
	userhandler "github.com/ak-repo/ecommerce-gin/internal/handlers/userHandler"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
	userservice "github.com/ak-repo/ecommerce-gin/internal/services/userService"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// control all user routes
func RegisterUserRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	userRoute := r.Group("/user")
	userRoute.Use(middleware.AuthMiddleware(cfg))

	{
		// profile
		userRepo := userrepository.NewUserRepo(db.DB)
		userService := userservice.NewUserService(userRepo, cfg)
		userHandler := userhandler.NewUserHandler(userService)
		userRoute.GET("/profile", userHandler.UserProfileHandler)
		userRoute.GET("/address/:address_id", userHandler.ShowAddressForm)
		userRoute.POST("/address/update/:address_id",userHandler.UserAddressUpdateHandler)

		userRoute.POST("/logout", userHandler.UserLogout)

	}

}
