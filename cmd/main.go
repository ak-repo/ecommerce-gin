package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/ak-repo/ecommerce-gin/routes"

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

	database, err := db.NewDB(cfg.GetDSN())
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// r.Static("/web/static", "./web/static")

	r.HTMLRender = createMyRender("web/templates")

	routes.RegisterRoute(r, database, cfg)

	if err := r.Run(cfg.ServerAddress()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

func createMyRender(templatesDir string) multitemplate.Renderer {
	t := multitemplate.NewRenderer()

	adminLayouts, _ := filepath.Glob(filepath.Join(templatesDir, "layouts", "admin_base.html"))

	adminPages, _ := filepath.Glob(filepath.Join(templatesDir, "pages", "**", "*.html"))
	for _, page := range adminPages {
		name, _ := filepath.Rel(templatesDir, page)

		var files []string
		if strings.HasSuffix(name, "adminLogin.html") {
			// use admin login base
			files = []string{filepath.Join(templatesDir, "layouts", "admin_login_base.html"), page}
		} else {
			// use normal admin base with sidebar
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
		Status:        "Active",
		EmailVerified: true,
	}, models.User{Email: "admin@freshbox.com"})
}
