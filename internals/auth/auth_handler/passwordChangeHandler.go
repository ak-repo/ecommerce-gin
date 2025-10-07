package authhandler

import (
	"errors"
	"net/http"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) CustomerPasswordChange(ctx *gin.Context) {
	h.PasswordChangeHandler(ctx, RoleCustomer)
}
func (h *AuthHandler) AdminPasswordChange(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "pages/auth/passwordChange.html", gin.H{})
		return
	}
	h.PasswordChangeHandler(ctx, RoleAdmin)

}

// password chamge main function
func (h *AuthHandler) PasswordChangeHandler(ctx *gin.Context, role string) {
	var req auth.PasswordChange
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "invalid input", err)
		return
	}
	userID, exists := ctx.Get("userID")
	if !exists || userID == "" {
		utils.RenderError(ctx, http.StatusUnauthorized, role, "unauthorised ", errors.New("user id not found or not valid"))
		return
	}

	if err := h.authService.PasswordChangeService(userID.(uint), &req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, role, "password change failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, role, "password updated", nil)

}
