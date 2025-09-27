package main

import (
	"log"
	"path/filepath"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"github.com/ak-repo/ecommerce-gin/internal/routes"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load env file: %v", err)
	}

	db, err := db.NewDB(cfg.GetDSN())
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	// r.LoadHTMLGlob("web/templates/**/*.html")
	r.HTMLRender = createMyRender("web/templates")

	routes.RegisterRoute(r, db, cfg)

	if err := r.Run(cfg.ServerAddress()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

func createMyRender(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	publicLayouts, _ := filepath.Glob(filepath.Join(templatesDir, "layouts", "base.html"))
	adminLayouts, _ := filepath.Glob(filepath.Join(templatesDir, "layouts", "admin_base.html"))

	publicPages, _ := filepath.Glob(filepath.Join(templatesDir, "pages", "**", "*.html"))
	for _, page := range publicPages {
		name, _ := filepath.Rel(templatesDir, page)
		if filepath.HasPrefix(page, filepath.Join(templatesDir, "pages", "admin")) {
			continue
		}
		files := append(publicLayouts, page)
		r.AddFromFiles(name, files...)
	}

	// Admin pages
	adminPages, _ := filepath.Glob(filepath.Join(templatesDir, "pages", "admin", "**", "*.html"))
	for _, page := range adminPages {
		name, _ := filepath.Rel(templatesDir, page)
		files := append(adminLayouts, page)
		r.AddFromFiles(name, files...)
	}

	return r
}

func SeedAdmin(db *gorm.DB) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("1122"), bcrypt.DefaultCost)
	db.FirstOrCreate(&models.User{
		Email:        "ak@fresh.com",
		PasswordHash: string(hash),
		Role:         "admin",
		IsActive:     true,
	}, models.User{Email: "ak@fresh.com"}) // prevents duplicate
}
