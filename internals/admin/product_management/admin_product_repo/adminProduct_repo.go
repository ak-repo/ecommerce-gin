package adminproductrepo

import (
	productmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/product_management"
	adminproductinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type AdminProductRepo struct {
	DB *gorm.DB
}

func NewAdminProductRepo(db *gorm.DB) adminproductinterface.RepoInterface {
	return &AdminProductRepo{DB: db}
}

//----------------------------------------------- GET admin/products => all products ----------------------------------------------------------------

func (r *AdminProductRepo) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}

//----------------------------------------------- GET admin/products/:id => get product by id -----------------------------------------------------------

func (r *AdminProductRepo) GetProductByID(id uint) (*models.Product, error) {
	product := &models.Product{}

	err := r.DB.Preload("Category").First(product, id).Error
	return product, err
}

// ----------------------------------------------- POST admin/products/:id => update product info e.g.= stock,details etc ------------------------------------
func (r *AdminProductRepo) UpdateProductDetails(updatedProduct *productmanagement.UpdateProductRequest) error {
	updates := map[string]interface{}{}

	if updatedProduct.Title != nil {
		updates["title"] = *updatedProduct.Title
	}
	if updatedProduct.Description != nil {
		updates["description"] = *updatedProduct.Description
	}
	if updatedProduct.BasePrice != nil {
		updates["base_price"] = *updatedProduct.BasePrice
	}
	if updatedProduct.DiscountPrice != nil {
		updates["discount_price"] = *updatedProduct.DiscountPrice
	}
	if updatedProduct.Stock != nil {
		updates["stock"] = *updatedProduct.Stock
	}
	if updatedProduct.ImageURL != nil {
		updates["image_url"] = *updatedProduct.ImageURL
	}
	if updatedProduct.IsActive != nil {
		updates["is_active"] = *updatedProduct.IsActive
	}
	if updatedProduct.IsPublished != nil {
		updates["is_published"] = *updatedProduct.IsPublished
	}
	if updatedProduct.CategoryID != nil {
		updates["category_id"] = *updatedProduct.CategoryID
	}

	if len(updates) == 0 {
		return nil // nothing to update
	}

	// Use pointer for GORM
	product := &models.Product{ID: *updatedProduct.ID}

	return r.DB.Debug().Model(product).Updates(updates).Error
}

//----------------------------------------------- POST admin/products/delete => delete product -----------------------------------------------------------

func (r *AdminProductRepo) DeleteProduct(productID uint) error {
	return r.DB.Delete(&models.Product{}, productID).Error
}

//----------------------------------------------- GET admin/products/new => all products -----------------------------------------------------------

func (r *AdminProductRepo) AddNewProduct(newProduct *productmanagement.CreateProductRequest) error {
	product := models.Product{
		Title:         newProduct.Title,
		Description:   newProduct.Description,
		SKU:           newProduct.SKU,
		BasePrice:     newProduct.BasePrice,
		DiscountPrice: newProduct.DiscountPrice,
		Stock:         newProduct.Stock,
		ImageURL:      newProduct.ImageURL,
		CategoryID:    newProduct.CategoryID,
		IsActive:      newProduct.IsActive,
		IsPublished:   newProduct.IsPublished,
	}
	return r.DB.Create(&product).Error
}

//----------------------------------------------- GET admin/categories => all categories -----------------------------------------------------------

func (r *AdminProductRepo) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}
