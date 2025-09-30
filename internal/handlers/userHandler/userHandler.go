package userhandler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ak-repo/ecommerce-gin/internal/dto"
	userservice "github.com/ak-repo/ecommerce-gin/internal/services/userService"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService userservice.UserService
}

func NewUserHandler(userService userservice.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// ShowRegistrationForm renders the registration page.
func (h *UserHandler) ShowRegistrationForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/user/userRegisterShow.html", gin.H{
		"Title":       "Create Your Account",
		"CurrentYear": time.Now().Year(),
		"Form":        gin.H{},
		"Errors":      gin.H{},
	})
}

// RegistrationHandler processes a new user registration.
func (h *UserHandler) RegistrationHandler(ctx *gin.Context) {
	var input dto.RegisterRequest
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/registerFail.html", gin.H{
			"Title":       "Registration Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       "Invalid data submitted. Please check the fields and try again.",
		})
		return
	}

	err := h.userService.RegisterService(&input)
	if err != nil {
		ctx.HTML(http.StatusConflict, "pages/user/registerFail.html", gin.H{
			"Title":       "Registration Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusCreated, "pages/user/registerSuccess.html", gin.H{
		"Title":       "Registration Successful!",
		"CurrentYear": time.Now().Year(),
		"User":        input.Username,
	})
}

// ShowLoginForm renders the login page.
func (h *UserHandler) ShowLoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/user/userLogin.html", gin.H{
		"Title":       "Login to freshBox",
		"CurrentYear": time.Now().Year(),
		"Form":        gin.H{},
		"Errors":      gin.H{},
	})
}

// LoginHandler processes a user login attempt
func (h *UserHandler) LoginHandler(ctx *gin.Context) {
	var input dto.LoginRequest
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/loginFail.html", gin.H{
			"Title":       "Login Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       "Please provide both email and password.",
		})
		return
	}

	res, err := h.userService.LoginService(&input)
	if err != nil {
		ctx.HTML(http.StatusUnauthorized, "pages/user/loginFail.html", gin.H{
			"Title":       "Login Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       err.Error(),
		})
		return
	}

	// Set secure cookies
	ctx.SetCookie("accessToken", res.AccessToken, int(res.AccessExp), "/", "localhost", true, true)
	ctx.SetCookie("refreshToken", res.RefreshToken, int(res.RefreshExp), "/", "localhost", true, true)

	ctx.Redirect(http.StatusSeeOther, "/")
}

// GET Logout
func (h *UserHandler) UserLogout(ctx *gin.Context) {

	ctx.SetCookie("accessToken", "", 0, "/", "localhost", true, true)
	ctx.SetCookie("refreshToken", "", 0, "/", "localhost", true, true)

	ctx.Redirect(http.StatusSeeOther, "/")

}

// ------------------------------------------------------------------------------------------------------------------------------------------
// GET /home   => home page
func (h *UserHandler) HomePageHandler(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		ctx.HTML(http.StatusOK, "pages/home/home.html", gin.H{
			"User":        nil,
			"CurrentYear": time.Now().Year(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "pages/home/home.html", gin.H{
		"User":        email,
		"CurrentYear": time.Now().Year(),
	})
}

// User profile GET user/profile
func (h *UserHandler) UserProfileHandler(ctx *gin.Context) {

	email, exists := ctx.Get("email")
	if !exists {
		ctx.HTML(http.StatusInternalServerError, "pages/user/profile.html", gin.H{
			"Error": "no email on token",
		})
		return
	}

	emailStr, ok := email.(string)
	if !ok || emailStr == "" {
		ctx.HTML(http.StatusInternalServerError, "pages/user/profile.html", gin.H{
			"Error": "no email on token",
		})
		return
	}

	profile, err := h.userService.UserProfileService(emailStr)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/user/profile.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "pages/user/profile.html", gin.H{
		"User":    email,
		"Error":   nil,
		"Profile": profile,
	})

}

// GET user/address -> shop user address form
func (h *UserHandler) ShowAddressForm(ctx *gin.Context) {
	addressID := ctx.Param("address_id")
	email, _ := ctx.Get("email")
	emailStr := email.(string)

	address := dto.AddressDTO{}
	if addressID == "0" {
		ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
			"Address": address,
			"User":    email,
		})
		return
	}

	profile, err := h.userService.UserProfileService(emailStr)
	if err != nil {
		ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
			"Address": address,
			"User":    email,
			"Error":   err.Error(),
		})
		return
	}

	addID, _ := strconv.ParseUint(addressID, 10, 64)
	if profile.Address.ID == 0 || uint(addID) != profile.Address.ID {
		ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
			"Address": address,
			"User":    email,
		})
		return

	}

	ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
		"User":    email,
		"Address": profile.Address,
	})
}

// POST user/address/update -> add or update user address
func (h *UserHandler) UserAddressUpdateHandler(ctx *gin.Context) {
	addressID := ctx.Param("address_id")
	email, _ := ctx.Get("email")
	emailStr := email.(string)

	var address dto.AddressDTO
	if err := ctx.ShouldBind(&address); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/address.html", gin.H{
			"User":    email,
			"Address": address,
			"Error":   err.Error(),
		})
		return
	}
	log.Println("phone:", address.Phone)

	if err := h.userService.UserAddressUpdateService(&address, addressID, emailStr); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/address.html", gin.H{
			"User":    email,
			"Address": address,
			"Error":   err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/user/profile")
}
