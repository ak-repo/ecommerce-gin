package routes

import (
	"errors"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/ak-repo/ecommerce-gin/config"
	db "github.com/ak-repo/ecommerce-gin/config/database"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	authmiddleware "github.com/ak-repo/ecommerce-gin/pkg/middleware/auth"
	"github.com/ak-repo/ecommerce-gin/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Routes handler
func RegisterRoute(r *gin.Engine, db *db.Database, cfg *config.Config) {
	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // for local react
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger(), gin.Recovery())

	r.HTMLRender = createMyRender("web/templates")
	r.Static("uploads", "./web/uploads")

	// 404 handling
	r.NoRoute(func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		if role == "admin" {
			ctx.HTML(http.StatusNotFound, "pages/response/404.html", gin.H{
				"CurrentYear": time.Now(),
			})
			return
		}
		if role != nil {
			utils.RenderError(ctx, http.StatusNotFound, role.(string), "no page found", errors.New("invalid url"))
		}
		utils.RenderError(ctx, http.StatusNotFound, "no role", "no page found", errors.New("invalid url"))
	})
	r.Use(authmiddleware.AccessMiddleware(cfg))

	RegisterAdminRoutes(r, db, cfg)
	RegisterCustomerRoutes(r, db, cfg)

}

func createMyRender(templatesDir string) multitemplate.Renderer {
	t := multitemplate.NewRenderer()

	adminLayouts, _ := filepath.Glob(filepath.Join(templatesDir, "layouts", "admin_base.html"))
	adminPages, _ := filepath.Glob(filepath.Join(templatesDir, "pages", "**", "*.html"))

	for _, page := range adminPages {
		name, _ := filepath.Rel(templatesDir, page)

		var files []string

		if strings.HasSuffix(name, "adminLogin.html") {
			// Use admin login base
			files = []string{filepath.Join(templatesDir, "layouts", "admin_login_base.html"), page}

		} else if strings.HasSuffix(name, "success.html") || strings.HasSuffix(name, "error.html") || strings.HasSuffix(name, "404.html") {
			// Render without any base
			files = []string{page}

		} else {
			// Default: use normal admin base with sidebar
			files = append(adminLayouts, page)
		}

		t.AddFromFiles(name, files...)
	}

	return t
}

func SeedAdmin(db *gorm.DB) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("11"), bcrypt.DefaultCost)
	db.FirstOrCreate(&models.User{
		Username:      "super admin",
		Email:         "admin@freshbox.com",
		PasswordHash:  string(hash),
		Role:          "admin",
		Status:        "active",
		EmailVerified: true,
	}, models.User{Email: "admin@freshbox.com"})
}
