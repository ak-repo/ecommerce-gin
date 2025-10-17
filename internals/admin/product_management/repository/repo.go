package productrepo

import (
	productdto "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_dto"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewProductRepoMG(db *gorm.DB) productinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllProducts(req *productdto.ProductPagination) ([]models.Product, error) {
	db := r.DB.Model(&models.Product{})
	if req.Query != "" {
		db.Where("title ILIKE ?", "%"+req.Query+"%")
	}
	db.Count(&req.Total)

	offset := (req.Page - 1) * req.Limit

	var products []models.Product
	err := db.Preload("Category").
		Limit(req.Limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&products).Error
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


