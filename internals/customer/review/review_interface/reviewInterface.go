package custreviewinter

import (
	custreviewdto "github.com/ak-repo/ecommerce-gin/internals/customer/review/dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddReview(ctx *gin.Context)
}

type Service interface {
	AddReview(req *custreviewdto.CreateReviewRequest) error
}

type Repository interface {
	AddReview(review *models.Review) error
}
