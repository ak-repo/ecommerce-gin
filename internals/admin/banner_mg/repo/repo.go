package bannerrepo

import (
	bannerinter "github.com/ak-repo/ecommerce-gin/internals/admin/banner_mg/banner_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewBannerRepoMG(db *gorm.DB) bannerinter.Repository {

	return &repo{DB: db}
}

func (r *repo) Create(banner *models.Banner) error {
	err := r.DB.Create(banner).Error
	return err
}


func (r *repo) Update(banner *models.Banner) error {
	err := r.DB.Save(banner).Error
	return err
}

func (r *repo) Delete(bannerID uint) error {
	err := r.DB.Delete(&models.Banner{}, "id=?", bannerID).Error
	return err
}

func (r *repo) GetAllBanners() ([]models.Banner, error) {
	var banners []models.Banner
	err := r.DB.Find(&banners).Error
	return banners, err
}

func (r *repo) GetBannerByID(bannerID uint) (*models.Banner, error) {
	var banner models.Banner
	err := r.DB.Where("id=?", bannerID).First(&banner).Error
	return &banner, err

}
