package userhandler

import (
	"net/http"

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

// GET /user/register -> show form
func (h *UserAuthHandler) ShowRegistrationForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "userRegisterShow.html", gin.H{
		"Form":   gin.H{},
		"Errors": gin.H{},
		"Flash":  "",
	})
}

// POST user/register
func (h *UserAuthHandler) RegistrationHandler(ctx *gin.Context) {

	input := models.InputUser{}

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "registerFail.html", gin.H{
			"Form": gin.H{
				"Email": input.Email,
			},
			"Errors": map[string]string{
				"General": "Please fill all required fields correctly",
			},
			"Flash": "Invalid input",
		})
		return
	}

	user, err := h.userAuthService.Register(&input)
	if err != nil {
		ctx.HTML(http.StatusConflict, "registerFail.html", gin.H{
			"Form": gin.H{
				"Email": input.Email,
			},
			"Errors": gin.H{"general": err.Error()},
			"Flash":  "",
		})
		return
	}
	ctx.HTML(http.StatusCreated, "registerSuccess.html", gin.H{
		"Form": gin.H{
			"Email": user.Email,
		},
		"Errors": "",
		"Flash":  "",
	})

}


// GET /user/login -> show form
func (h *UserAuthHandler) ShowLoginForm(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "userLogin.html", gin.H{
		"Form":   gin.H{},
		"Errors": gin.H{},
		"Flash":  "",
	})
}

// POSt /user/login -> Login 
func (h *UserAuthHandler) LoginHandler(ctx *gin.Context) {
	input := models.InputUser{}

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "loginFail.html", gin.H{
			"Form": gin.H{
				"Email": input.Email,
			},
			"Errors": map[string]string{
				"General": "Please fill all required fields correctly",
			},
			"Flash": "Invalid input",
		})
		return
	}

	res, err := h.userAuthService.Login(&input)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "loginFail.html", gin.H{
			"Form": gin.H{
				"Email": input.Email,
			},
			"Errors": map[string]string{
				"General": err.Error(),
			},
			"Flash": "Somthing went wroug",
		})
		return

	}

	//cookie
	ctx.SetCookie("refreshToken", res.RefreshToken, int(res.RefreshExp), "/", "", true, true)
	ctx.SetCookie("accessToken", res.AccessToken, int(res.AccessExp), "/", "", true, true)

	ctx.HTML(http.StatusOK, "loginSuccess.html", gin.H{
		"Email": res.User.Email,
	})

}
