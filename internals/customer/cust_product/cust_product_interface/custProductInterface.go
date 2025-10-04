package custproductinterface

import (
	custproduct "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	CustomerAllProducstListHandler(ctx *gin.Context)
	CustomerProductListHandler(ctx *gin.Context)
}

type ServiceInterface interface {
	ListAllProductsService() ([]custproduct.CustomerProductListItem, error)
	ListProductByIDService(productID uint) (*custproduct.CustomerProductResponse, error)
}

type RepoInterface interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
}
