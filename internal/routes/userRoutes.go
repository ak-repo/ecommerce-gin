package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/middleware"
	carthandler "github.com/ak-repo/ecommerce-gin/internal/handlers/cartHandler"
	userhandler "github.com/ak-repo/ecommerce-gin/internal/handlers/userHandler"
	cartrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/cartRepository"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
	cartservice "github.com/ak-repo/ecommerce-gin/internal/services/cartService"
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
		userRoute.POST("/address/update/:address_id", userHandler.UserAddressUpdateHandler)
		userRoute.GET("/password", userHandler.UserPasswordChangeFormHandler)
		userRoute.POST("/password", userHandler.UserPasswordChangeHandler)
		userRoute.POST("/logout", userHandler.UserLogout)

		// cart
		cartRepo := cartrepository.NewCartRepository(db.DB)
		productRepo := productrepository.NewProductRepo(db.DB)
		cartService := cartservice.NewCartRepository(userRepo, cartRepo, productRepo)
		cartHandler := carthandler.NewCartHandler(cartService)

		userRoute.POST("/addToCart", cartHandler.AddtoCartHandler)
		userRoute.GET("/cart", cartHandler.ShowUserCartHandler)
		userRoute.POST("/cart/update", cartHandler.UpdateCartQuantityHandler)
		userRoute.POST("/cart/remove", cartHandler.RemoveCartItemHandler)

	}

}
