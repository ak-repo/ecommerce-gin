package productmanagementrepository

import (
	productinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) productinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}

func (r *repository) GetProductByID(id uint) (*models.Product, error) {
	product := &models.Product{}

	err := r.DB.Preload("Category").First(product, id).Error
	return product, err
}

func (r *repository) UpdateProduct(product *models.Product) error {

	return r.DB.Save(product).Error
}

func (r *repository) DeleteProduct(productID uint) error {
	return r.DB.Delete(&models.Product{}, productID).Error
}

func (r *repository) AddProduct(product *models.Product) error {
	return r.DB.Create(&product).Error
}
