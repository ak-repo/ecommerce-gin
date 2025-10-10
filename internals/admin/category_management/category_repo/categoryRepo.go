package categoryrepo

import (
	categoryinterface "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) categoryinterface.RepoInterface {
	return &CategoryRepo{DB: db}

}

func (r *CategoryRepo) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Preload("Products").
		Find(&categories).Error
	return categories, err
}

func (r *CategoryRepo) CategoryDetailedReponse(categoryID uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.
		Preload("Products").
		Where("id=?", categoryID).
		First(&category).Error
	return &category, err
}

// add new category
func (r *CategoryRepo) CreateNewCategory(category *models.Category) error {

	err := r.DB.Create(category).Error
	return err
}
