package userrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(user *models.User) error
	GetUserProfile(address *models.Address) error
}

type userRepo struct {
	DB *gorm.DB
}

// Init auth repo
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

// user registration
func (r *userRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// Check user email already register
func (r *userRepo) GetUserByEmail(user *models.User) error {
	return r.DB.Where("email=?", user.Email).First(user).Error

}

// Get user profile
func (r *userRepo) GetUserProfile(address *models.Address) error {

	return r.DB.Where("user_id", address.UserID).Find(address).Error
}
