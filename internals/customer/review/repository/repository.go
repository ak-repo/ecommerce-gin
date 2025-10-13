package reviewrepo

import (
	custreviewinter "github.com/ak-repo/ecommerce-gin/internals/customer/review/review_interface"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewReviewRepo(db *gorm.DB) custreviewinter.Repository {
	return &repository{DB: db}
}

func (r *repository) AddReview(review *models.Review) error {
	return r.DB.Create(review).Error
}
