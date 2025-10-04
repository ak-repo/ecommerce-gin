package adminproducthandler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	adminproductinterface "github.com/ak-repo/ecommerce-gin/internals/admin/product_management/admin_product_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminProductHandler struct {
	AdminProductService adminproductinterface.ServiceInterface
}

func NewAdminProductHandler(adminProductService adminproductinterface.ServiceInterface) adminproductinterface.HandlerInterface {
	return &AdminProductHandler{AdminProductService: adminProductService}
}

//----------------------------------------------- GET admin/products => all products ----------------------------------------------------------------

func (h *AdminProductHandler) AdminAllProducstListHandler(ctx *gin.Context) {
	products, err := h.AdminProductService.ListAllProductsService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "Failed to load products", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/product/adminProducts.html", gin.H{
		"Title":       "Products - Admin",
		"Products":    products,
		"CurrentYear": time.Now().Year(),
		"User":        nil,
	})
}

//----------------------------------------------- GET admin/products/:id => get product by id -----------------------------------------------------------

func (h *AdminProductHandler) AdminProductListHandler(ctx *gin.Context) {

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
	product, err := h.AdminProductService.ListProductByIDService(uint(uid))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "product db not responsding", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/product/adminProducts.html", gin.H{
		"Title":       "Products - Admin",
		"Products":    product,
		"CurrentYear": time.Now().Year(),
		"User":        nil,
	})

}

//----------------------------------------------- POST admin/products/:id => update product info e.g.= stock,details etc ------------------------------------

//----------------------------------------------- POST admin/products/delete => delete product -----------------------------------------------------------

//----------------------------------------------- GET admin/products/new => all products -----------------------------------------------------------

//----------------------------------------------- GET admin/products => all products -----------------------------------------------------------
