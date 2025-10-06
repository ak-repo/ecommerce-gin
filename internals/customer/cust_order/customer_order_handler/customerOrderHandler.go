package customerorderhandler

import orderservices "github.com/ak-repo/ecommerce-gin/internals/order/order_services"

//* | # | Functionality    | Endpoint              | Method | Typical Status |
// | - | ---------------- | --------------------- | ------ | -------------- |
// | 1 | Place order      | `/orders/place`       | POST   | Pending        |
// | 2 | List all orders  | `/orders`             | GET    | All            |
// | 3 | Order details    | `/orders/:id`         | GET    | Any            |
// | 4 | Cancel order     | `/orders/:id/cancel`  | POST   | Pending        |
// | 5 | Return order     | `/orders/:id/return`  | POST   | Delivered      |
// | 6 | Track order      | `/orders/:id/track`   | GET    | Any            |
// | 7 | Reorder          | `/orders/:id/reorder` | POST   | Delivered      |
// | 8 | Download invoice | `/orders/:id/invoice` | GET    | Delivered      |

type CustomerOrderHandler struct {
	OrderService orderservices.OrderService
}

func NewCustomerOrderHandler(orderService orderservices.OrderService) CustomerOrderHandler {
	return CustomerOrderHandler{OrderService: orderService}
}

