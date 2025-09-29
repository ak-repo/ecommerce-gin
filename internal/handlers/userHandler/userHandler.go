package userhandler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ak-repo/ecommerce-gin/internal/models"
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
	var input models.InputUser
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/registerFail.html", gin.H{
			"Title":       "Registration Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       "Invalid data submitted. Please check the fields and try again.",
		})
		return
	}

	user, err := h.userService.Register(&input)
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
		"User":        user,
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
	var input models.InputUser
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/loginFail.html", gin.H{
			"Title":       "Login Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       "Please provide both email and password.",
		})
		return
	}

	res, err := h.userService.Login(&input)
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

// home page
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

	emailI, exists := ctx.Get("email")
	if !exists {
		ctx.String(http.StatusBadRequest, "no email found")
		return
	}

	emailStr, ok := emailI.(string)
	if !ok || emailStr == "" {
		ctx.String(http.StatusBadRequest, "invalid email")
		return
	}

	user, err := h.userService.UserProfileService(emailStr)
	if err != nil || user == nil {
		// render template with nil User safely
		ctx.HTML(http.StatusOK, "pages/user/profile.html", gin.H{
			"User":    nil,
			"Address": nil,
		})
		return
	}

	var addr *models.Address
	if len(user.Addresses) > 0 {
		addr = &user.Addresses[0]
	}

	ctx.HTML(http.StatusOK, "pages/user/profile.html", gin.H{
		"User":    user,
		"Address": addr,
	})
}

// GET user/address -> shop user address form
func (h *UserHandler) ShowAddressForm(ctx *gin.Context) {
	log.Println("hoooiii")
	addressID := ctx.Param("address_id")
	email, _ := ctx.Get("email")

	emailStr := email.(string)

	user, err := h.userService.UserProfileService(emailStr)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "user not found")
		return
	}
	if addressID == "0" {
		ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
			"Address": nil,
			"User":    email,
		})
		return

	}

	var addr *models.Address = nil
	if addressID != "" && addressID != "0" {
		id, _ := strconv.Atoi(addressID)
		for _, v := range user.Addresses {
			if v.ID == uint(id) {
				addr = &v
				break
			}
		}
	}

	ctx.HTML(http.StatusOK, "pages/user/address.html", gin.H{
		"User":    user,
		"Address": addr,
	})
}

// POST user/address add user address into db
func (h *UserHandler) UserAddressHandler(ctx *gin.Context) {
	address_id := ctx.Param("address_id")
	email, _ := ctx.Get("email")
	emailStr, _ := email.(string)
	address := models.Address{}
	if err := ctx.ShouldBind(&address); err != nil {
		ctx.String(http.StatusBadRequest, "binding failed")
		return
	}

	if err := h.userService.UserProfileUpdateService(emailStr, address_id, &address); err != nil {
		ctx.String(http.StatusBadRequest, "adding failed")

		return
	}

	ctx.Redirect(http.StatusSeeOther, "/user/profile")

}
