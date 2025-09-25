# ecommerce-gin

A full-featured e-commerce backend application built with Golang Gin framework

### E-Commerce Project Structure

Here is a detailed folder-by-folder breakdown for your Go-based e-commerce platform:

```
/ecommerce-project
├── cmd/
│   └── web/
│       └── main.go              # Application entry point, server setup
├── configs/
│   └── config.yaml              # Configuration for database, server, etc.
├── internal/
│   ├── api/
│   │   ├── handlers/            # (Controllers) Handles HTTP requests
│   │   │   ├── admin/
│   │   │   │   ├── dashboard_handler.go
│   │   │   │   ├── product_mgmt_handler.go
│   │   │   │   ├── order_mgmt_handler.go
│   │   │   │   └── user_mgmt_handler.go
│   │   │   ├── auth/
│   │   │   │   └── auth_handler.go  # Login, register, logout
│   │   │   ├── user/
│   │   │   │   ├── profile_handler.go
│   │   │   │   └── address_handler.go
│   │   │   ├── product_handler.go   # View products, search
│   │   │   ├── cart_handler.go
│   │   │   ├── wishlist_handler.go
│   │   │   ├── order_handler.go
│   │   │   └── payment_handler.go
│   │   ├── middleware/
│   │   │   ├── auth_middleware.go   # Protects routes
│   │   │   └── admin_middleware.go  # Verifies admin privileges
│   │   └── router.go              # All route definitions
│   ├── models/                    # Data structures (structs)
│   │   ├── user.go
│   │   ├── product.go
│   │   ├── order.go
│   │   ├── cart.go
│   │   ├── wishlist.go
│   │   ├── payment.go
│   │   └── address.go
│   ├── repositories/              # Data access layer
│   │   ├── user_repository.go
│   │   ├── product_repository.go
│   │   ├── order_repository.go
│   │   └── ... (other repositories)
│   └── services/                  # Business logic
│       ├── auth_service.go
│       ├── user_service.go
│       ├── product_service.go
│       ├── cart_service.go
│       ├── order_service.go
│       ├── payment_service.go
│       └── admin_service.go
├── pkg/                         # Reusable packages (utilities)
│   ├── utils/
│   │   ├── password.go          # Hashing and validation
│   │   └── validator.go         # Input validation
│   └── tokens/
│       └── token_generator.go   # JWT generation and validation [23]
├── web/
│   ├── templates/
│   │   ├── admin/
│   │   │   ├── dashboard.html
│   │   │   ├── products.html
│   │   │   └── orders.html
│   │   ├── auth/
│   │   │   ├── login.html
│   │   │   └── register.html
│   │   ├── user/
│   │   │   ├── profile.html
│   │   │   ├── orders.html
│   │   │   └── wishlist.html
│   │   ├── products/
│   │   │   ├── index.html       # Product listing
│   │   │   └── show.html        # Product detail
│   │   ├── cart.html
│   │   ├── checkout.html
│   │   └── layouts/
│   │       ├── base.html
│   │       ├── header.html
│   │       └── footer.html
│   └── static/                    # CSS, JavaScript, images
│       ├── css/
│       ├── js/
│       └── images/
├── go.mod
└── go.sum
```

### Explanation of Key Directories

- **`cmd/web/main.go`**: This is the main entry point of your application. It is responsible for initializing the database connection, setting up the Gin router, and starting the web server.[2]

- **`configs/`**: Stores application configuration files. Using a file like `config.yaml` allows you to manage database credentials, server ports, and other settings outside of your code, which is a security best practice.

- **`internal/`**: This directory contains the core application logic. According to Go conventions, code in the `internal` directory can only be imported by code within the same parent directory, preventing it from being used in other projects.

  - **`api/handlers/` (Controllers)**: These are the Gin handlers that process incoming HTTP requests. Each feature has its own set of handlers.[3]
    - **`admin/`**: Contains all handlers for the admin dashboard, such as managing products, users, and orders.
    - **`auth/`**: Manages user authentication, including registration, login, and token generation.
    - **`user/`**: Handles user-specific actions like updating a profile or managing addresses.
  - **`api/middleware/`**: Contains middleware functions that can be chained to routes. `auth_middleware.go` would verify a user's JWT token before allowing access to protected routes.[2]
  - **`api/router.go`**: This file defines all the API routes, mapping URLs to their corresponding handler functions in a clean and organized way.
  - **`models/`**: Defines the Go structs that represent your database tables (e.g., `User`, `Product`, `Order`). These models are used across all layers of the application.[3]
  - **`repositories/`**: The repository layer is responsible for all database operations (CRUD - Create, Read, Update, Delete). It abstracts the database away from your business logic, making it easier to switch databases or write tests.
  - **`services/`**: The service layer holds the business logic. It orchestrates calls between the handlers and repositories. For example, the `OrderService` might take an order request, validate product availability via the `ProductRepository`, create an order with the `OrderRepository`, and trigger a payment process.

- **`pkg/`**: This directory is for code that can be safely used by external applications. It is ideal for generic utility functions.[2]

  - **`utils/`**: Contains helper functions like password hashing, input validation, or formatting.
  - **`tokens/`**: A dedicated package for creating and validating JWTs or other session tokens.[4]

- **`web/`**: This directory holds all front-end related files.
  - **`templates/`**: Contains HTML templates that Gin will render. Organizing them into subfolders (e.g., `admin`, `user`, `products`) keeps your views tidy.[5]
  - **`static/`**: Stores static assets like CSS stylesheets, JavaScript files, and images that will be served directly to the client.



client (HTTP request)
↓
routes (decide which handler)
↓
handler (parse input, call service, return response)
↓
service (business logic, validation, orchestration)
↓
repository (actual DB queries)
↓
database
