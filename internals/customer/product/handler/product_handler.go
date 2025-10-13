package producthandler

import (
	"errors"
	"net/http"
	"strconv"

	productinterface "github.com/ak-repo/ecommerce-gin/internals/customer/product/product_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	ProductService productinterface.Service
}

func NewProductHandler(service productinterface.Service) productinterface.Handler {
	return &handler{ProductService: service}
}

// GET - cust/products => all products
func (h *handler) GetAllProducts(ctx *gin.Context) {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "Failed to load products", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "products fetched successfully", products)
}

// GET - cust/product/:id => product by id
func (h *handler) GetProductByID(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "product id not found", errors.New("no id in the give url=> parameter missing"))
		return
	}

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "customer", "product id invalid", err)
		return
	}
	product, err := h.ProductService.GetProductByID(uint(uid))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "product db not responsding", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "product fetched successfully", product)

}
