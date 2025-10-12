package reviewinterface

import (
	reviewdto "github.com/ak-repo/ecommerce-gin/internals/review/review_dto"
	"github.com/ak-repo/ecommerce-gin/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddReview(ctx *gin.Context)
	GetAllReviews(ctx *gin.Context)
	ApporveReview(ctx *gin.Context)
	RejectReview(ctx *gin.Context)
}

type Service interface {
	AddReview(req *reviewdto.CreateReviewRequest) error
	GetAllReviews() ([]reviewdto.AdminReviewResponse, error)
	ApproveReview(reviewID uint) error
	RejectReview(reviewID uint) error
}

type Repository interface {
	AddReview(review *models.Review) error
	GetAllReviews() ([]models.Review, error)
	ApproveReview(id uint) error
	RejectReview(id uint) error
}
