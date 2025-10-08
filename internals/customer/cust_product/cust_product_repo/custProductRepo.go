package custproductrepo

import (
	custproductinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product/cust_product_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type CustomerProductRepo struct {
	DB *gorm.DB
}

func NewCustomerProductRepo(db *gorm.DB) custproductinterface.RepoInterface {
	return &CustomerProductRepo{DB: db}
}

func (r *CustomerProductRepo) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").
		Find(&products).Error
	return products, err
}

func (r *CustomerProductRepo) GetProductByID(id uint) (*models.Product, error) {
	product := &models.Product{}

	err := r.DB.Preload("Category").
		Preload("Reviews", "status=?", "APPROVED").
		First(product, id).Error
	return product, err
}
