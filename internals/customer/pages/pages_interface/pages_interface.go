package pagesinter

import (
	pagesdto "github.com/ak-repo/ecommerce-gin/internals/customer/pages/dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetBanners(ctx *gin.Context)
}

type Service interface {
	GetBanners() ([]pagesdto.BannerDTO, error)
}

type Repository interface {
	GetBanners() ([]models.Banner, error)
}
