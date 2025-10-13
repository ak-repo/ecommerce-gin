package orderhandler

import (
	"errors"
	"net/http"
	"strconv"

	orderdto "github.com/ak-repo/ecommerce-gin/internals/customer/order/order_dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/customer/order/order_interface"
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

type handler struct {
	OrderService orderinterface.Service
}

func NewOrderHandler(service orderinterface.Service) handler {
	return handler{OrderService: service}
}

// GET orders/:id => order by ID
func (h *handler) GetOrderByID(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "order id  not found", errors.New("np order at parameter"))
		return
	}
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "invalid order id", err)
		return
	}

	order, err := h.OrderService.GetOrderByID(uint(orderID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "order not found", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "order fetched successfully", order)
}

func (h *handler) GetOrderByCustomerID(ctx *gin.Context) {

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	orders, err := h.OrderService.GetOrderByCustomerID(uint(userID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "failed to fetch orders", err)
		return
	}
	// for admin
	utils.RenderSuccess(ctx, http.StatusOK, "customer", "orders", orders)
}

// POST cust/auth/orders/cancel
func (h *handler) CancelOrder(ctx *gin.Context) {

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	var req orderdto.CreateCancelRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.OrderService.CancelOrder(&req, userID); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "order cancellation failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusAccepted, "customer", "order cancellation request accepted", nil)
}

// GET - cust/auth/orders/cancel-response/:id
func (h *handler) CancellationResponse(ctx *gin.Context) {

	id := ctx.Param("id")
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || id == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	response, err := h.OrderService.CancellationResponse(uint(orderID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "no response", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "order cancellation response",
	 response)
}
