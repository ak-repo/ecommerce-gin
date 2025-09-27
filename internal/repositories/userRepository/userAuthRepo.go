package userrepository

import (
	"github.com/ak-repo/ecommerce-gin/internal/models"
	"gorm.io/gorm"
)

type UserAuthRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(user *models.User)bool
}

type userAuthRepo struct {
	DB *gorm.DB
}

// Init auth repo
func NewUserAuthRepo(db *gorm.DB) UserAuthRepo {
	return &userAuthRepo{DB: db}
}

// user registration
func (r *userAuthRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// Check user email already register
func (r *userAuthRepo) GetUserByEmail(user *models.User) bool {
	if err := r.DB.Where("email=?", user.Email).First(user).Error; err != nil {
		return false
	}
	return true
}
