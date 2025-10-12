package productrepository

import (
	productinterface "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) productinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").
		Find(&products).Error
	return products, err
}

func (r *repository) GetProductByID(id uint) (*models.Product, error) {
	product := &models.Product{}

	err := r.DB.Preload("Category").
		Preload("Reviews", "status=?", "APPROVED").
		First(product, id).Error
	return product, err
}
