package adminorderhandler

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

// GET admin/orders/:id => order by ID
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

// POST admin/orders/status/:id => update order status
func (h *AdminOrderHandler) UpdateOrderStatusHandler(ctx *gin.Context) {

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

	var req orderdto.AdminUpdateOrderStatusRequest
	req.OrderID = uint(orderID)
	req.Status = ctx.PostForm("status")
	if req.Status == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "status is required", nil)
		return
	}

	if err := h.OrderService.UpdateOrderStatusService(&req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "status update failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/admin/orders/%d", req.OrderID))
}

// GET admin/orders/cancels => get all cancel requests
func (h *AdminOrderHandler) ShowAllCancelRequestHandler(ctx *gin.Context) {

	requests, err := h.OrderService.OrderCancellationReqListingService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "cancel request not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/orders/orderCancelReq.html", gin.H{
		"CancelRequests": requests,
	})

}

// POST admin/orders/cancles/accept/:id => accept the request
func (h *AdminOrderHandler) OrderCancellationAcceptHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order id  not found", errors.New("np order at parameter"))
		return
	}
	reqID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid order id", err)
		return
	}

	if err := h.OrderService.AcceptCancellationReqService(uint(reqID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order cancellation request handling failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/orders/cancels")

}

// POST admin/orders/cancles/reject/:id => reject and give reason
func (h *AdminOrderHandler) OrderCancellationRejectHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order id  not found", errors.New("np order at parameter"))
		return
	}
	reqID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid order id", err)
		return
	}

	if err := h.OrderService.RejectCancellationReqService(uint(reqID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order cancellation request handling failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/orders/cancels")

}
