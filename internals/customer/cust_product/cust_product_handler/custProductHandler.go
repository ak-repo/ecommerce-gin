package custproducthandler

import (
	"errors"
	"net/http"
	"strconv"

	custproductinterface "github.com/ak-repo/ecommerce-gin/internals/customer/cust_product/cust_product_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CustomerProductHandler struct {
	CustomerProductService custproductinterface.ServiceInterface
}

func NewCustomerProductHandler(customerProductService custproductinterface.ServiceInterface) custproductinterface.HandlerInterface {
	return &CustomerProductHandler{CustomerProductService: customerProductService}
}

// GET - cust/products => all products
func (h *CustomerProductHandler) CustomerAllProducstListHandler(ctx *gin.Context) {
	products, err := h.CustomerProductService.ListAllProductsService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "Failed to load products", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "products fetched successfully", map[string]interface{}{
		"data": products,
	})
}

// GET - cust/product/:id => product by id
func (h *CustomerProductHandler) CustomerProductListHandler(ctx *gin.Context) {

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
	product, err := h.CustomerProductService.ListProductByIDService(uint(uid))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "customer", "product db not responsding", err)
		return
	}

	utils.RenderSuccess(ctx, http.StatusOK, "customer", "product fetched successfully", map[string]interface{}{
		"data": product,
	})

}
