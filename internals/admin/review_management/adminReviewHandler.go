package reviewmanagement

import (
	"errors"
	"net/http"
	"strconv"

	reviewinterface "github.com/ak-repo/ecommerce-gin/internals/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminReviewHandler struct {
	ReviewService reviewinterface.ServiceInterface
}

func NewAdminReviewService(service reviewinterface.ServiceInterface) AdminReviewHandler {
	return AdminReviewHandler{ReviewService: service}
}

// POST - admin/reviews
func (h *AdminReviewHandler) ListAllReviewsHandler(ctx *gin.Context) {

	reviews, err := h.ReviewService.ListAllReviewsForVerifyService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "no review found", err)
		return
	}
	ctx.HTML(http.StatusOK, "pages/reviews/reviews.html", gin.H{
		"Reviews": reviews,
	})

}

// POST - admin/reviews/approve/:id
func (h *AdminReviewHandler) ApproveReviewHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "no review id found", errors.New("no review id given in parameter"))
		return
	}
	reviewID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid id", err)
		return
	}

	if err := h.ReviewService.ApporveReviewService(uint(reviewID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "review approval failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/reviews")
}

// POST - admin/reviews/reject/:id
func (h *AdminReviewHandler) RejectReviewHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "no review id found", errors.New("no review id given in parameter"))
		return
	}
	reviewID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid id", err)
		return
	}

	if err := h.ReviewService.RejectReviewService(uint(reviewID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "review rejected failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/reviews")
}
