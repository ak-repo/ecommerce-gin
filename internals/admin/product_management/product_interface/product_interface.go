package productmanagementinterface

import (
	productdto "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllProducts(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	AddProduct(ctx *gin.Context)
	AddProductForm(ctx *gin.Context)
	EditProductForm(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type Service interface {
	GetAllProducts(req *productdto.ProductPagination) ([]productdto.ProductListItem, error)
	GetProductByID(productID uint) (*productdto.ProductResponse, error)
	AddProduct(newProduct *productdto.CreateProductRequest) error
	UpdateProduct(updatedProduct *productdto.UpdateProductRequest) error
	DeleteProduct(productID uint) error
}

type Repository interface {
	GetAllProducts(req *productdto.ProductPagination) ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	AddProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID uint) error
}


