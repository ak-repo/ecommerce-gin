package productrepository

import (
	productdto "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_dto"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
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

func (r *repository) FilterProducts(req *productdto.FilterParams) ([]models.Product, error) {

	db := r.DB.Model(&models.Product{})

	switch req.Category {
	case "smoothies":
		db = db.Where("category_id=1")
	case "pasta":
		db = db.Where("category_id=2")
	}

	if req.MinPrice > 0 && req.MaxPrice > 0 {
		db = db.Where("base_price BETWEEN ? AND ?", req.MinPrice, req.MaxPrice)
	} else if req.MinPrice > 0 {
		db = db.Where("base_price>=?", req.MinPrice)

	} else if req.MaxPrice > 0 {
		db = db.Where("base_price<=?", req.MaxPrice)
	}

	if req.Search != "" {
		db = db.Where("title ILIKE ?", "%"+req.Search+"%")
	}

	// sorting
	switch req.Sort {
	case "price_asc":
		db = db.Order("base_price ASC")
	case "price_desc":
		db = db.Order("base_price DESC")
	case "newest":
		db = db.Order("created_at DESC")
	}

	// pagination
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit
	db = db.Offset(offset).Limit(req.Limit)

	var products []models.Product
	if err := db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}
