package producthandler

import (
	"net/http"
	"time"

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

type ProductUpdateForm struct {
	Title       string
	Description string
	CategoryID  uint
	Price       float64
	Stock       int
	IsActive    bool
	ImageURL    string
}

// POST /admin/product/update/:id → handles update logic.
// func (h *AdminproductHandler) UpdateProduct(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	form := ProductUpdateForm{}

// 	if err := ctx.ShouldBind(&form); err != nil {
// 		ctx.String(http.StatusBadRequest, "Invalid form: %v", err)
// 		return
// 	}

// 	err := h.productService.UpdateProductService(id, &form)
// 	if err != nil {
// 		ctx.String(http.StatusInternalServerError, "Update failed: %v", err)
// 		return

// 	}
// 	// ctx.Redirect(http.StatusSeeOther, "/admin/product/"+id)
// 	ctx.String(http.StatusOK, "updated")
// }

// GET /admin/products/delete/:id → shows confirmation page.

// POST /admin/products/delete/:id → actually deletes the product.
