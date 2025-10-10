package categoryservice

import (
	categorymanagement "github.com/ak-repo/ecommerce-gin/internals/admin/category_management"
	categoryinterface "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type CategoryService struct {
	CategoryRepo categoryinterface.RepoInterface
}

func NewCategoryService(repo categoryinterface.RepoInterface) categoryinterface.ServiceInterface {
	return &CategoryService{CategoryRepo: repo}
}

func (s *CategoryService) ListAllCategoriesService() ([]categorymanagement.CategoryDTO, error) {

	data, err := s.CategoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	var categories []categorymanagement.CategoryDTO

	for _, v := range data {
		cate := categorymanagement.CategoryDTO{
			ID:           v.ID,
			Name:         v.Name,
			ProductCount: len(v.Products),
		}
		categories = append(categories, cate)
	}
	return categories, nil
}

func (s *CategoryService) CategoryDetailedResponse(ID uint) (*categorymanagement.CategoryDeatiledResponse, error) {

	data, err := s.CategoryRepo.CategoryDetailedReponse(ID)
	if err != nil {
		return nil, err
	}

	category := categorymanagement.CategoryDeatiledResponse{
		ID:           data.ID,
		Name:         data.Name,
		ProductCount: len(data.Products),
	}

	var products []categorymanagement.ProductListItem
	for _, p := range data.Products {
		product := categorymanagement.ProductListItem{
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
func (s *CategoryService) CreateNewCategoryService(name string) error {

	category := models.Category{
		Name: name,
	}
	return s.CategoryRepo.CreateNewCategory(&category)

}
