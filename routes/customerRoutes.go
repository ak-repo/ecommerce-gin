package routes

import (
	"github.com/ak-repo/ecommerce-gin/config"
	authhandler "github.com/ak-repo/ecommerce-gin/internals/auth/auth_handler"
	authrepo "github.com/ak-repo/ecommerce-gin/internals/auth/auth_repo"
	authservice "github.com/ak-repo/ecommerce-gin/internals/auth/auth_service"
	carthandler "github.com/ak-repo/ecommerce-gin/internals/customer/cart/handler"
	cartrepo "github.com/ak-repo/ecommerce-gin/internals/customer/cart/repo"
	cartservice "github.com/ak-repo/ecommerce-gin/internals/customer/cart/service"
	checkouthandler "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/handler"
	checkoutrepository "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/repository"
	checkoutservice "github.com/ak-repo/ecommerce-gin/internals/customer/checkout/service"
	producthandler "github.com/ak-repo/ecommerce-gin/internals/customer/product/handler"
	productrepository "github.com/ak-repo/ecommerce-gin/internals/customer/product/repository"
	productservice "github.com/ak-repo/ecommerce-gin/internals/customer/product/service"
	wishlisthandler "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/handler"
	wishlistrepository "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/repository"
	wishlistservice "github.com/ak-repo/ecommerce-gin/internals/customer/wishlist/service"
	orderhandler "github.com/ak-repo/ecommerce-gin/internals/order/handler"
	orderrepository "github.com/ak-repo/ecommerce-gin/internals/order/order_repos"
	orderservices "github.com/ak-repo/ecommerce-gin/internals/order/order_services"
	profilehandler "github.com/ak-repo/ecommerce-gin/internals/profile/handler"
	profilerepository "github.com/ak-repo/ecommerce-gin/internals/profile/repository"
	profileservice "github.com/ak-repo/ecommerce-gin/internals/profile/service"
	reviewhandler "github.com/ak-repo/ecommerce-gin/internals/review/handler"
	reviewrepository "github.com/ak-repo/ecommerce-gin/internals/review/repository"
	reviewservice "github.com/ak-repo/ecommerce-gin/internals/review/service"
	middleware "github.com/ak-repo/ecommerce-gin/middleware/auth"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
)

// RegisterCustomerRoutes sets up all customer-related routes
func RegisterCustomerRoutes(r *gin.Engine, db *db.Database, cfg *config.Config) {
	// --------------------------
	// INIT CORE REPOSITORIES
	// --------------------------
	authRepo := authrepo.NewAuthRepo(db.DB)
	productRepo := productrepository.NewProductRepository(db.DB)
	reviewRepo := reviewrepository.NewReviewRepo(db.DB)
	profileRepo := profilerepository.NewProfileRepository(db.DB)
	orderRepo := orderrepository.Newrepository(db.DB)
	wishlistRepo := wishlistrepository.NewWishlistRepo(db.DB)
	cartRepo := cartrepo.NewCartRepo(db.DB)
	checkoutRepo := checkoutrepository.NewCheckoutRepo(db.DB)

	// --------------------------
	// INIT SERVICES
	// --------------------------
	authService := authservice.NewAuthService(authRepo, cfg)
	productService := productservice.NewProductService(productRepo)
	reviewService := reviewservice.NewReviewService(reviewRepo)
	profileService := profileservice.NewProfileService(profileRepo, authRepo)
	cartService := cartservice.NewCartService(cartRepo, authRepo, productRepo)
	checkoutService := checkoutservice.NewCheckoutService(cartService, profileService, authRepo, checkoutRepo)
	orderService := orderservices.NewOrderService(orderRepo)
	wishlistService := wishlistservice.NewWishlistSevice(wishlistRepo)

	// --------------------------
	// INIT HANDLERS
	// --------------------------
	authHandler := authhandler.NewAuthHandler(authService)
	productHandler := producthandler.NewProductHandler(productService)
	reviewHandler := reviewhandler.NewReviewHandler(reviewService)
	profileHandler := profilehandler.NewProfileHandler(profileService)
	cartHandler := carthandler.NewCartHandler(cartService)
	checkoutHandler := checkouthandler.NewCheckoutHandler(checkoutService)
	orderHandler := orderhandler.NewOrderHandler(orderService)
	wishlistHandler := wishlisthandler.NewWishlistHandler(wishlistService)

	// PUBLIC ROUTES
	public := r.Group("/api/v1/customer")
	{
		// Auth
		public.POST("/login", authHandler.CustomerLogin)
		public.POST("/register", authHandler.CustomerRegister)

		// Product Browsing
		public.GET("/products", productHandler.GetAllProducts)
		public.GET("/products/:id", productHandler.GetProductByID)
	}

	// PROTECTED ROUTES (Authenticated Customers)
	protected := public.Group("/auth")
	protected.Use(middleware.AuthMiddleware(cfg), middleware.RoleMiddleware("customer"))
	{
		// Auth Management
		protected.POST("/password-change", authHandler.CustomerPasswordChange)
		protected.POST("/send-otp", authHandler.SendOTP)
		protected.POST("/verify-otp", authHandler.VerifyOTP)

		// Profile
		protected.GET("/profile", profileHandler.GetProfile)
		protected.GET("/profile/address", profileHandler.GetAddress)
		protected.PATCH("/profile/address", profileHandler.UpdateAddress)

		// Cart
		protected.GET("/cart", cartHandler.GetUserCart)
		protected.POST("/cart", cartHandler.AddItem)
		protected.PATCH("/cart", cartHandler.UpdateQuantity)
		protected.DELETE("/cart", cartHandler.RemoveItem)

		// Checkout
		protected.GET("/checkout", checkoutHandler.CheckoutSummary)
		protected.POST("/checkout", checkoutHandler.ProcessCheckout)

		// Orders
		protected.GET("/orders", orderHandler.GetOrderByCustomerIDForCustomer)
		protected.GET("/orders/:id", orderHandler.GetOrderByIDForCustomer)
		protected.POST("/orders/cancel", orderHandler.CancelOrder)
		protected.GET("/orders/cancel-response/:id", orderHandler.CancellationResponse)

		// Wishlist
		protected.GET("/wishlist", wishlistHandler.List)
		protected.POST("/wishlist/:id", wishlistHandler.Add)
		protected.DELETE("/wishlist/:id", wishlistHandler.Remove)

		// Reviews
		protected.POST("/review", reviewHandler.AddReview)
	}
}
