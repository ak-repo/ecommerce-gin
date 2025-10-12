package authhandler

import (
	"net/http"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *authHandler) CustomerPasswordChange(ctx *gin.Context) {
	h.PasswordChange(ctx, RoleCustomer)
}
func (h *authHandler) AdminPasswordChange(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "pages/auth/passwordChange.html", gin.H{})
		return
	}
	h.PasswordChange(ctx, RoleAdmin)

}

// password chamge main function
func (h *authHandler) PasswordChange(ctx *gin.Context, role string) {
	var req auth.PasswordChange
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "invalid input", err)
		return
	}
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "user id not found", err)
		return
	}

	if err := h.authService.PasswordChange(userID, &req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, role, "password change failed", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, role, "password updated", nil)

}
