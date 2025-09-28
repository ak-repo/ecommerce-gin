package adminrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type AdminAuthRepo interface {
	GetAdminInfo(admin *models.User) error
}

type adminAuthRepo struct {
	DB *gorm.DB
}

func NewAdminAuthRepo(db *gorm.DB) AdminAuthRepo {
	return &adminAuthRepo{DB: db}
}

// Get admin data from db
func (r *adminAuthRepo) GetAdminInfo(admin *models.User) error {

	return r.DB.Where("email=?", admin.Email).First(admin).Error
}
