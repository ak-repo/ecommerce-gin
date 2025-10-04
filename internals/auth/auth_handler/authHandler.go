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

type AuthHandler struct {
	authService authinterface.AuthServiceInterface
}

func NewAuthHandler(authServcie authinterface.AuthServiceInterface) authinterface.AuthHandlerInterface {
	return &AuthHandler{authService: authServcie}
}

//----------------------------------------------------- Register   => customer, store,  delivery-partner--------------------------------------------------------

// customer
func (h *AuthHandler) CustomerRegister(ctx *gin.Context) {
	h.RegisterationHandler(ctx, RoleCustomer)
}

// store
func (h *AuthHandler) StoreRegister(ctx *gin.Context) {
	h.RegisterationHandler(ctx, RoleStore)
}

// delivery-partner
func (h *AuthHandler) DeliveryPartnerRegister(ctx *gin.Context) {
	h.RegisterationHandler(ctx, RoleDeliveryPartner)
}

// main
func (h *AuthHandler) RegisterationHandler(ctx *gin.Context, role string) {

	var input auth.RegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "invalid input request", err)
		return
	}

	err := h.authService.RegisterService(&input, role)
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
func (h *AuthHandler) CustomerLogin(ctx *gin.Context) {
	h.LoginHandler(ctx, RoleCustomer)
}

// store
func (h *AuthHandler) StoreLogin(ctx *gin.Context) {
	h.LoginHandler(ctx, RoleStore)
}

// delivery-partner
func (h *AuthHandler) DeliveryPartnerLogin(ctx *gin.Context) {
	h.LoginHandler(ctx, RoleDeliveryPartner)
}

// Admin
func (h *AuthHandler) AdminLogin(ctx *gin.Context) {
	h.LoginHandler(ctx, RoleAdmin)
}

func (h *AuthHandler) AdminLoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/auth/adminLogin.html", gin.H{})
}
func (h *AuthHandler) LoginHandler(ctx *gin.Context, role string) {
	var input auth.LoginRequest
	if err := ctx.ShouldBind(&input); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "invalid inputes", err)
		return
	}

	res, err := h.authService.LoginService(&input, role)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, role, "invalid inputes", err)
		return
	}

	// Set secure cookies
	ctx.SetCookie("refreshToken", res.RefreshToken, int(res.RefreshExp), "/", "localhost", true, true)

	if role == RoleAdmin {
		ctx.SetCookie("accessToken", res.AccessToken, int(res.AccessExp), "/", "localhost", true, true)
		ctx.Redirect(http.StatusSeeOther, "/admin/dashboard")
	} else {
		utils.RenderSuccess(ctx, http.StatusOK, role, "Login successful", map[string]interface{}{
			"token":  res.AccessToken,
			"userID": res.User.ID,
		})
	}

}

// logout  => customer, store, delivery-partner, admin

// OTP
