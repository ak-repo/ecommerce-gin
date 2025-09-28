package productrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type ProductRepo interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	GetCategories() ([]models.Category, error)
	UpdateProductService(product *models.Product, updatedProduct *models.UpdateProductInput) error
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
func (s *productRepo) UpdateProductService(product *models.Product, updatedProduct *models.UpdateProductInput) error {
	updates := map[string]interface{}{}

	if updatedProduct.Title != nil {
		updates["name"] = *updatedProduct.Title
	}
	if updatedProduct.Description != nil {
		updates["description"] = *updatedProduct.Description
	}
	if updatedProduct.BasePrice != nil {
		updates["base_price "] = *updatedProduct.BasePrice
	}
	if updatedProduct.Stock != nil {
		updates["stock"] = *updatedProduct.Stock
	}
	if updatedProduct.IsActive != nil {
		updates["is_active"] = *updatedProduct.IsActive
	}

	if len(updates) == 0 {
		return nil // nothing to update
	}

	return s.DB.Model(product).Updates(updatedProduct).Error
}
