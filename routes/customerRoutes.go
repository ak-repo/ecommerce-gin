package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	authhandler "github.com/ak-repo/ecommerce-gin/internals/auth/auth_handler"
	authrepo "github.com/ak-repo/ecommerce-gin/internals/auth/auth_repo"
	authservice "github.com/ak-repo/ecommerce-gin/internals/auth/auth_service"
	carthandler "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart/cart_handler"
	cartrepo "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart/cart_repo"
	cartservice "github.com/ak-repo/ecommerce-gin/internals/customer/cust_cart/cart_service"
	custcheckouthandler "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout/cust_checkout_handler"
	custcheckoutrepo "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout/cust_checkout_repo"
	custcheckoutservice "github.com/ak-repo/ecommerce-gin/internals/customer/cust_checkout/cust_checkout_service"
	customerorderhandler "github.com/ak-repo/ecommerce-gin/internals/customer/cust_order/customer_order_handler"
	custproducthandler "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product/cust_product_handler"
	custproductrepo "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product/cust_product_repo"
	custproductservice "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product/cust_product_service"
	customerprofilehandler "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile/customer_profile_handler"
	customerprofilerepo "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile/customer_profile_repo"
	customerprofileservice "github.com/ak-repo/ecommerce-gin/internals/customer/cust_profile/customer_profile_service"
	orderrepos "github.com/ak-repo/ecommerce-gin/internals/order/order_repos"
	orderservices "github.com/ak-repo/ecommerce-gin/internals/order/order_services"
	middleware "github.com/ak-repo/ecommerce-gin/middleware/auth"

	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// control all user routes
func RegisterCustomerRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {

	publicRoute := r.Group("/cust")
	{
		// auth
		authRepo := authrepo.NewAuthRepo(db.DB)
		authService := authservice.NewAuthService(authRepo, cfg)
		authHandler := authhandler.NewAuthHandler(authService)

		publicRoute.POST("/login", authHandler.CustomerLogin)
		publicRoute.POST("/register", authHandler.CustomerRegister)

		// products
		productRepo := custproductrepo.NewCustomerProductRepo(db.DB)
		productService := custproductservice.NewCustomerProductService(productRepo)
		productHandler := custproducthandler.NewCustomerProductHandler(productService)

		publicRoute.GET("/products", productHandler.CustomerAllProducstListHandler)
		publicRoute.GET("/products/:id", productHandler.CustomerProductListHandler)

		// auth => secure routes
		custRoute := publicRoute.Group("/auth")
		custRoute.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("customer"))

		custRoute.POST("/password-change", authHandler.CustomerPasswordChange)
		custRoute.POST("/sent-OTP", authHandler.SendOTPHandler)
		custRoute.POST("/verify-OTP", authHandler.VerifyOTPHandler)

		// cart
		cartRepo := cartrepo.NewCartRepository(db.DB)
		cartService := cartservice.NewCartRepository(cartRepo, authRepo, productRepo)
		cartHandler := carthandler.NewCartHandler(cartService)

		custRoute.GET("/cart", cartHandler.ShowUserCartHandler)
		custRoute.POST("/cart", cartHandler.AddtoCartHandler)
		custRoute.PATCH("/cart/update", cartHandler.UpdateCartQuantityHandler)
		custRoute.DELETE("/cart/remove", cartHandler.RemoveCartItemHandler)

		// profile
		profileRepo := customerprofilerepo.NewCustomerProfileRepo(db.DB)
		profileService := customerprofileservice.NewCustomerProfileService(profileRepo, authRepo)
		profileHandler := customerprofilehandler.NewCustomerProfileHandler(profileService)

		custRoute.GET("/profile", profileHandler.CustomerProfileHandler)
		custRoute.GET("/profile/address", profileHandler.GetCustomerAddress)
		custRoute.PATCH("/profile/address", profileHandler.CustomerAddressUpdateHandler)

		// checkout
		checkoutRepo := custcheckoutrepo.NewCustomerCheckoutRepo(db.DB)
		checkoutService := custcheckoutservice.NewCustomerCheckoutService(cartService, profileService, authRepo, checkoutRepo)
		checkoutHandler := custcheckouthandler.NewCustomerCheckoutHandler(checkoutService)

		custRoute.GET("/checkout", checkoutHandler.CustomerShowCheckoutHandler)
		custRoute.POST("/checkout", checkoutHandler.CustomerCheckoutHandler)

		// orders
		orderRepo := orderrepos.NewOrderRepo(db.DB)
		orderService := orderservices.NewOrderService(orderRepo)
		orderHandler := customerorderhandler.NewCustomerOrderHandler(orderService)

		custRoute.GET("/orders", orderHandler.ListCustomerOrdersHandler)
		custRoute.GET("/orders/:id", orderHandler.CustomerOrderDetailHandler)
		custRoute.POST("/orders/cancel", orderHandler.CustomerOrderCancellationReqHandler)
		custRoute.GET("/orders/cancel-response/:id", orderHandler.CustomerOrderCancellationReqResponseHandler)

	}

}
