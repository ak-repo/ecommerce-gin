package orderhandler

import (
	"fmt"
	"net/http"

	orderservice "github.com/ak-repo/ecommerce-gin/internal/services/orderService"
	"github.com/gin-gonic/gin"
)

type AdminOrderHandler struct {
	orderService orderservice.OrderService
}

func NewAdminOrderHandler(orderService orderservice.OrderService) *AdminOrderHandler {
	return &AdminOrderHandler{orderService: orderService}
}

// GET admin/orders   => get all orders
func (h *AdminOrderHandler) ShowAllOrderHandler(ctx *gin.Context) {

	orders, err := h.orderService.GetAllOrdersService()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/admin/orders/orders.html", gin.H{
			"Error":  err.Error(),
			"Orders": nil,
		})
	}

	fmt.Println("orders:", orders)

	ctx.HTML(http.StatusOK, "pages/admin/orders/orders.html", gin.H{
		"Orders": orders,
		"Error":  nil,
	})

}

func (h *AdminOrderHandler) ShowOrderByIDHandler(ctx *gin.Context) {

	orderID := ctx.Param("orderID")
	fmt.Println("orderID", orderID)
	if orderID == "" {
		ctx.HTML(http.StatusBadRequest, "pages/admin/orders/orderShow.html", gin.H{
			"Order": nil,
			"Error": "no order id found ",
		})
		return
	}

	order, err := h.orderService.GetOrderByIDService(orderID)
	fmt.Println("order:", order)
	if err != nil {
		ctx.HTML(http.StatusOK, "pages/admin/orders/orderShow.html", gin.H{
			"Order": nil,
			"Error": "error while getting order info: " + err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "pages/admin/orders/orderShow.html", gin.H{
		"Order": order,
		"Error": nil,
	})
}
