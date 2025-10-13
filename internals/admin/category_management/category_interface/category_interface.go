package categoryinterface

import (
	categorydto "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllCategories(ctx *gin.Context)
	GetCategoryDetails(ctx *gin.Context)
	ShowNewCategoryForm(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
}

type Service interface {
	GetAllCategories() ([]categorydto.CategoryDTO, error)
	GetCategoryDetails(id uint) (*categorydto.CategoryDeatiledResponse, error)
	CreateCategory(name string) error
}

type Repository interface {
	FindAll() ([]models.Category, error)
	FindByID(categoryID uint) (*models.Category, error)
	Create(category *models.Category) error
}
