package userhandler

import (
	"net/http"
	"time"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	userservice "github.com/ak-repo/ecommerce-gin/internal/services/userService"
	"github.com/gin-gonic/gin"
)

type UserAuthHandler struct {
	userAuthService userservice.UserAuthService
}

func NewUserAuthHandler(userAuthService userservice.UserAuthService) *UserAuthHandler {
	return &UserAuthHandler{userAuthService: userAuthService}
}

// ShowRegistrationForm renders the registration page.
func (h *UserAuthHandler) ShowRegistrationForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/user/userRegisterShow.html", gin.H{
		"Title":       "Create Your Account",
		"CurrentYear": time.Now().Year(),
		"Form":        gin.H{},
		"Errors":      gin.H{},
	})
}

// RegistrationHandler processes a new user registration.
func (h *UserAuthHandler) RegistrationHandler(ctx *gin.Context) {
	var input models.InputUser
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/registerFail.html", gin.H{
			"Title":       "Registration Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       "Invalid data submitted. Please check the fields and try again.",
		})
		return
	}

	user, err := h.userAuthService.Register(&input)
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
func (h *UserAuthHandler) ShowLoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/user/userLogin.html", gin.H{
		"Title":       "Login to freshBox",
		"CurrentYear": time.Now().Year(),
		"Form":        gin.H{},
		"Errors":      gin.H{},
	})
}

// LoginHandler processes a user login attempt
func (h *UserAuthHandler) LoginHandler(ctx *gin.Context) {
	var input models.InputUser
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "pages/user/loginFail.html", gin.H{
			"Title":       "Login Failed",
			"CurrentYear": time.Now().Year(),
			"Flash":       "Please provide both email and password.",
		})
		return
	}

	res, err := h.userAuthService.Login(&input)
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

func (h *UserAuthHandler) HomePageHandler(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/home/home.html", gin.H{})
}
