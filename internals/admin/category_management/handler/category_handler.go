package categoryhandler

import (
	"errors"
	"net/http"
	"strconv"

	categoryinterface "github.com/ak-repo/ecommerce-gin/internals/admin/category_management/category_interface"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	CategoryService categoryinterface.Service
}

func NewCategoryHandler(service categoryinterface.Service) categoryinterface.Handler {
	return &handler{CategoryService: service}
}

// GET admin/categories
func (h *handler) GetAllCategories(ctx *gin.Context) {

	categories, err := h.CategoryService.GetAllCategories()
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "categories fetching failed", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/category/categories.html", gin.H{
		"Categories": categories,
	})
}

// GET admin/categories/:id
func (h *handler) GetCategoryDetails(ctx *gin.Context) {

	id := ctx.Param("id")
	catID, err := strconv.ParseUint(id, 10, 64)
	if id == "" || err != nil {
		if err == nil {
			err = errors.New("id not found")
		}
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "id not valid", err)
		return
	}
	category, err := h.CategoryService.GetCategoryDetails(uint(catID))
	if err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "category fetching failed", err)
		return
	}

	ctx.HTML(http.StatusOK, "pages/category/category.html", gin.H{
		"Category": category,
	})
}

func (h *handler) ShowNewCategoryForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "pages/category/addCategory.html", nil)

}

func (h *handler) CreateCategory(ctx *gin.Context) {

	name := ctx.PostForm("name")
	if name == "" {
		utils.RenderError(ctx, http.StatusBadRequest, "admin", "invalid input", nil)
		return
	}
	if err := h.CategoryService.CreateCategory(name); err != nil {
		utils.RenderError(ctx, http.StatusInternalServerError, "admin", "category creation failed", err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/admin/categories")
}
