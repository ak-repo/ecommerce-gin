package reviewinterface

import (
	reviewdto "github.com/ak-repo/ecommerce-gin/internals/review/review_DTO"
	"github.com/ak-repo/ecommerce-gin/models"
)

type ServiceInterface interface {
	CreateNewReviewService(req *reviewdto.CreateReviewRequest) error
	ListAllReviewsForVerifyService() ([]reviewdto.AdminReviewResponse, error)
	ApporveReviewService(reviewID uint) error
	RejectReviewService(reviewID uint) error
}

type RepoInterface interface {
	CreateNewReview(review *models.Review) error
	GetAllReviews() ([]models.Review, error)
	ApproveReview(id uint) error
	RejectReview(id uint) error
}
