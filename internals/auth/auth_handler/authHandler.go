package authhandler

import (
	"net/http"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

const (
	RoleCustomer        = "customer"
	RoleStore           = "store"
	RoleDeliveryPartner = "delivery_partner"
	RoleAdmin           = "admin"
)

type authHandler struct {
	authService authinterface.Service
}

func NewAuthHandler(servcie authinterface.Service) authinterface.Handler {
	return &authHandler{authService: servcie}
}

//----------------------------------------------------- Register   => customer, store,  delivery-partner--------------------------------------------------------

// customer
func (h *authHandler) CustomerRegister(ctx *gin.Context) {
	h.Registeration(ctx, RoleCustomer)
}

// main
func (h *authHandler) Registeration(ctx *gin.Context, role string) {

	var input auth.RegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "invalid input request", err)
		return
	}

	err := h.authService.Registeration(&input, role)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, role, "something went wrong while registering", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusCreated, role, role+"'s registration successful", map[string]interface{}{
		"username": input.Username,
		"email":    input.Email,
		"role":     role,
	})

}

//-------------------------------------------------------------- login   => customer, store, delivery-partner, admin ---------------------------------------

// customer
func (h *authHandler) CustomerLogin(ctx *gin.Context) {
	h.Login(ctx, RoleCustomer)
}

// Admin
func (h *authHandler) AdminLogin(ctx *gin.Context) {
	h.Login(ctx, RoleAdmin)
}

func (h *authHandler) AdminLoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/auth/adminLogin.html", gin.H{})
}
func (h *authHandler) Login(ctx *gin.Context, role string) {
	var input auth.LoginRequest
	if err := ctx.ShouldBind(&input); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "invalid inputes", err)
		return
	}

	res, err := h.authService.Login(&input, role)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "error while login", err)
		return
	}

	ctx.SetCookie("refreshToken", res.RefreshToken, int(res.RefreshExp), "/", "localhost", true, true)

	if role == RoleAdmin {
		ctx.SetCookie("accessToken", res.AccessToken, int(res.AccessExp), "/", "localhost", true, true)
		ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/dashboard")
	} else {
		utils.RenderSuccess(ctx, http.StatusOK, role, "Login successful", map[string]interface{}{
			"token":  res.AccessToken,
			"userID": res.User.ID,
		})
	}

}

// logout  => customer, store, delivery-partner, admin

// Admin
func (h *authHandler) AdminLogout(ctx *gin.Context) {
	h.Logout(ctx, RoleAdmin)
}

func (h *authHandler) Logout(ctx *gin.Context, role string) {

	ctx.SetCookie("accessToken", "", -1, "/", "localhost", true, true)
	ctx.SetCookie("refreshToken", "", -1, "/", "localhost", true, true)

	if role == "admin" {
		ctx.Redirect(http.StatusSeeOther, "/login")

	}

}
