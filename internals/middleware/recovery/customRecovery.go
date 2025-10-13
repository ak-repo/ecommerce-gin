package authmiddleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v\n", err)
				ctx.HTML(http.StatusInternalServerError, "pages/404/500.html", gin.H{})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
