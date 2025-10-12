package productservice

import (
	"errors"

	productdto "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_dto"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_interface"
	reviewdto "github.com/ak-repo/ecommerce-gin/internals/review/review_dto"
)

type service struct {
	ProductRepo productinterface.Repository
}

func NewProductService(repo productinterface.Repository) productinterface.Service {
	return &service{ProductRepo: repo}
}

func (s *service) GetAllProducts() ([]productdto.CustomerProductListItem, error) {
	data, err := s.ProductRepo.GetAllProducts()

	if data == nil || err != nil {
		return nil, errors.New("products not found")
	}

	var products []productdto.CustomerProductListItem

	for _, item := range data {
		product := productdto.CustomerProductListItem{
			Title:         item.Title,
			ID:            item.ID,
			SKU:           item.SKU,
			BasePrice:     item.BasePrice,
			DiscountPrice: item.DiscountPrice,
			ImageURL:      item.ImageURL,
		}
		products = append(products, product)

	}
	return products, nil
}

func (s *service) GetProductByID(productID uint) (*productdto.CustomerProductResponse, error) {

	data, err := s.ProductRepo.GetProductByID(productID)
	if data == nil || err != nil {
		return nil, errors.New("product not found")
	}

	listproduct := productdto.CustomerProductResponse{
		Title:         data.Title,
		ID:            data.ID,
		Description:   data.Description,
		SKU:           data.SKU,
		BasePrice:     data.BasePrice,
		DiscountPrice: data.DiscountPrice,
		Stock:         data.Stock,
		ImageURL:      data.ImageURL,
		IsPublished:   data.IsPublished,
		Category: productdto.CategoryDTO{
			ID:   data.Category.ID,
			Name: data.Category.Name,
		},
	}
	var reviews []reviewdto.ReviewResponse
	for _, r := range data.Reviews {
		review := reviewdto.ReviewResponse{
			ID:        r.ID,
			ProductID: r.ProductID,
			UserID:    r.UserID,
			Rating:    r.Rating,
			Comment:   r.Comment,
			CreatedAt: r.CreatedAt,
		}
		reviews = append(reviews, review)
	}
	listproduct.Reviews = reviews

	return &listproduct, nil
}
