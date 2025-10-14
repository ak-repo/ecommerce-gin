package pagesrepo

import (
	pagesinter "github.com/ak-repo/ecommerce-gin/internals/customer/pages/pages_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPagesRepo(db *gorm.DB) pagesinter.Repository {
	return &repo{DB: db}
}

func (r *repo) GetBanners() ([]models.Banner, error) {

	var banners []models.Banner
	err := r.DB.Where("is_active", true).Find(&banners).Error
	return banners, err

}
