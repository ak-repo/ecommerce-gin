package adminproducthandler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	productmanagement "github.com/ak-repo/ecommerce-gin/internals/admin/product_management"
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

//----------------------------------------------- GET admin/products => all products && search ----------------------------------------------------------------

func (h *AdminProductHandler) AdminAllProducstListHandler(ctx *gin.Context) {
	query := ctx.Query("q")
	products, err := h.AdminProductService.ListAllProductsService(query)
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

	log.Println("pro: ", product)
	ctx.HTML(http.StatusOK, "pages/product/product.html", gin.H{
		"Product":     product,
		"CurrentYear": time.Now().Year(),
	})

}

//----------------------------------------------- POST admin/products/:id => update product info e.g.= stock,details etc ------------------------------------

func (h *AdminProductHandler) AdminShowProductEditPageHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := strconv.ParseUint(id, 10, 64)
	if id == "" || uid == 0 || err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "product id not found", errors.New("no id in the give url=> parameter missing"))
		return
	}
	product, err := h.AdminProductService.ListProductByIDService(uint(uid))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "product not found", err)
		return
	}

	categories, err := h.AdminProductService.GetCategoriesService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "categories not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/product/editProduct.html", gin.H{
		"Product":    product,
		"Categories": categories,
	})
}

func (h *AdminProductHandler) AdminUpdateProductHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	uid65, err := strconv.ParseUint(id, 10, 64)
	if id == "" || uid65 == 0 || err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "product id not found", errors.New("no id in the give url=> parameter missing"))
		return
	}
	var updatedProduct productmanagement.UpdateProductRequest
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
	if err := h.AdminProductService.UpdateProductDetailsService(&updatedProduct); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "update unsuccessful", err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/admin/products/%d", uid))

}

//----------------------------------------------- POST admin/products/delete => delete product -----------------------------------------------------------

//----------------------------------------------- GET admin/products/new => new products -----------------------------------------------------------

func (h *AdminProductHandler) AdminAddProductPageHandler(ctx *gin.Context) {
	categories, err := h.AdminProductService.GetCategoriesService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "categories not found", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/product/addProduct.html", gin.H{
		"Categories": categories,
	})

}

func (h *AdminProductHandler) AdminAddProductHandler(ctx *gin.Context) {
	var product productmanagement.CreateProductRequest
	if err := ctx.ShouldBind(&product); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", err)
		return
	}

	product.IsActive = ctx.PostForm("is_active") == "on"
	product.IsPublished = ctx.PostForm("is_published") == "on"

	if err := h.AdminProductService.AddNewProductService(&product); err != nil {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "failed to add new product", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/products")
}

//----------------------------------------------- GET admin/products => all products -----------------------------------------------------------
