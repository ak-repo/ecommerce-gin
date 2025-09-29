package producthandler

import (
	"log"
	"net/http"
	"time"

	"github.com/ak-repo/ecommerce-gin/internal/models"
	productservice "github.com/ak-repo/ecommerce-gin/internal/services/productService"
	"github.com/gin-gonic/gin"
)

type AdminproductHandler struct {
	productService productservice.ProductService
}

// New product Handler
func NewAdminproductHandler(productService productservice.ProductService) *AdminproductHandler {

	return &AdminproductHandler{productService: productService}
}

// GET /products
func (h *AdminproductHandler) GetAllProductHandler(ctx *gin.Context) {
	products, err := h.productService.GetAllProductService()

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/admin/product/adminProducts.html", gin.H{
			"Error": "Failed to load products: " + err.Error(),
			"User":  "",
		})
		return

	}
	ctx.HTML(http.StatusOK, "pages/admin/product/adminProducts.html", gin.H{
		"Title":       "Products - freshBox",
		"Products":    products,
		"CurrentYear": time.Now().Year(),
		"User":        nil, // Pass the user object (or nil)
	})
}

// GET  /product/:id
func (h *AdminproductHandler) GetProductByIDHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		ctx.HTML(http.StatusBadRequest, "pages/admin/product/adminProduct.html", gin.H{
			"Error": "No id found",
		})
	}

	product, err := h.productService.GetOneProductService(id)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "pages/admin/product/adminProduct.html", gin.H{
			"Error": "No product found: " + err.Error(),
		})
	}

	ctx.HTML(http.StatusOK, "pages/admin/product/adminProduct.html", gin.H{
		"Product":     product,
		"CurrentYear": time.Now().Year(),
		"User":        nil,
		"Error":       "",
	})
}

// GET /admin/products/edit/:id → shows Edit form (prefilled).
func (h *AdminproductHandler) ShowProductEdit(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.HTML(http.StatusBadRequest, "pages/admin/product/productEdit.html", gin.H{
			"Error": "No ID found",
		})
		return
	}

	// Fetch product (with category)
	product, err := h.productService.GetOneProductService(idStr)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "pages/admin/product/productEdit.html", gin.H{
			"Error": "Product not found: " + err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "pages/admin/product/productEdit.html", gin.H{
		"Product": product,
	})
}

// POST /admin/product/update/:id → handles update logic.
func (h *AdminproductHandler) UpdateProductHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	product, err := h.productService.GetOneProductService(id)
	if err != nil {
		ctx.String(http.StatusConflict, "not found in db"+err.Error())
		return

	}

	updates := models.UpdateProductInput{}

	if err := ctx.ShouldBind(&updates); err != nil {
		ctx.String(http.StatusConflict, "not bind"+err.Error())
		return

	}
	log.Println("updates:", updates)
	if err := h.productService.UpdateProductService(product, &updates); err != nil {
		ctx.String(http.StatusConflict, "DB error"+err.Error())
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/product/"+id) 

}

// GET /admin/products/delete/:id → shows confirmation page.

// POST /admin/products/delete/:id → actually deletes the product.


