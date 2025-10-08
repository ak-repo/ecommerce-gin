package custreview

import (
	"errors"
	"net/http"

	reviewdto "github.com/ak-repo/ecommerce-gin/internals/review/review_DTO"
	reviewinterface "github.com/ak-repo/ecommerce-gin/internals/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CustomerReviewHandler struct {
	ReviewService reviewinterface.ServiceInterface
}

func NewCustomerReviewHandler(service reviewinterface.ServiceInterface) CustomerReviewHandler {
	return CustomerReviewHandler{ReviewService: service}
}

func (h *CustomerReviewHandler) ReviewCreateCustomerHandler(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "unauthorised", errors.New("customer id not found or unauthorised access"))
		return
	}
	var req reviewdto.CreateReviewRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}
	req.UserID = userID.(uint)
	if err := h.ReviewService.CreateNewReviewService(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "review creation failed", err)
		return
	}
	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "review added", nil)
}
