ecommerce-gin/
├── cmd/
│   └── main.go                  # Entry point of the app
├── config/
│   └── config.go                # App configuration & env loader
├── internal/                    # Core business logic (private)
│   ├── auth/                    # Authentication & user login
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repo.go
│   │   ├── model.go
│   │   └── middleware.go
│   ├── user/                    # User-facing modules
│   │   ├── cart/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   ├── wishlist/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   ├── profile/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   ├── order/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   └── address/
│   │       ├── handler.go
│   │       ├── service.go
│   │       ├── repo.go
│   │       └── model.go
│   ├── admin/                   # Admin-facing modules
│   │   ├── dashboard/
│   │   │   ├── handler.go
│   │   │   └── service.go
│   │   ├── user_management/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   └── repo.go
│   │   ├── store_management/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   └── repo.go
│   │   ├── product_management/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   └── repo.go
│   │   ├── order_management/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   └── repo.go
│   │   └── analytics/
│   │       ├── handler.go
│   │       └── service.go
│   ├── store/                   # Store-specific modules
│   │   ├── inventory/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   ├── profile/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   ├── order/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   └── shipment/
│   │       ├── handler.go
│   │       ├── service.go
│   │       ├── repo.go
│   │       └── model.go
│   ├── deliverypartner/         # Delivery partner modules
│   │   ├── profile/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   ├── assignment/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   ├── history/
│   │   │   ├── handler.go
│   │   │   ├── service.go
│   │   │   ├── repo.go
│   │   │   └── model.go
│   │   └── status/
│   │       ├── handler.go
│   │       ├── service.go
│   │       ├── repo.go
│   │       └── model.go
│   ├── order/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repo.go
│   │   └── model.go
│   ├── payment/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repo.go
│   │   └── model.go
│   └── search/
│       ├── handler.go
│       ├── service.go
│       ├── repo.go
│       └── model.go
├── middleware/                  # Global middleware
│   ├── jwt.go
│   └── logging.go
├── infra/                       # Infrastructure setup
│   ├── db.go
│   └── logger.go
├── routes/
│   └── router.go                 # Route registration
├── pkg/                          # Reusable utilities
│   ├── utils.go
│   └── errors.go
└── docs/                         # Documentation
    ├── api.md
    └── arch.md


Cart (Add/Update/Delete)
   ↓
View Cart
   ↓
Select Address
   ↓
Checkout (Order Summary)
   ↓
Payment (Optional)
   ↓
Place Order
   ↓
Clear Cart → Show Confirmation
