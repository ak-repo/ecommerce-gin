package adminproductinterface

import (
	productmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/product_management"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AdminAllProducstListHandler(ctx *gin.Context)
	AdminProductListHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	ListAllProductsService() ([]productmanagement.ProductListItem, error)
	ListProductByIDService(productID uint) (*productmanagement.ProductResponse, error)
	UpdateProductDetailsService(updatedProduct *productmanagement.UpdateProductRequest, productID uint) error
	DeleteProductService(productID uint) error
	AddNewProductService(newProduct *productmanagement.CreateProductRequest) error
}

type RepoInterface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProductDetails(productID uint, updatedProduct *productmanagement.UpdateProductRequest) error
	DeleteProduct(productID uint) error
	AddNewProduct(newProduct *productmanagement.CreateProductRequest) error
	GetCategories() ([]models.Category, error)
}
