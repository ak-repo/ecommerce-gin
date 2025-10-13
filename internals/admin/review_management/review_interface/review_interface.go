package reviewinterface

import (
	reviewdto "github.com/ak-repo/ecommerce-gin/internals/admin/review_management/dto"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllReviews(ctx *gin.Context)
	ApporveReview(ctx *gin.Context)
	RejectReview(ctx *gin.Context)
}

type Service interface {
	GetAllReviews() ([]reviewdto.AdminReviewResponse, error)
	ApproveReview(reviewID uint) error
	RejectReview(reviewID uint) error
}

type Repository interface {
	GetAllReviews() ([]models.Review, error)
	ApproveReview(id uint) error
	RejectReview(id uint) error
}
