package pageshandler

import (
	"net/http"

	pagesinter "github.com/ak-repo/ecommerce-gin/internals/customer/pages/pages_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	PagesService pagesinter.Service
}

func NewPagesHanlder(service pagesinter.Service) pagesinter.Handler {
	return &handler{PagesService: service}
}

func (h *handler) GetBanners(ctx *gin.Context) {

	banners, err := h.PagesService.GetBanners()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "failed to fetch banners", err)
		return

	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "banners", banners)

}
