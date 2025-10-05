package adminproductservice

import (
	"errors"
	"strings"

	productmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/product_management"
	adminproductinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_interface"
)

type AdminProductService struct {
	AdminProductRepo adminproductinterface.RepoInterface
}

func NewAdminProductService(adminProductRepo adminproductinterface.RepoInterface) adminproductinterface.ServiceInterface {
	return &AdminProductService{AdminProductRepo: adminProductRepo}
}

//----------------------------------------------- GET admin/products => all products ----------------------------------------------------------------

func (s *AdminProductService) ListAllProductsService(query string) ([]productmanagement.ProductListItem, error) {
	products, err := s.AdminProductRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	if products == nil {
		return nil, errors.New("products not found")
	}

	var listProducts []productmanagement.ProductListItem

	for _, item := range products {
		if strings.Contains(strings.ToLower(item.Title), strings.ToLower(query)) {
			product := productmanagement.ProductListItem{
				Title:         item.Title,
				ID:            item.ID,
				SKU:           item.SKU,
				BasePrice:     item.BasePrice,
				DiscountPrice: item.DiscountPrice,
				Stock:         item.Stock,
				ImageURL:      item.ImageURL,
				IsActive:      item.IsActive,
				IsPublished:   item.IsPublished,
			}
			listProducts = append(listProducts, product)
		}

	}
	return listProducts, nil
}

// ----------------------------------------------- GET admin/products/:id => get product by id -----------------------------------------------------------
func (s *AdminProductService) ListProductByIDService(productID uint) (*productmanagement.ProductResponse, error) {

	product, err := s.AdminProductRepo.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	listproduct := productmanagement.ProductResponse{
		Title:         product.Title,
		ID:            product.ID,
		Description:   product.Description,
		SKU:           product.SKU,
		BasePrice:     product.BasePrice,
		DiscountPrice: product.DiscountPrice,
		Stock:         product.Stock,
		ImageURL:      product.ImageURL,
		IsActive:      product.IsActive,
		IsPublished:   product.IsPublished,
		Category: productmanagement.CategoryDTO{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
	}

	return &listproduct, nil
}

// ----------------------------------------------- POST admin/products/:id => update product info e.g.= stock,details etc ------------------------------------
func (s *AdminProductService) UpdateProductDetailsService(updatedProduct *productmanagement.UpdateProductRequest) error {
	return s.AdminProductRepo.UpdateProductDetails(updatedProduct)
}

// ----------------------------------------------- POST admin/products/delete => delete product -----------------------------------------------------------
func (s *AdminProductService) DeleteProductService(productID uint) error {
	return s.AdminProductRepo.DeleteProduct(productID)
}

//----------------------------------------------- GET admin/products/new => create new product -----------------------------------------------------------

func (s *AdminProductService) AddNewProductService(newProduct *productmanagement.CreateProductRequest) error {
	return s.AdminProductRepo.AddNewProduct(newProduct)
}

// ----------------------------------------------- Get categories-----------------------------------------------------------
func (s *AdminProductService) GetCategoriesService() ([]productmanagement.CategoryDTO, error) {

	data, err := s.AdminProductRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	var categories []productmanagement.CategoryDTO
	for _, cat := range data {
		category := productmanagement.CategoryDTO{
			ID:   cat.ID,
			Name: cat.Name,
		}
		categories = append(categories, category)
	}
	return categories, nil
}
