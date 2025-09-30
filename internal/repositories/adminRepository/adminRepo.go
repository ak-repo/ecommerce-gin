package adminrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type AdminRepo interface {
	GetAdminInfo(email string) (*models.User, error)
	GetAdminAddress(userID uint) (*models.Address, error)
}

type adminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) AdminRepo {
	return &adminRepo{DB: db}
}

// Get admin data from db
func (r *adminRepo) GetAdminInfo(email string) (*models.User, error) {
	var admin models.User
	err := r.DB.Where("email=?", email).First(&admin).Error
	return &admin, err
}

// Get admin address
func (r *adminRepo) GetAdminAddress(userID uint) (*models.Address, error) {

	var address models.Address
	err := r.DB.Where("user_id=?", userID).First(&address).Error
	return &address, err
}
