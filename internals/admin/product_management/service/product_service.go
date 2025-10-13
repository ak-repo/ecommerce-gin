package productmanagementservice

import (
	"errors"
	"strings"

	productdto "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_dto"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
)

type service struct {
	ProductRepo productinterface.Repository
}

func Newservice(repo productinterface.Repository) productinterface.Service {
	return &service{ProductRepo: repo}
}

func (s *service) GetAllProducts(query string) ([]productdto.ProductListItem, error) {
	data, err := s.ProductRepo.GetAllProducts()
	if data == nil || err != nil {
		return nil, errors.New("products not found")
	}

	var products []productdto.ProductListItem

	for _, item := range data {
		if strings.Contains(strings.ToLower(item.Title), strings.ToLower(query)) {
			product := productdto.ProductListItem{
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
			products = append(products, product)
		}

	}
	return products, nil
}

func (s *service) GetProductByID(productID uint) (*productdto.ProductResponse, error) {

	data, err := s.ProductRepo.GetProductByID(productID)

	if data == nil || err != nil {
		return nil, errors.New("product not found")
	}

	product := productdto.ProductResponse{
		Title:         data.Title,
		ID:            data.ID,
		Description:   data.Description,
		SKU:           data.SKU,
		BasePrice:     data.BasePrice,
		DiscountPrice: data.DiscountPrice,
		Stock:         data.Stock,
		ImageURL:      data.ImageURL,
		IsActive:      data.IsActive,
		IsPublished:   data.IsPublished,
		Category: productdto.CategoryDTO{
			ID:   data.Category.ID,
			Name: data.Category.Name,
		},
	}

	return &product, nil
}

func (s *service) UpdateProduct(req *productdto.UpdateProductRequest) error {

	product, err := s.ProductRepo.GetProductByID(*req.ID)

	if product == nil || err != nil {
		return errors.New("product not found")
	}

	if req.Title != nil {
		product.Title = *req.Title
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.BasePrice != nil {
		product.BasePrice = *req.BasePrice
	}
	if req.DiscountPrice != nil {
		product.DiscountPrice = *req.DiscountPrice
	}
	if req.Stock != nil {
		product.Stock = *req.Stock
	}
	if req.ImageURL != nil {
		product.ImageURL = *req.ImageURL
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}
	if req.IsPublished != nil {
		product.IsPublished = *req.IsPublished
	}
	if req.CategoryID != nil {
		product.CategoryID = *req.CategoryID
	}

	// Optional: Skip update if nothing is provided
	if req.Title == nil &&
		req.Description == nil &&
		req.BasePrice == nil &&
		req.DiscountPrice == nil &&
		req.Stock == nil &&
		req.ImageURL == nil &&
		req.IsActive == nil &&
		req.IsPublished == nil &&
		req.CategoryID == nil {
		return nil
	}

	return s.ProductRepo.UpdateProduct(product)
}

func (s *service) DeleteProduct(productID uint) error {
	return s.ProductRepo.DeleteProduct(productID)
}

func (s *service) AddProduct(req *productdto.CreateProductRequest) error {

	product := models.Product{
		Title:         req.Title,
		Description:   req.Description,
		SKU:           req.SKU,
		BasePrice:     req.BasePrice,
		DiscountPrice: req.DiscountPrice,
		Stock:         req.Stock,
		ImageURL:      req.ImageURL,
		CategoryID:    req.CategoryID,
		IsActive:      req.IsActive,
		IsPublished:   req.IsPublished,
	}
	return s.ProductRepo.AddProduct(&product)
}
