package reviewservice

import (
	reviewdto "github.com/ak-repo/ecommerce-gin/internals/review/review_DTO"
	reviewinterface "github.com/ak-repo/ecommerce-gin/internals/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/models"
)

type ReviewService struct {
	ReviewRepo reviewinterface.RepoInterface
}

func NewReviewService(repo reviewinterface.RepoInterface) reviewinterface.ServiceInterface {
	return &ReviewService{ReviewRepo: repo}
}

// -------------admin------------------
func (s *ReviewService) ListAllReviewsForVerifyService() ([]reviewdto.AdminReviewResponse, error) {
	data, err := s.ReviewRepo.GetAllReviews()
	if err != nil {
		return nil, err

	}

	var reviews []reviewdto.AdminReviewResponse
	for _, r := range data {
		review := reviewdto.AdminReviewResponse{
			ID:          r.ID,
			ProductID:   r.ProductID,
			ProductName: r.Product.Title,
			UserID:      r.UserID,
			UserName:    r.User.Username,
			Rating:      r.Rating,
			Comment:     r.Comment,
			Status:      r.Status,
			CreatedAt:   r.CreatedAt,
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

// approve
func (s *ReviewService) ApporveReviewService(reviewID uint) error {

	return s.ReviewRepo.ApproveReview(reviewID)

}

// approve
func (s *ReviewService) RejectReviewService(reviewID uint) error {

	return s.ReviewRepo.RejectReview(reviewID)

}

// customer review add
func (s *ReviewService) CreateNewReviewService(req *reviewdto.CreateReviewRequest) error {

	review := models.Review{
		ProductID: req.ProductID,
		UserID:    req.UserID,
		Comment:   req.Comment,
		Rating:    req.Rating,
	}
	return s.ReviewRepo.CreateNewReview(&review)
}
