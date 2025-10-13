package authrepo

import (
	"errors"
	"time"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/internals/models"
	"gorm.io/gorm"
)

// otp  sent and verifcation
func (r *authRepo) CreateOTP(record *models.EmailOTP) error {
	return r.DB.Create(record).Error
}

func (r *authRepo) DeleteOTP(record *models.EmailOTP) error {
	return r.DB.Delete(record).Error
}

func (r *authRepo) VerifyOTP(req *auth.VerifyOTPRequest) (*models.EmailOTP, error) {
	var record models.EmailOTP
	err := r.DB.Where("email=? AND used=? AND expires_at>=?", req.Email, false, time.Now()).Order("created_at desc").First(&record).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("no valid OTP found or it expired")
	}
	return &record, err
}

func (r *authRepo) UpdateOTP(record *models.EmailOTP) error {
	return r.DB.Save(record).Error
}

func (r *authRepo) UserEmailVerified(userID uint) error {

	return r.DB.Model(&models.User{}).Where("id=?", userID).Update("email_verified", true).Error
}
