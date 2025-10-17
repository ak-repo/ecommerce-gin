package usershandler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	userdto "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/user_dto"
	usersinterface "github.com/ak-repo/ecommerce-gin/internals/admin/users_mg/user_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	UsersService usersinterface.Service
}

func NewAdminUserHandler(service usersinterface.Service) usersinterface.Handler {
	return &handler{UsersService: service}
}

// GET admin/users   => display all users & users search
func (h *handler) GetAllUsers(ctx *gin.Context) {

	var req userdto.UsersPagination
	var err error
	req.Query = ctx.Query("q")
	req.Role = ctx.Query("role")
	req.Status = ctx.Query("status")
	req.Limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "6"))
	if err != nil {
		req.Limit = 6
	}
	req.Page, err = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		req.Page = 1
	}

	users, err := h.UsersService.GetAllUsers(&req)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "users not found => DP not found", err)
		return
	}

	req.TotalPages = int(math.Ceil(float64(req.Total) / float64(req.Limit)))

	ctx.HTML(http.StatusOK, "pages/users/users.html", gin.H{
		"Users":        users,
		"Query":        req.Query,
		"FilterStatus": req.Status,
		"FilterRole":   req.Role,
		"Page":         req.Page,
		"Limit":        req.Limit,
		"TotalPages":   req.TotalPages,
		"CurrentYear":  time.Now().Year(),
	})
}

// GET admin/users/userID
func (h *handler) GetUserByID(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "user id not valid", err)
		return
	}

	user, err := h.UsersService.GetUserByID(uint(userID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to fetch user", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/users/user.html", gin.H{
		"User": user,
	})
}

// POST admin/users/role/:id
func (h *handler) ChangeUserRole(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid user ID", err)
		return
	}

	var req userdto.AdminUserRoleChange
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input data", err)
		return
	}

	req.ID = uint(userID)

	if err := h.UsersService.ChangeUserRole(&req); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to update user role", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/v1/admin/users/%d", userID))
}

// POST admin/users/block/:id
func (h *handler) BlockUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid user ID", err)
		return
	}

	if err := h.UsersService.BlockUser(uint(userID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to update user status", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/v1/admin/users/%d", userID))
}

// GET admin/users/add =>
func (h *handler) ShowUserAddForm(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/users/addUser.html", nil)
}

// POST admin/users/add
func (h *handler) CreateUser(ctx *gin.Context) {
	var req userdto.CreateUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}

	req.EmailVerified = ctx.PostForm("email_verified") == "1"

	id, err := h.UsersService.CreateUser(&req)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "user creation failed", err)
		return
	}

	if id == 0 {
		ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/users/")
	} else {
		ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/v1/admin/users/%d", id))
	}
}
