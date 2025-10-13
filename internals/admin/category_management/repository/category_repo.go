package categoryrepository

import (
	categoryinterface "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) categoryinterface.Repository {
	return &repository{DB: db}

}

func (r *repository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Preload("Products").
		Find(&categories).Error
	return categories, err
}

func (r *repository) FindByID(categoryID uint) (*models.Category, error) {
	var category models.Category
	err := r.DB.
		Preload("Products").
		Where("id=?", categoryID).
		First(&category).Error
	return &category, err
}

func (r *repository) Create(category *models.Category) error {
	err := r.DB.Create(category).Error
	return err
}
