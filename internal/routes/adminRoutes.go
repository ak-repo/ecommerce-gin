package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/common/middleware"
	producthandler "github.com/ak-repo/ecommerce-gin/internal/handlers/productHandler"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
	productservice "github.com/ak-repo/ecommerce-gin/internal/services/productService"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("admin"))
	{
		// products
		productRepo := productrepository.NewProductRepo(db.DB)
		productService := productservice.NewProductService(cfg, productRepo)
		productHandler := producthandler.NewAdminproductHandler(productService)

		adminRoute.GET("/products", productHandler.GetAllProductHandler)
		adminRoute.GET("/product/:id", productHandler.GetProductByIDHandler)
		adminRoute.GET("/product/edit/:id", productHandler.ShowProductEdit)
		// adminRoute.POST("/product/update/:id", productHandler.UpdateProduct)
	}
}
