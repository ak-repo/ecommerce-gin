package reviewrepository

import (
	reviewinterface "github.com/ak-repo/ecommerce-gin/internals/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewReviewRepo(db *gorm.DB) reviewinterface.Repository {
	return &repository{DB: db}
}

func (r *repository) AddReview(review *models.Review) error {
	return r.DB.Create(review).Error
}

func (r *repository) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.
		Preload("User").
		Preload("Product").
		Find(&reviews).Error
	return reviews, err
}

func (r *repository) ApproveReview(id uint) error {

	return r.DB.Model(&models.Review{}).Where("id=?", id).Update("status", "APPROVED").Error
}

func (r *repository) RejectReview(id uint) error {

	return r.DB.Model(&models.Review{}).Where("id=?", id).Update("status", "REJECTED").Error
}
