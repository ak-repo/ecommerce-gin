package productservice

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
)

type ProductService interface {
	GetAllProductService() ([]models.Product, error)
	GetOneProductService(id string) (*models.Product, error)
	GetCategoriesService() ([]models.Category, error)
	// UpdateProductService(id string, form *producthandler.ProductUpdateForm) error
}

type productService struct {
	cfg         *config.Config
	productRepo productrepository.ProductRepo
}

func NewProductService(cfg *config.Config, productRepo productrepository.ProductRepo) ProductService {

	return &productService{cfg: cfg, productRepo: productRepo}
}

func (s *productService) GetAllProductService() ([]models.Product, error) {
	products, err := s.productRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) GetOneProductService(id string) (*models.Product, error) {
	product, err := s.productRepo.GetProductByID(id)

	if err != nil {
		return nil, err
	}
	return product, err
}

// category fetch
func (s *productService) GetCategoriesService() ([]models.Category, error) {

	categories, err := s.productRepo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, err
}

// // Update product details
// func (s *productService) UpdateProductService(id string, form *producthandler.ProductUpdateForm) error {
// 	return s.productRepo.UpdateProduct(id, form)
// }
