package reviewhandler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

// POST - admin/reviews
func (h *handler) GetAllReviews(ctx *gin.Context) {

	reviews, err := h.ReviewService.GetAllReviews()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "no review found", err)
		return
	}
	ctx.HTML(http.StatusOK, "pages/reviews/reviews.html", gin.H{
		"Reviews": reviews,
	})

}

// POST - admin/reviews/approve/:id
func (h *handler) ApporveReview(ctx *gin.Context) {

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

	if err := h.ReviewService.ApproveReview(uint(reviewID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "review approval failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/reviews")
}

// POST - admin/reviews/reject/:id
func (h *handler) RejectReview(ctx *gin.Context) {

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

	if err := h.ReviewService.RejectReview(uint(reviewID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "review rejected failed", err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/reviews")
}
