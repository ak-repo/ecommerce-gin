package productrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type ProductRepo interface {
	GetAllProducts() ([]models.Product, error)
}

type productRepo struct {
	DB *gorm.DB
}

// New repo
func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{DB: db}

}


// Get all products from DB
func (r *productRepo) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}
