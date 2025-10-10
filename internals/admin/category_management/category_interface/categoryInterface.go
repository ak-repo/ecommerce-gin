package categoryinterface

import (
	categorymanagement "github.com/ak-repo/ecommerce-gin/internals/admin/category_management"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	ListAllCategoriesHandler(ctx *gin.Context)
	DetailedCategoryHandler(ctx *gin.Context)
	NewCategoryFormHandler(ctx *gin.Context)
	CreateNewcategoryHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	ListAllCategoriesService() ([]categorymanagement.CategoryDTO, error)
	CategoryDetailedResponse(ID uint) (*categorymanagement.CategoryDeatiledResponse, error)
	CreateNewCategoryService(name string) error
}

type RepoInterface interface {
	GetAllCategories() ([]models.Category, error)
	CategoryDetailedReponse(categoryID uint) (*models.Category, error)
	CreateNewCategory(category *models.Category) error
}
