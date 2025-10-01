package producthandler

import (
	"net/http"
	"time"

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

// GET /products
func (h *ProductHandler) GetAllProductHandler(ctx *gin.Context) {
	products, err := h.productService.GetAllProductService()

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/products/products.html", gin.H{
			"Error": "Failed to load products: " + err.Error(),
			"User":  "",
		})
		return

	}

	email, ok := ctx.Get("email")
	if ok {
		ctx.HTML(http.StatusOK, "pages/products/products.html", gin.H{
			"Title":       "Products - freshBox",
			"Products":    products,
			"CurrentYear": time.Now().Year(),
			"User":        email,
		})
		return

	}
	ctx.HTML(http.StatusOK, "pages/products/products.html", gin.H{
		"Title":       "Products - freshBox",
		"Products":    products,
		"CurrentYear": time.Now().Year(),
		"User":        nil,
	})

}

// GET  /product/:id
func (h *ProductHandler) GetProductByIDHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.HTML(http.StatusBadRequest, "pages/products/product.html", gin.H{
			"Error": "No id found",
		})
	}

	product, err := h.productService.GetOneProductService(id)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/products/product.html", gin.H{
			"Error": "No product found: " + err.Error(),
		})
	}

	email, _ := ctx.Get("email")

	msg, _ := ctx.Cookie("flash")
	ctx.SetCookie("flash", "", -1, "/", "localhost", false, true)

	ctx.HTML(http.StatusOK, "pages/products/product.html", gin.H{
		"Product":     product,
		"CurrentYear": time.Now().Year(),
		"Message":     msg,
		"User":        email,
	})
}
