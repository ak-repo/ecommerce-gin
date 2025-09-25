package main

import (
	"log"

	"github.com/ak-repo/ecommerce-gin/config"
	"github.com/ak-repo/ecommerce-gin/internal/routes"
	db "github.com/ak-repo/ecommerce-gin/pkg/database"
	"github.com/gin-gonic/gin"
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

	routes.RegisterRoute(r, db, cfg)

	if err := r.Run(cfg.ServerAddress()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
