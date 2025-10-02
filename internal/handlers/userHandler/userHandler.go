package userhandler

import (
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

	userIDAny, exists := ctx.Get("userID")
	if !exists {
		ctx.HTML(http.StatusInternalServerError, "pages/user/profile.html", gin.H{
			"Error": "no email on token",
		})
		return
	}

	userUID, ok := userIDAny.(uint)
	if !ok || userUID == 0 {
		ctx.HTML(http.StatusInternalServerError, "pages/user/profile.html", gin.H{
			"Error": "no email on token",
		})
		return
	}

	profile, err := h.userService.UserProfileService(userUID)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/user/profile.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "pages/user/profile.html", gin.H{
		"User":    userUID,
		"Error":   nil,
		"Profile": profile,
	})

}

// GET user/address -> shop user address form
func (h *UserHandler) ShowAddressForm(ctx *gin.Context) {
	addressID := ctx.Param("address_id")
	userIDAny, _ := ctx.Get("userID")
	userUID := userIDAny.(uint)

	address := dto.AddressDTO{}
	if addressID == "0" {
		ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
			"Address": address,
			"User":    userUID,
		})
		return
	}

	profile, err := h.userService.UserProfileService(userUID)
	if err != nil {
		ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
			"Address": address,
			"User":    userUID,
			"Error":   err.Error(),
		})
		return
	}

	addID, _ := strconv.ParseUint(addressID, 10, 64)
	if profile.Address.ID == 0 || uint(addID) != profile.Address.ID {
		ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
			"Address": address,
			"User":    userUID,
		})
		return

	}

	ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
		"User":    userUID,
		"Address": profile.Address,
	})
}

// POST user/address/update -> add or update user address
func (h *UserHandler) UserAddressUpdateHandler(ctx *gin.Context) {
	addressID := ctx.Param("address_id")
	userIDAny, _ := ctx.Get("userID")
	userUID := userIDAny.(uint)

	var address dto.AddressDTO
	if err := ctx.ShouldBind(&address); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/address.html", gin.H{
			"User":    userUID,
			"Address": address,
			"Error":   err.Error(),
		})
		return
	}

	addressUID, err := strconv.ParseUint(addressID, 10, 64)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/user/address.html", gin.H{
			"User":    userUID,
			"Address": address,
			"Error":   "address id error: " + err.Error(),
		})
		return

	}
	if err := h.userService.UserAddressUpdateService(&address, uint(addressUID), userUID); err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/user/address.html", gin.H{
			"User":    userUID,
			"Address": address,
			"Error":   err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/user/profile")
}

// GET user/password -> return password change form
func (h *UserHandler) UserPasswordChangeFormHandler(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/user/passwordChange.html", gin.H{})
}

// POST user/password -> change password
func (h *UserHandler) UserPasswordChangeHandler(ctx *gin.Context) {

	userIDAny, exists := ctx.Get("userID")
	userUID := userIDAny.(uint)
	if !exists || userUID == 0 {
		ctx.HTML(http.StatusBadRequest, "pages/user/passwordChange.html", gin.H{
			"Error": "token not found, login required",
		})
		return

	}
	var passform dto.PasswordChange

	if err := ctx.ShouldBind(&passform); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/passwordChange.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	if passform.ConfirmPassword != passform.NewPassword {
		ctx.HTML(http.StatusBadRequest, "pages/user/passwordChange.html", gin.H{
			"Error": "confirm password not matching",
		})
		return
	}

	if passform.Password == passform.NewPassword {
		ctx.HTML(http.StatusBadRequest, "pages/user/passwordChange.html", gin.H{
			"Error": "old password same as new password",
		})
		return
	}

	err := h.userService.UserPasswordChangeService(passform.NewPassword, passform.Password, userUID)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/user/passwordChange.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "pages/user/passwordChange.html", gin.H{
		"Success": "Password changed",
	})
}
