package productmanagementhandler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	productdto "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_dto"
	productinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/product_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ProductService productinterface.Service
}

func Newhandler(service productinterface.Service) productinterface.Handler {
	return &handler{ProductService: service}
}

// GET admin/products
func (h *handler) GetAllProducts(ctx *gin.Context) {
	query := ctx.Query("q")
	products, err := h.ProductService.GetAllProducts(query)
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "Failed to load products", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/product/products.html", gin.H{
		"Products":    products,
		"Query":       query,
		"CurrentYear": time.Now().Year(),
	})
}

// GET admin/products/:id
func (h *handler) GetProductByID(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "product id not found", errors.New("no id in the give url=> parameter missing"))
		return
	}

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "product id invalid", err)
		return
	}
	product, err := h.ProductService.GetProductByID(uint(uid))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "failed to fetch product", err)
		return
	}

	log.Println("pro: ", product)
	ctx.HTML(http.StatusOK, "pages/product/product.html", gin.H{
		"Product":     product,
		"CurrentYear": time.Now().Year(),
	})

}

// GET => admin/products/update/:id
func (h *handler) EditProductForm(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := strconv.ParseUint(id, 10, 64)
	if id == "" || uid == 0 || err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "product id not found", errors.New("no id in the give url=> parameter missing"))
		return
	}
	product, err := h.ProductService.GetProductByID(uint(uid))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "product not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/product/editProduct.html", gin.H{
		"Product":    product,
		"Categories": product.Category,
	})
}

// POST => admin/products/update/:id
func (h *handler) UpdateProduct(ctx *gin.Context) {

	id := ctx.Param("id")
	uid65, err := strconv.ParseUint(id, 10, 64)
	if id == "" || uid65 == 0 || err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "product id not found", errors.New("no id in the give url=> parameter missing"))
		return
	}
	var updatedProduct productdto.UpdateProductRequest
	if err := ctx.ShouldBind(&updatedProduct); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "invalid or incompleted inputes", err)
		return

	}
	isActive := ctx.PostForm("is_active") == "on"
	updatedProduct.IsActive = &isActive

	isPublished := ctx.PostForm("is_published") == "on"
	updatedProduct.IsPublished = &isPublished
	uid := uint(uid65)
	updatedProduct.ID = &uid
	if updatedProduct.ID == nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "no id in request", errors.New("product id not valid"))
	}
	if err := h.ProductService.UpdateProduct(&updatedProduct); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "update unsuccessful", err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/admin/products/%d", uid))

}

// GET => admin/products/add
func (h *handler) AddProductForm(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/product/addProduct.html", nil)

}

// POST => admin/products/add
func (h *handler) AddProduct(ctx *gin.Context) {
	var req productdto.CreateProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}

	req.IsActive = ctx.PostForm("is_active") == "on"
	req.IsPublished = ctx.PostForm("is_published") == "on"

	if err := h.ProductService.AddProduct(&req); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "failed to add new product", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/products")
}
