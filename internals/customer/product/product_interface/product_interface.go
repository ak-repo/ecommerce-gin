package productinterface

import (
	productdto "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_dto"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllProducts(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
}

type Service interface {
	GetAllProducts() ([]productdto.CustomerProductListItem, error)
	GetProductByID(productID uint) (*productdto.CustomerProductResponse, error)
}

type Repository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
}
