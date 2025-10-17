package main

import (
	"log"

	"github.com/ak-repo/ecommerce-gin/config"
	db "github.com/ak-repo/ecommerce-gin/config/database"
	"github.com/ak-repo/ecommerce-gin/internals/routes"

	"github.com/gin-gonic/gin"
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

	routes.RegisterRoute(r, database, cfg)

	if err := r.Run(cfg.ServerAddress()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}


}

