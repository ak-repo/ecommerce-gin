package custproductservice

import (
	"errors"

	custproduct "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product"
	custproductinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product/cust_product_interface"
	reviewdto "github.com/ak-repo/ecommerce-gin/internals/review/review_DTO"
)

type CustomerProductService struct {
	CustomerProductRepo custproductinterface.RepoInterface
}

func NewCustomerProductService(custProductRepo custproductinterface.RepoInterface) custproductinterface.ServiceInterface {
	return &CustomerProductService{CustomerProductRepo: custProductRepo}
}

//----------------------------------------------- GET admin/products => all products ----------------------------------------------------------------

func (s *CustomerProductService) ListAllProductsService() ([]custproduct.CustomerProductListItem, error) {
	products, err := s.CustomerProductRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	if products == nil {
		return nil, errors.New("products not found")
	}

	var listProducts []custproduct.CustomerProductListItem

	for _, item := range products {
		product := custproduct.CustomerProductListItem{
			Title:         item.Title,
			ID:            item.ID,
			SKU:           item.SKU,
			BasePrice:     item.BasePrice,
			DiscountPrice: item.DiscountPrice,
			ImageURL:      item.ImageURL,
		}
		listProducts = append(listProducts, product)

	}
	return listProducts, nil
}

// ----------------------------------------------- GET admin/products/:id => get product by id -----------------------------------------------------------
func (s *CustomerProductService) ListProductByIDService(productID uint) (*custproduct.CustomerProductResponse, error) {

	product, err := s.CustomerProductRepo.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	listproduct := custproduct.CustomerProductResponse{
		Title:         product.Title,
		ID:            product.ID,
		Description:   product.Description,
		SKU:           product.SKU,
		BasePrice:     product.BasePrice,
		DiscountPrice: product.DiscountPrice,
		Stock:         product.Stock,
		ImageURL:      product.ImageURL,
		IsPublished:   product.IsPublished,
		Category: custproduct.CategoryDTO{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
	}
	var reviews []reviewdto.ReviewResponse
	for _, r := range product.Reviews {
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
