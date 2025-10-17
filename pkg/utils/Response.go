package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type JSONResponse struct {
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	Details     string      `json:"details,omitempty"`
	Role        string      `json:"role,omitempty"`
	CurrentYear int         `json:"current_year"`
	Data        interface{} `json:"data,omitempty"`
}
type TemplateData struct {
	Status      string
	Message     string
	Details     string
	Role        string
	CurrentYear int
	Data        interface{}
}

// RenderError returns JSON error response
func RenderError(ctx *gin.Context, status int, role string, message string, err error) {

	if role == "admin" {
		ctx.HTML(status, "pages/response/error.html", TemplateData{
			Status:      "error",
			Message:     message,
			Details:     err.Error(),
			Role:        role,
			CurrentYear: time.Now().Year(),
		})

	} else {
		ctx.JSON(status, JSONResponse{
			Status:      "error",
			Message:     message,
			Details:     err.Error(),
			Role:        role,
			CurrentYear: time.Now().Year(),
		})
	}

}

// RenderSuccess returns JSON success response
func RenderSuccess(ctx *gin.Context, status int, role string, message string, data interface{}) {

	if role == "admin" {
		ctx.HTML(status, "pages/response/success.html", TemplateData{
			Status:      "success",
			Message:     message,
			Role:        role,
			CurrentYear: time.Now().Year(),
			Data:        data,
		})

	} else {
		ctx.JSON(http.StatusOK,JSONResponse{
			Status:      "success",
			Message:     message,
			Role:        role,
			CurrentYear: time.Now().Year(),
			Data:        data,
		})
	}

}
