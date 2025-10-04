package adminuserhandler

import (
	"net/http"
	"strconv"

	adminuserinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_management/admin_user_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminUserHandler struct {
	AdminUserService adminuserinterface.ServiceInterface
}

func NewAdminUserHandler(adminUserService adminuserinterface.ServiceInterface) adminuserinterface.HandlerInterface {
	return &AdminUserHandler{AdminUserService: adminUserService}
}

// GET admin/users   => display all users
func (h *AdminUserHandler) ListAllUsersHandler(ctx *gin.Context) {

	users, err := h.AdminUserService.AdminAllUsersService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "users not found => DP not found", err)
		return
	}
	ctx.HTML(http.StatusOK, "pages/users/users.html", gin.H{
		"Users": users,
	})
}

// GET admin/users/userID   => get user full info by id
func (h *AdminUserHandler) ListUserByIDHandler(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "user id not valid", err)
		return
	}

	user, err := h.AdminUserService.AdminGetUserByIDService(uint(userID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "user db error", err)
		return
	}

	if user == nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "user is null", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/users/user.html", gin.H{
		"User": user,
	})
}
