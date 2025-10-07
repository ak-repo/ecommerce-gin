package authrepo

import "github.com/ak-repo/ecommerce-gin/models"

func (r *AuthRepo) PasswordChange(userID uint, password string) error {
	return r.DB.Model(&models.User{}).Where("id=?", userID).Update("password_hash", password).Error
}
