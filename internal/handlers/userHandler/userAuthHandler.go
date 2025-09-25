package userhandler

import (
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

// user registration service
func (h *UserAuthHandler) RegistrationHandler(ctx *gin.Context) {

	input := models.InputUser{}

	if err := ctx.ShouldBind(&input); err != nil {

		// error
	}

	// respose

}
