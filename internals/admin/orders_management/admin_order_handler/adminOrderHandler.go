package adminorderhandler

import (
	"errors"
	"net/http"
	"strconv"

	orderinterface "github.com/ak-repo/ecommerce-gin/internals/order/order_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminOrderHandler struct {
	OrderService orderinterface.OrderServiceInterface
}

func NewAdminOrderHandler(orderService orderinterface.OrderServiceInterface) AdminOrderHandler {
	return AdminOrderHandler{OrderService: orderService}
}

// GET admin/orders   => get all orders
func (h *AdminOrderHandler) ShowAllOrderHandler(ctx *gin.Context) {

	orders, err := h.OrderService.GetAllOrdersService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "orders not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/orders/orders.html", gin.H{
		"Orders": orders,
		"Error":  nil,
	})

}

func (h *AdminOrderHandler) ShowOrderByIDHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order id  not found", errors.New("np order at parameter"))
		return
	}
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid order id", err)
		return
	}

	order, err := h.OrderService.GetOrderByIDService(uint(orderID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/orders/orderShow.html", gin.H{
		"Order": order,
		"Error": nil,
	})
}
