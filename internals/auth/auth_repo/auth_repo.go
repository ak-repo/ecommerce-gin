package authrepo

import (
	"errors"

	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) authinterface.AuthRepoInterface {
	return &AuthRepo{DB: db}
}

// user registration
func (r *AuthRepo) CreateUser(username, email, password, role string) error {
	user := models.User{
		Email:         email,
		Username:      username,
		PasswordHash:  password,
		Role:          role,
		Status:        "active",
		EmailVerified: false,
	}
	return r.DB.Create(&user).Error
}

// return user details
func (r *AuthRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// user by id
func (r *AuthRepo) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}

// password change
func (r *AuthRepo) UpdatePassword(userID uint, hashPassword string) error {
	return r.DB.Model(&models.User{}).Where("id=?", userID).Update("password_hash", hashPassword).Error

}
