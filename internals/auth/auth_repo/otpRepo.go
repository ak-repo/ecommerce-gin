package authrepo

import (
	"errors"
	"time"

	"github.com/ak-repo/ecommerce-gin/internals/auth"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

// otp  sent and verifcation
func (r *AuthRepo) CreateOTP(record *models.EmailOTP) error {
	return r.DB.Create(record).Error
}

func (r *AuthRepo) DeleteOTP(record *models.EmailOTP) error {
	return r.DB.Delete(record).Error
}

func (r *AuthRepo) VerifyOTP(req *auth.VerifyOTPRequest) (*models.EmailOTP, error) {
	var record models.EmailOTP
	err := r.DB.Where("email=? AND used=? AND expires_at>=?", req.Email, false, time.Now()).Order("created_at desc").First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no valid OTP found or it expired")
		}
		return nil, err
	}

	return &record, nil
}

func (r *AuthRepo) UpdateOTP(record *models.EmailOTP) error {
	return r.DB.Save(record).Error
}

// email verification true
func (r *AuthRepo) UserEmailVerified(userID uint) error {

	return r.DB.Model(&models.User{}).Where("id=?", userID).Update("email_verified", true).Error
}
