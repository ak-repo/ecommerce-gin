# ecommerce-gin

http://localhost:8080/login

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

rotes

## **1️⃣ Public Routes** (accessible to anyone)

| Method | Path               | Description               |
| ------ | ------------------ | ------------------------- |
| GET    | `/`                | Home page                 |
| GET    | `/about`           | About us page             |
| GET    | `/products`        | List all products         |
| GET    | `/products/:id`    | Product detail page       |
| GET    | `/login`           | Show login form           |
| POST   | `/login`           | Submit login form         |
| GET    | `/register`        | Show registration form    |
| POST   | `/register`        | Submit registration form  |
| GET    | `/forgot-password` | Show forgot password form |
| POST   | `/forgot-password` | Submit forgot password    |
| GET    | `/contact`         | Contact page              |

## **2️⃣ User Routes** (protected — requires login)

| Method | Path                    | Description                  |
| ------ | ----------------------- | ---------------------------- |
| POST   | `/user/logout`          | Logout user                  |
| GET    | `/user/profile`         | Show user profile            |
| POST   | `/user/profile`         | Update user profile          |
| GET    | `/user/orders`          | List user orders             |
| GET    | `/user/orders/:id`      | Order detail                 |
| POST   | `/user/cart/add`        | Add product to cart          |
| POST   | `/user/cart/remove`     | Remove product from cart     |
| GET    | `/user/cart`            | Show current cart            |
| POST   | `/user/wishlist/add`    | Add product to wishlist      |
| POST   | `/user/wishlist/remove` | Remove product from wishlist |
| GET    | `/user/wishlist`        | Show wishlist                |












#### Customer-Facing Features (Front-end)

These are the features that your shoppers will directly interact with. The goal is to provide a smooth, intuitive, and secure shopping experience.

*   **Homepage and Landing Pages** A visually appealing homepage is crucial for making a good first impression and communicating your brand's identity. This includes featured products, promotions, and clear navigation.[1]
*   **User Accounts** Customers should be able to create an account to view their order history, manage shipping addresses, and save payment information for faster checkout.[2]
*   **Product Catalog and Categories** Products should be organized into logical categories and subcategories for easy browsing.[3]
*   **Powerful Search and Filtering** A robust search bar helps users find specific items quickly. Filtering options (by price, size, color, brand) allow customers to narrow down their choices.[4][2]
*   **Detailed Product Pages** Each product needs its own page with high-quality images and videos, a detailed description, price, and clear "Add to Cart" button.[4][1]
*   **Shopping Cart** A persistent shopping cart that users can easily view and modify is essential. It should clearly display items, quantities, and the subtotal.[5]
*   **Secure Checkout Process** The checkout process should be as simple as possible, with multiple payment options (credit/debit cards, digital wallets) and a secure payment gateway.[3][5]
*   **Ratings and Reviews** Social proof like customer reviews and ratings builds trust and helps other shoppers make informed decisions.[5][4]
*   **Mobile-Friendly Design** With a significant portion of online shopping done on mobile devices, your website must be fully responsive and easy to use on a smartphone.[5][4]

#### Administrative Features (Back-end)

These are the tools you'll need in your admin dashboard to manage the store effectively.

*   **Product Management** You need the ability to add, edit, and delete products from your catalog. This includes managing descriptions, images, prices, and inventory levels.[6][3]
*   **Order Management** A centralized system to view incoming orders, update their status (e.g., "Processing," "Shipped"), print invoices, and manage returns is critical.[6][3]
*   **Customer Management** This involves viewing customer information, their order history, and managing user accounts.[6]
*   **Inventory Management** The platform should automatically track stock levels as products are sold and provide low-stock alerts to prevent overselling.[4]
*   **Discount and Promotion Engine** The ability to create coupon codes, run sales (e.g., "20% off all shirts"), and offer special promotions is a key marketing tool.[6]
*   **Content Management System (CMS)** You'll need a simple way to edit content on pages like the "About Us" or "Contact" page without needing to code.[6]
*   **Analytics and Reporting** The dashboard should provide insights into sales trends, top-selling products, and customer behavior to help you make informed business decisions.

### Guidance for Building Your E-commerce Platform

Building a platform from scratch is a significant learning opportunity. Here’s a step-by-step approach to guide you:

1.  **Start with Market Research**: Before writing any code, understand your target audience and look at what your competitors are doing. This will help you identify unique features that can set your platform apart.[7][8]
2.  **Focus on the Foundation (Back-end First)**:
    *   **Database Design**: This is the most critical first step. Design your database schemas for `users`, `products`, `orders`, and `cart_items`. Think about the relationships between them. A well-designed database will save you countless headaches later.
    *   **API Development**: Use the API documentation you created earlier as your blueprint. Start building the back-end endpoints using Go and the Gin framework. Begin with user authentication, as it's fundamental to most other features. Then move on to product and order management.
    *   **Testing**: Test each endpoint thoroughly as you build it. Use a tool like Postman to ensure that your API behaves exactly as expected.

3.  **Build the Admin Dashboard**: Once you have your core back-end APIs working, you can build the admin dashboard. Since you are using Server-Side Rendering (SSR) with Go's HTML templates, you can create pages for:
    *   Viewing a list of all orders.
    *   Viewing a list of all products, with buttons to "Add New Product."
    *   A form to add or edit a product.
    This will give you a functional interface to manage your store's data.

4.  **Develop the Customer-Facing Front-end**: With the back-end and admin panel in place, you can now focus on the customer experience.
    *   **SSR Pages**: Create the main pages using your HTML templates: the homepage, product pages, and category pages.
    *   **Client-Side Interactivity**: For features that require a dynamic experience without a full page reload (like adding an item to the cart), use JavaScript to call the APIs you built. The JavaScript will send a request to your back-end and then update the page with the response.

5.  **Implement Security Measures**: Security is non-negotiable in e-commerce. Ensure you are following best practices for:[5]
    *   **Password Hashing**: Never store passwords in plain text.
    *   **Input Validation**: Protect against injection attacks by validating all user input.
    *   **Secure Payment Integration**: Use a trusted and PCI-compliant payment processor like Stripe or PayPal. Do not handle or store credit card information on your servers.[7]

