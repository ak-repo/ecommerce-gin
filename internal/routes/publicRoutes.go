package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	producthandler "github.com/ak-repo/ecommerce-gin/internal/handlers/productHandler"
	userhandler "github.com/ak-repo/ecommerce-gin/internal/handlers/userHandler"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
	userrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/userRepository"
	productservice "github.com/ak-repo/ecommerce-gin/internal/services/productService"
	userservice "github.com/ak-repo/ecommerce-gin/internal/services/userService"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// public routes  "/"
func PublicRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	publicRoute := r.Group("/")
	{

		//
		userAuthRepo := userrepository.NewUserAuthRepo(db.DB)
		userAuthService := userservice.NewUserAuthService(userAuthRepo, cfg)
		userAuthHandler := userhandler.NewUserAuthHandler(userAuthService)

		publicRoute.GET("/register", userAuthHandler.ShowRegistrationForm)
		publicRoute.POST("/register", userAuthHandler.RegistrationHandler)
		publicRoute.GET("/login", userAuthHandler.ShowLoginForm)
		publicRoute.POST("/login", userAuthHandler.LoginHandler)

		publicRoute.GET("/")

		{
			productRepo := productrepository.NewProductRepo(db.DB)
			productService := productservice.NewProductService(cfg, productRepo)
			productHandler := producthandler.NewProductHandler(productService)

			publicRoute.GET("/products", productHandler.GetAllProductHandler)
			publicRoute.GET("/product/:id", productHandler.GetProductByIDHandler)

		}

	}

}
