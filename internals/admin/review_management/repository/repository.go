package reviewrepo

import (
	reviewinter "github.com/ak-repo/ecommerce-gin/internals/admin/review_management/review_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewReviewRepoMG(db *gorm.DB) reviewinter.Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.
		Preload("User").
		Preload("Product").
		Order("created_at DESC").
		Find(&reviews).Error
	return reviews, err
}

func (r *repository) ApproveReview(id uint) error {

	return r.DB.Model(&models.Review{}).Where("id=?", id).Update("status", "APPROVED").Error
}

func (r *repository) RejectReview(id uint) error {

	return r.DB.Model(&models.Review{}).Where("id=?", id).Update("status", "REJECTED").Error
}
