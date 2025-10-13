package reviewhander

import (
	"errors"
	"net/http"

	custreviewdto "github.com/ak-repo/ecommerce-gin/internals/customer/review/dto"
	custreviewinter "github.com/ak-repo/ecommerce-gin/internals/customer/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ReviewService custreviewinter.Service
}

func NewReviewHandler(service custreviewinter.Service) custreviewinter.Handler {
	return &handler{ReviewService: service}
}

func (h *handler) AddReview(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusUnauthorized, "customer", "unauthorised", errors.New("customer id not found or unauthorised access"))
		return
	}
	var req custreviewdto.CreateReviewRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}
	req.UserID = userID
	if err := h.ReviewService.AddReview(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "review creation failed", err)
		return
	}
	utils.RenderSuccess(ctx, http.StatusCreated, "customer", "review added", nil)
}
