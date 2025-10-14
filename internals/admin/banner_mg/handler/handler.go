package bannerhandler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	bannerinter "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/banner_interface"
	bannerdto "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/dto"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	BannerService bannerinter.Service
}

func NewBannerHandlerMG(service bannerinter.Service) bannerinter.Handler {

	return &handler{BannerService: service}
}

// GET => /admin/banners/add
func (h *handler) CreateForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/banners/add_banner.html", nil)
}

// POST => /admin/banners/add
func (h *handler) Create(ctx *gin.Context) {
	var input bannerdto.CreateBannerRequest
	if err := ctx.ShouldBind(&input); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}
	//isActive
	input.IsActive = ctx.PostForm("is_active") == "on"
	// image
	imageURL, err := bannerUpload(ctx)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to upload image", err)
		return
	}
	input.ImageURL = imageURL

	id, err := h.BannerService.Create(&input)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to create banner", err)
		return
	}
	if id == 0 {
		ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/banners")
	} else {
		ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/v1/admin/banners/%d", id))
	}
}

// GET => /admin/banners/:id/update
func (h *handler) UpdateForm(ctx *gin.Context) {

	id := ctx.Param("id")
	bannerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {

		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}

	banner, err := h.BannerService.GetBannerByID(uint(bannerID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to fetch banner", err)
		return
	}
	ctx.HTML(http.StatusOK, "pages/banners/edit_banner.html", gin.H{"Banner": banner})
}

// POST => /admin/banners/:id/update
func (h *handler) Update(ctx *gin.Context) {

	var input bannerdto.UpdateBannerRequest
	if err := ctx.ShouldBind(&input); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}

	fmt.Println("ac:", input.IsActive)
	id := ctx.Param("id")
	bannerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {

		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}
	input.ID = uint(bannerID)

	// image
	imageURL, err := bannerUpload(ctx)
	if err == nil {
		// utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to upload image", err)
		input.ImageURL = imageURL
		return
	}
	if err := h.BannerService.Update(&input); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to update banner", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/v1/admin/banners/%d", input.ID))

}

// helper function
func bannerUpload(ctx *gin.Context) (string, error) {
	file, err := ctx.FormFile("image")
	if err != nil {
		return "", errors.New("image file not supported" + err.Error())
	}

	timestamp := time.Now().Unix()
	ext := filepath.Ext(file.Filename)
	NewFileName := fmt.Sprintf("banner_%d%s", timestamp, ext)
	saveDir := "./web/uploads/banners"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return "", errors.New("failed to create upload directory" + err.Error())
	}
	savePath := filepath.Join(saveDir, NewFileName)
	// save to disk
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		return "", errors.New("failed to upload image" + err.Error())

	}

	imageURL := filepath.Join("banners", NewFileName)
	return imageURL, nil

}

// GET => /admin/banners/:id/delete
func (h *handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	bannerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || id == "" {
		if err == nil {
			err = errors.New("invalid banner id")
		}
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}

	if err := h.BannerService.Delete(uint(bannerID)); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to delete banner", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/api/v1/admin/banners")
}

// GET => /admin/banners
func (h *handler) GetAllBanners(ctx *gin.Context) {
	banners, err := h.BannerService.GetAllBanners()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to fetch banners", err)
		return
	}
	ctx.HTML(http.StatusOK, "pages/banners/banners.html", gin.H{"Banners": banners})
}

// GET => /admin/banners/:id
func (h *handler) GetBannerByID(ctx *gin.Context) {
	id := ctx.Param("id")
	bannerID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || id == "" {
		if err == nil {
			err = errors.New("invalid banner id")
		}
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}
	banner, err := h.BannerService.GetBannerByID(uint(bannerID))

	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to fetch banner", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/banners/banner.html", gin.H{"Banner": banner})

}
