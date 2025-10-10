package categoryhandler

import (
	"errors"
	"net/http"
	"strconv"

	categoryinterface "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryService categoryinterface.ServiceInterface
}

func NewCategoryHandler(service categoryinterface.ServiceInterface) categoryinterface.HandlerInterface {
	return &CategoryHandler{CategoryService: service}
}

// GET admin/categories
func (h *CategoryHandler) ListAllCategoriesHandler(ctx *gin.Context) {

	categories, err := h.CategoryService.ListAllCategoriesService()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "categories fetching failed", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/category/categories.html", gin.H{
		"Categories": categories,
	})
}

// GET admin/categories/:id
func (h *CategoryHandler) DetailedCategoryHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	catID, err := strconv.ParseUint(id, 10, 64)
	if id == "" || err != nil {
		if err == nil {
			err = errors.New("id not found")
		}
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "id not valid", err)
		return
	}
	category, err := h.CategoryService.CategoryDetailedResponse(uint(catID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "category fetching failed", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/category/category.html", gin.H{
		"Category": category,
	})
}

func (h *CategoryHandler) NewCategoryFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/category/addCategory.html", nil)

}

func (h *CategoryHandler) CreateNewcategoryHandler(ctx *gin.Context) {

	name := ctx.PostForm("name")
	if name == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", nil)
		return
	}
	if err := h.CategoryService.CreateNewCategoryService(name); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "category creation failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/categories")
}
