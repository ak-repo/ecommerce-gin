package reviewsvc

import (
	reviewdto "github.com/ak-repo/ecommerce-gin/internals/admin/review_management/dto"
	reviewinterface "github.com/ak-repo/ecommerce-gin/internals/admin/review_management/review_interface"
)

type service struct {
	ReviewRepo reviewinterface.Repository
}

func NewReviewServiceMG(repo reviewinterface.Repository) reviewinterface.Service {
	return &service{ReviewRepo: repo}
}

func (s *service) GetAllReviews() ([]reviewdto.AdminReviewResponse, error) {
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

func (s *service) ApproveReview(reviewID uint) error {

	return s.ReviewRepo.ApproveReview(reviewID)

}

func (s *service) RejectReview(reviewID uint) error {

	return s.ReviewRepo.RejectReview(reviewID)

}
