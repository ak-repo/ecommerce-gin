package authhandler

import (
	"fmt"
	"net/http"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

// sent otp
func (h *AuthHandler) SendOTPHandler(ctx *gin.Context) {
	var req auth.SendOTPRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}

	if err := h.authService.SentOTPService(&req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "otp generation failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", fmt.Sprintf("OTP sented into: %s", req.Email), map[string]interface{}{
		"expired_in": "1 minute",
	})
}

// verify
func (h *AuthHandler) VerifyOTPHandler(ctx *gin.Context) {

	var req auth.VerifyOTPRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "invalid input", err)
		return
	}
	if err := h.authService.VerifyOTPService(&req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "otp verification failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "Email verified", nil)
}
