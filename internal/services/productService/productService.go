package productservice

import (
	"strconv"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
)

type ProductService interface {
	GetAllProductService() ([]models.Product, error)
	GetOneProductService(id string) (*models.Product, error)
	GetCategoriesService() ([]models.Category, error)
	UpdateProductService(product *models.Product, updatedProduct *models.UpdateProductInput) error
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
	idUID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	product, err := s.productRepo.GetProductByID(uint(idUID))

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

// Update product details

func (s *productService) UpdateProductService(product *models.Product, updatedProduct *models.UpdateProductInput) error {

	return s.productRepo.UpdateProductService(product, updatedProduct)

}
