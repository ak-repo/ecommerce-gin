package reviewsvc

import (
	custreviewdto "github.com/ak-repo/ecommerce-gin/internals/customer/review/dto"
	custreviewinter "github.com/ak-repo/ecommerce-gin/internals/customer/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
)

type service struct {
	Repo custreviewinter.Repository
}

func NewReviewService(repo custreviewinter.Repository) custreviewinter.Service {
	return &service{Repo: repo}
}

func (s *service) AddReview(req *custreviewdto.CreateReviewRequest) error {

	review := models.Review{
		ProductID: req.ProductID,
		UserID:    req.UserID,
		Comment:   req.Comment,
		Rating:    req.Rating,
	}
	return s.Repo.AddReview(&review)
}
