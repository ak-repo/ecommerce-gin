package categoryservice

import (
	categorydto "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_dto"
	categoryinterface "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type service struct {
	CategoryRepo categoryinterface.Repository
}

func NewCategoryService(repo categoryinterface.Repository) categoryinterface.Service {
	return &service{CategoryRepo: repo}
}

func (s *service) GetAllCategories() ([]categorydto.CategoryDTO, error) {

	data, err := s.CategoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []categorydto.CategoryDTO
	for _, v := range data {
		cate := categorydto.CategoryDTO{
			ID:           v.ID,
			Name:         v.Name,
			ProductCount: len(v.Products),
		}
		categories = append(categories, cate)
	}
	return categories, nil
}

func (s *service) GetCategoryDetails(ID uint) (*categorydto.CategoryDeatiledResponse, error) {

	data, err := s.CategoryRepo.FindByID(ID)
	if err != nil {
		return nil, err
	}

	category := categorydto.CategoryDeatiledResponse{
		ID:           data.ID,
		Name:         data.Name,
		ProductCount: len(data.Products),
	}

	var products []categorydto.ProductListItem
	for _, p := range data.Products {
		product := categorydto.ProductListItem{
			ID:            p.ID,
			Title:         p.Title,
			SKU:           p.SKU,
			BasePrice:     p.BasePrice,
			DiscountPrice: p.DiscountPrice,
			Stock:         p.Stock,
			IsActive:      p.IsActive,
			IsPublished:   p.IsPublished,
			ImageURL:      p.ImageURL,
		}
		products = append(products, product)
	}

	category.Products = products
	return &category, nil
}

// create new category
func (s *service) CreateCategory(name string) error {

	category := models.Category{
		Name: name,
	}
	return s.CategoryRepo.Create(&category)

}
