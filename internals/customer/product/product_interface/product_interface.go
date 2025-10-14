package productinterface

import (
	productdto "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllProducts(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	FilterProducts(ctx *gin.Context)
}

type Service interface {
	GetAllProducts() ([]productdto.CustomerProductListItem, error)
	GetProductByID(productID uint) (*productdto.CustomerProductResponse, error)
	FilterProducts(req *productdto.FilterParams) ([]productdto.CustomerProductListItem, error)
}

type Repository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	FilterProducts(req *productdto.FilterParams) ([]models.Product, error)
}
