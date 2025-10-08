package customerorderhandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	orderdto "github.com/ak-repo/ecommerce-gin/internals/order/order_dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/order/order_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

// | 1 | Place order      | `/orders/place`       | POST   | Pending        | => handled in checkout
// | 2 | List all orders  | `/orders`             | GET    | All            |
// | 3 | Order details    | `/orders/:id`         | GET    | Any            |
// | 4 | Cancel order     | `/orders/:id/cancel`  | POST   | Pending        |
// | 5 | Return order     | `/orders/:id/return`  | POST   | Delivered      |
// | 6 | Track order      | `/orders/:id/track`   | GET    | Any            |
// | 7 | Reorder          | `/orders/:id/reorder` | POST   | Delivered      |
// | 8 | Download invoice | `/orders/:id/invoice` | GET    | Delivered      |

type CustomerOrderHandler struct {
	OrderService orderinterface.OrderServiceInterface
}

func NewCustomerOrderHandler(orderService orderinterface.OrderServiceInterface) CustomerOrderHandler {
	return CustomerOrderHandler{OrderService: orderService}
}

// GET - cust/auth/orders  => get all ordes by customer id
func (h *CustomerOrderHandler) ListCustomerOrdersHandler(ctx *gin.Context) {

	id, exists := ctx.Get("userID")
	userID := id.(uint)
	if !exists || id == 0 {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "User authentication failed", errors.New("user ID not found or invalid"))
		return
	}
	orders, err := h.OrderService.GetCustomerOrdersService(userID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "orders not found", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "orders", map[string]interface{}{
		"data": orders,
	})
}

// GET -cust/auth/orders/:id => get single order details by order ID
func (h *CustomerOrderHandler) CustomerOrderDetailHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || orderID == 0 {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "order id not valid", err)
		return
	}

	order, err := h.OrderService.GetCustomerOrderbyOrderIDService(uint(orderID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "order not found", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", fmt.Sprintf("order: %d", orderID), map[string]interface{}{
		"data": order,
	})

}

// POST cust/auth/orders/cancel
func (h *CustomerOrderHandler) CustomerOrderCancellationReqHandler(ctx *gin.Context) {

	id, exists := ctx.Get("userID")
	userID := id.(uint)
	if !exists || id == 0 {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "User authentication failed", errors.New("user ID not found or invalid"))
		return
	}

	var req orderdto.CreateCancelRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.OrderService.CancelOrderByCustomerService(&req, userID); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "order cancellation failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusAccepted, "customer", "order cancellation request accepted", nil)
}

// GET - cust/auth/orders/cancel-response/:id
func (h *CustomerOrderHandler) CustomerOrderCancellationReqResponseHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || id == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	response, err := h.OrderService.CancellationResponseForCustomerService(uint(orderID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "no response", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "order cancellation response", map[string]interface{}{
		"data": response,
	})
}
