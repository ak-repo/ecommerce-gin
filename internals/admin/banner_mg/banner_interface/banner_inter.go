package bannerinter

import (
	bannerdto "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(ctx *gin.Context)
	CreateForm(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdateForm(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetAllBanners(ctx *gin.Context)
	GetBannerByID(ctx *gin.Context)
}

type Service interface {
	Create(req *bannerdto.CreateBannerRequest) (uint, error)
	Update(req *bannerdto.UpdateBannerRequest) error
	Delete(bannerID uint) error
	GetAllBanners() ([]bannerdto.BannerResponse, error)
	GetBannerByID(bannerID uint) (*bannerdto.BannerResponse, error)
}

type Repository interface {
	GetAllBanners() ([]models.Banner, error)
	GetBannerByID(bannerID uint) (*models.Banner, error)
	Create(banner *models.Banner) error
	Update(banner *models.Banner) error
	Delete(bannerID uint) error
}
