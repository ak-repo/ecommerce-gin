package productrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type ProductRepo interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	GetCategories() ([]models.Category, error)
	// UpdateProduct(id string, form *producthandler.ProductUpdateForm) error
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

// Get one product from DB by ID
func (r *productRepo) GetProductByID(id string) (*models.Product, error) {
	product := &models.Product{}

	err := r.DB.Preload("Category").First(product, id).Error
	return product, err
}

// Get categories
func (r *productRepo) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

// Update product details
// func (r *productRepo) UpdateProduct(id string, form *ProductRepo) error {
// 	return r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(models.Product{
// 		Title:       form.Title,
// 		Description: form.Description,
// 		CategoryID:  form.CategoryID,
// 		BasePrice:   form.Price,
// 		Stock:       form.Stock,
// 		IsActive:    form.IsActive,
// 		ImageURL:    form.ImageURL,
// 	}).Error
// }
