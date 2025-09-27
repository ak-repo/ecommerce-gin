package producthandler

import (
	"net/http"

	productservice "github.com/ak-repo/ecommerce-gin/internal/services/productService"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService productservice.ProductService
}

// New product Handler

func NewProductHandler(productService productservice.ProductService) *ProductHandler {

	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetAllProductHandler(ctx *gin.Context) {
	products, err := h.productService.GetAllProductService()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "products", gin.H{
			"Error": "Failed to load products: " + err.Error(),
			"User":  "",
		})
		return

	}
	ctx.HTML(http.StatusOK, "products", gin.H{
		"User":     "",
		"Products": products,
	})
}
