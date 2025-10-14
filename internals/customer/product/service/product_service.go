package productservice

import (
	"errors"

	productdto "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_dto"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_interface"
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
			BasePrice:     item.BasePrice,
			DiscountPrice: item.DiscountPrice,
			ImageURL:      item.ImageURL,
			CategoryName:  item.Category.Name,
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

	product := productdto.CustomerProductResponse{
		Title:         data.Title,
		ID:            data.ID,
		Description:   data.Description,
		BasePrice:     data.BasePrice,
		DiscountPrice: data.DiscountPrice,
		Stock:         data.Stock,
		ImageURL:      data.ImageURL,
		Category: productdto.CategoryDTO{
			ID:   data.Category.ID,
			Name: data.Category.Name,
		},
	}
	var reviews []productdto.Reviews
	for _, r := range data.Reviews {
		review := productdto.Reviews{
			ID:        r.ID,
			UserID:    r.UserID,
			Rating:    r.Rating,
			Comment:   r.Comment,
			CreatedAt: r.CreatedAt,
		}
		reviews = append(reviews, review)
	}
	product.Reviews = reviews

	// similar products
	var similarProducts []productdto.SimilarProducts
	similars, err := s.ProductRepo.FilterProducts(&productdto.FilterParams{
		Category: product.Category.Name,
		MinPrice: product.BasePrice - 10,
		MaxPrice: product.BasePrice + 10,
		Limit:    5,
	})
	if err != nil {
		return nil, err
	}

	for _, i := range similars {
		product := productdto.SimilarProducts{
			ID:            i.ID,
			Title:         i.Title,
			BasePrice:     i.BasePrice,
			DiscountPrice: i.DiscountPrice,
			ImageURL:      i.ImageURL,
			CategoryName:  i.Category.Name,
		}
		similarProducts = append(similarProducts, product)
	}
	product.SimilarProducts = similarProducts

	return &product, nil
}

func (s *service) FilterProducts(req *productdto.FilterParams) ([]productdto.CustomerProductListItem, error) {
	data, err := s.ProductRepo.FilterProducts(req)

	if data == nil || err != nil {
		return nil, errors.New("products not found")
	}

	var products []productdto.CustomerProductListItem

	for _, item := range data {
		product := productdto.CustomerProductListItem{
			Title:         item.Title,
			ID:            item.ID,
			BasePrice:     item.BasePrice,
			DiscountPrice: item.DiscountPrice,
			ImageURL:      item.ImageURL,
			CategoryName:  item.Category.Name,
		}
		products = append(products, product)

	}
	return products, nil
}
