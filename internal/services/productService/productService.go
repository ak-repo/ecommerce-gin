package productservice

import (
	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	productrepository "github.com/ak-repo/ecommerce-gin/internal/repositories/productRepository"
)

type ProductService interface {
	GetAllProductService() ([]models.Product, error)
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
