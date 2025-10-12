package authrepo

import (
	"errors"

	authinterface "github.com/ak-repo/ecommerce-gin/internals/auth/auth_interface"
	"github.com/ak-repo/ecommerce-gin/models"
	"gorm.io/gorm"
)

type authRepo struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) authinterface.Repository {
	return &authRepo{DB: db}
}

// user registration
func (r *authRepo) Registeration(user *models.User) error {

	return r.DB.Create(user).Error
}

// return user details
func (r *authRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email=?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}

// user by id
func (r *authRepo) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil

}
