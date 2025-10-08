package reviewrepo

import (
	reviewinterface "github.com/ak-repo/ecommerce-gin/internals/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type ReviewRepo struct {
	DB *gorm.DB
}

func NewReviewRepo(db *gorm.DB) reviewinterface.RepoInterface {
	return &ReviewRepo{DB: db}
}

// create review
func (r *ReviewRepo) CreateNewReview(review *models.Review) error {
	return r.DB.Create(review).Error
}

// get all reviews
func (r *ReviewRepo) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.
		Preload("User").
		Preload("Product").
		Find(&reviews).Error
	return reviews, err
}

// approve
func (r *ReviewRepo) ApproveReview(id uint) error {

	return r.DB.Model(&models.Review{}).Where("id=?", id).Update("status", "APPROVED").Error
}

func (r *ReviewRepo) RejectReview(id uint) error {

	return r.DB.Model(&models.Review{}).Where("id=?", id).Update("status", "REJECTED").Error
}
