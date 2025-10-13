package orderhandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	orderdto "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/dto"
	orderinterface "github.com/ak-repo/ecommerce-gin/internals/admin/order_mg/order_interface"

	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	OrderService orderinterface.Service
}

func NewOrderHandlerMG(service orderinterface.Service) handler {
	return handler{OrderService: service}
}

// GET admin/orders
func (h *handler) GetAllOrders(ctx *gin.Context) {

	orders, err := h.OrderService.GetAllOrders()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "orders not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/orders/orders.html", gin.H{
		"Orders": orders,
		"Error":  nil,
	})

}

// GET orders/:id => order by ID
func (h *handler) GetOrderByID(ctx *gin.Context) {

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

	order, err := h.OrderService.GetOrderByID(uint(orderID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/orders/orderShow.html", gin.H{
		"Order": order,
		"Error": nil,
	})
}

// GET admin/orders/users/:id =>

func (h *handler) GetOrderByCustomerID(ctx *gin.Context) {

	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}

	orders, err := h.OrderService.GetOrderByCustomerID(uint(userID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to fetch orders", err)
		return
	}
	// for admin
	utils.RenderSuccess(ctx, http.StatusOK, "admin", "orders", map[string]interface{}{
		"data": orders,
	})
}

// POST admin/orders/status/:id =>
func (h *handler) UpdateStatus(ctx *gin.Context) {

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
	if err := h.OrderService.UpdateStatus(&req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "status update failed", err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/admin/orders/%d", req.OrderID))
}

// GET admin/orders/cancels
func (h *handler) GetAllCancels(ctx *gin.Context) {
	requests, err := h.OrderService.GetAllCancels()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to fetch calcellation order requests", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/orders/orderCancelReq.html", gin.H{
		"CancelRequests": requests,
	})

}

// POST admin/orders/cancles/accept/:id =>
func (h *handler) AcceptCancel(ctx *gin.Context) {
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
	if err := h.OrderService.AcceptCancel(uint(reqID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order cancellation request handling failed", err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/orders/cancel-requests")
}

// POST admin/orders/cancles/reject/:id
func (h *handler) RejectCancel(ctx *gin.Context) {
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
	if err := h.OrderService.RejectCancel(uint(reqID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "order cancellation request handling failed", err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/orders/cancel-requests")
}
