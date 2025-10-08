package adminuserhandler

import (
	"fmt"
	"net/http"
	"strconv"

	usersmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/users_management"
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

// GET admin/users   => display all users & users search
func (h *AdminUserHandler) ListAllUsersHandler(ctx *gin.Context) {
	search := ctx.Query("q")
	role := ctx.Query("role")
	status := ctx.Query("status")
	users, err := h.AdminUserService.AdminAllUsersService(search)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "users not found => DP not found", err)
		return
	}
	var filterd []usersmanagement.AdminUserListDTO
	if role != "" {
		// clear(filered)
		for _, u := range users {

			if u.Role == role {
				filterd = append(filterd, u)
			}
		}
	}

	if status != "" {
		// clear(filterd )
		for _, u := range users {

			if u.Status == status {
				filterd = append(filterd, u)
			}
		}
	}
	if len(filterd) == 0 {
		filterd = users
	}
	ctx.HTML(http.StatusOK, "pages/users/users.html", gin.H{
		"Users":        filterd,
		"Query":        search,
		"FilterStatus": status,
		"FilterRole":   role,
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

// POST admin/users/role/:id   => user role change
func (h *AdminUserHandler) AdminUserRoleChangeHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid user ID", err)
		return
	}

	var req usersmanagement.AdminUserRoleChange
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input data", err)
		return
	}

	req.ID = uint(userID)

	if err := h.AdminUserService.AdminUserRoleChangeService(&req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to update user role", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/admin/users/%d", userID))
}

// POST admin/users/block/:id => user block and unblock
func (h *AdminUserHandler) AdminUserBlockHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid user ID", err)
		return
	}

	if err := h.AdminUserService.AdminUserBlockService(uint(userID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to update user status", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/admin/users/%d", userID))
}
