package userhandler

import (
	"net/http"
	"strconv"

	userservice "github.com/ak-repo/ecommerce-gin/internal/services/userService"
	"github.com/gin-gonic/gin"
)

type AdminUserHandler struct {
	AdminUserService userservice.AdminUserService
}

func NewAdminUserHandler(adminUserService userservice.AdminUserService) *AdminUserHandler {
	return &AdminUserHandler{AdminUserService: adminUserService}
}

// GET admin/users   => display all users
func (h *AdminUserHandler) ListAllUsersHandler(ctx *gin.Context) {

	users, err := h.AdminUserService.AdminAllUsersService()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/admin/users/users.html", gin.H{
			"Errors": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "pages/admin/users/users.html", gin.H{
		"Users": users,
	})
}

// GET admin/users/userID   => get user full info by id
func (h *AdminUserHandler) ListUserByIDHandler(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/admin/users/user.html", gin.H{
			"Error": "Invalid user ID: " + err.Error(),
		})
		return
	}

	user, err := h.AdminUserService.AdminGetUserByIDService(uint(userID))
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/admin/users/user.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	if user == nil {
		ctx.HTML(http.StatusNotFound, "pages/admin/users/user.html", gin.H{
			"Error": "User not found",
		})
		return
	}

	ctx.HTML(http.StatusOK, "pages/admin/users/user.html", gin.H{
		"User": user,
	})
}
