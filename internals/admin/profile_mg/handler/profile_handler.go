package profilehandler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	profiledto "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/profile_dto"
	profileinterface "github.com/ak-repo/ecommerce-gin/internals/admin/profile_mg/profile_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ProfileService profileinterface.Service
}

func NewProfileHandlerMG(profileService profileinterface.Service) profileinterface.Handler {
	return &handler{ProfileService: profileService}
}

func (h *handler) GetProfile(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "user id not found", err)
		return
	}

	profile, err := h.ProfileService.GetProfile(userID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "DB issue", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/profile/profile.html", gin.H{
		"Profile": profile,
	})
}

func (h *handler) GetAddress(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "user id not found", err)
		return
	}

	address, err := h.ProfileService.GetAddress(userID)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "address not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/profile/address.html", gin.H{
		"Address": address,
	})

}

func (h *handler) UpdateAddress(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "user id not found", err)
		return
	}

	var address profiledto.AddressDTO
	if err := ctx.ShouldBind(&address); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "binding isssue", err)
		return
	}

	if err := h.ProfileService.UpdateAddress(&address, userID); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "address update failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/profile")
}

func (h *handler) UploadPicture(ctx *gin.Context) {
	file, err := ctx.FormFile("profile_pic")
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "image file not supported", err)
		return
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "user id not found", err)
		return
	}

	timestamp := time.Now().Unix()
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("admin_%d_%d%s", userID, timestamp, ext)
	saveDir := "./web/uploads/profile"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to create upload directory", err)
		return
	}
	savePath := filepath.Join(saveDir, newFileName)

	// disk
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to upload image", err)
		return
	}

	//   DB 
	relativePath := filepath.Join("profile", newFileName)
	if err := h.ProfileService.UploadPicture(userID, relativePath); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to update profile picture", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/profile")
}
